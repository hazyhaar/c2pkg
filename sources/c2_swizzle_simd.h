#ifndef C2_SWIZZLE_SIMD_H
#define C2_SWIZZLE_SIMD_H

#include <stdint.h>
#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Swizzle unitaire sur mot 32-bit (swap R <-> B, conserve G et A) */
uint32_t c2_swizzle_pixel(uint32_t pixel);

/* Swizzle unitaire sur mot 64-bit (2 pixels 32-bit simultanés) */
uint64_t c2_swizzle_pair(uint64_t pair);

/* Swizzle de blocs fixes alignés */
void c2_swizzle_block4(const uint8_t in[16], uint8_t out[16]);
void c2_swizzle_block8(const uint8_t in[32], uint8_t out[32]);
void c2_swizzle_block16(const uint8_t in[64], uint8_t out[64]);

/* Transposition vectorielle déroulée RGBA <-> BGRA (out-of-place) */
void c2_swizzle_rgba_bgra(const uint8_t *src, uint8_t *dst, size_t num_pixels);

/* Transposition vectorielle déroulée RGBA <-> BGRA (in-place) */
void c2_swizzle_rgba_bgra_inplace(uint8_t *data, size_t num_pixels);

/* Conversions ARGB <-> RGBA / BGRA */
void c2_swizzle_argb_to_rgba(const uint8_t *src, uint8_t *dst, size_t num_pixels);
void c2_swizzle_rgba_to_argb(const uint8_t *src, uint8_t *dst, size_t num_pixels);
void c2_swizzle_argb_to_bgra(const uint8_t *src, uint8_t *dst, size_t num_pixels);

/* Transposition 2D avec pas de ligne (stride) */
void c2_swizzle_rgba_bgra_stride(const uint8_t *src, uint8_t *dst, int width, int height, int src_stride, int dst_stride);

#if defined(__x86_64__) || defined(_M_X64)
/* Implémentations vectorielles AVX2 saturant le bus mémoire */
void c2_swizzle_rgba_bgra_avx2(const uint8_t *src, uint8_t *dst, size_t num_pixels);
void c2_swizzle_rgba_bgra_inplace_avx2(uint8_t *data, size_t num_pixels);
void c2_swizzle_rgba_bgra_stream_avx2(const uint8_t *src, uint8_t *dst, size_t num_pixels);
#endif

#ifdef __cplusplus
}
#endif

#endif /* C2_SWIZZLE_SIMD_H */
