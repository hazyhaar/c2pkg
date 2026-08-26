/*
 * dot_product_l2.c - Noyau arithmétique C23 pour Produit Scalaire et Distance L2
 * Conçu spécifiquement pour la transpilation déterministe C-vers-Go 1.27 (c2simd).
 * 
 * Règles d'optimisation transpilation :
 * - Pointeurs const float * restrict stricts (zéro void*, zéro aliasing).
 * - Boucles à pas fixe multiples de 16 (64 octets / alignement ligne de cache).
 * - Accumulateurs scalaires locaux pour éliminer les écritures mémoire intermédiaires.
 */

#include <stddef.h>
#include <stdint.h>

/*
 * c2simd_dot_product_f32 - Produit scalaire f32 avec déroulement 4x (16 éléments/boucle)
 */
float c2simd_dot_product_f32(const float * restrict a, const float * restrict b, size_t count) {
    float sum0 = 0.0f;
    float sum1 = 0.0f;
    float sum2 = 0.0f;
    float sum3 = 0.0f;

    size_t i = 0;
    size_t limit = count & ~((size_t)15);

    for (; i < limit; i += 16) {
        sum0 += a[i + 0] * b[i + 0];
        sum0 += a[i + 1] * b[i + 1];
        sum0 += a[i + 2] * b[i + 2];
        sum0 += a[i + 3] * b[i + 3];

        sum1 += a[i + 4] * b[i + 4];
        sum1 += a[i + 5] * b[i + 5];
        sum1 += a[i + 6] * b[i + 6];
        sum1 += a[i + 7] * b[i + 7];

        sum2 += a[i + 8] * b[i + 8];
        sum2 += a[i + 9] * b[i + 9];
        sum2 += a[i + 10] * b[i + 10];
        sum2 += a[i + 11] * b[i + 11];

        sum3 += a[i + 12] * b[i + 12];
        sum3 += a[i + 13] * b[i + 13];
        sum3 += a[i + 14] * b[i + 14];
        sum3 += a[i + 15] * b[i + 15];
    }

    float total = (sum0 + sum1) + (sum2 + sum3);
    for (; i < count; i++) {
        total += a[i] * b[i];
    }
    return total;
}

/*
 * c2simd_l2_squared_f32 - Distance L2 carrée f32 directe (somme des (a_i - b_i)^2)
 */
float c2simd_l2_squared_f32(const float * restrict a, const float * restrict b, size_t count) {
    float sum0 = 0.0f;
    float sum1 = 0.0f;
    float sum2 = 0.0f;
    float sum3 = 0.0f;

    size_t i = 0;
    size_t limit = count & ~((size_t)15);

    for (; i < limit; i += 16) {
        float d0 = a[i + 0] - b[i + 0];
        float d1 = a[i + 1] - b[i + 1];
        float d2 = a[i + 2] - b[i + 2];
        float d3 = a[i + 3] - b[i + 3];
        sum0 += d0 * d0 + d1 * d1 + d2 * d2 + d3 * d3;

        float d4 = a[i + 4] - b[i + 4];
        float d5 = a[i + 5] - b[i + 5];
        float d6 = a[i + 6] - b[i + 6];
        float d7 = a[i + 7] - b[i + 7];
        sum1 += d4 * d4 + d5 * d5 + d6 * d6 + d7 * d7;

        float d8 = a[i + 8] - b[i + 8];
        float d9 = a[i + 9] - b[i + 9];
        float d10 = a[i + 10] - b[i + 10];
        float d11 = a[i + 11] - b[i + 11];
        sum2 += d8 * d8 + d9 * d9 + d10 * d10 + d11 * d11;

        float d12 = a[i + 12] - b[i + 12];
        float d13 = a[i + 13] - b[i + 13];
        float d14 = a[i + 14] - b[i + 14];
        float d15 = a[i + 15] - b[i + 15];
        sum3 += d12 * d12 + d13 * d13 + d14 * d14 + d15 * d15;
    }

    float total = (sum0 + sum1) + (sum2 + sum3);
    for (; i < count; i++) {
        float diff = a[i] - b[i];
        total += diff * diff;
    }
    return total;
}
