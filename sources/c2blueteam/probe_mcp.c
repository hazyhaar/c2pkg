/*
 * probe_mcp.c — Proxy Unix transparent et filtrage synchrone des outils JSON-RPC.
 */

#define _GNU_SOURCE
#include "c2blueteam.h"
#include "ring_buffer.h"
#include <sys/socket.h>
#include <sys/un.h>
#include <unistd.h>
#include <string.h>
#include <stdio.h>

int c2bt_probe_mcp_inspect(const char *json_buf, size_t len, char *out_tool, size_t out_max) {
    if (!json_buf || len == 0 || !out_tool || out_max == 0) return -1;
    
    out_tool[0] = '\0';
    
    const char *p = strstr(json_buf, "\"name\":");
    if (!p) {
        p = strstr(json_buf, "\"method\":");
    }
    if (!p) return 0;
    
    p = strchr(p, ':');
    if (!p) return 0;
    p++;
    while (*p == ' ' || *p == '\t' || *p == '\"') p++;
    
    size_t i = 0;
    while (*p && *p != '\"' && *p != ',' && *p != '}' && i + 1 < out_max) {
        out_tool[i++] = *p++;
    }
    out_tool[i] = '\0';
    return (int)i;
}

int c2bt_probe_mcp_extract_deep(const char *json_buf, size_t len, char *out_payload, size_t max_payload) {
    if (!json_buf || len == 0 || !out_payload || max_payload == 0) return -1;
    
    char tool_name[32] = {0};
    c2bt_probe_mcp_inspect(json_buf, len, tool_name, sizeof(tool_name));
    
    /* Recherche des arguments JSON */
    const char *args_ptr = strstr(json_buf, "\"arguments\":");
    if (!args_ptr) {
        args_ptr = strstr(json_buf, "\"params\":");
    }
    
    if (tool_name[0] != '\0' && args_ptr != NULL) {
        snprintf(out_payload, max_payload, "%s: %.*s", tool_name, (int)(max_payload - strlen(tool_name) - 3), args_ptr);
    } else if (tool_name[0] != '\0') {
        snprintf(out_payload, max_payload, "%.*s", (int)(max_payload - 1), json_buf);
    } else {
        snprintf(out_payload, max_payload, "%.*s", (int)(max_payload - 1), json_buf);
    }
    
    out_payload[max_payload - 1] = '\0';
    return 0;
}

/*
 * Extraction de l'identifiant JSON-RPC (id) sans allocation dynamique.
 */
static int c2bt_probe_mcp_extract_id(const char *json_in, size_t in_len, char *out_id, size_t max_id) {
    if (!out_id || max_id == 0) return -1;
    out_id[0] = '\0';
    if (!json_in || in_len == 0) {
        snprintf(out_id, max_id, "null");
        return 0;
    }
    
    const char *p = json_in;
    const char *end = json_in + in_len;
    const char *id_pos = NULL;
    
    while (p + 4 <= end) {
        if (p[0] == '"' && p[1] == 'i' && p[2] == 'd' && p[3] == '"') {
            if (p == json_in || *(p - 1) == '{' || *(p - 1) == ',' || *(p - 1) == ' ' || *(p - 1) == '\t' || *(p - 1) == '\n' || *(p - 1) == '\r') {
                const char *colon = p + 4;
                while (colon < end && (*colon == ' ' || *colon == '\t' || *colon == '\r' || *colon == '\n')) {
                    colon++;
                }
                if (colon < end && *colon == ':') {
                    id_pos = colon + 1;
                    break;
                }
            }
        }
        p++;
    }
    
    if (!id_pos) {
        snprintf(out_id, max_id, "null");
        return 0;
    }
    
    while (id_pos < end && (*id_pos == ' ' || *id_pos == '\t' || *id_pos == '\r' || *id_pos == '\n')) {
        id_pos++;
    }
    if (id_pos >= end) {
        snprintf(out_id, max_id, "null");
        return 0;
    }
    
    if (*id_pos == '"') {
        size_t idx = 0;
        out_id[idx++] = *id_pos++;
        while (id_pos < end && idx + 1 < max_id) {
            char c = *id_pos++;
            out_id[idx++] = c;
            if (c == '"' && *(id_pos - 2) != '\\') {
                break;
            }
        }
        out_id[idx] = '\0';
        return (int)idx;
    } else {
        size_t idx = 0;
        while (id_pos < end && idx + 1 < max_id) {
            char c = *id_pos;
            if (c == ',' || c == '}' || c == ']' || c == ' ' || c == '\t' || c == '\r' || c == '\n') {
                break;
            }
            out_id[idx++] = c;
            id_pos++;
        }
        out_id[idx] = '\0';
        if (idx == 0) {
            snprintf(out_id, max_id, "null");
        }
        return (int)idx;
    }
}

int c2bt_probe_mcp_filter_call(const char *json_in, size_t in_len, char *json_out, size_t max_out, uint32_t flags, const c2bt_config_t *cfg) {
    if (!json_out || max_out == 0) {
        return -1;
    }
    
    if (!json_in || in_len == 0) {
        if (cfg && cfg->fail_safe_policy == C2BT_POLICY_FAIL_CLOSED) {
            int n = snprintf(json_out, max_out,
                "{\"jsonrpc\":\"2.0\",\"id\":null,\"error\":{\"code\":-32003,\"message\":\"Blocked by c2blueteam security doctrine\"}}");
            return (n > 0 && (size_t)n < max_out) ? 1 : -1;
        }
        json_out[0] = '\0';
        return -1;
    }
    
    int should_block = 0;
    if (cfg && (cfg->enforce_mode == C2BT_MODE_ACTIVE) && (flags & C2BT_FLAG_BLOCKED)) {
        should_block = 1;
    }
    
    if (should_block) {
        char id_val[64] = {0};
        c2bt_probe_mcp_extract_id(json_in, in_len, id_val, sizeof(id_val));
        
        int written = snprintf(json_out, max_out,
            "{\"jsonrpc\":\"2.0\",\"id\":%s,\"error\":{\"code\":-32003,\"message\":\"Blocked by c2blueteam security doctrine\"}}",
            id_val);
            
        if (written < 0 || (size_t)written >= max_out) {
            return -1;
        }
        return 1; /* Veto actif : réponse de rejet JSON-RPC 2.0 générée */
    }
    
    /* Observation seule / flux autorisé : transit transparent */
    if (max_out <= in_len) {
        if (cfg && cfg->fail_safe_policy == C2BT_POLICY_FAIL_CLOSED) {
            snprintf(json_out, max_out,
                "{\"jsonrpc\":\"2.0\",\"id\":null,\"error\":{\"code\":-32003,\"message\":\"Blocked by c2blueteam security doctrine\"}}");
            return 1;
        }
        return -1;
    }
    
    memcpy(json_out, json_in, in_len);
    json_out[in_len] = '\0';
    return 0; /* Flux autorisé / Pass-through */
}

int c2bt_probe_mcp_filter_request(int client_sock, probe_channel_t *ch) {
    if (client_sock < 0 || !ch) return -1;
    
    char peek_buf[2048];
    ssize_t peeked = recv(client_sock, peek_buf, sizeof(peek_buf) - 1, MSG_PEEK | MSG_DONTWAIT);
    if (peeked <= 0) return 0;
    peek_buf[peeked] = '\0';
    
    struct ucred cred;
    socklen_t cred_len = sizeof(cred);
    pid_t sender_pid = 0;
    if (getsockopt(client_sock, SOL_SOCKET, SO_PEERCRED, &cred, &cred_len) == 0) {
        sender_pid = cred.pid;
    }
    
    probe_event_t ev = {0};
    ev.ts_ns = c2bt_time_ns_raw();
    ev.pid = (uint32_t)sender_pid;
    ev.subsystem = C2BT_SUB_MCP;
    ev.action = C2BT_ACT_TOOL_CALL;
    
    /* Inspection profonde : préservation de l'outil et des arguments JSON */
    c2bt_probe_mcp_extract_deep(peek_buf, (size_t)peeked, (char *)ev.payload, sizeof(ev.payload));
    
    /* Profilage de charge utile sur le tampon brut de 2048 octets (Proposition 4 de l'audit 08) */
    c2bt_entropy_profile_t prof;
    c2bt_profile_payload((const uint8_t *)peek_buf, (size_t)peeked, &prof);
    ev.src = (uint64_t)prof.entropy_q8 | ((uint64_t)prof.payload_class << 32);

    if (prof.payload_class == C2BT_PAYLOAD_CLASS_CRYPTO_COMPRESSED) {
        ev.flags |= (C2BT_FLAG_CRYPTO_PAYLOAD | C2BT_FLAG_ANOMALY);
    } else if (prof.payload_class == C2BT_PAYLOAD_CLASS_BASE64) {
        ev.flags |= C2BT_FLAG_BASE64_PAYLOAD;
    } else if (prof.payload_class == C2BT_PAYLOAD_CLASS_HEX) {
        ev.flags |= C2BT_FLAG_HEX_PAYLOAD;
    }
    
    c2bt_channel_write(ch, &ev);
    return 1;
}
