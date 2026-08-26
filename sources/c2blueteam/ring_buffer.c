/*
 * ring_buffer.c — Implémentation des canaux annulaires SPSC atomiques sans CGO.
 */

#include "ring_buffer.h"
#include <string.h>

void c2bt_channel_init(probe_channel_t *ch) {
    if (!ch) return;
    __atomic_store_n(&ch->head, 0, __ATOMIC_RELAXED);
    __atomic_store_n(&ch->drops, 0, __ATOMIC_RELAXED);
    __atomic_store_n(&ch->tail, 0, __ATOMIC_RELAXED);
    memset(ch->slots, 0, sizeof(ch->slots));
}

int c2bt_channel_write(probe_channel_t *ch, const probe_event_t *ev) {
    if (!ch || !ev) return -1;
    
    uint64_t head = __atomic_load_n(&ch->head, __ATOMIC_RELAXED);
    uint64_t tail = __atomic_load_n(&ch->tail, __ATOMIC_ACQUIRE);
    
    if ((head - tail) >= C2BT_RING_CAP) {
        /* Anneau saturé : perte bornée (fail-safe non-bloquant) et télémétrie de drop atomique */
        __atomic_fetch_add(&ch->drops, 1, __ATOMIC_RELAXED);
        return -2;
    }
    
    uint64_t slot_idx = head & C2BT_RING_MASK;
    ch->slots[slot_idx] = *ev;
    
    /* Publication de l'événement avec barrière release */
    __atomic_store_n(&ch->head, head + 1, __ATOMIC_RELEASE);
    return 0;
}

int c2bt_channel_read(probe_channel_t *ch, probe_event_t *ev) {
    if (!ch || !ev) return -1;
    
    uint64_t tail = __atomic_load_n(&ch->tail, __ATOMIC_RELAXED);
    uint64_t head = __atomic_load_n(&ch->head, __ATOMIC_ACQUIRE);
    
    if (tail >= head) {
        /* Anneau vide */
        return 0;
    }
    
    uint64_t slot_idx = tail & C2BT_RING_MASK;
    *ev = ch->slots[slot_idx];
    
    /* Libération de la place avec barrière release */
    __atomic_store_n(&ch->tail, tail + 1, __ATOMIC_RELEASE);
    return 1;
}

int c2bt_channel_read_batch(probe_channel_t *ch, probe_event_t *out_events, int max_events) {
    if (!ch || !out_events || max_events <= 0) return 0;
    
    uint64_t tail = __atomic_load_n(&ch->tail, __ATOMIC_RELAXED);
    uint64_t head = __atomic_load_n(&ch->head, __ATOMIC_ACQUIRE);
    
    uint64_t avail = head - tail;
    if (avail == 0) return 0;
    
    int count = (int)avail;
    if (count > max_events) count = max_events;
    
    for (int i = 0; i < count; i++) {
        uint64_t slot_idx = (tail + (uint64_t)i) & C2BT_RING_MASK;
        out_events[i] = ch->slots[slot_idx];
    }
    
    __atomic_store_n(&ch->tail, tail + (uint64_t)count, __ATOMIC_RELEASE);
    return count;
}

uint64_t c2bt_channel_get_drops(const probe_channel_t *ch) {
    if (!ch) return 0;
    return __atomic_load_n(&ch->drops, __ATOMIC_RELAXED);
}
