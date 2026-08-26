#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include "c2quic.h"

static int fail(const char *msg) {
    printf("[FAIL] %s\n", msg);
    return 1;
}

int main(void) {
    c2quic_vint_t r;
    c2quic_hp_io_t io;
    c2quic_frame_slab_t slab;
    uint8_t *tiny;
    uint8_t pkt[4];
    uint8_t first;
    uint32_t st;
    uint8_t ping2[2];

    tiny = (uint8_t *)malloc(8);
    if (tiny == 0) {
        return fail("malloc");
    }
    memset(tiny, 0, 8);

    st = c2quic_vint_encode(37ull, tiny, 8, 100u, &r);
    if (st != C2QUIC_TRUNC) {
        return fail("encode off>len");
    }

    memset(&io, 0, sizeof(io));
    io.mask[0] = 0x1F;
    pkt[0] = 0x40;
    pkt[1] = 0x01;
    pkt[2] = 0x02;
    pkt[3] = 0x03;
    first = pkt[0];
    st = c2quic_hp_remove(pkt, 3, 99u, &io);
    if (st != C2QUIC_TRUNC) {
        return fail("hp_remove far pn_off");
    }
    if (pkt[0] != first) {
        return fail("hp_remove mutated on TRUNC");
    }

    st = c2quic_hp_read_pn(tiny, 4, 1000u, 4u, &io);
    if (st != C2QUIC_TRUNC) {
        return fail("hp_read_pn far off");
    }

    ping2[0] = 0x40;
    ping2[1] = 0x01;
    memset(&slab, 0, sizeof(slab));
    st = c2quic_unpack_frames(ping2, 2, &slab);
    if (st != C2QUIC_OK || slab.n != 1 || slab.frames[0].typ != 0x01) {
        printf("st=%u n=%u typ=%u\n", st, slab.n, slab.frames[0].typ);
        return fail("PING vint 0x4001");
    }

    free(tiny);
    printf("=== C2QUIC BOUNDS HARNESS PASSED ===\n");
    return 0;
}
