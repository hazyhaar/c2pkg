/*
 * rules_simd.c — Moteur d'évaluation de règles par lots sans branchement.
 */

#include "c2blueteam.h"
#include <string.h>

/*
 * Extraction et normalisation déterministe LOLBAS :
 * 1. Délimitation du premier token (avant espace / tabulation / saut de ligne).
 * 2. Extraction du nom de base (basename) après le dernier slash '/'.
 * 3. Filtrage des binaires d'attaque, interpréteurs versionnés et shells.
 */
static int check_lolbas_comm(const char *comm) {
    if (!comm || comm[0] == '\0') return 0;
    
    int start = 0;
    for (; comm[start] != '\0'; start++) {
        char c = comm[start];
        if (c != ' ' && c != '\t') {
            break;
        }
    }
    if (comm[start] == '\0') return 0;
    
    int end = start;
    for (; comm[end] != '\0'; end++) {
        char c = comm[end];
        if (c == ' ' || c == '\t' || c == '\r' || c == '\n') {
            break;
        }
    }
    
    int base_start = start;
    for (int i = start; i < end; i++) {
        if (comm[i] == '/') {
            base_start = i + 1;
        }
    }
    
    int base_len = end - base_start;
    if (base_len <= 0) return 0;
    
    const char *base = &comm[base_start];
    
    /* 1. Outils de transfert réseau / exfiltration / sockets */
    if (base_len == 4 && strncmp(base, "curl", 4) == 0) return 1;
    if (base_len == 4 && strncmp(base, "wget", 4) == 0) return 1;
    if (base_len == 2 && strncmp(base, "nc", 2) == 0) return 1;
    if (base_len == 4 && strncmp(base, "ncat", 4) == 0) return 1;
    if (base_len == 6 && strncmp(base, "netcat", 6) == 0) return 1;
    if (base_len == 5 && strncmp(base, "socat", 5) == 0) return 1;
    
    /* 2. Outils d'encodage */
    if (base_len == 6 && strncmp(base, "base64", 6) == 0) return 1;
    
    /* 3. Interpréteurs de scripts (exacts ou versionnés: python, python3, python3.11, etc.) */
    if (base_len >= 6 && strncmp(base, "python", 6) == 0) return 1;
    if (base_len == 4 && strncmp(base, "perl", 4) == 0) return 1;
    if (base_len == 4 && strncmp(base, "ruby", 4) == 0) return 1;
    if (base_len >= 3 && strncmp(base, "php", 3) == 0) return 1;
    if (base_len >= 3 && strncmp(base, "lua", 3) == 0) return 1;
    
    /* 4. Shells système */
    if (base_len == 2 && strncmp(base, "sh", 2) == 0) return 1;
    if (base_len == 4 && strncmp(base, "bash", 4) == 0) return 1;
    if (base_len == 3 && strncmp(base, "zsh", 3) == 0) return 1;
    if (base_len == 4 && strncmp(base, "dash", 4) == 0) return 1;
    if (base_len == 3 && strncmp(base, "ash", 3) == 0) return 1;
    
    return 0;
}

int c2bt_eval_rules_batch(const probe_event_t *in_events, probe_event_t *out_events, int count) {
    if (count <= 0) return 0;
    if (!in_events) return 0;
    if (!out_events) return 0;
    
    for (int i = 0; i < count; i++) {
        out_events[i] = in_events[i];
        uint32_t flags = out_events[i].flags;
        const char *pl = (const char *)out_events[i].payload;
        
        switch (out_events[i].subsystem) {
        case C2BT_SUB_PROC:
            if (out_events[i].action == C2BT_ACT_EXEC) {
                if (check_lolbas_comm(pl)) {
                    flags |= (C2BT_FLAG_LOLBAS | C2BT_FLAG_ANOMALY);
                } else {
                    flags |= C2BT_FLAG_VERDICT_OK;
                }
            }
            break;
            
        case C2BT_SUB_MCP:
            if (out_events[i].action == C2BT_ACT_TOOL_CALL) {
                /* Vérification des commandes destructrices, stagers réseau et altérations de doctrine */
                int has_net_dl = (strstr(pl, "curl") != NULL || strstr(pl, "wget") != NULL);
                int has_shell_pipe = (strstr(pl, "| sh") != NULL || strstr(pl, "| bash") != NULL ||
                                      strstr(pl, "|sh") != NULL || strstr(pl, "|bash") != NULL);
                if (strstr(pl, "rm -rf") != NULL ||
                    strstr(pl, "rm -r ") != NULL ||
                    strstr(pl, "chmod 777") != NULL ||
                    strstr(pl, "mkfs") != NULL ||
                    strstr(pl, "dd if=") != NULL ||
                    (has_net_dl && has_shell_pipe) ||
                    strstr(pl, ".claude") != NULL ||
                    strstr(pl, "CLAUDE.md") != NULL ||
                    strstr(pl, "AGENTS.md") != NULL ||
                    (flags & C2BT_FLAG_CRYPTO_PAYLOAD) != 0) {
                    flags |= (C2BT_FLAG_BLOCKED | C2BT_FLAG_ANOMALY);
                } else {
                    flags |= C2BT_FLAG_VERDICT_OK;
                }
            }
            break;
            
        case C2BT_SUB_ENTROPY:
            {
                uint32_t ent_q8 = (uint32_t)(out_events[i].src & 0xFFFFFFFFU);
                uint32_t pclass = (uint32_t)(out_events[i].src >> 32);
                if (pclass == C2BT_PAYLOAD_CLASS_CRYPTO_COMPRESSED || ent_q8 >= 1920 || (flags & C2BT_FLAG_CRYPTO_PAYLOAD)) {
                    flags |= (C2BT_FLAG_ANOMALY | C2BT_FLAG_BLOCKED | C2BT_FLAG_CRYPTO_PAYLOAD);
                } else if (pclass == C2BT_PAYLOAD_CLASS_BASE64 || (flags & C2BT_FLAG_BASE64_PAYLOAD)) {
                    flags |= (C2BT_FLAG_ANOMALY | C2BT_FLAG_BASE64_PAYLOAD);
                } else if (pclass == C2BT_PAYLOAD_CLASS_HEX || (flags & C2BT_FLAG_HEX_PAYLOAD)) {
                    flags |= (C2BT_FLAG_ANOMALY | C2BT_FLAG_HEX_PAYLOAD);
                } else {
                    flags |= C2BT_FLAG_VERDICT_OK;
                }
            }
            break;
            
        case C2BT_SUB_FS:
            /* Protection du harnais */
            if (strstr(pl, ".claude") != NULL ||
                strstr(pl, "CLAUDE.md") != NULL ||
                strstr(pl, "AGENTS.md") != NULL) {
                if (out_events[i].action == C2BT_ACT_WRITE || out_events[i].action == C2BT_ACT_OPEN) {
                    flags |= (C2BT_FLAG_BLOCKED | C2BT_FLAG_ANOMALY);
                }
            } else {
                flags |= C2BT_FLAG_VERDICT_OK;
            }
            break;
            
        default:
            flags |= C2BT_FLAG_VERDICT_OK;
            break;
        }
        
        out_events[i].flags = flags;
    }
    
    return count;
}
