/*
 * c2blueteam.c — Orchestrateur principal et API publique libc2blueteam.
 */

#define _GNU_SOURCE
#include "c2blueteam.h"
#include "ring_buffer.h"
#include "correlator.h"
#include <time.h>
#include <string.h>
#include <unistd.h>

uint64_t c2bt_time_ns_raw(void) {
    struct timespec ts;
    clock_gettime(CLOCK_MONOTONIC_RAW, &ts);
    return (uint64_t)ts.tv_sec * 1000000000ULL + (uint64_t)ts.tv_nsec;
}

int c2bt_init_inplace(c2bt_ctx_t *ctx, const c2bt_config_t *cfg) {
    if (!ctx) return -1;
    
    memset(ctx, 0, sizeof(*ctx));
    if (cfg) {
        ctx->config = *cfg;
    }
    
    c2bt_channel_init(&ctx->chan_proc);
    c2bt_channel_init(&ctx->chan_file);
    c2bt_channel_init(&ctx->chan_net);
    c2bt_channel_init(&ctx->chan_mcp);
    c2bt_tracker_init(&ctx->tracker_table);
    
    ctx->sealed_mem_fd = -1;
    ctx->fanotify_fd = -1;
    ctx->mcp_proxy_fd = -1;
    ctx->running = 0;
    
    return 0;
}

int c2bt_start(c2bt_ctx_t *ctx) {
    if (!ctx) return -1;
    ctx->running = 1;
    return 0;
}

int c2bt_stop(c2bt_ctx_t *ctx) {
    if (!ctx) return -1;
    ctx->running = 0;
    
    if (ctx->sealed_mem_fd >= 0) {
        close(ctx->sealed_mem_fd);
        ctx->sealed_mem_fd = -1;
    }
    if (ctx->fanotify_fd >= 0) {
        close(ctx->fanotify_fd);
        ctx->fanotify_fd = -1;
    }
    if (ctx->mcp_proxy_fd >= 0) {
        close(ctx->mcp_proxy_fd);
        ctx->mcp_proxy_fd = -1;
    }
    return 0;
}

int c2bt_poll_batch(c2bt_ctx_t *ctx, probe_event_t *out_batch, int max_events) {
    if (!ctx || !out_batch || max_events <= 0 || !ctx->running) return 0;

    int total_collected = 0;
    int has_events = 1;

    /* Relève équitable (Fair-Share Round-Robin) pour éliminer toute famine (Vecteur B) */
    while (total_collected < max_events && has_events) {
        has_events = 0;
        probe_event_t ev;

        if (total_collected < max_events && c2bt_channel_read(&ctx->chan_proc, &ev) == 1) {
            out_batch[total_collected++] = ev;
            has_events = 1;
        }
        if (total_collected < max_events && c2bt_channel_read(&ctx->chan_file, &ev) == 1) {
            out_batch[total_collected++] = ev;
            has_events = 1;
        }
        if (total_collected < max_events && c2bt_channel_read(&ctx->chan_net, &ev) == 1) {
            out_batch[total_collected++] = ev;
            has_events = 1;
        }
        if (total_collected < max_events && c2bt_channel_read(&ctx->chan_mcp, &ev) == 1) {
            out_batch[total_collected++] = ev;
            has_events = 1;
        }
    }

    if (total_collected == 0) return 0;

    /* Évaluation globale de tous les événements collectés (élimine la troncature à 8) */
    c2bt_eval_rules_batch(out_batch, out_batch, total_collected);

    /* Corrélation temporelle multi-flux en fenêtre glissante 1s */
    for (int i = 0; i < total_collected; i++) {
        uint32_t flags_res = out_batch[i].flags;
        probe_event_t ev_item = out_batch[i];
        c2bt_correlate_event(&ctx->tracker_table, &ev_item,
                             &flags_res,
                             C2BT_CORRELATION_DEFAULT_WINDOW_NS);
        out_batch[i].flags = flags_res;
    }

    return total_collected;
}
