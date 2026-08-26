#ifndef C2ARCHSIMD_H
#define C2ARCHSIMD_H

#include <stdint.h>
#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

#define C2ARCHSIMD_OK 0
#define C2ARCHSIMD_INVALID_HEX 1

typedef struct {
    uint8_t b[16];
} c2archsimd_table16_t;

typedef struct {
    uint8_t b[256];
} c2archsimd_table256_t;

// 1. LUT16 Scalaire (blocs de 32 octets)
void c2archsimd_lut16(const uint8_t in[32], const c2archsimd_table16_t *t16, uint8_t out[32]);
#if defined(__x86_64__) || defined(_M_X64)
void c2archsimd_lut16_avx2(const uint8_t in[32], const c2archsimd_table16_t *t16, uint8_t out[32]);
#endif

// 2. LUT256 Directe Scalaire (blocs de 32 octets)
void c2archsimd_lut256(const uint8_t in[32], const c2archsimd_table256_t *t256, uint8_t out[32]);
#if defined(__x86_64__) || defined(_M_X64)
void c2archsimd_lut256_avx2(const uint8_t in[32], const c2archsimd_table16_t *t_lo, const c2archsimd_table16_t *t_hi, uint8_t out[32]);
#endif

// 3. Hex Encode 16B -> 32B Scalaire
void c2archsimd_hex_encode16(const uint8_t in[16], uint8_t out[32]);
#if defined(__x86_64__) || defined(_M_X64)
void c2archsimd_hex_encode16_avx2(const uint8_t in[16], uint8_t out[32]);
#endif

// 4. Hex Decode 32B -> 16B Atomique (écriture directe, 0-fill sur erreur)
uint32_t c2archsimd_hex_decode32(const uint8_t in[32], uint8_t out[16]);

// 5. VarInt Batch Lengths (32 longueurs)
void c2archsimd_vint_lens32(const uint8_t in[32], uint8_t out[32]);
#if defined(__x86_64__) || defined(_M_X64)
void c2archsimd_vint_lens32_avx2(const uint8_t in[32], uint8_t out[32]);
#endif

// 6. Photometric Solid Blend (8 pixels uint32 = 32 octets)
typedef struct {
    uint32_t sa;
    uint32_t inv_a;
    uint32_t sr_scaled;
    uint32_t sg_scaled;
    uint32_t sb_scaled;
    const uint16_t *lut_srgb_to_lin; // 256 uint16
    const uint8_t *lut_lin_to_srgb;   // 4096 uint8
} c2archsimd_solid_blend_ctx_t;

void c2archsimd_blend_solid_photometric8(const uint32_t in[8], const c2archsimd_solid_blend_ctx_t *ctx, uint32_t out[8]);
#if defined(__x86_64__) || defined(_M_X64)
void c2archsimd_blend_solid_photometric8_avx2(const uint32_t in[8], const c2archsimd_solid_blend_ctx_t *ctx, uint32_t out[8]);
#endif

#ifdef __cplusplus
}
#endif

#endif // C2ARCHSIMD_H
