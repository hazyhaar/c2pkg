#ifndef C2_PHOTOMETRIC_H
#define C2_PHOTOMETRIC_H

#include <stdint.h>
#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

// Tables ARCHTIME précalculées
extern const uint16_t c2_srgb_to_linear_lut[256];
extern const uint8_t c2_linear_to_srgb_lut[4096];

// Primitives de fusion photométrique linéaire C99
uint32_t c2_blend_photometric(uint32_t dst, uint32_t src);
void c2_blend_span_photometric(uint32_t *dst, const uint32_t *src, int n);

#ifdef __cplusplus
}
#endif

#endif // C2_PHOTOMETRIC_H
