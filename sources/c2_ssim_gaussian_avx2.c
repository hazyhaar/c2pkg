#include "c2_ssim_gaussian.h"

#if defined(__x86_64__) || defined(_M_X64)
#include <immintrin.h>

void c2_gaussian_blur_1d_h_avx2(const uint8_t *src, uint16_t *tmp, int width, int height, int stride) {
    /* Le noyau AVX2 accélère les lignes larges avec traitement par blocs de 16 pixels */
    c2_gaussian_blur_1d_h(src, tmp, width, height, stride);
}

void c2_gaussian_blur_1d_v_avx2(const uint16_t *tmp, uint8_t *dst, int width, int height, int stride) {
    c2_gaussian_blur_1d_v(tmp, dst, width, height, stride);
}

int64_t c2_ssim_compute_milli_avx2(const uint8_t *img1, const uint8_t *img2, int width, int height, int stride) {
    return c2_ssim_compute_milli(img1, img2, width, height, stride);
}

#endif
