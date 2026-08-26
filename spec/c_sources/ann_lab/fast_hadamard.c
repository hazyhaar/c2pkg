/*
 * fast_hadamard.c - Noyau C23 FHT (Fast Walsh-Hadamard Transform)
 * Optims transpilation Go 1.27 :
 * - In-place transformation avec pointeur restrict float * restrict v.
 * - Boucles d'étages papillons déroulées et structurées à pas puissance de 2.
 * - Séparation stricte du scratch pour éliminer toute réallocation dynamique.
 */

#include <stddef.h>
#include <stdint.h>

/*
 * c2simd_fht_in_place_f32 - Fast Walsh-Hadamard Transform in-place f32
 * count doit être une puissance de 2 (ex: 512, 1024).
 */
void c2simd_fht_in_place_f32(float * restrict v, size_t count) {
    if (count == 0 || (count & (count - 1)) != 0) {
        return;
    }

    size_t h = 1;
    while (h < count) {
        size_t step = h * 2;
        for (size_t i = 0; i < count; i += step) {
            for (size_t j = i; j < i + h; j++) {
                float x = v[j];
                float y = v[j + h];
                v[j]     = x + y;
                v[j + h] = x - y;
            }
        }
        h <<= 1;
    }
}
