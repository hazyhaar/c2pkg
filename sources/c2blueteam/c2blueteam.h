/*
 * c2blueteam.h — Spécification et contrat binaire de la bibliothèque de défense système.
 * Conforme C99 strict, zéro CGO au runtime, structure fixe 128 octets.
 */

#ifndef C2BLUETEAM_H
#define C2BLUETEAM_H

#include <stdint.h>
#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Capacité par défaut de chaque canal SPSC (puissance de deux obligatoire) */
#define C2BT_RING_CAP 1024
#define C2BT_RING_MASK (C2BT_RING_CAP - 1)
#define C2BT_BATCH_SIZE 8

/* Sous-systèmes d'émission (subsystem) */
#define C2BT_SUB_PROC    1  /* Processus : fork, exec, exit, ptrace */
#define C2BT_SUB_FS      2  /* Fichiers : fanotify open, modify, delete */
#define C2BT_SUB_NET     3  /* Réseau : Netlink connect, bind, beaconing */
#define C2BT_SUB_MCP     4  /* Outils d'agents : JSON-RPC Unix socket */
#define C2BT_SUB_HARNESS 5  /* Intégrité de la doctrine (AGENTS.md, etc.) */
#define C2BT_SUB_ENTROPY 6  /* Analyse mathématique de flux/tokens */
#define C2BT_SUB_GPU     7  /* VRAM et allocations de tenseurs */

/* Actions (action) */
#define C2BT_ACT_EXEC      1
#define C2BT_ACT_OPEN      2
#define C2BT_ACT_CONNECT   3
#define C2BT_ACT_TOOL_CALL 4
#define C2BT_ACT_READ      5
#define C2BT_ACT_WRITE     6
#define C2BT_ACT_MMAP_EXEC 7

/* Drapeaux et verdicts (flags) */
#define C2BT_FLAG_NONE           0x0000
#define C2BT_FLAG_VERDICT_OK     0x0001
#define C2BT_FLAG_BLOCKED        0x0002
#define C2BT_FLAG_ANOMALY        0x0004
#define C2BT_FLAG_LOLBAS         0x0008
#define C2BT_FLAG_BEACONING      0x0010
#define C2BT_FLAG_DRIFT          0x0020
#define C2BT_FLAG_HEX_PAYLOAD    0x0040  /* Charge utile hexadécimale détectée */
#define C2BT_FLAG_BASE64_PAYLOAD 0x0080  /* Charge utile Base64 détectée */
#define C2BT_FLAG_CRYPTO_PAYLOAD 0x0100  /* Données chiffrées / haute entropie */
#define C2BT_FLAG_CORRELATED_THREAT 0x0200 /* Menace corrélée multi-flux / multi-événements */

/* Masques de classes de caractères ARCHTIME (8 bits) */
#define C2BT_CHAR_IS_PRINTABLE   0x01U  /* Caractères imprimables ASCII (0x20..0x7E, \t, \n, \r) */
#define C2BT_CHAR_IS_HEX         0x02U  /* Alphabet hexadécimal ('0'..'9', 'a'..'f', 'A'..'F') */
#define C2BT_CHAR_IS_BASE64      0x04U  /* Alphabet Base64/Base64URL ('A'..'Z', 'a'..'z', '0'..'9', '+', '/', '=', '-', '_') */
#define C2BT_CHAR_IS_CONTROL     0x08U  /* Octets de contrôle (0x00..0x1F hors whitespace, 0x7F) */
#define C2BT_CHAR_IS_HIGH_BYTE   0x10U  /* Octets de poids fort non-ASCII (0x80..0xFF) */
#define C2BT_CHAR_IS_WHITESPACE  0x20U  /* Espaces blancs (' ', '\t', '\n', '\r') */
#define C2BT_CHAR_IS_DOT         0x40U  /* Séparateur point '.' (0x2E) pour JWT / structuré */

/* Alias de conformité avec les spécifications du taskbook */
#define IS_PRINTABLE   C2BT_CHAR_IS_PRINTABLE
#define IS_HEX         C2BT_CHAR_IS_HEX
#define IS_BASE64      C2BT_CHAR_IS_BASE64
#define IS_CONTROL     C2BT_CHAR_IS_CONTROL
#define IS_HIGH_BYTE   C2BT_CHAR_IS_HIGH_BYTE
#define IS_WHITESPACE  C2BT_CHAR_IS_WHITESPACE
#define IS_DOT         C2BT_CHAR_IS_DOT

/* Classes de classification conjointe de charge utile (payload_class) */
#define C2BT_PAYLOAD_CLASS_UNKNOWN           0  /* Échantillon indéterminé ou trop court (< 8 octets) */
#define C2BT_PAYLOAD_CLASS_PROSE             1  /* Texte naturel, prose, documentation ou code source */
#define C2BT_PAYLOAD_CLASS_HEX               2  /* Obfuscation, clé ou charge utile hexadécimale (Base16) */
#define C2BT_PAYLOAD_CLASS_BASE64            3  /* Charge utile encodée en Base64 / Base64URL */
#define C2BT_PAYLOAD_CLASS_CRYPTO_COMPRESSED 4  /* Données chiffrées (AES, ChaCha20) ou compressées (gzip, zstd) */
#define C2BT_PAYLOAD_CLASS_JWT               5  /* Jeton d'authentification structuré (JSON Web Token) */

/* Profil métrologique complet de charge utile (sans allocation dynamique) */
typedef struct c2bt_entropy_profile {
    uint32_t entropy_q8;      /* Entropie de Shannon en virgule fixe Q8.8 (0..2048) */
    uint32_t char_mask;       /* Masque OU combiné de toutes les classes d'octets présentes */
    uint16_t distinct_count;  /* Nombre de symboles uniques rencontrés (0..256) */
    uint16_t payload_class;   /* Classification déterministe (C2BT_PAYLOAD_CLASS_*) */
    size_t   len;             /* Longueur analysée en octets */
} c2bt_entropy_profile_t;


/*
 * Structure universelle d'événement (probe_event) :
 * Strictement 128 octets, sans aucun pointeur, alignée sur 128 octets.
 */
typedef struct probe_event {
    uint64_t ts_ns;       /* 8 octets @0  : Horodatage nanoseconde (CLOCK_MONOTONIC_RAW) */
    uint32_t pid;         /* 4 octets @8  : PID du processus */
    uint32_t tid;         /* 4 octets @12 : TID / Thread ID */
    uint16_t subsystem;   /* 2 octets @16 : Sous-système source */
    uint16_t action;      /* 2 octets @18 : Action captée */
    uint32_t flags;       /* 4 octets @20 : Drapeaux de verdict et classification */
    uint64_t src;         /* 8 octets @24 : Identifiant source numérique ou descripteur */
    uint8_t  payload[96]; /* 96 octets @32: Charge utile fixe (chemin, nom outil, extrait) */
} probe_event_t;

/* Gardes statiques de compilation C99 : vérification bit-exacte de la taille et des offsets */
typedef char _c2bt_assert_size[(sizeof(probe_event_t) == 128) ? 1 : -1];
typedef char _c2bt_assert_off_ts[(offsetof(probe_event_t, ts_ns) == 0) ? 1 : -1];
typedef char _c2bt_assert_off_pid[(offsetof(probe_event_t, pid) == 8) ? 1 : -1];
typedef char _c2bt_assert_off_src[(offsetof(probe_event_t, src) == 24) ? 1 : -1];
typedef char _c2bt_assert_off_payload[(offsetof(probe_event_t, payload) == 32) ? 1 : -1];

/* Canal SPSC lock-free à alignement cache-line strict (évite le cache-bouncing) */
typedef struct probe_channel {
    probe_event_t slots[C2BT_RING_CAP];
    uint64_t head;            /* Index d'écriture producteur (aligné 64 octets) */
    uint64_t drops;           /* Compteur atomique de rejet sur saturation (producteur) */
    uint8_t  pad1[48];        /* Padding cache-line producteur (64 - 8 - 8 = 48 octets) */
    uint64_t tail;            /* Index de lecture consommateur (aligné 64 octets) */
    uint8_t  pad2[56];        /* Padding cache-line consommateur (64 - 8 = 56 octets) */
} probe_channel_t;

/* Gardes statiques de compilation C99 pour probe_channel_t */
typedef char _c2bt_assert_chan_size[(sizeof(probe_channel_t) == 131200) ? 1 : -1];
typedef char _c2bt_assert_chan_off_head[(offsetof(probe_channel_t, head) == 131072) ? 1 : -1];
typedef char _c2bt_assert_chan_off_drops[(offsetof(probe_channel_t, drops) == 131080) ? 1 : -1];
typedef char _c2bt_assert_chan_off_tail[(offsetof(probe_channel_t, tail) == 131136) ? 1 : -1];

/* Modes d'application et politiques de sécurité (enforce_mode & fail_safe_policy) */
#define C2BT_MODE_PASSIVE       0  /* AUDIT_ONLY : observation seule, pas de veto actif */
#define C2BT_MODE_ACTIVE        1  /* ENFORCE : application effective du veto actif */

#define C2BT_POLICY_FAIL_OPEN   0  /* Laisser passer (FAN_ALLOW / flux RPC) en cas d'erreur/saturation */
#define C2BT_POLICY_FAIL_CLOSED 1  /* Bloquer (FAN_DENY / rejet RPC) en cas d'erreur/saturation */

/* Configuration statique */
typedef struct c2bt_config {
    uint8_t enable_proc;      /* 1 = actif, 0 = inactif */
    uint8_t enable_file;
    uint8_t enable_net;
    uint8_t enable_mcp;
    uint8_t enable_entropy;
    uint8_t enable_gpu;
    uint8_t enforce_mode;     /* 0 = PASSIVE / AUDIT_ONLY, 1 = ACTIVE / ENFORCE */
    uint8_t fail_safe_policy; /* 0 = FAIL_OPEN, 1 = FAIL_CLOSED */
    const char *mcp_socket_path;  /* Chemin socket Unix cible (ex: /tmp/mcp.sock) */
    const char *harness_dir;      /* Répertoire doctrine protégé (ex: .claude/) */
} c2bt_config_t;

/* Capacité fixe de la table de suivi de processus (puissance de deux) */
#define C2BT_TRACKER_CAP 1024
#define C2BT_TRACKER_MASK (C2BT_TRACKER_CAP - 1)

#define C2BT_CORRELATION_DEFAULT_WINDOW_NS 1000000000ULL /* Fenêtre 1s */
#define C2BT_CORRELATION_THRESHOLD_CRITICAL 100

/* Structure ARCHTIME d'entrée de suivi de processus (32 octets) */
typedef struct c2bt_process_tracker {
    uint64_t last_ts_ns;        /* 8 octets @0  : Dernier horodatage capté */
    uint32_t pid;               /* 4 octets @8  : Identifiant PID du processus */
    uint32_t subsystems_mask;   /* 4 octets @12 : Masque binaire des sous-systèmes */
    uint32_t accumulated_score; /* 4 octets @16 : Score de menace pondéré cumulé */
    uint16_t last_action;       /* 2 octets @20 : Dernière action */
    uint16_t event_count;       /* 2 octets @22 : Nombre d'événements */
    uint32_t reserved;          /* 4 octets @24 : Remplissage d'alignement */
} c2bt_process_tracker_t;

typedef char _c2bt_assert_tracker_size[(sizeof(c2bt_process_tracker_t) == 32) ? 1 : -1];

/* Table fixe pré-allouée de 1024 entrées (32 KiB) */
typedef struct c2bt_tracker_table {
    c2bt_process_tracker_t entries[C2BT_TRACKER_CAP];
} c2bt_tracker_table_t;

typedef char _c2bt_assert_table_size[(sizeof(c2bt_tracker_table_t) == (1024 * 32)) ? 1 : -1];

/* Contexte complet sans allocation dynamique (in-place) */
typedef struct c2bt_ctx {
    c2bt_config_t        config;
    probe_channel_t      chan_proc;
    probe_channel_t      chan_file;
    probe_channel_t      chan_net;
    probe_channel_t      chan_mcp;
    c2bt_tracker_table_t tracker_table; /* Table ARCHTIME 32 KiB pour corrélation temporelle */
    probe_event_t        batch_buf[C2BT_BATCH_SIZE];
    volatile int         running;
    int                  sealed_mem_fd; /* Descripteur memfd_create scellé */
    int                  fanotify_fd;   /* Descripteur fanotify */
    int                  mcp_proxy_fd;  /* Descripteur socket proxy MCP */
} c2bt_ctx_t;

/* API publique C2BLUETEAM */
int  c2bt_init_inplace(c2bt_ctx_t *ctx, const c2bt_config_t *cfg);
int  c2bt_start(c2bt_ctx_t *ctx);
int  c2bt_stop(c2bt_ctx_t *ctx);
int  c2bt_poll_batch(c2bt_ctx_t *ctx, probe_event_t *out_batch, int max_events);

/* Fonctions utilitaires rapides et déterministes */
uint64_t c2bt_time_ns_raw(void);
uint32_t c2bt_calc_entropy_8_8(const uint8_t *data, size_t len);
int      c2bt_profile_payload(const uint8_t *data, size_t len, c2bt_entropy_profile_t *out_profile);
uint16_t c2bt_classify_payload(const uint8_t *data, size_t len, uint32_t *out_entropy_q8, uint32_t *out_char_mask);
int      c2bt_eval_rules_batch(const probe_event_t *in_events, probe_event_t *out_events, int count);
int      c2bt_probe_mcp_inspect(const char *json_buf, size_t len, char *out_tool, size_t out_max);
int      c2bt_probe_mcp_extract_deep(const char *json_buf, size_t len, char *out_payload, size_t max_payload);
int      c2bt_probe_mcp_filter_call(const char *json_in, size_t in_len, char *json_out, size_t max_out, uint32_t flags, const c2bt_config_t *cfg);
int      c2bt_probe_fs_init(int *out_fan_fd, const char *protected_dir);
int      c2bt_probe_fs_init_perm(int *out_fan_fd, const char *protected_dir);
int      c2bt_probe_fs_verdict(int fanotify_fd, int fd_event, uint32_t flags, const c2bt_config_t *cfg);
int      c2bt_probe_fs_poll(int fan_fd, probe_channel_t *ch);
uint64_t c2bt_channel_get_drops(const probe_channel_t *ch);
void     c2bt_tracker_init(c2bt_tracker_table_t *table);
int      c2bt_correlate_event(c2bt_tracker_table_t *table, const probe_event_t *ev, uint32_t *out_flags, uint64_t window_ns);

#ifdef __cplusplus
}
#endif

#endif /* C2BLUETEAM_H */
