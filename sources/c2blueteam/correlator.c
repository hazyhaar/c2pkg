/*
 * correlator.c — Implémentation du moteur de corrélation multi-événements fenêtré.
 * Conforme C99 strict, zéro allocation (0 B/op), table ARCHTIME 1024 entrées.
 */

#include "correlator.h"
#include <string.h>

void c2bt_tracker_init(c2bt_tracker_table_t *table) {
    if (!table) return;
    memset(table, 0, sizeof(*table));
}

void c2bt_tracker_reset_entry(c2bt_process_tracker_t *entry, uint32_t pid, uint64_t ts_ns) {
    if (!entry) return;
    entry->last_ts_ns = ts_ns;
    entry->pid = pid;
    entry->subsystems_mask = 0;
    entry->accumulated_score = 0;
    entry->last_action = 0;
    entry->event_count = 0;
    entry->reserved = 0;
}

int c2bt_correlate_event(c2bt_tracker_table_t *table, const probe_event_t *ev, uint32_t *out_flags, uint64_t window_ns) {
    if (!table || !ev || !out_flags) {
        return -1;
    }

    if (window_ns == 0) {
        window_ns = C2BT_CORRELATION_DEFAULT_WINDOW_NS;
    }

    /* Indexation déterministe ARCHTIME par PID ou condensé de session */
    uint32_t idx = (ev->pid != 0) ? (ev->pid & C2BT_TRACKER_MASK) : ((uint32_t)ev->src & C2BT_TRACKER_MASK);

    /* Évaluation de la fenêtre temporelle glissante */
    int reset = 0;
    if (table->entries[idx].last_ts_ns == 0 || table->entries[idx].pid != ev->pid) {
        reset = 1;
    } else if (ev->ts_ns < table->entries[idx].last_ts_ns) {
        /* Horloge non-monotone ou saut temporel négatif */
        reset = 1;
    } else {
        uint64_t delta_ns = ev->ts_ns - table->entries[idx].last_ts_ns;
        if (delta_ns > window_ns) {
            reset = 1;
        }
    }

    if (reset) {
        c2bt_tracker_reset_entry(&table->entries[idx], ev->pid, ev->ts_ns);
    }

    /* Mise à jour de l'état du suiveur pour l'événement courant */
    table->entries[idx].last_ts_ns = ev->ts_ns;
    table->entries[idx].event_count = table->entries[idx].event_count + 1;
    table->entries[idx].last_action = ev->action;

    if (ev->subsystem > 0 && ev->subsystem <= 7) {
        table->entries[idx].subsystems_mask = table->entries[idx].subsystems_mask | (1U << ev->subsystem);
    }

    /* 1. Évaluation et cumul du score individuel de l'événement */
    uint32_t event_score = C2BT_SCORE_BASE_EVENT;
    if (ev->flags & C2BT_FLAG_LOLBAS) {
        event_score += C2BT_SCORE_LOLBAS;
    }
    if (ev->flags & C2BT_FLAG_ANOMALY) {
        event_score += C2BT_SCORE_ANOMALY;
    }
    if (ev->flags & C2BT_FLAG_BLOCKED) {
        event_score += C2BT_SCORE_BLOCKED;
    }
    if (ev->flags & C2BT_FLAG_CRYPTO_PAYLOAD) {
        event_score += C2BT_SCORE_CRYPTO_PAYLOAD;
    }
    if (ev->flags & C2BT_FLAG_BASE64_PAYLOAD) {
        event_score += C2BT_SCORE_BASE64_PAYLOAD;
    }
    if (ev->flags & C2BT_FLAG_HEX_PAYLOAD) {
        event_score += C2BT_SCORE_HEX_PAYLOAD;
    }

    /* Pondération selon l'action et le sous-système */
    if (ev->subsystem == C2BT_SUB_MCP && ev->action == C2BT_ACT_TOOL_CALL) {
        event_score += 15;
    }
    if (ev->subsystem == C2BT_SUB_FS && (ev->action == C2BT_ACT_WRITE || (ev->flags & C2BT_FLAG_BLOCKED))) {
        event_score += C2BT_SCORE_FS_MUTATION;
    }
    if (ev->subsystem == C2BT_SUB_NET && ev->action == C2BT_ACT_CONNECT) {
        event_score += C2BT_SCORE_NET_ACTIVITY;
    }
    if (ev->subsystem == C2BT_SUB_ENTROPY && ev->src >= 1400) {
        event_score += 25;
    }

    table->entries[idx].accumulated_score = table->entries[idx].accumulated_score + event_score;

    /* 2. Détection des motifs combinés croisés multi-flux */
    uint32_t mask = table->entries[idx].subsystems_mask;

    /* Combinaison 1 : Outil MCP suspect / stager + Spawning LOLBAS / Processus */
    if ((mask & C2BT_MASK_SUB_MCP) && (mask & C2BT_MASK_SUB_PROC)) {
        if ((ev->flags & (C2BT_FLAG_LOLBAS | C2BT_FLAG_ANOMALY | C2BT_FLAG_BLOCKED)) || (table->entries[idx].accumulated_score >= 50)) {
            table->entries[idx].accumulated_score = table->entries[idx].accumulated_score + C2BT_BONUS_MCP_PROC;
        }
    }

    /* Combinaison 2 : Charge obfusquée / forte entropie + Exécution d'outil ou commande */
    if ((mask & C2BT_MASK_SUB_ENTROPY) && ((mask & C2BT_MASK_SUB_PROC) || (mask & C2BT_MASK_SUB_MCP))) {
        table->entries[idx].accumulated_score = table->entries[idx].accumulated_score + C2BT_BONUS_ENTROPY_EXEC;
    }

    /* Combinaison 3 : Altération doctrine / FS sensible + Exfiltration / Spawning */
    if (mask & C2BT_MASK_SUB_FS) {
        if (mask & C2BT_MASK_SUB_PROC) {
            table->entries[idx].accumulated_score = table->entries[idx].accumulated_score + C2BT_BONUS_FS_PROC;
        }
        if (mask & C2BT_MASK_SUB_NET) {
            table->entries[idx].accumulated_score = table->entries[idx].accumulated_score + C2BT_BONUS_FS_NET;
        }
    }

    /* Combinaison 4 : Multi-flux étendu (>= 3 sous-systèmes distincts mobilisés) */
    uint32_t sub_count = 0;
    for (uint32_t b = 1; b <= 7; b++) {
        if (mask & (1U << b)) {
            sub_count++;
        }
    }
    if (sub_count >= 3) {
        table->entries[idx].accumulated_score = table->entries[idx].accumulated_score + C2BT_BONUS_MULTI_3SUB;
    }

    /* 3. Détermination du verdict et levée des drapeaux de menace corrélée */
    uint32_t flags = ev->flags;
    if (table->entries[idx].accumulated_score >= C2BT_CORRELATION_THRESHOLD_CRITICAL) {
        flags |= (C2BT_FLAG_CORRELATED_THREAT | C2BT_FLAG_ANOMALY | C2BT_FLAG_BLOCKED);
        flags &= ~C2BT_FLAG_VERDICT_OK;
    }

    *out_flags = flags;
    return 0;
}
