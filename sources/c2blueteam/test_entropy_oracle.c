/*
 * test_entropy_oracle.c — Harnais de validation mathématique de probe_entropy.c
 * Compare bit à bit l'implémentation Q8.8 contre la formule flottante de Shannon.
 */

#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include <math.h>
#include <assert.h>
#include "c2blueteam.h"

static double reference_shannon_entropy(const uint8_t *data, size_t len) {
    if (data == NULL || len == 0) return 0.0;
    
    uint32_t freq[256] = {0};
    for (size_t i = 0; i < len; i++) {
        freq[data[i]]++;
    }
    
    double entropy = 0.0;
    for (int i = 0; i < 256; i++) {
        if (freq[i] > 0) {
            double p = (double)freq[i] / (double)len;
            entropy -= p * log2(p);
        }
    }
    return entropy;
}

int main(void) {
    printf("=== DÉBUT DU TEST ORACLE C : SHANNON ARCHTIME Q8.8 ===\n");

    // 1. Test Null et Vide
    assert(c2bt_calc_entropy_8_8(NULL, 0) == 0);
    assert(c2bt_calc_entropy_8_8(NULL, 100) == 0);
    uint8_t dummy[1] = {0};
    assert(c2bt_calc_entropy_8_8(dummy, 0) == 0);
    printf("[PASS] Null / Empty guards\n");

    // 2. Test Zéro Entropie (tampon uniforme 'A' * 256, 'A' * 1024, 'A' * 65536)
    size_t sizes[] = {1, 16, 64, 128, 256, 512, 1024, 4096, 65536};
    for (size_t s = 0; s < sizeof(sizes)/sizeof(sizes[0]); s++) {
        size_t n = sizes[s];
        uint8_t *buf = (uint8_t *)malloc(n);
        memset(buf, 0x42, n);
        uint32_t ent = c2bt_calc_entropy_8_8(buf, n);
        assert(ent == 0);
        free(buf);
    }
    printf("[PASS] Zero Entropy on uniform buffers (1..65536)\n");

    // 3. Test Max Entropie (256 octets distincts 0..255)
    uint8_t max_buf[256];
    for (int i = 0; i < 256; i++) max_buf[i] = (uint8_t)i;
    uint32_t max_ent = c2bt_calc_entropy_8_8(max_buf, 256);
    assert(max_ent == 2048); // 8.0 * 256 = 2048
    printf("[PASS] Max Entropy on 256 distinct bytes == 2048 (8.0 bits)\n");

    // 4. Test Grandes Tailles & Comparaison Flottante (N = 512, 1024, 4096, 65536, 1MB)
    size_t large_sizes[] = {512, 1024, 2048, 4096, 8192, 16384, 65536, 1048576};
    for (size_t s = 0; s < sizeof(large_sizes)/sizeof(large_sizes[0]); s++) {
        size_t n = large_sizes[s];
        uint8_t *buf = (uint8_t *)malloc(n);
        // Pseudo-random pseudo-crypto noise
        uint32_t state = 123456789 + (uint32_t)n;
        for (size_t i = 0; i < n; i++) {
            state = state * 1664525 + 1013904223;
            buf[i] = (uint8_t)(state >> 24);
        }
        
        double ref_float = reference_shannon_entropy(buf, n);
        uint32_t ent_q8 = c2bt_calc_entropy_8_8(buf, n);
        double float_from_q8 = (double)ent_q8 / 256.0;
        
        double diff = fabs(ref_float - float_from_q8);
        printf("Size %8zu : Ref Float = %.6f b/o | ARCHTIME Q8.8 = %.6f b/o (raw=%5u) | Diff = %.6f\n",
               n, ref_float, float_from_q8, ent_q8, diff);
        
        assert(diff < 0.05); // Erreur absolue < 0.05 bit par octet
        assert(ent_q8 > 1920); // Doit être qualifié comme haute entropie (> 7.5 b/o)
        free(buf);
    }
    printf("[PASS] Continuous Shannon accuracy on large buffers (512 to 1MB)\n");

    // 5. Test Texte ASCII / Markdown (Entropie moyenne ~3.5 à 4.8)
    const char *prose = "The quick brown fox jumps over the lazy dog. In computer science and information theory, "
                        "the entropy of a random variable is the average level of 'information', 'surprise', or "
                        "'uncertainty' inherent to the variable's possible outcomes. Given a discrete random variable X.";
    size_t prose_len = strlen(prose);
    double prose_ref = reference_shannon_entropy((const uint8_t *)prose, prose_len);
    uint32_t prose_q8 = c2bt_calc_entropy_8_8((const uint8_t *)prose, prose_len);
    double prose_float = (double)prose_q8 / 256.0;
    double prose_diff = fabs(prose_ref - prose_float);
    printf("Prose len %zu : Ref Float = %.6f b/o | ARCHTIME Q8.8 = %.6f b/o (raw=%5u) | Diff = %.6f\n",
           prose_len, prose_ref, prose_float, prose_q8, prose_diff);
    assert(prose_diff < 0.05);
    assert(prose_q8 < 1500); // Doit être < 5.8 b/o (pas d'alerte de chiffrement)
    printf("[PASS] Natural text / ASCII classification\n");

    printf("=== TOUS LES TESTS ORACLE C PASSENT AVEC SUCCÈS (100%%) ===\n");
    return 0;
}
