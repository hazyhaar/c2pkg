// SPDX-License-Identifier: Apache-2.0 OR MIT
// c2ordered_swiss: Table de hachage ordonnée SIMD (Tags H2 16 octets + Vecteur dense ordonné)
//
// Invariants de la table :
//   - capacity_hash est une puissance de deux >= 16 ; ctrl porte capacity_hash + 16 octets,
//     les 16 derniers dupliquent les 16 premiers (miroir) pour les lectures de groupe non alignées.
//   - Mode packed (FLAG_PACKED) : clés entières 0..n_used-1 stockées à leur propre indice dense,
//     aucun plan de contrôle utilisé. Toute insertion hors séquence bascule en mode hashed.
//   - Mode hashed : n_tombstones compte les trous denses (val.type == UNDEF sous n_used) et les
//     slots de contrôle CTRL_DELETED. swiss_compact remet les deux à zéro.
//   - Facteur de charge : un slot EMPTY neuf n'est pris que si (n_used + n_tombstones + 1) * 8
//     < capacity_hash * 7, soit strictement sous 87,5 % après insertion ; réoccuper un slot DELETED
//     est toujours accepté (occupation constante). Au-delà, l'insertion est refusée (false) :
//     l'appelant compacte ou réalloue, la table n'alloue jamais.
//   - Les clés chaîne ne sont pas copiées : la table conserve le pointeur fourni.

#include <stdint.h>
#include <stddef.h>
#include <stdbool.h>
#include <string.h>

#define CTRL_EMPTY     ((uint8_t)0x80)
#define CTRL_DELETED   ((uint8_t)0xFE)
#define CTRL_SENTINEL  ((uint8_t)0xFF)
#define H2_MASK        ((uint8_t)0x7F)

#define FLAG_INITIALIZED  (1U << 0)
#define FLAG_PACKED       (1U << 1)

#define ELEM_TYPE_UNDEF   0
#define ELEM_TYPE_INT     1
#define ELEM_TYPE_STR     2

typedef struct swiss_val_s {
    uint32_t type;
    uint32_t flags;
    int64_t  i64;
} swiss_val_t;

typedef struct swiss_element_s {
    swiss_val_t val;
    uint64_t    hash;
    char       *str_key;
    uint32_t    str_len;
    int64_t     int_key;
} swiss_element_t;

typedef struct swiss_table_s {
    uint8_t         *ctrl;
    uint32_t        *ctrl_to_dense;
    swiss_element_t *entries;
    uint32_t         capacity_hash;
    uint32_t         capacity_data;
    uint32_t         n_elements;
    uint32_t         n_used;
    uint32_t         n_tombstones;
    uint32_t         flags;
} swiss_table_t;

// Hash DJBX33A pour chaînes
uint64_t swiss_hash_str(const char *str, uint32_t len) {
    uint64_t h = 5381;
    for (uint32_t i = 0; i < len; ++i) {
        h = ((h << 5) + h) + (uint8_t)str[i];
    }
    return h | (h == 0);
}

// Hash entier 64-bit splitmix
uint64_t swiss_hash_int(int64_t val) {
    uint64_t z = (uint64_t)val + 0x9e3779b97f4a7c15ULL;
    z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9ULL;
    z = (z ^ (z >> 27)) * 0x94d049bb133111ebULL;
    return z ^ (z >> 31);
}

// Recherche scalaire de tag H2 dans un groupe de 16 octets
static inline uint32_t swiss_match_tag_16(const uint8_t *ctrl_group, uint8_t h2) {
    uint32_t mask = 0;
    for (uint32_t i = 0; i < 16; ++i) {
        if (ctrl_group[i] == h2) {
            mask |= (1U << i);
        }
    }
    return mask;
}

// Recherche de slot vide dans un groupe de 16 octets
static inline uint32_t swiss_match_empty_16(const uint8_t *ctrl_group) {
    uint32_t mask = 0;
    for (uint32_t i = 0; i < 16; ++i) {
        if (ctrl_group[i] == CTRL_EMPTY) {
            mask |= (1U << i);
        }
    }
    return mask;
}

// Recherche de slot vide OU supprimé (réutilisable) dans un groupe de 16 octets
static inline uint32_t swiss_match_free_16(const uint8_t *ctrl_group) {
    uint32_t mask = 0;
    for (uint32_t i = 0; i < 16; ++i) {
        if (ctrl_group[i] == CTRL_EMPTY || ctrl_group[i] == CTRL_DELETED) {
            mask |= (1U << i);
        }
    }
    return mask;
}

static inline uint32_t swiss_lowest_bit(uint32_t mask) {
    uint32_t bit_pos = 0;
    while ((mask & (1U << bit_pos)) == 0) bit_pos++;
    return bit_pos;
}

// Écriture d'un octet de contrôle avec entretien du miroir des 16 premiers slots
static inline void swiss_ctrl_set(swiss_table_t *ht, uint32_t slot_idx, uint8_t c) {
    ht->ctrl[slot_idx] = c;
    if (slot_idx < 16) {
        ht->ctrl[ht->capacity_hash + slot_idx] = c;
    }
}

// Vérifie que le plan de contrôle est utilisable : puissance de deux >= 16
static inline bool swiss_ctrl_usable(const swiss_table_t *ht) {
    const uint32_t cap = ht->capacity_hash;
    return ht->ctrl != NULL && ht->ctrl_to_dense != NULL && cap >= 16 && (cap & (cap - 1)) == 0;
}

// Facteur de charge : refuse si l'occupation après insertion atteindrait ou dépassait 87,5 %
static inline bool swiss_load_ok(const swiss_table_t *ht) {
    const uint64_t occupied = (uint64_t)ht->n_used + (uint64_t)ht->n_tombstones + 1;
    return occupied * 8 < (uint64_t)ht->capacity_hash * 7;
}

// Cherche le premier slot libre (EMPTY ou DELETED) de la séquence de sondage d'un hash.
// Retourne false si aucun slot libre n'existe (impossible sous le facteur de charge).
static bool swiss_ctrl_find_free(const swiss_table_t *ht, uint64_t hash, uint32_t *slot_out) {
    const uint32_t mask = ht->capacity_hash - 1;
    uint32_t group_idx = (uint32_t)((hash >> 7) & mask);
    uint32_t step = 0;

    while (step <= mask) {
        uint32_t free_mask = swiss_match_free_16(ht->ctrl + group_idx);
        if (free_mask != 0) {
            *slot_out = (group_idx + swiss_lowest_bit(free_mask)) & mask;
            return true;
        }
        step += 16;
        group_idx = (group_idx + step) & mask;
    }
    return false;
}

// Inscrit dense_idx dans un slot libre ; un slot DELETED réoccupé retire un tombstone.
static void swiss_ctrl_claim(swiss_table_t *ht, uint32_t slot_idx, uint64_t hash, uint32_t dense_idx) {
    if (ht->ctrl[slot_idx] == CTRL_DELETED) {
        ht->n_tombstones--;
    }
    swiss_ctrl_set(ht, slot_idx, (uint8_t)(hash & H2_MASK));
    ht->ctrl_to_dense[slot_idx] = dense_idx;
}

// Trouve un slot libre pour un hash donné et y inscrit dense_idx, sans consulter le seuil de charge.
static bool swiss_ctrl_place(swiss_table_t *ht, uint64_t hash, uint32_t dense_idx) {
    uint32_t slot_idx = 0;
    if (!swiss_ctrl_find_free(ht, hash, &slot_idx)) {
        return false;
    }
    swiss_ctrl_claim(ht, slot_idx, hash, dense_idx);
    return true;
}

// Marque CTRL_DELETED le slot de contrôle qui référence dense_idx pour ce hash.
static void swiss_ctrl_unlink(swiss_table_t *ht, uint64_t hash, uint32_t dense_idx) {
    const uint32_t mask = ht->capacity_hash - 1;
    const uint8_t h2 = (uint8_t)(hash & H2_MASK);
    uint32_t group_idx = (uint32_t)((hash >> 7) & mask);
    uint32_t step = 0;

    while (step <= mask) {
        const uint8_t *ctrl_group = ht->ctrl + group_idx;
        uint32_t match_mask = swiss_match_tag_16(ctrl_group, h2);
        while (match_mask != 0) {
            uint32_t slot_idx = (group_idx + swiss_lowest_bit(match_mask)) & mask;
            if (ht->ctrl_to_dense[slot_idx] == dense_idx) {
                swiss_ctrl_set(ht, slot_idx, CTRL_DELETED);
                ht->n_tombstones++;
                return;
            }
            match_mask &= (match_mask - 1);
        }
        if (swiss_match_empty_16(ctrl_group) != 0) {
            return;
        }
        step += 16;
        group_idx = (group_idx + step) & mask;
    }
}

// Reconstruit intégralement le plan de contrôle depuis le vecteur dense
static void swiss_ctrl_rebuild(swiss_table_t *ht) {
    memset(ht->ctrl, CTRL_EMPTY, ht->capacity_hash + 16);
    for (uint32_t i = 0; i < ht->n_used; ++i) {
        if (ht->entries[i].val.type == ELEM_TYPE_UNDEF) {
            continue;
        }
        uint64_t hash = ht->entries[i].hash;
        if (ht->entries[i].str_key == NULL) {
            hash = swiss_hash_int(ht->entries[i].int_key);
            ht->entries[i].hash = hash;
        }
        swiss_ctrl_place(ht, hash, i);
    }
}

// Initialisation d'une table
void swiss_init(
    swiss_table_t *ht,
    uint8_t *ctrl_buf,
    uint32_t *slots_buf,
    swiss_element_t *entries_buf,
    uint32_t cap_hash,
    uint32_t cap_data
) {
    ht->ctrl = ctrl_buf;
    ht->ctrl_to_dense = slots_buf;
    ht->entries = entries_buf;
    ht->capacity_hash = cap_hash;
    ht->capacity_data = cap_data;
    ht->n_elements = 0;
    ht->n_used = 0;
    ht->n_tombstones = 0;
    ht->flags = FLAG_INITIALIZED | FLAG_PACKED;

    if (ctrl_buf != NULL && cap_hash > 0) {
        memset(ctrl_buf, CTRL_EMPTY, cap_hash + 16);
    }
}

// Bascule packed -> hashed : construit le plan de contrôle. Idempotent en mode hashed.
bool swiss_unpack(swiss_table_t *ht) {
    if (!(ht->flags & FLAG_PACKED)) {
        return true;
    }
    if (!swiss_ctrl_usable(ht)) {
        return false;
    }
    ht->flags &= ~FLAG_PACKED;
    ht->n_tombstones = 0;
    swiss_ctrl_rebuild(ht);
    return true;
}

// Recherche par clé chaîne : retourne l'indice dense ou -1
int32_t swiss_find_str(
    const swiss_table_t *ht,
    const char *key,
    uint32_t key_len,
    uint64_t hash
) {
    if (ht->n_elements == 0 || (ht->flags & FLAG_PACKED)) {
        return -1;
    }

    const uint32_t mask = ht->capacity_hash - 1;
    const uint8_t h2 = (uint8_t)(hash & H2_MASK);
    uint32_t group_idx = (uint32_t)((hash >> 7) & mask);
    uint32_t step = 0;

    while (step <= mask) {
        const uint8_t *ctrl_group = ht->ctrl + group_idx;
        uint32_t match_mask = swiss_match_tag_16(ctrl_group, h2);

        while (match_mask != 0) {
            uint32_t slot_idx = (group_idx + swiss_lowest_bit(match_mask)) & mask;
            uint32_t dense_idx = ht->ctrl_to_dense[slot_idx];

            if (ht->entries[dense_idx].val.type != ELEM_TYPE_UNDEF &&
                ht->entries[dense_idx].hash == hash &&
                ht->entries[dense_idx].str_len == key_len &&
                ht->entries[dense_idx].str_key != NULL &&
                memcmp(ht->entries[dense_idx].str_key, key, key_len) == 0) {
                return (int32_t)dense_idx;
            }

            match_mask &= (match_mask - 1);
        }

        if (swiss_match_empty_16(ctrl_group) != 0) {
            return -1;
        }

        step += 16;
        group_idx = (group_idx + step) & mask;
    }
    return -1;
}

// Recherche par clé entière : retourne l'indice dense ou -1
int32_t swiss_find_int(
    const swiss_table_t *ht,
    int64_t key
) {
    if (ht->n_elements == 0) {
        return -1;
    }

    // Fast Path Packed : comparaison en 64 bits, aucune troncature de la clé
    if (ht->flags & FLAG_PACKED) {
        if (key >= 0 && key < (int64_t)ht->n_used) {
            uint32_t ukey = (uint32_t)key;
            if (ht->entries[ukey].val.type != ELEM_TYPE_UNDEF) {
                return (int32_t)ukey;
            }
        }
        return -1;
    }

    // Slow Path Hashed
    uint64_t hash = swiss_hash_int(key);
    const uint32_t mask = ht->capacity_hash - 1;
    const uint8_t h2 = (uint8_t)(hash & H2_MASK);
    uint32_t group_idx = (uint32_t)((hash >> 7) & mask);
    uint32_t step = 0;

    while (step <= mask) {
        const uint8_t *ctrl_group = ht->ctrl + group_idx;
        uint32_t match_mask = swiss_match_tag_16(ctrl_group, h2);

        while (match_mask != 0) {
            uint32_t slot_idx = (group_idx + swiss_lowest_bit(match_mask)) & mask;
            uint32_t dense_idx = ht->ctrl_to_dense[slot_idx];

            if (ht->entries[dense_idx].val.type != ELEM_TYPE_UNDEF &&
                ht->entries[dense_idx].str_key == NULL &&
                ht->entries[dense_idx].int_key == key) {
                return (int32_t)dense_idx;
            }

            match_mask &= (match_mask - 1);
        }

        if (swiss_match_empty_16(ctrl_group) != 0) {
            return -1;
        }

        step += 16;
        group_idx = (group_idx + step) & mask;
    }
    return -1;
}

// Insertion ou mise à jour en mode Packed (clé == n_used uniquement)
bool swiss_insert_packed_int(
    swiss_table_t *ht,
    int64_t key,
    uint32_t val_type,
    int64_t val_i64
) {
    if (!(ht->flags & FLAG_PACKED)) {
        return false;
    }

    if (key == (int64_t)ht->n_used && ht->n_used < ht->capacity_data) {
        uint32_t idx = ht->n_used;
        ht->n_used++;
        ht->entries[idx].val.type = val_type;
        ht->entries[idx].val.flags = 0;
        ht->entries[idx].val.i64 = val_i64;
        ht->entries[idx].hash = 0;
        ht->entries[idx].str_key = NULL;
        ht->entries[idx].str_len = 0;
        ht->entries[idx].int_key = key;
        ht->n_elements++;
        return true;
    }

    return false;
}

// Ajout d'un élément en fin de vecteur dense en mode hashed, avec liaison de contrôle.
static bool swiss_append_hashed(
    swiss_table_t *ht,
    uint64_t hash,
    char *str_key,
    uint32_t str_len,
    int64_t int_key,
    uint32_t val_type,
    int64_t val_i64
) {
    if (ht->n_used >= ht->capacity_data) {
        return false;
    }
    uint32_t slot_idx = 0;
    if (!swiss_ctrl_find_free(ht, hash, &slot_idx)) {
        return false;
    }
    // Réoccuper un slot DELETED laisse l'occupation (n_used + n_tombstones) inchangée :
    // seul un slot EMPTY neuf est soumis au seuil de charge.
    if (ht->ctrl[slot_idx] != CTRL_DELETED && !swiss_load_ok(ht)) {
        return false;
    }
    uint32_t idx = ht->n_used;
    swiss_ctrl_claim(ht, slot_idx, hash, idx);
    ht->n_used++;
    ht->entries[idx].val.type = val_type;
    ht->entries[idx].val.flags = 0;
    ht->entries[idx].val.i64 = val_i64;
    ht->entries[idx].hash = hash;
    ht->entries[idx].str_key = str_key;
    ht->entries[idx].str_len = str_len;
    ht->entries[idx].int_key = int_key;
    ht->n_elements++;
    return true;
}

// Insertion ou mise à jour par clé entière, tout mode.
// Mode packed : append si key == n_used, mise à jour si key déjà présente, sinon bascule hashed.
bool swiss_insert_int(
    swiss_table_t *ht,
    int64_t key,
    uint32_t val_type,
    int64_t val_i64
) {
    if (val_type == ELEM_TYPE_UNDEF) {
        return false;
    }

    int32_t found = swiss_find_int(ht, key);
    if (found >= 0) {
        ht->entries[found].val.type = val_type;
        ht->entries[found].val.flags = 0;
        ht->entries[found].val.i64 = val_i64;
        return true;
    }

    if (ht->flags & FLAG_PACKED) {
        if (key == (int64_t)ht->n_used) {
            return swiss_insert_packed_int(ht, key, val_type, val_i64);
        }
        if (!swiss_unpack(ht)) {
            return false;
        }
    }

    return swiss_append_hashed(ht, swiss_hash_int(key), NULL, 0, key, val_type, val_i64);
}

// Insertion ou mise à jour par clé chaîne. Bascule en mode hashed si nécessaire.
// Le pointeur key est conservé tel quel (aucune copie) ; il doit survivre à la table.
bool swiss_insert_str(
    swiss_table_t *ht,
    char *key,
    uint32_t key_len,
    uint32_t val_type,
    int64_t val_i64
) {
    if (key == NULL || val_type == ELEM_TYPE_UNDEF) {
        return false;
    }

    uint64_t hash = swiss_hash_str(key, key_len);

    if (ht->flags & FLAG_PACKED) {
        if (!swiss_unpack(ht)) {
            return false;
        }
    }

    int32_t found = swiss_find_str(ht, key, key_len, hash);
    if (found >= 0) {
        ht->entries[found].val.type = val_type;
        ht->entries[found].val.flags = 0;
        ht->entries[found].val.i64 = val_i64;
        return true;
    }

    return swiss_append_hashed(ht, hash, key, key_len, 0, val_type, val_i64);
}

// Compaction in-place éliminant les tombstones
void swiss_compact(swiss_table_t *ht) {
    if (ht->n_tombstones == 0) {
        return;
    }

    uint32_t write_idx = 0;
    const uint32_t total_used = ht->n_used;

    for (uint32_t read_idx = 0; read_idx < total_used; ++read_idx) {
        if (ht->entries[read_idx].val.type != ELEM_TYPE_UNDEF) {
            if (write_idx != read_idx) {
                ht->entries[write_idx] = ht->entries[read_idx];
            }
            write_idx++;
        }
    }

    ht->n_used = write_idx;
    ht->n_tombstones = 0;

    // Régénération du plan de contrôle si non packed
    if (!(ht->flags & FLAG_PACKED) && ht->ctrl != NULL) {
        swiss_ctrl_rebuild(ht);
    }
}

// Opération array_pop en O(1) amorti ; en mode hashed, libère le slot de contrôle
bool swiss_pop(swiss_table_t *ht) {
    if (ht->n_elements == 0 || ht->n_used == 0) {
        return false;
    }

    while (ht->n_used > 0) {
        uint32_t last_idx = ht->n_used - 1;
        if (ht->entries[last_idx].val.type != ELEM_TYPE_UNDEF) {
            if (!(ht->flags & FLAG_PACKED)) {
                swiss_ctrl_unlink(ht, ht->entries[last_idx].hash, last_idx);
            }
            ht->entries[last_idx].val.type = ELEM_TYPE_UNDEF;
            ht->n_used--;
            ht->n_elements--;
            return true;
        }
        // Trou dense en fin de vecteur : son tombstone disparaît avec lui
        if (ht->n_tombstones > 0) {
            ht->n_tombstones--;
        }
        ht->n_used--;
    }

    return false;
}
