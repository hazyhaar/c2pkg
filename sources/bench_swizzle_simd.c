#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include <time.h>
#include "c2_swizzle_simd.h"

#if defined(__x86_64__) || defined(_M_X64)
#include <immintrin.h>
#endif

static double get_time_sec(void) {
    struct timespec ts;
    clock_gettime(CLOCK_MONOTONIC, &ts);
    return (double)ts.tv_sec + (double)ts.tv_nsec * 1e-9;
}

int main(void) {
    const size_t test_sizes[] = {
        512 * 1024,         /* 512 Ko (Cache L2) */
        8 * 1024 * 1024,    /* 8 Mo (Cache L3 / 1080p) */
        64 * 1024 * 1024    /* 64 Mo (Bus DRAM) */
    };
    const char *labels[] = {"L2 (512 Ko)", "L3 (8 Mo / 1080p)", "DRAM (64 Mo)"};
    const int iters[] = {20000, 1000, 100};

    for (int s = 0; s < 3; s++) {
        size_t byte_len = test_sizes[s];
        size_t num_pixels = byte_len / 4;
        int iterations = iters[s];

        uint8_t *src = NULL;
        uint8_t *dst = NULL;
        posix_memalign((void **)&src, 32, byte_len);
        posix_memalign((void **)&dst, 32, byte_len);
        memset(src, 0x5A, byte_len);

        printf("\n--- Palier %s (%d itérations) ---\n", labels[s], iterations);

        /* 1. Scalaire */
        double t0 = get_time_sec();
        for (int i = 0; i < iterations; i++) {
            c2_swizzle_rgba_bgra(src, dst, num_pixels);
        }
        double t1 = get_time_sec();
        double scalar_gbps = ((double)byte_len * iterations * 2.0 / (t1 - t0)) / 1e9;
        printf("Scalaire déroulé : %.3f µs/op | Débit : %.2f Go/s\n", (t1 - t0) * 1e6 / iterations, scalar_gbps);

#if defined(__x86_64__) || defined(_M_X64)
        /* 2. AVX2 vpshufb */
        t0 = get_time_sec();
        for (int i = 0; i < iterations; i++) {
            c2_swizzle_rgba_bgra_avx2(src, dst, num_pixels);
        }
        t1 = get_time_sec();
        double avx2_gbps = ((double)byte_len * iterations * 2.0 / (t1 - t0)) / 1e9;
        printf("AVX2 vpshufb     : %.3f µs/op | Débit : %.2f Go/s\n", (t1 - t0) * 1e6 / iterations, avx2_gbps);

        /* 3. AVX2 Stream */
        t0 = get_time_sec();
        for (int i = 0; i < iterations; i++) {
            c2_swizzle_rgba_bgra_stream_avx2(src, dst, num_pixels);
        }
        t1 = get_time_sec();
        double stream_gbps = ((double)byte_len * iterations * 2.0 / (t1 - t0)) / 1e9;
        printf("AVX2 Stream NT   : %.3f µs/op | Débit : %.2f Go/s\n", (t1 - t0) * 1e6 / iterations, stream_gbps);
#endif
        free(src);
        free(dst);
    }
    return 0;
}
