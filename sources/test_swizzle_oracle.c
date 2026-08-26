#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include <time.h>
#include <assert.h>
#include "c2_swizzle_simd.h"

#if defined(__x86_64__) || defined(_M_X64)
#include <immintrin.h>
#endif

static uint64_t hash_bytes(const uint8_t *data, size_t len) {
    uint64_t h = 0xcbf29ce484222325ULL;
    for (size_t i = 0; i < len; i++) {
        h ^= data[i];
        h *= 0x100000001b3ULL;
    }
    return h;
}

int main(int argc, char **argv) {
    (void)argc;
    (void)argv;
    setvbuf(stdout, NULL, _IONBF, 0);
    setvbuf(stderr, NULL, _IONBF, 0);

    printf("[ORACLE] Démarrage du banc de test de parité c2_swizzle_simd...\n");

    /* 1. Test unitaire mot 32-bit et 64-bit */
    uint32_t px_rgba = 0xAA332211U; /* A=0xAA, B=0x33, G=0x22, R=0x11 */
    uint32_t px_bgra = c2_swizzle_pixel(px_rgba);
    uint32_t px_expected = 0xAA112233U; /* A=0xAA, R=0x11, G=0x22, B=0x33 */
    if (px_bgra != px_expected) {
        fprintf(stderr, "FAIL: c2_swizzle_pixel 0x%08X != 0x%08X (got 0x%08X)\n", px_rgba, px_expected, px_bgra);
        return 1;
    }
    if (c2_swizzle_pixel(px_bgra) != px_rgba) {
        fprintf(stderr, "FAIL: c2_swizzle_pixel involution non respectée\n");
        return 1;
    }

    uint64_t pair_rgba = 0xFF443322ULL << 32 | 0xAA332211ULL;
    uint64_t pair_bgra = c2_swizzle_pair(pair_rgba);
    uint64_t pair_expected = 0xFF223344ULL << 32 | 0xAA112233ULL;
    if (pair_bgra != pair_expected) {
        fprintf(stderr, "FAIL: c2_swizzle_pair 0x%016llX != 0x%016llX\n", (unsigned long long)pair_bgra, (unsigned long long)pair_expected);
        return 1;
    }

    printf("[ORACLE] Test unitaire 32/64 bit validé.\n");

    /* 2. Test sur différentes tailles (cas limites : 0, 1, 2, 3, 4, 7, 8, 15, 16, 31, 32, 63, 64, 127, 128, 1024, 65536) */
    static const size_t test_sizes[] = {
        0, 1, 2, 3, 4, 5, 7, 8, 9, 15, 16, 17, 31, 32, 33, 63, 64, 65, 127, 128, 255, 256, 1024, 65536
    };
    size_t num_sizes = sizeof(test_sizes) / sizeof(test_sizes[0]);

    for (size_t s = 0; s < num_sizes; s++) {
        size_t n = test_sizes[s];
        size_t byte_len = n * 4;
        if (n == 0) continue;

        uint8_t *src = (uint8_t *)malloc(byte_len);
        uint8_t *dst_scalar = (uint8_t *)malloc(byte_len);
        uint8_t *dst_inplace = (uint8_t *)malloc(byte_len);
#if defined(__x86_64__) || defined(_M_X64)
        uint8_t *dst_avx2 = (uint8_t *)malloc(byte_len);
        uint8_t *dst_stream = (uint8_t *)malloc(byte_len);
#endif

        assert(src && dst_scalar && dst_inplace);

        /* Initialisation déterministe pseudo-aléatoire */
        for (size_t i = 0; i < byte_len; i++) {
            src[i] = (uint8_t)((i * 37 + 101) & 0xFF);
        }

        /* Test scalaire déroulé */
        c2_swizzle_rgba_bgra(src, dst_scalar, n);

        /* Vérification bit-exacte pixel par pixel */
        for (size_t i = 0; i < n; i++) {
            size_t off = i * 4;
            if (dst_scalar[off + 0] != src[off + 2] ||
                dst_scalar[off + 1] != src[off + 1] ||
                dst_scalar[off + 2] != src[off + 0] ||
                dst_scalar[off + 3] != src[off + 3]) {
                fprintf(stderr, "FAIL: Non-concordance scalaire taille %zu pixel %zu\n", n, i);
                return 1;
            }
        }

        /* Test in-place */
        memcpy(dst_inplace, src, byte_len);
        c2_swizzle_rgba_bgra_inplace(dst_inplace, n);
        if (memcmp(dst_scalar, dst_inplace, byte_len) != 0) {
            fprintf(stderr, "FAIL: Divergence in-place vs out-of-place taille %zu\n", n);
            return 1;
        }

        /* Test double application in-place (retour à l'identique) */
        c2_swizzle_rgba_bgra_inplace(dst_inplace, n);
        if (memcmp(src, dst_inplace, byte_len) != 0) {
            fprintf(stderr, "FAIL: Involution in-place échouée taille %zu\n", n);
            return 1;
        }

#if defined(__x86_64__) || defined(_M_X64)
        /* Test AVX2 vectoriel */
        c2_swizzle_rgba_bgra_avx2(src, dst_avx2, n);
        if (memcmp(dst_scalar, dst_avx2, byte_len) != 0) {
            fprintf(stderr, "FAIL: Divergence AVX2 vs Scalaire taille %zu\n", n);
            return 1;
        }

        /* Test AVX2 Stream */
        c2_swizzle_rgba_bgra_stream_avx2(src, dst_stream, n);
        if (memcmp(dst_scalar, dst_stream, byte_len) != 0) {
            fprintf(stderr, "FAIL: Divergence AVX2 Stream vs Scalaire taille %zu\n", n);
            return 1;
        }
#endif

        free(src);
        free(dst_scalar);
        free(dst_inplace);
#if defined(__x86_64__) || defined(_M_X64)
        free(dst_avx2);
        free(dst_stream);
#endif
    }

    /* 3. Test de conversions ARGB */
    size_t n_argb = 1000;
    uint8_t *src_argb = (uint8_t *)malloc(n_argb * 4);
    uint8_t *dst_rgba = (uint8_t *)malloc(n_argb * 4);
    uint8_t *dst_back_argb = (uint8_t *)malloc(n_argb * 4);
    for (size_t i = 0; i < n_argb * 4; i++) {
        src_argb[i] = (uint8_t)(i & 0xFF);
    }
    c2_swizzle_argb_to_rgba(src_argb, dst_rgba, n_argb);
    c2_swizzle_rgba_to_argb(dst_rgba, dst_back_argb, n_argb);
    if (memcmp(src_argb, dst_back_argb, n_argb * 4) != 0) {
        fprintf(stderr, "FAIL: ARGB <-> RGBA roundtrip non identique\n");
        return 1;
    }
    free(src_argb);
    free(dst_rgba);
    free(dst_back_argb);

    /* 4. Empreinte KAT opposable sur 100 000 pixels */
    size_t kat_pixels = 100000;
    size_t kat_bytes = kat_pixels * 4;
    uint8_t *kat_buf = (uint8_t *)malloc(kat_bytes);
    uint8_t *kat_out = (uint8_t *)malloc(kat_bytes);
    for (size_t i = 0; i < kat_bytes; i++) {
        kat_buf[i] = (uint8_t)((i * 131 + 17) & 0xFF);
    }
    c2_swizzle_rgba_bgra(kat_buf, kat_out, kat_pixels);
    uint64_t kat_hash = hash_bytes(kat_out, kat_bytes);
    printf("[ORACLE] KAT RGBA->BGRA 100k px Hash = 0x%016llX\n", (unsigned long long)kat_hash);
    free(kat_buf);
    free(kat_out);

    printf("[ORACLE] SUCCÈS : 100%% des vérifications de parité bit-exacte validées.\n");
    fflush(stdout);
    return 0;
}
