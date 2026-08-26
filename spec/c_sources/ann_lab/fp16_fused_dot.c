/*
 * fp16_fused_dot.c - Décodage fp16 et accumulation L2 / Produit scalaire fusionné
 * Optims transpilation Go 1.27 :
 * - Lecture directe d'octets Little-Endian sans structures complexes.
 * - Table de correspondance float32 précalculée lut65536 pour la parité archtime.
 * - Accumulation L2 directe sans matérialisation intermédiaire.
 */

#include <stddef.h>
#include <stdint.h>

/*
 * c2simd_l2_fused_fp16_f32 - Accumulation L2 entre requête f32 et vecteur fp16
 */
float c2simd_l2_fused_fp16_f32(const float * restrict q, const uint16_t * restrict row_fp16, const float * restrict lut65536, size_t count) {
    float sum0 = 0.0f;
    float sum1 = 0.0f;
    float sum2 = 0.0f;
    float sum3 = 0.0f;

    size_t i = 0;
    size_t limit = count & ~((size_t)3);

    for (; i < limit; i += 4) {
        float f0 = lut65536[row_fp16[i + 0]];
        float f1 = lut65536[row_fp16[i + 1]];
        float f2 = lut65536[row_fp16[i + 2]];
        float f3 = lut65536[row_fp16[i + 3]];

        float d0 = q[i + 0] - f0;
        float d1 = q[i + 1] - f1;
        float d2 = q[i + 2] - f2;
        float d3 = q[i + 3] - f3;

        sum0 += d0 * d0;
        sum1 += d1 * d1;
        sum2 += d2 * d2;
        sum3 += d3 * d3;
    }

    float total = (sum0 + sum1) + (sum2 + sum3);
    for (; i < count; i++) {
        float f = lut65536[row_fp16[i]];
        float diff = q[i] - f;
        total += diff * diff;
    }
    return total;
}
