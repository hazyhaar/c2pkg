#include <immintrin.h>
#include <string.h>
#include "c2chacha4.h"

/* Masques pshufb 256 bits : masque 16 octets de c2chacha1 dupliqué sur les
 * deux moitiés 128 bits (shuffle_epi8 AVX2 est intra-voie). */
static const uint8_t c2chacha4_rot16_mask[32] = {
    2, 3, 0, 1, 6, 7, 4, 5, 10, 11, 8, 9, 14, 15, 12, 13,
    2, 3, 0, 1, 6, 7, 4, 5, 10, 11, 8, 9, 14, 15, 12, 13};
static const uint8_t c2chacha4_rot8_mask[32] = {
    3, 0, 1, 2, 7, 4, 5, 6, 11, 8, 9, 10, 15, 12, 13, 14,
    3, 0, 1, 2, 7, 4, 5, 6, 11, 8, 9, 10, 15, 12, 13, 14};

/* Constantes "expand 32-byte k" dupliquées : une ligne sigma par moitié. */
static const uint8_t c2chacha4_sigma2[32] = {
    0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x20, 0x33,
    0x32, 0x2d, 0x62, 0x79, 0x74, 0x65, 0x20, 0x6b,
    0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x20, 0x33,
    0x32, 0x2d, 0x62, 0x79, 0x74, 0x65, 0x20, 0x6b};

/* Quatre blocs à `counter`…`counter+3`, au plus 256 octets. */
static size_t c2chacha4_xor256(uint8_t *out, const uint8_t *in, size_t n,
                               const uint8_t key[32], const uint8_t nonce[12],
                               uint32_t counter) {
    if (n > 256) {
        n = 256;
    }
    if (n == 0) {
        return 0;
    }

    __m256i rot16 = _mm256_loadu_si256((const __m256i *)c2chacha4_rot16_mask);
    __m256i rot8 = _mm256_loadu_si256((const __m256i *)c2chacha4_rot8_mask);

    uint8_t krow1[32], krow2[32], row3a[32], row3b[32];
    memcpy(krow1, key, 16);
    memcpy(krow1 + 16, key, 16);
    memcpy(krow2, key + 16, 16);
    memcpy(krow2 + 16, key + 16, 16);
    row3a[0] = (uint8_t)counter;
    row3a[1] = (uint8_t)(counter >> 8);
    row3a[2] = (uint8_t)(counter >> 16);
    row3a[3] = (uint8_t)(counter >> 24);
    memcpy(row3a + 4, nonce, 12);
    row3a[16] = (uint8_t)(counter + 1u);
    row3a[17] = (uint8_t)((counter + 1u) >> 8);
    row3a[18] = (uint8_t)((counter + 1u) >> 16);
    row3a[19] = (uint8_t)((counter + 1u) >> 24);
    memcpy(row3a + 20, nonce, 12);
    row3b[0] = (uint8_t)(counter + 2u);
    row3b[1] = (uint8_t)((counter + 2u) >> 8);
    row3b[2] = (uint8_t)((counter + 2u) >> 16);
    row3b[3] = (uint8_t)((counter + 2u) >> 24);
    memcpy(row3b + 4, nonce, 12);
    row3b[16] = (uint8_t)(counter + 3u);
    row3b[17] = (uint8_t)((counter + 3u) >> 8);
    row3b[18] = (uint8_t)((counter + 3u) >> 16);
    row3b[19] = (uint8_t)((counter + 3u) >> 24);
    memcpy(row3b + 20, nonce, 12);

    /* Groupe A : blocs counter et counter+1. Groupe B : +2 et +3.
     * Levier c_source : chaque vecteur est chargé depuis la mémoire, aucune
     * copie registre à registre (MOVUPS SSE au milieu de VEX). */
    __m256i a0 = _mm256_loadu_si256((const __m256i *)c2chacha4_sigma2);
    __m256i a1 = _mm256_loadu_si256((const __m256i *)krow1);
    __m256i a2 = _mm256_loadu_si256((const __m256i *)krow2);
    __m256i a3 = _mm256_loadu_si256((const __m256i *)row3a);
    __m256i b0 = _mm256_loadu_si256((const __m256i *)c2chacha4_sigma2);
    __m256i b1 = _mm256_loadu_si256((const __m256i *)krow1);
    __m256i b2 = _mm256_loadu_si256((const __m256i *)krow2);
    __m256i b3 = _mm256_loadu_si256((const __m256i *)row3b);

    for (int i = 0; i < 10; i++) {
        /* Tour de colonnes, groupes A et B entrelacés. */
        a0 = _mm256_add_epi32(a0, a1);
        b0 = _mm256_add_epi32(b0, b1);
        a3 = _mm256_xor_si256(a3, a0);
        b3 = _mm256_xor_si256(b3, b0);
        a3 = _mm256_shuffle_epi8(a3, rot16);
        b3 = _mm256_shuffle_epi8(b3, rot16);
        a2 = _mm256_add_epi32(a2, a3);
        b2 = _mm256_add_epi32(b2, b3);
        a1 = _mm256_xor_si256(a1, a2);
        b1 = _mm256_xor_si256(b1, b2);
        a1 = _mm256_or_si256(_mm256_slli_epi32(a1, 12), _mm256_srli_epi32(a1, 20));
        b1 = _mm256_or_si256(_mm256_slli_epi32(b1, 12), _mm256_srli_epi32(b1, 20));
        a0 = _mm256_add_epi32(a0, a1);
        b0 = _mm256_add_epi32(b0, b1);
        a3 = _mm256_xor_si256(a3, a0);
        b3 = _mm256_xor_si256(b3, b0);
        a3 = _mm256_shuffle_epi8(a3, rot8);
        b3 = _mm256_shuffle_epi8(b3, rot8);
        a2 = _mm256_add_epi32(a2, a3);
        b2 = _mm256_add_epi32(b2, b3);
        a1 = _mm256_xor_si256(a1, a2);
        b1 = _mm256_xor_si256(b1, b2);
        a1 = _mm256_or_si256(_mm256_slli_epi32(a1, 7), _mm256_srli_epi32(a1, 25));
        b1 = _mm256_or_si256(_mm256_slli_epi32(b1, 7), _mm256_srli_epi32(b1, 25));

        /* Diagonalisation : rotation des lignes 1, 2, 3 de 1, 2, 3 mots. */
        a1 = _mm256_shuffle_epi32(a1, 0x39);
        b1 = _mm256_shuffle_epi32(b1, 0x39);
        a2 = _mm256_shuffle_epi32(a2, 0x4e);
        b2 = _mm256_shuffle_epi32(b2, 0x4e);
        a3 = _mm256_shuffle_epi32(a3, 0x93);
        b3 = _mm256_shuffle_epi32(b3, 0x93);

        /* Tour de diagonales. */
        a0 = _mm256_add_epi32(a0, a1);
        b0 = _mm256_add_epi32(b0, b1);
        a3 = _mm256_xor_si256(a3, a0);
        b3 = _mm256_xor_si256(b3, b0);
        a3 = _mm256_shuffle_epi8(a3, rot16);
        b3 = _mm256_shuffle_epi8(b3, rot16);
        a2 = _mm256_add_epi32(a2, a3);
        b2 = _mm256_add_epi32(b2, b3);
        a1 = _mm256_xor_si256(a1, a2);
        b1 = _mm256_xor_si256(b1, b2);
        a1 = _mm256_or_si256(_mm256_slli_epi32(a1, 12), _mm256_srli_epi32(a1, 20));
        b1 = _mm256_or_si256(_mm256_slli_epi32(b1, 12), _mm256_srli_epi32(b1, 20));
        a0 = _mm256_add_epi32(a0, a1);
        b0 = _mm256_add_epi32(b0, b1);
        a3 = _mm256_xor_si256(a3, a0);
        b3 = _mm256_xor_si256(b3, b0);
        a3 = _mm256_shuffle_epi8(a3, rot8);
        b3 = _mm256_shuffle_epi8(b3, rot8);
        a2 = _mm256_add_epi32(a2, a3);
        b2 = _mm256_add_epi32(b2, b3);
        a1 = _mm256_xor_si256(a1, a2);
        b1 = _mm256_xor_si256(b1, b2);
        a1 = _mm256_or_si256(_mm256_slli_epi32(a1, 7), _mm256_srli_epi32(a1, 25));
        b1 = _mm256_or_si256(_mm256_slli_epi32(b1, 7), _mm256_srli_epi32(b1, 25));

        /* Retour à la forme colonnes. */
        a1 = _mm256_shuffle_epi32(a1, 0x93);
        b1 = _mm256_shuffle_epi32(b1, 0x93);
        a2 = _mm256_shuffle_epi32(a2, 0x4e);
        b2 = _mm256_shuffle_epi32(b2, 0x4e);
        a3 = _mm256_shuffle_epi32(a3, 0x39);
        b3 = _mm256_shuffle_epi32(b3, 0x39);
    }

    a0 = _mm256_add_epi32(a0, _mm256_loadu_si256((const __m256i *)c2chacha4_sigma2));
    a1 = _mm256_add_epi32(a1, _mm256_loadu_si256((const __m256i *)krow1));
    a2 = _mm256_add_epi32(a2, _mm256_loadu_si256((const __m256i *)krow2));
    a3 = _mm256_add_epi32(a3, _mm256_loadu_si256((const __m256i *)row3a));
    b0 = _mm256_add_epi32(b0, _mm256_loadu_si256((const __m256i *)c2chacha4_sigma2));
    b1 = _mm256_add_epi32(b1, _mm256_loadu_si256((const __m256i *)krow1));
    b2 = _mm256_add_epi32(b2, _mm256_loadu_si256((const __m256i *)krow2));
    b3 = _mm256_add_epi32(b3, _mm256_loadu_si256((const __m256i *)row3b));

    /* Un vecteur = 32 octets contigus d'un bloc (lo = rangées 0–1, hi = 2–3). */
    __m256i blo0 = _mm256_permute2x128_si256(a0, a1, 0x20);
    __m256i bhi0 = _mm256_permute2x128_si256(a2, a3, 0x20);
    __m256i blo1 = _mm256_permute2x128_si256(a0, a1, 0x31);
    __m256i bhi1 = _mm256_permute2x128_si256(a2, a3, 0x31);
    __m256i blo2 = _mm256_permute2x128_si256(b0, b1, 0x20);
    __m256i bhi2 = _mm256_permute2x128_si256(b2, b3, 0x20);
    __m256i blo3 = _mm256_permute2x128_si256(b0, b1, 0x31);
    __m256i bhi3 = _mm256_permute2x128_si256(b2, b3, 0x31);

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
    if (nfull >= 3) {
        _mm256_storeu_si256((__m256i *)(out + 128), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 128)), blo2));
        _mm256_storeu_si256((__m256i *)(out + 160), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 160)), bhi2));
    }
    if (nfull >= 4) {
        _mm256_storeu_si256((__m256i *)(out + 192), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 192)), blo3));
        _mm256_storeu_si256((__m256i *)(out + 224), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 224)), bhi3));
    }

    if (rem != 0) {
        uint8_t ks[64];
        switch (nfull) {
        case 0:
            _mm256_storeu_si256((__m256i *)(ks + 0), blo0);
            _mm256_storeu_si256((__m256i *)(ks + 32), bhi0);
            break;
        case 1:
            _mm256_storeu_si256((__m256i *)(ks + 0), blo1);
            _mm256_storeu_si256((__m256i *)(ks + 32), bhi1);
            break;
        case 2:
            _mm256_storeu_si256((__m256i *)(ks + 0), blo2);
            _mm256_storeu_si256((__m256i *)(ks + 32), bhi2);
            break;
        default:
            _mm256_storeu_si256((__m256i *)(ks + 0), blo3);
            _mm256_storeu_si256((__m256i *)(ks + 32), bhi3);
            break;
        }
        size_t base = nfull * 64;
        for (size_t k = 0; k < rem; k++) {
            out[base + k] = in[base + k] ^ ks[k];
        }
    }
    return n;
}

size_t c2chacha4_xor_blocks(uint8_t *out, const uint8_t *in, size_t n,
                            const uint8_t key[32], const uint8_t nonce[12],
                            uint32_t counter) {
    if (n > 512) {
        n = 512;
    }
    if (n == 0) {
        return 0;
    }
    size_t n1 = n;
    if (n1 > 256) {
        n1 = 256;
    }
    c2chacha4_xor256(out, in, n1, key, nonce, counter);
    if (n > 256) {
        c2chacha4_xor256(out + 256, in + 256, n - 256, key, nonce, counter + 4u);
    }
    return n;
}
