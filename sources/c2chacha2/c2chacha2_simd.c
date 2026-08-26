#include <immintrin.h>
#include <string.h>
#include "c2chacha2.h"

/* Masques pshufb 256 bits : masque 16 octets de c2chacha1 dupliqué sur les
 * deux moitiés 128 bits (shuffle_epi8 AVX2 est intra-voie). */
static const uint8_t c2chacha2_rot16_mask[32] = {
    2, 3, 0, 1, 6, 7, 4, 5, 10, 11, 8, 9, 14, 15, 12, 13,
    2, 3, 0, 1, 6, 7, 4, 5, 10, 11, 8, 9, 14, 15, 12, 13};
static const uint8_t c2chacha2_rot8_mask[32] = {
    3, 0, 1, 2, 7, 4, 5, 6, 11, 8, 9, 10, 15, 12, 13, 14,
    3, 0, 1, 2, 7, 4, 5, 6, 11, 8, 9, 10, 15, 12, 13, 14};

/* Constantes "expand 32-byte k" dupliquées : une ligne sigma par moitié. */
static const uint8_t c2chacha2_sigma2[32] = {
    0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x20, 0x33,
    0x32, 0x2d, 0x62, 0x79, 0x74, 0x65, 0x20, 0x6b,
    0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x20, 0x33,
    0x32, 0x2d, 0x62, 0x79, 0x74, 0x65, 0x20, 0x6b};

size_t c2chacha2_xor_blocks(uint8_t *out, const uint8_t *in, size_t n,
                            const uint8_t key[32], const uint8_t nonce[12],
                            uint32_t counter) {
    if (n > 128) {
        n = 128;
    }
    if (n == 0) {
        return 0;
    }

    __m256i rot16 = _mm256_loadu_si256((const __m256i *)c2chacha2_rot16_mask);
    __m256i rot8 = _mm256_loadu_si256((const __m256i *)c2chacha2_rot8_mask);

    uint8_t krow1[32], krow2[32];
    memcpy(krow1, key, 16);
    memcpy(krow1 + 16, key, 16);
    memcpy(krow2, key + 16, 16);
    memcpy(krow2 + 16, key + 16, 16);

    uint32_t n0 = (uint32_t)nonce[0] | ((uint32_t)nonce[1] << 8) |
                  ((uint32_t)nonce[2] << 16) | ((uint32_t)nonce[3] << 24);
    uint32_t n1 = (uint32_t)nonce[4] | ((uint32_t)nonce[5] << 8) |
                  ((uint32_t)nonce[6] << 16) | ((uint32_t)nonce[7] << 24);
    uint32_t n2 = (uint32_t)nonce[8] | ((uint32_t)nonce[9] << 8) |
                  ((uint32_t)nonce[10] << 16) | ((uint32_t)nonce[11] << 24);

    /* Levier c_source : chaque vecteur est chargé depuis la mémoire, aucune
     * copie registre à registre (MOVUPS SSE au milieu de VEX). Compteur :
     * _mm256_set_epi32 avec +1 dans la moitié haute (bloc counter+1). */
    __m256i a0 = _mm256_loadu_si256((const __m256i *)c2chacha2_sigma2);
    __m256i a1 = _mm256_loadu_si256((const __m256i *)krow1);
    __m256i a2 = _mm256_loadu_si256((const __m256i *)krow2);
    __m256i a3 = _mm256_set_epi32(n2, n1, n0, counter + 1u, n2, n1, n0, counter);

    uint8_t row3[32];
    _mm256_storeu_si256((__m256i *)row3, a3);

    for (int i = 0; i < 10; i++) {
        a0 = _mm256_add_epi32(a0, a1);
        a3 = _mm256_xor_si256(a3, a0);
        a3 = _mm256_shuffle_epi8(a3, rot16);
        a2 = _mm256_add_epi32(a2, a3);
        a1 = _mm256_xor_si256(a1, a2);
        a1 = _mm256_or_si256(_mm256_slli_epi32(a1, 12), _mm256_srli_epi32(a1, 20));
        a0 = _mm256_add_epi32(a0, a1);
        a3 = _mm256_xor_si256(a3, a0);
        a3 = _mm256_shuffle_epi8(a3, rot8);
        a2 = _mm256_add_epi32(a2, a3);
        a1 = _mm256_xor_si256(a1, a2);
        a1 = _mm256_or_si256(_mm256_slli_epi32(a1, 7), _mm256_srli_epi32(a1, 25));

        a1 = _mm256_shuffle_epi32(a1, 0x39);
        a2 = _mm256_shuffle_epi32(a2, 0x4e);
        a3 = _mm256_shuffle_epi32(a3, 0x93);

        a0 = _mm256_add_epi32(a0, a1);
        a3 = _mm256_xor_si256(a3, a0);
        a3 = _mm256_shuffle_epi8(a3, rot16);
        a2 = _mm256_add_epi32(a2, a3);
        a1 = _mm256_xor_si256(a1, a2);
        a1 = _mm256_or_si256(_mm256_slli_epi32(a1, 12), _mm256_srli_epi32(a1, 20));
        a0 = _mm256_add_epi32(a0, a1);
        a3 = _mm256_xor_si256(a3, a0);
        a3 = _mm256_shuffle_epi8(a3, rot8);
        a2 = _mm256_add_epi32(a2, a3);
        a1 = _mm256_xor_si256(a1, a2);
        a1 = _mm256_or_si256(_mm256_slli_epi32(a1, 7), _mm256_srli_epi32(a1, 25));

        a1 = _mm256_shuffle_epi32(a1, 0x93);
        a2 = _mm256_shuffle_epi32(a2, 0x4e);
        a3 = _mm256_shuffle_epi32(a3, 0x39);
    }

    a0 = _mm256_add_epi32(a0, _mm256_loadu_si256((const __m256i *)c2chacha2_sigma2));
    a1 = _mm256_add_epi32(a1, _mm256_loadu_si256((const __m256i *)krow1));
    a2 = _mm256_add_epi32(a2, _mm256_loadu_si256((const __m256i *)krow2));
    a3 = _mm256_add_epi32(a3, _mm256_loadu_si256((const __m256i *)row3));

    __m256i blo0 = _mm256_permute2x128_si256(a0, a1, 0x20);
    __m256i bhi0 = _mm256_permute2x128_si256(a2, a3, 0x20);
    __m256i blo1 = _mm256_permute2x128_si256(a0, a1, 0x31);
    __m256i bhi1 = _mm256_permute2x128_si256(a2, a3, 0x31);

    size_t nfull = n / 64;
    size_t rem = n % 64;

    if (nfull >= 1) {
        _mm256_storeu_si256((__m256i *)(out + 0), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 0)), blo0));
        _mm256_storeu_si256((__m256i *)(out + 32), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 32)), bhi0));
    }
    if (nfull >= 2) {
        _mm256_storeu_si256((__m256i *)(out + 64), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 64)), blo1));
        _mm256_storeu_si256((__m256i *)(out + 96), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 96)), bhi1));
    }

    if (rem != 0) {
        uint8_t ks[64];
        if (nfull == 0) {
            _mm256_storeu_si256((__m256i *)(ks + 0), blo0);
            _mm256_storeu_si256((__m256i *)(ks + 32), bhi0);
        } else {
            _mm256_storeu_si256((__m256i *)(ks + 0), blo1);
            _mm256_storeu_si256((__m256i *)(ks + 32), bhi1);
        }
        size_t base = nfull * 64;
        for (size_t k = 0; k < rem; k++) {
            out[base + k] = in[base + k] ^ ks[k];
        }
    }
    return n;
}
