#include "c2_swizzle_simd.h"

#if defined(__x86_64__) || defined(_M_X64)
#include <immintrin.h>

void c2_swizzle_rgba_bgra_avx2(const uint8_t *src, uint8_t *dst, size_t num_pixels) {
    size_t i = 0;
    /* Masque de transposition RGBA <-> BGRA pour vpshufb sur deux voies 128-bit */
    const __m256i v_shuf = _mm256_setr_epi8(
        2, 1, 0, 3,  6, 5, 4, 7,  10, 9, 8, 11,  14, 13, 12, 15,
        2, 1, 0, 3,  6, 5, 4, 7,  10, 9, 8, 11,  14, 13, 12, 15
    );

    /* Déroulage 4x (128 octets = 32 pixels par itération) pour saturer le bus mémoire */
    size_t blocks32 = num_pixels >> 5;
    size_t limit32 = blocks32 << 5;

    for (i = 0; i < limit32; i += 32) {
        size_t off = i * 4;
        __m256i v0 = _mm256_loadu_si256((const __m256i *)(src + off + 0));
        __m256i v1 = _mm256_loadu_si256((const __m256i *)(src + off + 32));
        __m256i v2 = _mm256_loadu_si256((const __m256i *)(src + off + 64));
        __m256i v3 = _mm256_loadu_si256((const __m256i *)(src + off + 96));

        v0 = _mm256_shuffle_epi8(v0, v_shuf);
        v1 = _mm256_shuffle_epi8(v1, v_shuf);
        v2 = _mm256_shuffle_epi8(v2, v_shuf);
        v3 = _mm256_shuffle_epi8(v3, v_shuf);

        _mm256_storeu_si256((__m256i *)(dst + off + 0), v0);
        _mm256_storeu_si256((__m256i *)(dst + off + 32), v1);
        _mm256_storeu_si256((__m256i *)(dst + off + 64), v2);
        _mm256_storeu_si256((__m256i *)(dst + off + 96), v3);
    }

    /* Traitement des blocs de 8 pixels résiduels (32 octets) */
    size_t blocks8 = num_pixels >> 3;
    size_t limit8 = blocks8 << 3;

    for (; i < limit8; i += 8) {
        size_t off = i * 4;
        __m256i v = _mm256_loadu_si256((const __m256i *)(src + off));
        v = _mm256_shuffle_epi8(v, v_shuf);
        _mm256_storeu_si256((__m256i *)(dst + off), v);
    }

    /* Résidu scalaire pixel par pixel */
    for (; i < num_pixels; i++) {
        size_t off = i * 4;
        dst[off + 0] = src[off + 2];
        dst[off + 1] = src[off + 1];
        dst[off + 2] = src[off + 0];
        dst[off + 3] = src[off + 3];
    }
}

void c2_swizzle_rgba_bgra_inplace_avx2(uint8_t *data, size_t num_pixels) {
    size_t i = 0;
    const __m256i v_shuf = _mm256_setr_epi8(
        2, 1, 0, 3,  6, 5, 4, 7,  10, 9, 8, 11,  14, 13, 12, 15,
        2, 1, 0, 3,  6, 5, 4, 7,  10, 9, 8, 11,  14, 13, 12, 15
    );

    size_t blocks32 = num_pixels >> 5;
    size_t limit32 = blocks32 << 5;

    for (i = 0; i < limit32; i += 32) {
        size_t off = i * 4;
        __m256i v0 = _mm256_loadu_si256((const __m256i *)(data + off + 0));
        __m256i v1 = _mm256_loadu_si256((const __m256i *)(data + off + 32));
        __m256i v2 = _mm256_loadu_si256((const __m256i *)(data + off + 64));
        __m256i v3 = _mm256_loadu_si256((const __m256i *)(data + off + 96));

        v0 = _mm256_shuffle_epi8(v0, v_shuf);
        v1 = _mm256_shuffle_epi8(v1, v_shuf);
        v2 = _mm256_shuffle_epi8(v2, v_shuf);
        v3 = _mm256_shuffle_epi8(v3, v_shuf);

        _mm256_storeu_si256((__m256i *)(data + off + 0), v0);
        _mm256_storeu_si256((__m256i *)(data + off + 32), v1);
        _mm256_storeu_si256((__m256i *)(data + off + 64), v2);
        _mm256_storeu_si256((__m256i *)(data + off + 96), v3);
    }

    size_t blocks8 = num_pixels >> 3;
    size_t limit8 = blocks8 << 3;

    for (; i < limit8; i += 8) {
        size_t off = i * 4;
        __m256i v = _mm256_loadu_si256((const __m256i *)(data + off));
        v = _mm256_shuffle_epi8(v, v_shuf);
        _mm256_storeu_si256((__m256i *)(data + off), v);
    }

    for (; i < num_pixels; i++) {
        size_t off = i * 4;
        uint8_t tmp = data[off + 0];
        data[off + 0] = data[off + 2];
        data[off + 2] = tmp;
    }
}

void c2_swizzle_rgba_bgra_stream_avx2(const uint8_t *src, uint8_t *dst, size_t num_pixels) {
    size_t i = 0;
    const __m256i v_shuf = _mm256_setr_epi8(
        2, 1, 0, 3,  6, 5, 4, 7,  10, 9, 8, 11,  14, 13, 12, 15,
        2, 1, 0, 3,  6, 5, 4, 7,  10, 9, 8, 11,  14, 13, 12, 15
    );

    /* Écriture directe en mémoire par flux non-temporel si destination alignée 32B */
    if ((((uintptr_t)dst) & 31) != 0) {
        c2_swizzle_rgba_bgra_avx2(src, dst, num_pixels);
        return;
    }

    size_t blocks32 = num_pixels >> 5;
    size_t limit32 = blocks32 << 5;

    for (i = 0; i < limit32; i += 32) {
        size_t off = i * 4;
        __m256i v0 = _mm256_loadu_si256((const __m256i *)(src + off + 0));
        __m256i v1 = _mm256_loadu_si256((const __m256i *)(src + off + 32));
        __m256i v2 = _mm256_loadu_si256((const __m256i *)(src + off + 64));
        __m256i v3 = _mm256_loadu_si256((const __m256i *)(src + off + 96));

        v0 = _mm256_shuffle_epi8(v0, v_shuf);
        v1 = _mm256_shuffle_epi8(v1, v_shuf);
        v2 = _mm256_shuffle_epi8(v2, v_shuf);
        v3 = _mm256_shuffle_epi8(v3, v_shuf);

        _mm256_stream_si256((__m256i *)(dst + off + 0), v0);
        _mm256_stream_si256((__m256i *)(dst + off + 32), v1);
        _mm256_stream_si256((__m256i *)(dst + off + 64), v2);
        _mm256_stream_si256((__m256i *)(dst + off + 96), v3);
    }
    _mm_sfence();

    for (; i < num_pixels; i++) {
        size_t off = i * 4;
        dst[off + 0] = src[off + 2];
        dst[off + 1] = src[off + 1];
        dst[off + 2] = src[off + 0];
        dst[off + 3] = src[off + 3];
    }
}

#endif
