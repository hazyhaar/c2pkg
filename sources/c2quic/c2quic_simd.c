#include "c2quic.h"

#if defined(__AVX2__)
#include <immintrin.h>
#include <string.h>

void c2quic_vint_decode_batch8_avx2(const uint8_t *buf, uint32_t len, c2quic_batch8_t *b) {
    int i;
    int all_wide = 1;
    uint64_t raw[8];
    if (buf == 0 || b == 0) {
        c2quic_vint_decode_batch8_scalar(buf, len, b);
        return;
    }
    for (i = 0; i < 8; i++) {
        if (b->offs[i] > len || len - b->offs[i] < 8u) {
            all_wide = 0;
            break;
        }
    }
    if (!all_wide) {
        c2quic_vint_decode_batch8_scalar(buf, len, b);
        return;
    }
    for (i = 0; i < 8; i++) {
        memcpy(&raw[i], buf + b->offs[i], 8);
    }
    {
        __m256i v0 = _mm256_loadu_si256((const __m256i *)&raw[0]);
        __m256i v1 = _mm256_loadu_si256((const __m256i *)&raw[4]);
        const __m256i shuf = _mm256_setr_epi8(
            7, 6, 5, 4, 3, 2, 1, 0,
            15, 14, 13, 12, 11, 10, 9, 8,
            7, 6, 5, 4, 3, 2, 1, 0,
            15, 14, 13, 12, 11, 10, 9, 8);
        v0 = _mm256_shuffle_epi8(v0, shuf);
        v1 = _mm256_shuffle_epi8(v1, shuf);
        _mm256_storeu_si256((__m256i *)&raw[0], v0);
        _mm256_storeu_si256((__m256i *)&raw[4], v1);
    }
    for (i = 0; i < 8; i++) {
        uint64_t w = raw[i];
        uint32_t prefix = (uint32_t)(w >> 62);
        uint32_t n = 1u << prefix;
        uint32_t shift = 64u - (n << 3);
        uint64_t mask;
        if (n == 8u) {
            mask = 0x3FFFFFFFFFFFFFFFull;
        } else {
            mask = (1ull << ((n << 3) - 2u)) - 1ull;
        }
        b->vals[i] = (w >> shift) & mask;
        b->nbytes[i] = n;
        b->status[i] = C2QUIC_OK;
    }
}
#endif

#if defined(__ARM_NEON) || defined(__ARM_NEON__)
#include <arm_neon.h>
#include <string.h>

void c2quic_vint_decode_batch8_neon(const uint8_t *buf, uint32_t len, c2quic_batch8_t *b) {
    int i;
    int all_wide = 1;
    uint64_t raw[8];
    if (buf == 0 || b == 0) {
        c2quic_vint_decode_batch8_scalar(buf, len, b);
        return;
    }
    for (i = 0; i < 8; i++) {
        if (b->offs[i] > len || len - b->offs[i] < 8u) {
            all_wide = 0;
            break;
        }
    }
    if (!all_wide) {
        c2quic_vint_decode_batch8_scalar(buf, len, b);
        return;
    }
    for (i = 0; i < 8; i++) {
        memcpy(&raw[i], buf + b->offs[i], 8);
        raw[i] = __builtin_bswap64(raw[i]);
    }
    for (i = 0; i < 8; i++) {
        uint64_t w = raw[i];
        uint32_t prefix = (uint32_t)(w >> 62);
        uint32_t n = 1u << prefix;
        uint32_t shift = 64u - (n << 3);
        uint64_t mask;
        if (n == 8u) {
            mask = 0x3FFFFFFFFFFFFFFFull;
        } else {
            mask = (1ull << ((n << 3) - 2u)) - 1ull;
        }
        b->vals[i] = (w >> shift) & mask;
        b->nbytes[i] = n;
        b->status[i] = C2QUIC_OK;
    }
}
#endif
