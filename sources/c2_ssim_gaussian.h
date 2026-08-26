#ifndef C2_SSIM_GAUSSIAN_H
#define C2_SSIM_GAUSSIAN_H

#include <stdint.h>
#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

#define C2_SSIM_KERNEL_SIZE 11
#define C2_SSIM_KERNEL_RADIUS 5
#define C2_SSIM_SCALE_Q16 65536

/* Constantes SSIM standard (K1 = 0.01, K2 = 0.03, L = 255) */
/* C1 = (0.01 * 255)^2 = 6.5025  -> en Q16 : 6.5025 * 65536 = 426154 */
/* C2 = (0.03 * 255)^2 = 58.5225 -> en Q16 : 58.5225 * 65536 = 3835208 */
#define C2_SSIM_C1_Q16 426154LL
#define C2_SSIM_C2_Q16 3835208LL

/* Filtre 1D horizontal gaussien séparable 11x11 en virgule fixe Q16 (src 8-bit -> tmp 16-bit) */
void c2_gaussian_blur_1d_h(const uint8_t *src, uint16_t *tmp, int width, int height, int stride);

/* Filtre 1D vertical gaussien séparable 11x11 en virgule fixe Q16 (tmp 16-bit -> dst 8-bit) */
void c2_gaussian_blur_1d_v(const uint16_t *tmp, uint8_t *dst, int width, int height, int stride);

/* Filtre 2D gaussien complet 11x11 séparable (utilise un buffer temporaire 16-bit de taille width * height) */
void c2_gaussian_blur_2d(const uint8_t *src, uint8_t *dst, uint16_t *tmp, int width, int height, int stride);

/* Convolution 1D horizontale et verticale pour buffers 32-bit (utilisés pour les moments X^2, Y^2, XY) */
void c2_gaussian_blur_1d_h_u32(const uint32_t *src, uint32_t *tmp, int width, int height, int stride);
void c2_gaussian_blur_1d_v_u32(const uint32_t *tmp, uint32_t *dst, int width, int height, int stride);
void c2_gaussian_blur_2d_u32(const uint32_t *src, uint32_t *dst, uint32_t *tmp, int width, int height, int stride);

/* Calcul instantané du score SSIM global en millionièmes (0..1000000, 1000000 = 1.000000) */
int64_t c2_ssim_compute_milli(const uint8_t *img1, const uint8_t *img2, int width, int height, int stride);

/* Calcul instantané du score SSIM global en Q16 (0..65536, 65536 = 1.000000) */
int64_t c2_ssim_compute_q16(const uint8_t *img1, const uint8_t *img2, int width, int height, int stride);

#if defined(__x86_64__) || defined(_M_X64)
/* Versions accélérées AVX2 */
void c2_gaussian_blur_1d_h_avx2(const uint8_t *src, uint16_t *tmp, int width, int height, int stride);
void c2_gaussian_blur_1d_v_avx2(const uint16_t *tmp, uint8_t *dst, int width, int height, int stride);
int64_t c2_ssim_compute_milli_avx2(const uint8_t *img1, const uint8_t *img2, int width, int height, int stride);
#endif

#ifdef __cplusplus
}
#endif

#endif /* C2_SSIM_GAUSSIAN_H */
