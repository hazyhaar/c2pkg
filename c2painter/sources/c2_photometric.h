#ifndef C2_PHOTOMETRIC_H
#define C2_PHOTOMETRIC_H

#include <stdint.h>
#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

extern const uint16_t c2_srgb_to_linear_lut[256];
extern const uint8_t c2_linear_to_srgb_lut[4096];

static inline uint32_t c2_div255_u32(uint32_t t) {
    return (t + 127) / 255;
}

uint32_t c2_blend_photometric(uint32_t dst, uint32_t src);
uint32_t c2_blend_pixel_cov_photometric(uint32_t dst, uint32_t src, uint32_t cov);
void c2_blend_span_photometric(uint32_t *dst, const uint32_t *src, int n);
void c2_blend_solid_span_photometric(uint32_t *dst, uint32_t color, int n);

#ifdef __cplusplus
}
#endif

#endif
