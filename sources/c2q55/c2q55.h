#ifndef C2Q55_H
#define C2Q55_H

#include <stdint.h>
#include <stddef.h>
#include <stdbool.h>

#define C2Q_BLOCK_SLOTS 8
#define C2Q_RING_CAP    1024
#define C2Q_RING_MASK   (C2Q_RING_CAP - 1)

#define C2Q_STATE_FREE      0x00000000U
#define C2Q_STATE_READY     0x00000001U
#define C2Q_STATE_RESERVED  0x00000002U
#define C2Q_STATE_COMMITTED 0x00000003U
#define C2Q_STATE_ACKED     0x00000004U
#define C2Q_STATE_POISON    0x00000005U
#define C2Q_STATE_DLQ       0x00000006U

/* Bloc de métadonnées AoSoA de 128 octets (2 lignes de cache L1 de 64B) */
typedef struct c2q_meta_block {
    uint64_t visible_at[C2Q_BLOCK_SLOTS]; /* 8 * 8 = 64 octets */
    uint32_t state[C2Q_BLOCK_SLOTS];      /* 8 * 4 = 32 octets */
    uint32_t flags[C2Q_BLOCK_SLOTS];      /* 8 * 4 = 32 octets */
} c2q_meta_block_t;

/* Slot individuel de 64 octets aligné sur 1 ligne de cache L1 */
typedef struct c2q_slot {
    uint64_t id_low;          /* 0..7   - UUIDv7 low 64 bits */
    uint64_t id_high;         /* 8..15  - UUIDv7 high 64 bits */
    uint64_t visible_at_ns;   /* 16..23 - Timestamp nanoseconde d'échéance */
    uint64_t lease_expire_ns; /* 24..31 - Timestamp nanoseconde d'expiration de bail */
    uint32_t topic_id;        /* 32..35 - Identifiant de sujet */
    uint32_t state;           /* 36..39 - État atomique du slot */
    uint32_t payload_len;     /* 40..43 - Longueur effective de la donnée (0..16) */
    uint32_t flags;           /* 44..47 - Flags et compteur retry */
    uint8_t  payload[16];     /* 48..63 - Donnée inlinée 0-alloc */
} c2q_slot_t;

/* Anneau SPSC circulaire avec isolation 128B anti-false sharing */
typedef struct c2q_ring {
    c2q_meta_block_t meta[C2Q_RING_CAP / C2Q_BLOCK_SLOTS];
    c2q_slot_t       slots[C2Q_RING_CAP];
    
    uint64_t head;
    uint64_t drops;
    uint8_t  pad_head[112]; /* Isolation 128B pour curseur d'écriture */
    
    uint64_t tail;
    uint64_t pad_tail_val;
    uint8_t  pad_tail[112]; /* Isolation 128B pour curseur de lecture */
} c2q_ring_t;

/* Prototypes des fonctions SIMD & branchless */
uint32_t c2q_scan_block_avx2(const c2q_meta_block_t *block, uint64_t now);
uint32_t c2q_scan_block_scalar(const c2q_meta_block_t *block, uint64_t now);
uint32_t c2q_scan_block(const c2q_meta_block_t *block, uint64_t now);

uint32_t c2q_trailing_zeros32(uint32_t mask);
uint32_t c2q_count_ones32(uint32_t mask);
uint64_t c2q_byteswap64(uint64_t val);
uint32_t c2q_crc32c_hw(uint32_t crc, const uint8_t *data, size_t len);

uint64_t c2q_select64(uint64_t condition, uint64_t val_true, uint64_t val_false);
void c2q_consume_mask(c2q_meta_block_t *block, uint32_t mask, uint32_t new_state);

#endif /* C2Q55_H */
