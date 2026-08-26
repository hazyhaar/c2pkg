#include "c2q55.h"
#include "c2q55_crc_archtime.h"
#include <string.h>

#if defined(__AVX2__)
#include <immintrin.h>

uint32_t c2q_scan_block_avx2(const c2q_meta_block_t *block, uint64_t now) {
    const __m256i v_sign_flip = _mm256_set1_epi64x((int64_t)0x8000000000000000ULL);
    const __m256i v_now_biased = _mm256_xor_si256(_mm256_set1_epi64x((int64_t)now), v_sign_flip);
    const __m256i v_state_ready = _mm256_set1_epi32((int)C2Q_STATE_READY);

    /* 1. Chargement aligné 256-bit des 8 timestamps (2 x 4 x u64) */
    __m256i v_time0 = _mm256_loadu_si256((const __m256i *)&block->visible_at[0]);
    __m256i v_time1 = _mm256_loadu_si256((const __m256i *)&block->visible_at[4]);

    v_time0 = _mm256_xor_si256(v_time0, v_sign_flip);
    v_time1 = _mm256_xor_si256(v_time1, v_sign_flip);

    /* Comparaison : now_biased >= time_biased => !(time_biased > now_biased) */
    __m256i m_time0 = _mm256_xor_si256(_mm256_cmpgt_epi64(v_time0, v_now_biased), _mm256_set1_epi32(-1));
    __m256i m_time1 = _mm256_xor_si256(_mm256_cmpgt_epi64(v_time1, v_now_biased), _mm256_set1_epi32(-1));

    /* 2. Chargement des 8 états (1 x 8 x u32) et comparaison */
    __m256i v_state = _mm256_loadu_si256((const __m256i *)block->state);
    __m256i m_state = _mm256_cmpeq_epi32(v_state, v_state_ready);

    /* Dépaquetage exact 32-bit vers 64-bit via cvtepi32_epi64 */
    __m256i m_state0 = _mm256_cvtepi32_epi64(_mm256_castsi256_si128(m_state));
    __m256i m_state1 = _mm256_cvtepi32_epi64(_mm256_extracti128_si256(m_state, 1));

    /* Conjonction logique */
    __m256i res0 = _mm256_and_si256(m_time0, m_state0);
    __m256i res1 = _mm256_and_si256(m_time1, m_state1);

    uint32_t mask0 = (uint32_t)_mm256_movemask_pd((__m256d)res0);
    uint32_t mask1 = (uint32_t)_mm256_movemask_pd((__m256d)res1);

    return mask0 | (mask1 << 4);
}
#endif

uint32_t c2q_scan_block_scalar(const c2q_meta_block_t *block, uint64_t now) {
    uint32_t mask = 0;
    for (int i = 0; i < C2Q_BLOCK_SLOTS; i++) {
        if (block->visible_at[i] <= now && block->state[i] == C2Q_STATE_READY) {
            mask |= (1U << i);
        }
    }
    return mask;
}

uint32_t c2q_scan_block(const c2q_meta_block_t *block, uint64_t now) {
    return c2q_scan_block_scalar(block, now);
}

uint32_t c2q_trailing_zeros32(uint32_t mask) {
    if (mask == 0) {
        return 32;
    }
#if defined(__GNUC__) || defined(__clang__)
    return (uint32_t)__builtin_ctz(mask);
#else
    uint32_t n = 0;
    if ((mask & 0x0000FFFFU) == 0) { n += 16; mask >>= 16; }
    if ((mask & 0x000000FFU) == 0) { n += 8;  mask >>= 8;  }
    if ((mask & 0x0000000FU) == 0) { n += 4;  mask >>= 4;  }
    if ((mask & 0x00000003U) == 0) { n += 2;  mask >>= 2;  }
    if ((mask & 0x00000001U) == 0) { n += 1; }
    return n;
#endif
}

uint32_t c2q_count_ones32(uint32_t mask) {
#if defined(__GNUC__) || defined(__clang__)
    return (uint32_t)__builtin_popcount(mask);
#else
    mask = mask - ((mask >> 1) & 0x55555555U);
    mask = (mask & 0x33333333U) + ((mask >> 2) & 0x33333333U);
    return (((mask + (mask >> 4)) & 0x0F0F0F0FU) * 0x01010101U) >> 24;
#endif
}

uint64_t c2q_byteswap64(uint64_t val) {
#if defined(__GNUC__) || defined(__clang__)
    return (uint64_t)__builtin_bswap64(val);
#else
    val = ((val << 8) & 0xFF00FF00FF00FF00ULL) | ((val >> 8) & 0x00FF00FF00FF00FFULL);
    val = ((val << 16) & 0xFFFF0000FFFF0000ULL) | ((val >> 16) & 0x0000FFFF0000FFFFULL);
    return (val << 32) | (val >> 32);
#endif
}

uint64_t c2q_select64(uint64_t condition, uint64_t val_true, uint64_t val_false) {
    uint64_t mask = (uint64_t)(-(int64_t)(condition != 0));
    return (val_true & mask) | (val_false & ~mask);
}

void c2q_consume_mask(c2q_meta_block_t *block, uint32_t mask, uint32_t new_state) {
    mask &= 0xFFU;
    while (mask != 0) {
        uint32_t idx = c2q_trailing_zeros32(mask);
        block->state[idx] = new_state;
        mask &= (mask - 1);
    }
}

/* Calcul CRC32-C Castagnoli ARCHTIME Table 256 (1 Ko L1D) */
uint32_t c2q_crc32c_hw(uint32_t crc, const uint8_t *data, size_t len) {
    crc = ~crc;
    for (size_t i = 0; i < len; i++) {
        crc = (crc >> 8) ^ crc32c_table256[(crc ^ data[i]) & 0xFF];
    }
    return ~crc;
}
