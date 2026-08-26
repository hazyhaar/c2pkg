#include <stdint.h>
#include <stddef.h>

/* SimSIMD L2 squared distance float32 kernel */
float simsimd_l2sq_f32(const float *a, const float *b, uint64_t n) {
    float sum = 0.0f;
    uint64_t i;
    for (i = 0; i < n; i++) {
        float diff = a[i] - b[i];
        sum += diff * diff;
    }
    return sum;
}
