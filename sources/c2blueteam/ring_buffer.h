/*
 * ring_buffer.h — Canaux annulaires SPSC sans verrou et zéro contention.
 */

#ifndef C2BT_RING_BUFFER_H
#define C2BT_RING_BUFFER_H

#include "c2blueteam.h"

#ifdef __cplusplus
extern "C" {
#endif

void c2bt_channel_init(probe_channel_t *ch);
int  c2bt_channel_write(probe_channel_t *ch, const probe_event_t *ev);
int  c2bt_channel_read(probe_channel_t *ch, probe_event_t *ev);
int      c2bt_channel_read_batch(probe_channel_t *ch, probe_event_t *out_events, int max_events);
uint64_t c2bt_channel_get_drops(const probe_channel_t *ch);

#ifdef __cplusplus
}
#endif

#endif /* C2BT_RING_BUFFER_H */
