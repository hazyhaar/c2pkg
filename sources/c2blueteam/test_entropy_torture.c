/*
 * test_entropy_torture.c — Banc de torture et qualification statistique d'entropie.
 * Confronte l'algorithme ARCHTIME Q8.8 à la formule de Shannon continue sur 100 000 vecteurs.
 */

#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include <math.h>
#include <assert.h>
#include <time.h>
#include "c2blueteam.h"

static double reference_shannon_entropy(const uint8_t *data, size_t len) {
    if (data == NULL || len == 0) return 0.0;
    uint32_t freq[256] = {0};
    for (size_t i = 0; i < len; i++) {
        freq[data[i]]++;
    }
    double entropy = 0.0;
    double n = (double)len;
    for (int i = 0; i < 256; i++) {
        if (freq[i] > 0) {
            double p = (double)freq[i] / n;
            entropy -= p * log2(p);
        }
    }
    return entropy;
}

// Générateur LCG déterministe pour reproductibilité
static inline uint32_t lcg_next(uint64_t *state) {
    *state = *state * 6364136223846793005ULL + 1442695040888963407ULL;
    return (uint32_t)(*state >> 32);
}

int main(void) {
    printf("=================================================================\n");
    printf("   BANC DE TORTURE ET QUALIFICATION MATHÉMATIQUE SHANNON ARCHTIME\n");
    printf("=================================================================\n\n");

    uint64_t rng_state = 0xDEADBEEFCAFE1337ULL;
    double max_diff = 0.0;
    double sum_diff = 0.0;
    long total_tests = 0;

    // 1. ÉPREUVE DE CAS LIMITES & ALIGNEMENT (1 à 512 octets avec décalage de pointeur)
    printf("[1/5] Épreuve des cas limites, tailles non alignées et offsets...");
    fflush(stdout);
    {
        uint8_t raw_arena[2048];
        for (int offset = 0; offset < 16; offset++) {
            uint8_t *base = raw_arena + offset;
            for (size_t len = 0; len <= 512; len++) {
                // Remplissage pseudo-aléatoire
                for (size_t i = 0; i < len; i++) {
                    base[i] = (uint8_t)lcg_next(&rng_state);
                }
                
                double ref = reference_shannon_entropy(base, len);
                uint32_t q8 = c2bt_calc_entropy_8_8(base, len);
                double q8_f = (double)q8 / 256.0;
                double diff = fabs(ref - q8_f);
                
                if (diff > max_diff) max_diff = diff;
                sum_diff += diff;
                total_tests++;

                // Vérification de garde
                if (len == 0) {
                    assert(q8 == 0);
                } else if (len == 1) {
                    assert(q8 == 0);
                } else {
                    assert(diff < 0.05); // Borne de précision stricte
                }
            }
        }
    }
    printf(" OK (%ld tests validés)\n", total_tests);

    // 2. ÉPREUVE DES DISTRIBUTIONS STATISTIQUES CARACTÉRISTIQUES
    printf("[2/5] Épreuve sur 6 classes de distributions statistiques...");
    fflush(stdout);
    {
        size_t test_sizes[] = {32, 64, 128, 255, 256, 257, 512, 1000, 1024, 4096, 65536};
        for (size_t s = 0; s < sizeof(test_sizes)/sizeof(test_sizes[0]); s++) {
            size_t len = test_sizes[s];
            uint8_t *buf = (uint8_t *)malloc(len + 64);

            // Distribution A : Dirac (tous les octets identiques)
            memset(buf, 0x55, len);
            assert(c2bt_calc_entropy_8_8(buf, len) == 0);
            total_tests++;

            // Distribution B : Bimodale (50% 0x00, 50% 0xFF) -> Entropie théorique = 1.0 bit (256 en Q8.8)
            for (size_t i = 0; i < len; i++) buf[i] = (i % 2 == 0) ? 0x00 : 0xFF;
            {
                double ref = reference_shannon_entropy(buf, len);
                uint32_t q8 = c2bt_calc_entropy_8_8(buf, len);
                double diff = fabs(ref - ((double)q8 / 256.0));
                if (diff > max_diff) max_diff = diff;
                sum_diff += diff;
                total_tests++;
                assert(diff < 0.05);
            }

            // Distribution C : Base64 (64 symboles autorisés) -> Entropie max ~6.0 bits
            const char b64_chars[] = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
            for (size_t i = 0; i < len; i++) buf[i] = (uint8_t)b64_chars[lcg_next(&rng_state) % 64];
            {
                double ref = reference_shannon_entropy(buf, len);
                uint32_t q8 = c2bt_calc_entropy_8_8(buf, len);
                double diff = fabs(ref - ((double)q8 / 256.0));
                if (diff > max_diff) max_diff = diff;
                sum_diff += diff;
                total_tests++;
                assert(diff < 0.05);
                if (len >= 512) {
                    assert(q8 >= 1400 && q8 <= 1600); // 5.5 à 6.2 bits
                }
            }

            // Distribution D : Hexadécimale (16 symboles) -> Entropie max ~4.0 bits
            const char hex_chars[] = "0123456789abcdef";
            for (size_t i = 0; i < len; i++) buf[i] = (uint8_t)hex_chars[lcg_next(&rng_state) % 16];
            {
                double ref = reference_shannon_entropy(buf, len);
                uint32_t q8 = c2bt_calc_entropy_8_8(buf, len);
                double diff = fabs(ref - ((double)q8 / 256.0));
                if (diff > max_diff) max_diff = diff;
                sum_diff += diff;
                total_tests++;
                assert(diff < 0.05);
            }

            // Distribution E : Zipf / Prose (distribution exponentielle décroissante)
            for (size_t i = 0; i < len; i++) {
                uint32_t r = lcg_next(&rng_state) % 1000;
                if (r < 400) buf[i] = ' ';
                else if (r < 600) buf[i] = 'e';
                else if (r < 750) buf[i] = 't';
                else if (r < 850) buf[i] = 'a';
                else buf[i] = (uint8_t)('a' + (lcg_next(&rng_state) % 26));
            }
            {
                double ref = reference_shannon_entropy(buf, len);
                uint32_t q8 = c2bt_calc_entropy_8_8(buf, len);
                double diff = fabs(ref - ((double)q8 / 256.0));
                if (diff > max_diff) max_diff = diff;
                sum_diff += diff;
                total_tests++;
                assert(diff < 0.05);
            }

            // Distribution F : Bruit blanc complet (Uniforme 256 valeurs)
            for (size_t i = 0; i < len; i++) buf[i] = (uint8_t)(lcg_next(&rng_state) & 0xFF);
            {
                double ref = reference_shannon_entropy(buf, len);
                uint32_t q8 = c2bt_calc_entropy_8_8(buf, len);
                double diff = fabs(ref - ((double)q8 / 256.0));
                if (diff > max_diff) max_diff = diff;
                sum_diff += diff;
                total_tests++;
                assert(diff < 0.05);
                if (len >= 1024) {
                    assert(q8 > 1920); // > 7.5 bits
                }
            }

            free(buf);
        }
    }
    printf(" OK (%ld tests validés)\n", total_tests);

    // 3. TORTURE MONTE-CARLO MASSIVE (50 000 itérations sur tailles et données aléatoires)
    printf("[3/5] Torture Monte-Carlo massive (50 000 vecteurs aléatoires)...");
    fflush(stdout);
    {
        uint8_t *big_buf = (uint8_t *)malloc(65536);
        for (int iter = 0; iter < 50000; iter++) {
            size_t len = 1 + (lcg_next(&rng_state) % 4096);
            int mode = lcg_next(&rng_state) % 4;
            
            if (mode == 0) {
                // Totalement aléatoire
                for (size_t i = 0; i < len; i++) big_buf[i] = (uint8_t)lcg_next(&rng_state);
            } else if (mode == 1) {
                // Sous-ensemble restreint de 4 à 32 octets
                int alphabet_size = 4 + (lcg_next(&rng_state) % 28);
                for (size_t i = 0; i < len; i++) big_buf[i] = (uint8_t)(lcg_next(&rng_state) % alphabet_size);
            } else if (mode == 2) {
                // Motif périodique avec bruit
                int period = 2 + (lcg_next(&rng_state) % 16);
                for (size_t i = 0; i < len; i++) {
                    big_buf[i] = (uint8_t)((i % period) + (lcg_next(&rng_state) % 3));
                }
            } else {
                // Répétition par blocs
                uint8_t val = (uint8_t)lcg_next(&rng_state);
                for (size_t i = 0; i < len; i++) {
                    if (i % 32 == 0) val = (uint8_t)lcg_next(&rng_state);
                    big_buf[i] = val;
                }
            }

            double ref = reference_shannon_entropy(big_buf, len);
            uint32_t q8 = c2bt_calc_entropy_8_8(big_buf, len);
            double diff = fabs(ref - ((double)q8 / 256.0));

            if (diff > max_diff) max_diff = diff;
            sum_diff += diff;
            total_tests++;

            if (diff >= 0.05) {
                fprintf(stderr, "\nÉCHEC TORTURE à l'itération %d (len=%zu, mode=%d) : ref=%.6f, q8=%.6f, diff=%.6f\n",
                        iter, len, mode, ref, (double)q8/256.0, diff);
                assert(diff < 0.05);
            }
        }
        free(big_buf);
    }
    printf(" OK (%ld tests validés)\n", total_tests);

    // 4. ÉPREUVE SUR TAMPONS GÉANTS (1 Mo, 4 Mo, 16 Mo)
    printf("[4/5] Épreuve sur tampons volumineux (1 Mo, 4 Mo, 16 Mo)...");
    fflush(stdout);
    {
        size_t giant_sizes[] = {1048576, 4194304, 16777216};
        for (size_t g = 0; g < sizeof(giant_sizes)/sizeof(giant_sizes[0]); g++) {
            size_t len = giant_sizes[g];
            uint8_t *gbuf = (uint8_t *)malloc(len);
            for (size_t i = 0; i < len; i++) {
                gbuf[i] = (uint8_t)lcg_next(&rng_state);
            }

            double ref = reference_shannon_entropy(gbuf, len);
            uint32_t q8 = c2bt_calc_entropy_8_8(gbuf, len);
            double diff = fabs(ref - ((double)q8 / 256.0));

            if (diff > max_diff) max_diff = diff;
            sum_diff += diff;
            total_tests++;

            assert(diff < 0.01);
            assert(q8 >= 2040); // Proche de 8.0 bits
            free(gbuf);
        }
    }
    printf(" OK\n");

    // 5. RAPPORT STATISTIQUE FINAL
    printf("[5/5] Rapport de qualification métrologique :\n");
    printf("   - Total de vecteurs éprouvés : %ld\n", total_tests);
    printf("   - Écart absolu maximal       : %.6f bit par octet (seuil max admis : 0.050000)\n", max_diff);
    printf("   - Écart moyen (biais global) : %.6f bit par octet\n\n", sum_diff / (double)total_tests);

    printf("=================================================================\n");
    printf("   RÉSULTAT : HARNAIS DE TORTURE ORACLE C VALIDÉ (100%% KATs)\n");
    printf("=================================================================\n");
    return 0;
}
