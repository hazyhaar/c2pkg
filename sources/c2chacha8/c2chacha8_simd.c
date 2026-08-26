#include <immintrin.h>
#include "c2chacha8.h"

/* Masques pshufb 256 bits : masque 16 octets de c2chacha1 dupliqué sur les
 * deux moitiés 128 bits (shuffle_epi8 AVX2 est intra-voie). */
static const uint8_t c2chacha8_rot16_mask[32] = {
    2, 3, 0, 1, 6, 7, 4, 5, 10, 11, 8, 9, 14, 15, 12, 13,
    2, 3, 0, 1, 6, 7, 4, 5, 10, 11, 8, 9, 14, 15, 12, 13};
static const uint8_t c2chacha8_rot8_mask[32] = {
    3, 0, 1, 2, 7, 4, 5, 6, 11, 8, 9, 10, 15, 12, 13, 14,
    3, 0, 1, 2, 7, 4, 5, 6, 11, 8, 9, 10, 15, 12, 13, 14};

/* Constantes "expand 32-byte k", octets petit-boutistes (levier c_source). */
static const uint8_t c2chacha8_sigma[16] = {0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x20, 0x33,
                                            0x32, 0x2d, 0x62, 0x79, 0x74, 0x65, 0x20, 0x6b};

/* Grille de compteurs : voie j = counter + j. */
static const uint32_t c2chacha8_lane_ctr[8] = {0, 1, 2, 3, 4, 5, 6, 7};

static uint32_t load32_le(const uint8_t *p) {
    return (uint32_t)p[0]
        | ((uint32_t)p[1] << 8)
        | ((uint32_t)p[2] << 16)
        | ((uint32_t)p[3] << 24);
}

size_t c2chacha8_xor_blocks(uint8_t *out, const uint8_t *in, size_t n,
                            const uint8_t key[32], const uint8_t nonce[12],
                            uint32_t counter) {
    if (n > 512) {
        n = 512;
    }
    if (n == 0) {
        return 0;
    }

    __m256i rot16 = _mm256_loadu_si256((const __m256i *)c2chacha8_rot16_mask);
    __m256i rot8 = _mm256_loadu_si256((const __m256i *)c2chacha8_rot8_mask);

    /* Disposition un mot par registre : le registre k contient le mot k
     * des huit blocs. Diffusion par set1 depuis load32_le. Levier c_source :
     * l'état initial est rechargé depuis la mémoire pour l'addition finale. */
    __m256i x0 = _mm256_set1_epi32((int)load32_le(c2chacha8_sigma + 0));
    __m256i x1 = _mm256_set1_epi32((int)load32_le(c2chacha8_sigma + 4));
    __m256i x2 = _mm256_set1_epi32((int)load32_le(c2chacha8_sigma + 8));
    __m256i x3 = _mm256_set1_epi32((int)load32_le(c2chacha8_sigma + 12));
    __m256i x4 = _mm256_set1_epi32((int)load32_le(key + 0));
    __m256i x5 = _mm256_set1_epi32((int)load32_le(key + 4));
    __m256i x6 = _mm256_set1_epi32((int)load32_le(key + 8));
    __m256i x7 = _mm256_set1_epi32((int)load32_le(key + 12));
    __m256i x8 = _mm256_set1_epi32((int)load32_le(key + 16));
    __m256i x9 = _mm256_set1_epi32((int)load32_le(key + 20));
    __m256i x10 = _mm256_set1_epi32((int)load32_le(key + 24));
    __m256i x11 = _mm256_set1_epi32((int)load32_le(key + 28));
    __m256i x12 = _mm256_add_epi32(_mm256_set1_epi32((int)counter),
                                   _mm256_loadu_si256((const __m256i *)c2chacha8_lane_ctr));
    __m256i x13 = _mm256_set1_epi32((int)load32_le(nonce + 0));
    __m256i x14 = _mm256_set1_epi32((int)load32_le(nonce + 4));
    __m256i x15 = _mm256_set1_epi32((int)load32_le(nonce + 8));

    for (int i = 0; i < 10; i++) {
        /* Tour de colonnes : (0,4,8,12), (1,5,9,13), (2,6,10,14), (3,7,11,15). */
        x0 = _mm256_add_epi32(x0, x4);
        x12 = _mm256_xor_si256(x12, x0);
        x12 = _mm256_shuffle_epi8(x12, rot16);
        x8 = _mm256_add_epi32(x8, x12);
        x4 = _mm256_xor_si256(x4, x8);
        x4 = _mm256_or_si256(_mm256_slli_epi32(x4, 12), _mm256_srli_epi32(x4, 20));
        x0 = _mm256_add_epi32(x0, x4);
        x12 = _mm256_xor_si256(x12, x0);
        x12 = _mm256_shuffle_epi8(x12, rot8);
        x8 = _mm256_add_epi32(x8, x12);
        x4 = _mm256_xor_si256(x4, x8);
        x4 = _mm256_or_si256(_mm256_slli_epi32(x4, 7), _mm256_srli_epi32(x4, 25));

        x1 = _mm256_add_epi32(x1, x5);
        x13 = _mm256_xor_si256(x13, x1);
        x13 = _mm256_shuffle_epi8(x13, rot16);
        x9 = _mm256_add_epi32(x9, x13);
        x5 = _mm256_xor_si256(x5, x9);
        x5 = _mm256_or_si256(_mm256_slli_epi32(x5, 12), _mm256_srli_epi32(x5, 20));
        x1 = _mm256_add_epi32(x1, x5);
        x13 = _mm256_xor_si256(x13, x1);
        x13 = _mm256_shuffle_epi8(x13, rot8);
        x9 = _mm256_add_epi32(x9, x13);
        x5 = _mm256_xor_si256(x5, x9);
        x5 = _mm256_or_si256(_mm256_slli_epi32(x5, 7), _mm256_srli_epi32(x5, 25));

        x2 = _mm256_add_epi32(x2, x6);
        x14 = _mm256_xor_si256(x14, x2);
        x14 = _mm256_shuffle_epi8(x14, rot16);
        x10 = _mm256_add_epi32(x10, x14);
        x6 = _mm256_xor_si256(x6, x10);
        x6 = _mm256_or_si256(_mm256_slli_epi32(x6, 12), _mm256_srli_epi32(x6, 20));
        x2 = _mm256_add_epi32(x2, x6);
        x14 = _mm256_xor_si256(x14, x2);
        x14 = _mm256_shuffle_epi8(x14, rot8);
        x10 = _mm256_add_epi32(x10, x14);
        x6 = _mm256_xor_si256(x6, x10);
        x6 = _mm256_or_si256(_mm256_slli_epi32(x6, 7), _mm256_srli_epi32(x6, 25));

        x3 = _mm256_add_epi32(x3, x7);
        x15 = _mm256_xor_si256(x15, x3);
        x15 = _mm256_shuffle_epi8(x15, rot16);
        x11 = _mm256_add_epi32(x11, x15);
        x7 = _mm256_xor_si256(x7, x11);
        x7 = _mm256_or_si256(_mm256_slli_epi32(x7, 12), _mm256_srli_epi32(x7, 20));
        x3 = _mm256_add_epi32(x3, x7);
        x15 = _mm256_xor_si256(x15, x3);
        x15 = _mm256_shuffle_epi8(x15, rot8);
        x11 = _mm256_add_epi32(x11, x15);
        x7 = _mm256_xor_si256(x7, x11);
        x7 = _mm256_or_si256(_mm256_slli_epi32(x7, 7), _mm256_srli_epi32(x7, 25));

        /* Tour de diagonales : (0,5,10,15), (1,6,11,12), (2,7,8,13), (3,4,9,14). */
        x0 = _mm256_add_epi32(x0, x5);
        x15 = _mm256_xor_si256(x15, x0);
        x15 = _mm256_shuffle_epi8(x15, rot16);
        x10 = _mm256_add_epi32(x10, x15);
        x5 = _mm256_xor_si256(x5, x10);
        x5 = _mm256_or_si256(_mm256_slli_epi32(x5, 12), _mm256_srli_epi32(x5, 20));
        x0 = _mm256_add_epi32(x0, x5);
        x15 = _mm256_xor_si256(x15, x0);
        x15 = _mm256_shuffle_epi8(x15, rot8);
        x10 = _mm256_add_epi32(x10, x15);
        x5 = _mm256_xor_si256(x5, x10);
        x5 = _mm256_or_si256(_mm256_slli_epi32(x5, 7), _mm256_srli_epi32(x5, 25));

        x1 = _mm256_add_epi32(x1, x6);
        x12 = _mm256_xor_si256(x12, x1);
        x12 = _mm256_shuffle_epi8(x12, rot16);
        x11 = _mm256_add_epi32(x11, x12);
        x6 = _mm256_xor_si256(x6, x11);
        x6 = _mm256_or_si256(_mm256_slli_epi32(x6, 12), _mm256_srli_epi32(x6, 20));
        x1 = _mm256_add_epi32(x1, x6);
        x12 = _mm256_xor_si256(x12, x1);
        x12 = _mm256_shuffle_epi8(x12, rot8);
        x11 = _mm256_add_epi32(x11, x12);
        x6 = _mm256_xor_si256(x6, x11);
        x6 = _mm256_or_si256(_mm256_slli_epi32(x6, 7), _mm256_srli_epi32(x6, 25));

        x2 = _mm256_add_epi32(x2, x7);
        x13 = _mm256_xor_si256(x13, x2);
        x13 = _mm256_shuffle_epi8(x13, rot16);
        x8 = _mm256_add_epi32(x8, x13);
        x7 = _mm256_xor_si256(x7, x8);
        x7 = _mm256_or_si256(_mm256_slli_epi32(x7, 12), _mm256_srli_epi32(x7, 20));
        x2 = _mm256_add_epi32(x2, x7);
        x13 = _mm256_xor_si256(x13, x2);
        x13 = _mm256_shuffle_epi8(x13, rot8);
        x8 = _mm256_add_epi32(x8, x13);
        x7 = _mm256_xor_si256(x7, x8);
        x7 = _mm256_or_si256(_mm256_slli_epi32(x7, 7), _mm256_srli_epi32(x7, 25));

        x3 = _mm256_add_epi32(x3, x4);
        x14 = _mm256_xor_si256(x14, x3);
        x14 = _mm256_shuffle_epi8(x14, rot16);
        x9 = _mm256_add_epi32(x9, x14);
        x4 = _mm256_xor_si256(x4, x9);
        x4 = _mm256_or_si256(_mm256_slli_epi32(x4, 12), _mm256_srli_epi32(x4, 20));
        x3 = _mm256_add_epi32(x3, x4);
        x14 = _mm256_xor_si256(x14, x3);
        x14 = _mm256_shuffle_epi8(x14, rot8);
        x9 = _mm256_add_epi32(x9, x14);
        x4 = _mm256_xor_si256(x4, x9);
        x4 = _mm256_or_si256(_mm256_slli_epi32(x4, 7), _mm256_srli_epi32(x4, 25));
    }

    x0 = _mm256_add_epi32(x0, _mm256_set1_epi32((int)load32_le(c2chacha8_sigma + 0)));
    x1 = _mm256_add_epi32(x1, _mm256_set1_epi32((int)load32_le(c2chacha8_sigma + 4)));
    x2 = _mm256_add_epi32(x2, _mm256_set1_epi32((int)load32_le(c2chacha8_sigma + 8)));
    x3 = _mm256_add_epi32(x3, _mm256_set1_epi32((int)load32_le(c2chacha8_sigma + 12)));
    x4 = _mm256_add_epi32(x4, _mm256_set1_epi32((int)load32_le(key + 0)));
    x5 = _mm256_add_epi32(x5, _mm256_set1_epi32((int)load32_le(key + 4)));
    x6 = _mm256_add_epi32(x6, _mm256_set1_epi32((int)load32_le(key + 8)));
    x7 = _mm256_add_epi32(x7, _mm256_set1_epi32((int)load32_le(key + 12)));
    x8 = _mm256_add_epi32(x8, _mm256_set1_epi32((int)load32_le(key + 16)));
    x9 = _mm256_add_epi32(x9, _mm256_set1_epi32((int)load32_le(key + 20)));
    x10 = _mm256_add_epi32(x10, _mm256_set1_epi32((int)load32_le(key + 24)));
    x11 = _mm256_add_epi32(x11, _mm256_set1_epi32((int)load32_le(key + 28)));
    x12 = _mm256_add_epi32(x12, _mm256_add_epi32(_mm256_set1_epi32((int)counter),
                                                _mm256_loadu_si256((const __m256i *)c2chacha8_lane_ctr)));
    x13 = _mm256_add_epi32(x13, _mm256_set1_epi32((int)load32_le(nonce + 0)));
    x14 = _mm256_add_epi32(x14, _mm256_set1_epi32((int)load32_le(nonce + 4)));
    x15 = _mm256_add_epi32(x15, _mm256_set1_epi32((int)load32_le(nonce + 8)));

    /* Transposition 8×8 des mots 0–7 : un vecteur = 32 premiers octets d'un bloc. */
    __m256i t0 = _mm256_unpacklo_epi32(x0, x1);
    __m256i t1 = _mm256_unpackhi_epi32(x0, x1);
    __m256i t2 = _mm256_unpacklo_epi32(x2, x3);
    __m256i t3 = _mm256_unpackhi_epi32(x2, x3);
    __m256i t4 = _mm256_unpacklo_epi32(x4, x5);
    __m256i t5 = _mm256_unpackhi_epi32(x4, x5);
    __m256i t6 = _mm256_unpacklo_epi32(x6, x7);
    __m256i t7 = _mm256_unpackhi_epi32(x6, x7);
    __m256i u0 = _mm256_unpacklo_epi64(t0, t2);
    __m256i u1 = _mm256_unpackhi_epi64(t0, t2);
    __m256i u2 = _mm256_unpacklo_epi64(t1, t3);
    __m256i u3 = _mm256_unpackhi_epi64(t1, t3);
    __m256i u4 = _mm256_unpacklo_epi64(t4, t6);
    __m256i u5 = _mm256_unpackhi_epi64(t4, t6);
    __m256i u6 = _mm256_unpacklo_epi64(t5, t7);
    __m256i u7 = _mm256_unpackhi_epi64(t5, t7);
    __m256i blo0 = _mm256_permute2x128_si256(u0, u4, 0x20);
    __m256i blo1 = _mm256_permute2x128_si256(u1, u5, 0x20);
    __m256i blo2 = _mm256_permute2x128_si256(u2, u6, 0x20);
    __m256i blo3 = _mm256_permute2x128_si256(u3, u7, 0x20);
    __m256i blo4 = _mm256_permute2x128_si256(u0, u4, 0x31);
    __m256i blo5 = _mm256_permute2x128_si256(u1, u5, 0x31);
    __m256i blo6 = _mm256_permute2x128_si256(u2, u6, 0x31);
    __m256i blo7 = _mm256_permute2x128_si256(u3, u7, 0x31);

    /* Transposition 8×8 des mots 8–15 : un vecteur = 32 derniers octets d'un bloc. */
    t0 = _mm256_unpacklo_epi32(x8, x9);
    t1 = _mm256_unpackhi_epi32(x8, x9);
    t2 = _mm256_unpacklo_epi32(x10, x11);
    t3 = _mm256_unpackhi_epi32(x10, x11);
    t4 = _mm256_unpacklo_epi32(x12, x13);
    t5 = _mm256_unpackhi_epi32(x12, x13);
    t6 = _mm256_unpacklo_epi32(x14, x15);
    t7 = _mm256_unpackhi_epi32(x14, x15);
    u0 = _mm256_unpacklo_epi64(t0, t2);
    u1 = _mm256_unpackhi_epi64(t0, t2);
    u2 = _mm256_unpacklo_epi64(t1, t3);
    u3 = _mm256_unpackhi_epi64(t1, t3);
    u4 = _mm256_unpacklo_epi64(t4, t6);
    u5 = _mm256_unpackhi_epi64(t4, t6);
    u6 = _mm256_unpacklo_epi64(t5, t7);
    u7 = _mm256_unpackhi_epi64(t5, t7);
    __m256i bhi0 = _mm256_permute2x128_si256(u0, u4, 0x20);
    __m256i bhi1 = _mm256_permute2x128_si256(u1, u5, 0x20);
    __m256i bhi2 = _mm256_permute2x128_si256(u2, u6, 0x20);
    __m256i bhi3 = _mm256_permute2x128_si256(u3, u7, 0x20);
    __m256i bhi4 = _mm256_permute2x128_si256(u0, u4, 0x31);
    __m256i bhi5 = _mm256_permute2x128_si256(u1, u5, 0x31);
    __m256i bhi6 = _mm256_permute2x128_si256(u2, u6, 0x31);
    __m256i bhi7 = _mm256_permute2x128_si256(u3, u7, 0x31);

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
    if (nfull >= 5) {
        _mm256_storeu_si256((__m256i *)(out + 256), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 256)), blo4));
        _mm256_storeu_si256((__m256i *)(out + 288), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 288)), bhi4));
    }
    if (nfull >= 6) {
        _mm256_storeu_si256((__m256i *)(out + 320), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 320)), blo5));
        _mm256_storeu_si256((__m256i *)(out + 352), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 352)), bhi5));
    }
    if (nfull >= 7) {
        _mm256_storeu_si256((__m256i *)(out + 384), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 384)), blo6));
        _mm256_storeu_si256((__m256i *)(out + 416), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 416)), bhi6));
    }
    if (nfull >= 8) {
        _mm256_storeu_si256((__m256i *)(out + 448), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 448)), blo7));
        _mm256_storeu_si256((__m256i *)(out + 480), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 480)), bhi7));
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
        case 3:
            _mm256_storeu_si256((__m256i *)(ks + 0), blo3);
            _mm256_storeu_si256((__m256i *)(ks + 32), bhi3);
            break;
        case 4:
            _mm256_storeu_si256((__m256i *)(ks + 0), blo4);
            _mm256_storeu_si256((__m256i *)(ks + 32), bhi4);
            break;
        case 5:
            _mm256_storeu_si256((__m256i *)(ks + 0), blo5);
            _mm256_storeu_si256((__m256i *)(ks + 32), bhi5);
            break;
        case 6:
            _mm256_storeu_si256((__m256i *)(ks + 0), blo6);
            _mm256_storeu_si256((__m256i *)(ks + 32), bhi6);
            break;
        default:
            _mm256_storeu_si256((__m256i *)(ks + 0), blo7);
            _mm256_storeu_si256((__m256i *)(ks + 32), bhi7);
            break;
        }
        size_t base = nfull * 64;
        for (size_t k = 0; k < rem; k++) {
            out[base + k] = in[base + k] ^ ks[k];
        }
    }
    return n;
}
