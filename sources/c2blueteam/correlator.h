/*
 * correlator.h — Moteur à état ARCHTIME et corrélation temporelle multi-flux.
 * Conforme C99 strict, zéro allocation dynamique, table fixe 1024 entrées.
 */

#ifndef C2BT_CORRELATOR_H
#define C2BT_CORRELATOR_H

#include "c2blueteam.h"
#include <stdint.h>
#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Masques binaires de sous-systèmes pour corrélation temporelle */
#define C2BT_MASK_SUB_PROC    (1U << C2BT_SUB_PROC)    /* Bit 1 (0x02) : Processus */
#define C2BT_MASK_SUB_FS      (1U << C2BT_SUB_FS)      /* Bit 2 (0x04) : Fichiers */
#define C2BT_MASK_SUB_NET     (1U << C2BT_SUB_NET)     /* Bit 3 (0x08) : Réseau */
#define C2BT_MASK_SUB_MCP     (1U << C2BT_SUB_MCP)     /* Bit 4 (0x10) : Outils d'agents MCP */
#define C2BT_MASK_SUB_HARNESS (1U << C2BT_SUB_HARNESS) /* Bit 5 (0x20) : Harnais / Doctrine */
#define C2BT_MASK_SUB_ENTROPY (1U << C2BT_SUB_ENTROPY) /* Bit 6 (0x40) : Entropie */
#define C2BT_MASK_SUB_GPU     (1U << C2BT_SUB_GPU)     /* Bit 7 (0x80) : GPU / VRAM */

/* Barème de pondération des événements individuels */
#define C2BT_SCORE_BASE_EVENT       5   /* Score d'activité de base */
#define C2BT_SCORE_LOLBAS           40  /* Binaire détourné / LOLBAS */
#define C2BT_SCORE_ANOMALY          30  /* Anomalie détectée */
#define C2BT_SCORE_BLOCKED          50  /* Événement bloqué */
#define C2BT_SCORE_CRYPTO_PAYLOAD   35  /* Charge utile chiffrée / aléatoire pure */
#define C2BT_SCORE_BASE64_PAYLOAD   25  /* Charge utile Base64 / obfusquée */
#define C2BT_SCORE_HEX_PAYLOAD      20  /* Charge utile hexadécimale */
#define C2BT_SCORE_SUSPICIOUS_MCP   40  /* Commande destructrice ou stager MCP */
#define C2BT_SCORE_FS_MUTATION      40  /* Altération ou écriture fichier sensible */
#define C2BT_SCORE_NET_ACTIVITY     25  /* Connexion ou flux réseau */

/* Bonus de corrélation croisée multi-flux */
#define C2BT_BONUS_MCP_PROC         40  /* Corrélation MCP + Spawning LOLBAS / PROC */
#define C2BT_BONUS_ENTROPY_EXEC     45  /* Corrélation Entropie élevée + Exécution */
#define C2BT_BONUS_FS_PROC          40  /* Corrélation Altération FS + LOLBAS / Spawn */
#define C2BT_BONUS_FS_NET           45  /* Corrélation Altération FS + Exfiltration Réseau */
#define C2BT_BONUS_MULTI_3SUB       50  /* Corrélation multi-flux (>= 3 sous-systèmes distincts) */

/* API du moteur de corrélation temporelle */
void c2bt_tracker_init(c2bt_tracker_table_t *table);
void c2bt_tracker_reset_entry(c2bt_process_tracker_t *entry, uint32_t pid, uint64_t ts_ns);
int  c2bt_correlate_event(c2bt_tracker_table_t *table, const probe_event_t *ev, uint32_t *out_flags, uint64_t window_ns);

#ifdef __cplusplus
}
#endif

#endif /* C2BT_CORRELATOR_H */
