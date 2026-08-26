#include "c2archsimd.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

static uint64_t xorshift_state = 0x853c49e6748fea9bULL;

static inline uint64_t xorshift64(void) {
    uint64_t x = xorshift_state;
    x ^= x >> 12;
    x ^= x << 25;
    x ^= x >> 27;
    xorshift_state = x;
    return x * 0x2545F4914F6CDD1DULL;
}

static uint64_t fold_mix(uint64_t acc, uint64_t val) {
    return acc ^ (val + 0x9E3779B97F4A7C15ULL + (acc << 6) + (acc >> 2));
}

static uint16_t oracle_srgb_to_lin_lut[256];
static uint8_t  oracle_lin_to_srgb_lut[4096];

static void init_oracle_luts(void) {
    for (int i = 0; i < 256; i++) {
        double c = i / 255.0;
        double lin = (c <= 0.04045) ? (c / 12.92) : pow((c + 0.055) / 1.055, 2.4);
        oracle_srgb_to_lin_lut[i] = (uint16_t)(lin * 65535.0 + 0.5);
    }
    for (int i = 0; i < 4096; i++) {
        double lin = (i << 4) / 65535.0;
        double srgb = (lin <= 0.0031308) ? (lin * 12.92) : (1.055 * pow(lin, 1.0 / 2.4) - 0.055);
        int val = (int)(srgb * 255.0 + 0.5);
        if (val < 0) val = 0;
        if (val > 255) val = 255;
        oracle_lin_to_srgb_lut[i] = (uint8_t)val;
    }
}

int main(void) {
    printf("=== C2ARCHSIMD C ORACLE HARNESS (AVX2 vs SCALAR) ===\n");
    init_oracle_luts();

    uint64_t fold_lut16 = 0;
    uint64_t fold_lut256 = 0;
    uint64_t fold_hex = 0;
    uint64_t fold_vint = 0;
    uint64_t fold_photo = 0;

    // 1. KAT Vectors Tests
    uint8_t kat_uuid_bin[16] = {
        0x01, 0x91, 0x7F, 0x8B, 0xC9, 0xA0, 0x71, 0x11,
        0x80, 0x00, 0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC
    };
    uint8_t hex_out_scal[32];
    uint8_t hex_out_avx2[32];

    c2archsimd_hex_encode16(kat_uuid_bin, hex_out_scal);
#if defined(__x86_64__) || defined(_M_X64)
    c2archsimd_hex_encode16_avx2(kat_uuid_bin, hex_out_avx2);
    if (memcmp(hex_out_scal, hex_out_avx2, 32) != 0) {
        printf("KAT Hex Encode AVX2 mismatch! [FAIL]\n");
        return 1;
    }
#endif
    if (memcmp(hex_out_scal, "01917f8bc9a071118000123456789abc", 32) != 0) {
        printf("KAT Hex Encode value mismatch! [FAIL]\n");
        return 1;
    }
    printf("KAT Hex Encode: PASS\n");

    uint8_t decoded_bin[16];
    uint32_t st = c2archsimd_hex_decode32(hex_out_scal, decoded_bin);
    if (st != C2ARCHSIMD_OK || memcmp(decoded_bin, kat_uuid_bin, 16) != 0) {
        printf("KAT Hex Decode mismatch! [FAIL]\n");
        return 1;
    }
    printf("KAT Hex Decode: PASS\n");

    // Invalid hex decode test (out must be 0-filled)
    uint8_t invalid_hex[32];
    memcpy(invalid_hex, "01917f8bc9a071118000123456789abg", 32);
    uint8_t invalid_out[16];
    memset(invalid_out, 0xAA, 16);
    if (c2archsimd_hex_decode32(invalid_hex, invalid_out) != C2ARCHSIMD_INVALID_HEX) {
        printf("KAT Invalid Hex Decode failed to return error! [FAIL]\n");
        return 1;
    }
    for (int k = 0; k < 16; k++) {
        if (invalid_out[k] != 0) {
            printf("KAT Invalid Hex Decode failed to sanitize out! [FAIL]\n");
            return 1;
        }
    }
    printf("KAT Invalid Hex Decode with 0-fill: PASS\n");

    // 2. Torture Campaign : 1,000,000 Random Vectors
    printf("--- [TEST] Torture Test 1,000,000 Random Vectors (Full Domain) ---\n");

    c2archsimd_table16_t table16;
    c2archsimd_table256_t table256;
    for (int k = 0; k < 16; k++) {
        table16.b[k] = (uint8_t)(k * 17);
    }
    for (int k = 0; k < 256; k++) {
        table256.b[k] = (uint8_t)(k * 31 + 7);
    }

    uint8_t in_buf[32];
    uint8_t out_scal[32];
    uint8_t out_avx2[32];

    for (int iter = 0; iter < 1000000; iter++) {
        for (int b = 0; b < 32; b += 8) {
            uint64_t r = xorshift64();
            memcpy(&in_buf[b], &r, 8);
        }

        // 2.1 LUT16
        c2archsimd_lut16(in_buf, &table16, out_scal);
#if defined(__x86_64__) || defined(_M_X64)
        c2archsimd_lut16_avx2(in_buf, &table16, out_avx2);
        if (memcmp(out_scal, out_avx2, 32) != 0) {
            printf("Torture LUT16 mismatch at iter %d! [FAIL]\n", iter);
            return 1;
        }
#endif
        for (int k = 0; k < 32; k += 8) {
            uint64_t val;
            memcpy(&val, &out_scal[k], 8);
            fold_lut16 = fold_mix(fold_lut16, val);
        }

        // 2.2 LUT256 Directe
        c2archsimd_lut256(in_buf, &table256, out_scal);
        for (int k = 0; k < 32; k += 8) {
            uint64_t val;
            memcpy(&val, &out_scal[k], 8);
            fold_lut256 = fold_mix(fold_lut256, val);
        }

        // 2.3 Hex Encode 16B -> 32B
        c2archsimd_hex_encode16(in_buf, out_scal);
#if defined(__x86_64__) || defined(_M_X64)
        c2archsimd_hex_encode16_avx2(in_buf, out_avx2);
        if (memcmp(out_scal, out_avx2, 32) != 0) {
            printf("Torture Hex Encode mismatch at iter %d! [FAIL]\n", iter);
            return 1;
        }
#endif
        for (int k = 0; k < 32; k += 8) {
            uint64_t val;
            memcpy(&val, &out_scal[k], 8);
            fold_hex = fold_mix(fold_hex, val);
        }

        // 2.4 VarInt Lens
        c2archsimd_vint_lens32(in_buf, out_scal);
#if defined(__x86_64__) || defined(_M_X64)
        c2archsimd_vint_lens32_avx2(in_buf, out_avx2);
        if (memcmp(out_scal, out_avx2, 32) != 0) {
            printf("Torture Vint Lens mismatch at iter %d! [FAIL]\n", iter);
            return 1;
        }
#endif
        for (int k = 0; k < 32; k += 8) {
            uint64_t val;
            memcpy(&val, &out_scal[k], 8);
            fold_vint = fold_mix(fold_vint, val);
        }

        // 2.5 Photometric Solid Blend 8-voies (32B) sur tables CEI 61966-2-1 réelles
        {
            c2archsimd_solid_blend_ctx_t blend_ctx = {
                .sa = (uint32_t)(in_buf[0]),
                .inv_a = 255 - (uint32_t)(in_buf[0]),
                .sr_scaled = (uint32_t)(oracle_srgb_to_lin_lut[in_buf[1]]) * (uint32_t)(in_buf[0]),
                .sg_scaled = (uint32_t)(oracle_srgb_to_lin_lut[in_buf[2]]) * (uint32_t)(in_buf[0]),
                .sb_scaled = (uint32_t)(oracle_srgb_to_lin_lut[in_buf[3]]) * (uint32_t)(in_buf[0]),
                .lut_srgb_to_lin = oracle_srgb_to_lin_lut,
                .lut_lin_to_srgb = oracle_lin_to_srgb_lut
            };

            uint32_t photo_in[8], photo_out_scal[8], photo_out_avx2[8];
            memcpy(photo_in, in_buf, 32);

            c2archsimd_blend_solid_photometric8(photo_in, &blend_ctx, photo_out_scal);
#if defined(__x86_64__) || defined(_M_X64)
            c2archsimd_blend_solid_photometric8_avx2(photo_in, &blend_ctx, photo_out_avx2);
            if (memcmp(photo_out_scal, photo_out_avx2, 32) != 0) {
                printf("Torture Photometric Blend AVX2 mismatch at iter %d! [FAIL]\n", iter);
                return 1;
            }
#endif
            for (int k = 0; k < 8; k++) {
                fold_photo = fold_mix(fold_photo, photo_out_scal[k]);
            }
        }
    }

    printf("1,000,000 vectors tested: 100%% bit-exact [PASS]\n");
    printf("FOLD lut16  0x%016llX\n", (unsigned long long)fold_lut16);
    printf("FOLD lut256 0x%016llX\n", (unsigned long long)fold_lut256);
    printf("FOLD hex    0x%016llX\n", (unsigned long long)fold_hex);
    printf("FOLD vint   0x%016llX\n", (unsigned long long)fold_vint);
    printf("FOLD photo  0x%016llX\n", (unsigned long long)fold_photo);
    printf("=== ALL C ORACLE TESTS PASSED ===\n");
    return 0;
}
