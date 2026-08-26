#include <stdio.h>
#include <stdint.h>
#include <string.h>
#include "c2quic.h"

#if defined(__AVX2__)
void c2quic_vint_decode_batch8_avx2(const uint8_t *buf, uint32_t len, c2quic_batch8_t *b);
#endif

static uint64_t rng_state = 0x853c49e6748fea9bULL;

static uint64_t next_u64(void) {
    rng_state ^= rng_state >> 12;
    rng_state ^= rng_state << 25;
    rng_state ^= rng_state >> 27;
    return rng_state * 0x2545F4914F6CDD1DULL;
}

static int fail(const char *msg) {
    printf("[FAIL] %s\n", msg);
    return 1;
}

static uint64_t fold_mix(uint64_t acc, uint64_t x) {
    acc ^= x + 0x9E3779B97F4A7C15ULL + (acc << 6) + (acc >> 2);
    return acc;
}

int main(void) {
    c2quic_vint_t r;
    uint8_t enc[8];
    uint8_t pkt[64];
    c2quic_hp_io_t io;
    c2quic_frame_slab_t slab;
    uint32_t st;
    uint64_t fold_vint = 0;
    uint64_t fold_hp = 0;
    uint64_t fold_fr = 0;
    int i;

    printf("=== C2QUIC C ORACLE HARNESS (gcc -O2) ===\n");

    enc[0] = 0x25;
    st = c2quic_vint_decode(enc, 1, 0, &r);
    if (st != C2QUIC_OK || r.value != 37ull || r.nbytes != 1) {
        return fail("KAT vint 37");
    }
    printf("KAT vint 37: PASS\n");

    enc[0] = 0x7b;
    enc[1] = 0xbd;
    st = c2quic_vint_decode(enc, 2, 0, &r);
    if (st != C2QUIC_OK || r.value != 15293ull || r.nbytes != 2) {
        return fail("KAT vint 15293");
    }
    printf("KAT vint 15293: PASS\n");

    enc[0] = 0x9d;
    enc[1] = 0x7f;
    enc[2] = 0x3e;
    enc[3] = 0x7d;
    st = c2quic_vint_decode(enc, 4, 0, &r);
    if (st != C2QUIC_OK || r.value != 494878333ull || r.nbytes != 4) {
        return fail("KAT vint 494878333");
    }
    printf("KAT vint 494878333: PASS\n");

    enc[0] = 0xc2;
    enc[1] = 0x19;
    enc[2] = 0x7c;
    enc[3] = 0x5e;
    enc[4] = 0xff;
    enc[5] = 0x14;
    enc[6] = 0xe8;
    enc[7] = 0x8c;
    st = c2quic_vint_decode(enc, 8, 0, &r);
    if (st != C2QUIC_OK || r.value != 151288809941952652ull || r.nbytes != 8) {
        return fail("KAT vint 151288809941952652");
    }
    printf("KAT vint 151288809941952652: PASS\n");

    st = c2quic_vint_encode(37ull, enc, 8, 0, &r);
    if (st != C2QUIC_OK || enc[0] != 0x25 || r.nbytes != 1) {
        return fail("KAT encode 37");
    }
    st = c2quic_vint_encode(15293ull, enc, 8, 0, &r);
    if (st != C2QUIC_OK || enc[0] != 0x7b || enc[1] != 0xbd || r.nbytes != 2) {
        return fail("KAT encode 15293");
    }

    for (i = 0; i < 1000000; i++) {
        uint64_t v = next_u64() >> 2;
        uint8_t buf[64];
        c2quic_batch8_t b;
        int k;
        memset(buf, 0, sizeof(buf));
        memset(&b, 0, sizeof(b));
        st = c2quic_vint_encode(v, buf, 64, 0, &r);
        if (st != C2QUIC_OK) {
            return fail("random encode");
        }
        st = c2quic_vint_decode(buf, 64, 0, &r);
        if (st != C2QUIC_OK || r.value != v) {
            return fail("random decode mismatch");
        }
        fold_vint = fold_mix(fold_vint, r.value);
        fold_vint = fold_mix(fold_vint, r.nbytes);
        for (k = 0; k < 8; k++) {
            uint64_t vk = next_u64() >> 2;
            b.offs[k] = (uint32_t)(k * 8);
            c2quic_vint_encode(vk, buf, 64, b.offs[k], &r);
        }
        c2quic_vint_decode_batch8_scalar(buf, 64, &b);
#if defined(__AVX2__)
        {
            c2quic_batch8_t av;
            av = b;
            c2quic_vint_decode_batch8_avx2(buf, 64, &av);
            for (k = 0; k < 8; k++) {
                if (av.vals[k] != b.vals[k] || av.nbytes[k] != b.nbytes[k] || av.status[k] != b.status[k]) {
                    return fail("avx2 vs scalar batch");
                }
            }
        }
#endif
        for (k = 0; k < 8; k++) {
            fold_vint = fold_mix(fold_vint, b.vals[k]);
            fold_vint = fold_mix(fold_vint, b.nbytes[k]);
            fold_vint = fold_mix(fold_vint, b.status[k]);
        }
    }
    printf("FOLD vint 0x%016llX\n", (unsigned long long)fold_vint);

    rng_state = 0x853c49e6748fea9bULL;
    for (i = 0; i < 1000000; i++) {
        uint32_t pn_len = (uint32_t)((next_u64() & 3ull) + 1ull);
        uint32_t pn_off = 1;
        uint64_t pn;
        uint64_t pn_mask;
        int k;
        memset(pkt, 0, sizeof(pkt));
        memset(&io, 0, sizeof(io));
        if ((next_u64() & 1ull) != 0) {
            pkt[0] = (uint8_t)(0xC0u | (uint8_t)(pn_len - 1u));
        } else {
            pkt[0] = (uint8_t)(0x40u | (uint8_t)(pn_len - 1u));
        }
        pn = next_u64() & 0xFFFFFFFFull;
        for (k = 0; k < (int)pn_len; k++) {
            pkt[pn_off + (uint32_t)k] = (uint8_t)(pn >> (8 * (pn_len - 1u - (uint32_t)k)));
        }
        for (k = 0; k < 5; k++) {
            io.mask[k] = (uint8_t)next_u64();
        }
        memset(pkt + 5, 0xA5, 32);
        st = c2quic_hp_sample(pkt, 64, pn_off, &io);
        if (st != C2QUIC_OK) {
            return fail("hp sample");
        }
        st = c2quic_hp_apply(pkt, 64, pn_off, pn_len, &io);
        if (st != C2QUIC_OK) {
            return fail("hp apply");
        }
        st = c2quic_hp_remove(pkt, 64, pn_off, &io);
        if (st != C2QUIC_OK || io.pn_len != pn_len) {
            return fail("hp remove");
        }
        pn_mask = (1ull << (8u * pn_len)) - 1ull;
        st = c2quic_hp_read_pn(pkt, 64, pn_off, pn_len, &io);
        if (st != C2QUIC_OK || io.pn != (pn & pn_mask)) {
            return fail("hp pn restore");
        }
        fold_hp = fold_mix(fold_hp, pkt[0]);
        fold_hp = fold_mix(fold_hp, io.pn);
        fold_hp = fold_mix(fold_hp, io.sample[0]);
        fold_hp = fold_mix(fold_hp, io.sample[15]);
    }
    printf("FOLD hp 0x%016llX\n", (unsigned long long)fold_hp);

    {
        uint8_t body[128];
        uint32_t o = 0;
        memset(body, 0, sizeof(body));
        memset(&slab, 0, sizeof(slab));
        body[o++] = 0x00;
        body[o++] = 0x01;
        body[o++] = 0x0F;
        c2quic_vint_encode(7ull, body, 128, o, &r);
        o += r.nbytes;
        c2quic_vint_encode(100ull, body, 128, o, &r);
        o += r.nbytes;
        c2quic_vint_encode(4ull, body, 128, o, &r);
        o += r.nbytes;
        body[o++] = 'A';
        body[o++] = 'B';
        body[o++] = 'C';
        body[o++] = 'D';
        body[o++] = 0x04;
        c2quic_vint_encode(7ull, body, 128, o, &r);
        o += r.nbytes;
        c2quic_vint_encode(0x42ull, body, 128, o, &r);
        o += r.nbytes;
        c2quic_vint_encode(4ull, body, 128, o, &r);
        o += r.nbytes;
        body[o++] = 0x02;
        c2quic_vint_encode(20ull, body, 128, o, &r);
        o += r.nbytes;
        c2quic_vint_encode(1ull, body, 128, o, &r);
        o += r.nbytes;
        c2quic_vint_encode(0ull, body, 128, o, &r);
        o += r.nbytes;
        c2quic_vint_encode(5ull, body, 128, o, &r);
        o += r.nbytes;
        body[o++] = 0x31;
        c2quic_vint_encode(3ull, body, 128, o, &r);
        o += r.nbytes;
        body[o++] = 'x';
        body[o++] = 'y';
        body[o++] = 'z';
        st = c2quic_unpack_frames(body, o, &slab);
        if (st != C2QUIC_OK || slab.n != 5) {
            printf("frames st=%u n=%u\n", st, slab.n);
            return fail("unpack kat count");
        }
        if (slab.frames[0].typ != 0x01 || slab.frames[1].typ != 0x0F || slab.frames[1].id != 7 || slab.frames[1].offset != 100 || slab.frames[1].payload_len != 4) {
            return fail("unpack stream");
        }
        if (slab.frames[2].typ != 0x04 || slab.frames[2].extra != 0x42 || slab.frames[2].length != 4) {
            return fail("unpack reset");
        }
        if (slab.frames[3].typ != 0x02 || slab.frames[3].id != 20 || slab.frames[3].extra != 5) {
            return fail("unpack ack");
        }
        if (slab.frames[4].typ != 0x31 || slab.frames[4].payload_len != 3) {
            return fail("unpack datagram");
        }
        fold_fr = fold_mix(fold_fr, slab.n);
        fold_fr = fold_mix(fold_fr, slab.frames[1].id);
        fold_fr = fold_mix(fold_fr, slab.frames[4].payload_len);
        printf("KAT frames: PASS\n");
    }

    rng_state = 0x123456789abcdef0ULL;
    for (i = 0; i < 1000000; i++) {
        uint8_t body[256];
        uint32_t o = 0;
        uint64_t sid = next_u64() & 0xFFFFull;
        uint64_t plen = (next_u64() & 15ull) + 1ull;
        uint32_t k;
        memset(body, 0, sizeof(body));
        memset(&slab, 0, sizeof(slab));
        body[o++] = 0x0A;
        c2quic_vint_encode(sid, body, 256, o, &r);
        o += r.nbytes;
        c2quic_vint_encode(plen, body, 256, o, &r);
        o += r.nbytes;
        for (k = 0; k < (uint32_t)plen; k++) {
            body[o++] = (uint8_t)next_u64();
        }
        st = c2quic_unpack_frames(body, o, &slab);
        if (st != C2QUIC_OK || slab.n != 1 || slab.frames[0].id != sid || slab.frames[0].payload_len != (uint32_t)plen) {
            return fail("random stream unpack");
        }
        fold_fr = fold_mix(fold_fr, slab.frames[0].id);
        fold_fr = fold_mix(fold_fr, slab.frames[0].payload_len);
        fold_fr = fold_mix(fold_fr, slab.frames[0].payload_off);
    }
    printf("FOLD frames 0x%016llX\n", (unsigned long long)fold_fr);
    printf("=== ALL C ORACLE TESTS PASSED ===\n");
    return 0;
}
