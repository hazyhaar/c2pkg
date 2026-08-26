/*
 * bench_entropy_c.c — Benchmark comparatif C99 : ARCHTIME Q8.8 vs Standard Float (libm)
 * Mesure du débit (Mo/s, Go/s), de la latence (ns/op) et de la précision bit-exacte.
 */

#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <stddef.h>
#include <string.h>
#include <math.h>
#include <time.h>

#include "c2blueteam.h"

/* 1. Implémentation standard avec math.h (double / float FPU) */
static double calc_entropy_float_std(const uint8_t *data, size_t len) {
    if (!data || len == 0) return 0.0;
    uint32_t freq[256];
    memset(freq, 0, sizeof(freq));
    for (size_t i = 0; i < len; i++) {
        freq[data[i]]++;
    }
    double entropy = 0.0;
    double dlen = (double)len;
    for (int i = 0; i < 256; i++) {
        if (freq[i] > 0) {
            double p = (double)freq[i] / dlen;
            entropy -= p * log2(p);
        }
    }
    return entropy;
}

static uint64_t get_time_ns(void) {
    struct timespec ts;
    clock_gettime(CLOCK_MONOTONIC, &ts);
    return (uint64_t)ts.tv_sec * 1000000000ULL + (uint64_t)ts.tv_nsec;
}

int main(void) {
    printf("========================================================================================\n");
    printf("   BENCHMARK C99 : ENTROPIE DE SHANNON ARCHTIME Q8.8 VS STANDARD FLOAT (libm)\n");
    printf("   Processeur : Intel(R) Core(TM) i9-14900K | Options : gcc -O3 -march=native\n");
    printf("========================================================================================\n\n");

    const size_t sizes[] = {16, 64, 128, 256, 512, 1024, 4096, 16384, 65536, 1048576};
    const int num_sizes = sizeof(sizes) / sizeof(sizes[0]);

    printf("%-10s | %-18s | %-18s | %-14s | %-10s\n",
           "Taille", "ARCHTIME Q8.8", "Standard Float", "Gain Vitesse", "Ecart Max");
    printf("-----------|--------------------|--------------------|----------------|------------\n");

    for (int s = 0; s < num_sizes; s++) {
        size_t len = sizes[s];
        uint8_t *buf = (uint8_t *)malloc(len);
        if (!buf) continue;

        /* Remplissage avec données pseudo-aléatoires variées */
        for (size_t i = 0; i < len; i++) {
            buf[i] = (uint8_t)((i * 137 + 73) ^ (i >> 3));
        }

        /* 1. Mesure de précision */
        double ent_float = calc_entropy_float_std(buf, len);
        uint32_t ent_q8 = c2bt_calc_entropy_8_8(buf, len);
        double ent_q8_float = (double)ent_q8 / 256.0;
        double diff = fabs(ent_float - ent_q8_float);

        /* 2. Warmup & Itérations */
        size_t iters = 100000000ULL / (len + 32);
        if (iters < 500) iters = 500;
        if (iters > 5000000) iters = 5000000;

        /* Warmup */
        volatile uint32_t sink_q8 = 0;
        volatile double sink_flt = 0.0;
        for (size_t i = 0; i < iters / 10; i++) {
            sink_q8 += c2bt_calc_entropy_8_8(buf, len);
            sink_flt += calc_entropy_float_std(buf, len);
        }

        /* Benchmark ARCHTIME */
        uint64_t t0 = get_time_ns();
        for (size_t i = 0; i < iters; i++) {
            sink_q8 += c2bt_calc_entropy_8_8(buf, len);
        }
        uint64_t t1 = get_time_ns();
        double elapsed_arch_sec = (double)(t1 - t0) / 1e9;
        double ns_per_op_arch = (double)(t1 - t0) / (double)iters;
        double mb_s_arch = ((double)len * (double)iters) / (elapsed_arch_sec * 1024.0 * 1024.0);

        /* Benchmark Standard Float */
        uint64_t t2 = get_time_ns();
        for (size_t i = 0; i < iters; i++) {
            sink_flt += calc_entropy_float_std(buf, len);
        }
        uint64_t t3 = get_time_ns();
        double elapsed_flt_sec = (double)(t3 - t2) / 1e9;
        double ns_per_op_flt = (double)(t3 - t2) / (double)iters;
        double mb_s_flt = ((double)len * (double)iters) / (elapsed_flt_sec * 1024.0 * 1024.0);

        double speedup = ns_per_op_flt / ns_per_op_arch;

        char size_str[32];
        if (len < 1024) sprintf(size_str, "%zu B", len);
        else if (len < 1048576) sprintf(size_str, "%zu KB", len / 1024);
        else sprintf(size_str, "%zu MB", len / 1048576);

        char arch_str[32];
        if (mb_s_arch >= 1024.0) sprintf(arch_str, "%.2f GB/s (%.0fns)", mb_s_arch / 1024.0, ns_per_op_arch);
        else sprintf(arch_str, "%.1f MB/s (%.0fns)", mb_s_arch, ns_per_op_arch);

        char flt_str[32];
        if (mb_s_flt >= 1024.0) sprintf(flt_str, "%.2f GB/s (%.0fns)", mb_s_flt / 1024.0, ns_per_op_flt);
        else sprintf(flt_str, "%.1f MB/s (%.0fns)", mb_s_flt, ns_per_op_flt);

        printf("%-10s | %-18s | %-18s | %5.2fx rapide  | %-10.5f b/o\n",
               size_str, arch_str, flt_str, speedup, diff);

        free(buf);
    }
    printf("----------------------------------------------------------------------------------------\n\n");
    return 0;
}
