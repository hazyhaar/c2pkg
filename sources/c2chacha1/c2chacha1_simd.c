#include <immintrin.h>
#include <string.h>
#include "c2chacha1.h"

/* Masques pshufb : rotation gauche de 16 puis 8 bits dans chaque mot 32 bits. */
static const uint8_t c2chacha1_rot16_mask[16] = {2, 3, 0, 1, 6, 7, 4, 5, 10, 11, 8, 9, 14, 15, 12, 13};
static const uint8_t c2chacha1_rot8_mask[16] = {3, 0, 1, 2, 7, 4, 5, 6, 11, 8, 9, 10, 15, 12, 13, 14};

/* Constantes "expand 32-byte k", octets petit-boutistes (levier c_source :
 * les chargements 128 bits se font depuis des tableaux d'octets, forme que
 * l'émetteur type sans réinterprétation). */
static const uint8_t c2chacha1_sigma[16] = {0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x20, 0x33,
                                            0x32, 0x2d, 0x62, 0x79, 0x74, 0x65, 0x20, 0x6b};

size_t c2chacha1_xor_block(uint8_t *out, const uint8_t *in, size_t n,
                           const uint8_t key[32], const uint8_t nonce[12],
                           uint32_t counter) {
    if (n > 64) {
        n = 64;
    }

    /* État initial : quatre lignes de quatre mots ; la ligne 3 est
     * compteur (petit-boutiste) puis nonce, assemblée en octets. */
    uint8_t row3[16];
    row3[0] = (uint8_t)counter;
    row3[1] = (uint8_t)(counter >> 8);
    row3[2] = (uint8_t)(counter >> 16);
    row3[3] = (uint8_t)(counter >> 24);
    memcpy(row3 + 4, nonce, 12);

    __m128i rot16 = _mm_loadu_si128((const __m128i *)c2chacha1_rot16_mask);
    __m128i rot8 = _mm_loadu_si128((const __m128i *)c2chacha1_rot8_mask);

    /* Levier c_source : l'état de travail est chargé directement (aucune
     * copie registre à registre, que le compilateur Go encode en MOVUPS SSE
     * hérité au milieu de code VEX) ; l'état initial est rechargé depuis la
     * mémoire pour l'addition finale. */
    __m128i v0 = _mm_loadu_si128((const __m128i *)c2chacha1_sigma);
    __m128i v1 = _mm_loadu_si128((const __m128i *)(key + 0));
    __m128i v2 = _mm_loadu_si128((const __m128i *)(key + 16));
    __m128i v3 = _mm_loadu_si128((const __m128i *)row3);

    for (int i = 0; i < 10; i++) {
        /* Tour de colonnes : quarter-round vectorisé sur les quatre colonnes. */
        v0 = _mm_add_epi32(v0, v1);
        v3 = _mm_xor_si128(v3, v0);
        v3 = _mm_shuffle_epi8(v3, rot16);
        v2 = _mm_add_epi32(v2, v3);
        v1 = _mm_xor_si128(v1, v2);
        v1 = _mm_or_si128(_mm_slli_epi32(v1, 12), _mm_srli_epi32(v1, 20));
        v0 = _mm_add_epi32(v0, v1);
        v3 = _mm_xor_si128(v3, v0);
        v3 = _mm_shuffle_epi8(v3, rot8);
        v2 = _mm_add_epi32(v2, v3);
        v1 = _mm_xor_si128(v1, v2);
        v1 = _mm_or_si128(_mm_slli_epi32(v1, 7), _mm_srli_epi32(v1, 25));

        /* Diagonalisation : rotation des lignes 1, 2, 3 de 1, 2, 3 mots. */
        v1 = _mm_shuffle_epi32(v1, 0x39); /* 1,2,3,0 */
        v2 = _mm_shuffle_epi32(v2, 0x4e); /* 2,3,0,1 */
        v3 = _mm_shuffle_epi32(v3, 0x93); /* 3,0,1,2 */

        /* Tour de diagonales. */
        v0 = _mm_add_epi32(v0, v1);
        v3 = _mm_xor_si128(v3, v0);
        v3 = _mm_shuffle_epi8(v3, rot16);
        v2 = _mm_add_epi32(v2, v3);
        v1 = _mm_xor_si128(v1, v2);
        v1 = _mm_or_si128(_mm_slli_epi32(v1, 12), _mm_srli_epi32(v1, 20));
        v0 = _mm_add_epi32(v0, v1);
        v3 = _mm_xor_si128(v3, v0);
        v3 = _mm_shuffle_epi8(v3, rot8);
        v2 = _mm_add_epi32(v2, v3);
        v1 = _mm_xor_si128(v1, v2);
        v1 = _mm_or_si128(_mm_slli_epi32(v1, 7), _mm_srli_epi32(v1, 25));

        /* Retour à la forme colonnes. */
        v1 = _mm_shuffle_epi32(v1, 0x93);
        v2 = _mm_shuffle_epi32(v2, 0x4e);
        v3 = _mm_shuffle_epi32(v3, 0x39);
    }

    v0 = _mm_add_epi32(v0, _mm_loadu_si128((const __m128i *)c2chacha1_sigma));
    v1 = _mm_add_epi32(v1, _mm_loadu_si128((const __m128i *)(key + 0)));
    v2 = _mm_add_epi32(v2, _mm_loadu_si128((const __m128i *)(key + 16)));
    v3 = _mm_add_epi32(v3, _mm_loadu_si128((const __m128i *)row3));

    /* Bloc plein : XOR vectoriel direct, sans tampon de keystream (levier
     * c_source : la forme scalaire octet par octet n'est pas vectorisée par
     * le compilateur Go). */
    if (n == 64) {
        _mm_storeu_si128((__m128i *)(out + 0), _mm_xor_si128(_mm_loadu_si128((const __m128i *)(in + 0)), v0));
        _mm_storeu_si128((__m128i *)(out + 16), _mm_xor_si128(_mm_loadu_si128((const __m128i *)(in + 16)), v1));
        _mm_storeu_si128((__m128i *)(out + 32), _mm_xor_si128(_mm_loadu_si128((const __m128i *)(in + 32)), v2));
        _mm_storeu_si128((__m128i *)(out + 48), _mm_xor_si128(_mm_loadu_si128((const __m128i *)(in + 48)), v3));
        return n;
    }

    uint8_t ks[64];
    _mm_storeu_si128((__m128i *)(ks + 0), v0);
    _mm_storeu_si128((__m128i *)(ks + 16), v1);
    _mm_storeu_si128((__m128i *)(ks + 32), v2);
    _mm_storeu_si128((__m128i *)(ks + 48), v3);

    for (size_t k = 0; k < n; k++) {
        out[k] = in[k] ^ ks[k];
    }
    return n;
}
