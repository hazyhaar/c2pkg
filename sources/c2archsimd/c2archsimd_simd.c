#include "c2archsimd.h"

#if defined(__x86_64__) || defined(_M_X64)
#include <immintrin.h>

// 1. LUT16 AVX2 (32 lookups en 1 instruction vpshufb)
void c2archsimd_lut16_avx2(const uint8_t in[32], const c2archsimd_table16_t *t16, uint8_t out[32]) {
    __m128i t128 = _mm_loadu_si128((const __m128i *)t16->b);
    __m256i v_table = _mm256_broadcastsi128_si256(t128);
    __m256i v_in = _mm256_loadu_si256((const __m256i *)in);
    __m256i v_mask = _mm256_set1_epi8(0x0F);
    __m256i v_idx = _mm256_and_si256(v_in, v_mask);
    __m256i v_res = _mm256_shuffle_epi8(v_table, v_idx);
    _mm256_storeu_si256((__m256i *)out, v_res);
}

// 2. LUT256 Nibble AVX2 (32 classifications 8-bit en 3 instructions)
void c2archsimd_lut256_avx2(const uint8_t in[32], const c2archsimd_table16_t *t_lo, const c2archsimd_table16_t *t_hi, uint8_t out[32]) {
    __m256i v_tlo = _mm256_broadcastsi128_si256(_mm_loadu_si128((const __m128i *)t_lo->b));
    __m256i v_thi = _mm256_broadcastsi128_si256(_mm_loadu_si128((const __m128i *)t_hi->b));
    __m256i v_in = _mm256_loadu_si256((const __m256i *)in);
    __m256i v_mask = _mm256_set1_epi8(0x0F);

    __m256i v_lo_idx = _mm256_and_si256(v_in, v_mask);
    __m256i v_hi_idx = _mm256_and_si256(_mm256_srli_epi16(v_in, 4), v_mask);

    __m256i v_lo_res = _mm256_shuffle_epi8(v_tlo, v_lo_idx);
    __m256i v_hi_res = _mm256_shuffle_epi8(v_thi, v_hi_idx);

    __m256i v_res = _mm256_and_si256(v_lo_res, v_hi_res);
    _mm256_storeu_si256((__m256i *)out, v_res);
}

// 3. Hex Encode 16B -> 32B AVX2 / SSSE3
void c2archsimd_hex_encode16_avx2(const uint8_t in[16], uint8_t out[32]) {
    __m128i v_in = _mm_loadu_si128((const __m128i *)in);
    __m128i v_mask = _mm_set1_epi8(0x0F);
    __m128i v_hi = _mm_and_si128(_mm_srli_epi16(v_in, 4), v_mask);
    __m128i v_lo = _mm_and_si128(v_in, v_mask);

    __m128i v_unpck_lo = _mm_unpacklo_epi8(v_hi, v_lo);
    __m128i v_unpck_hi = _mm_unpackhi_epi8(v_hi, v_lo);

    __m128i v_hex = _mm_loadu_si128((const __m128i *)"0123456789abcdef");
    __m128i res0 = _mm_shuffle_epi8(v_hex, v_unpck_lo);
    __m128i res1 = _mm_shuffle_epi8(v_hex, v_unpck_hi);

    _mm_storeu_si128((__m128i *)out, res0);
    _mm_storeu_si128((__m128i *)(out + 16), res1);
}

// 4. VarInt Batch Lengths AVX2 (32 longueurs résolues en 1 cycle VPSHUFB)
void c2archsimd_vint_lens32_avx2(const uint8_t in[32], uint8_t out[32]) {
    __m128i t128 = _mm_setr_epi8(1, 2, 4, 8, 1, 2, 4, 8, 1, 2, 4, 8, 1, 2, 4, 8);
    __m256i v_table = _mm256_broadcastsi128_si256(t128);
    __m256i v_in = _mm256_loadu_si256((const __m256i *)in);
    __m256i v_idx = _mm256_and_si256(_mm256_srli_epi16(v_in, 6), _mm256_set1_epi8(0x03));
    __m256i v_res = _mm256_shuffle_epi8(v_table, v_idx);
    _mm256_storeu_si256((__m256i *)out, v_res);
}

// Helper vectoriel : division par 255 exacte pour 8 entiers 32-bit dans [0, 33423000]
// Multiplie par 33686019 (0x02020203) et décale à droite de 33 bits.
static inline __m256i c2_div255_epu32_avx2(__m256i v_t) {
    __m256i v_magic = _mm256_set1_epi64x(33686019ULL);
    __m256i even_mul = _mm256_mul_epu32(v_t, v_magic);
    __m256i even_res = _mm256_srli_epi64(even_mul, 33);

    __m256i odd_in = _mm256_srli_epi64(v_t, 32);
    __m256i odd_mul = _mm256_mul_epu32(odd_in, v_magic);
    __m256i odd_res = _mm256_srli_epi64(odd_mul, 33);

    return _mm256_or_si256(even_res, _mm256_slli_epi64(odd_res, 32));
}

// 5. Photometric Solid Blend AVX2 (8 pixels = 32 octets - 100% vectorisé en registres YMM)
void c2archsimd_blend_solid_photometric8_avx2(const uint32_t in[8], const c2archsimd_solid_blend_ctx_t *ctx, uint32_t out[8]) {
    __m256i v_in = _mm256_loadu_si256((const __m256i *)in);
    __m256i v_mask8 = _mm256_set1_epi32(0xFF);

    // Extraction 8 voies des 4 canaux
    __m256i v_dr = _mm256_and_si256(v_in, v_mask8);
    __m256i v_dg = _mm256_and_si256(_mm256_srli_epi32(v_in, 8), v_mask8);
    __m256i v_db = _mm256_and_si256(_mm256_srli_epi32(v_in, 16), v_mask8);
    __m256i v_da = _mm256_and_si256(_mm256_srli_epi32(v_in, 24), v_mask8);

    // Gathers 8 voies sur la LUT 256 (uint16 -> échelle 2 octets)
    __m256i v_mask16 = _mm256_set1_epi32(0xFFFF);
    __m256i v_dr_lin = _mm256_and_si256(_mm256_i32gather_epi32((const int*)ctx->lut_srgb_to_lin, v_dr, 2), v_mask16);
    __m256i v_dg_lin = _mm256_and_si256(_mm256_i32gather_epi32((const int*)ctx->lut_srgb_to_lin, v_dg, 2), v_mask16);
    __m256i v_db_lin = _mm256_and_si256(_mm256_i32gather_epi32((const int*)ctx->lut_srgb_to_lin, v_db, 2), v_mask16);

    __m256i v_sa = _mm256_set1_epi32(ctx->sa);
    __m256i v_inv_a = _mm256_set1_epi32(ctx->inv_a);
    __m256i v_sr_scaled = _mm256_set1_epi32(ctx->sr_scaled);
    __m256i v_sg_scaled = _mm256_set1_epi32(ctx->sg_scaled);
    __m256i v_sb_scaled = _mm256_set1_epi32(ctx->sb_scaled);
    __m256i v_127 = _mm256_set1_epi32(127);

    // Interpolation linéaire : t = (scaled + lin * inv_a + 127)
    __m256i v_tr = _mm256_add_epi32(_mm256_add_epi32(_mm256_mullo_epi32(v_dr_lin, v_inv_a), v_sr_scaled), v_127);
    __m256i v_tg = _mm256_add_epi32(_mm256_add_epi32(_mm256_mullo_epi32(v_dg_lin, v_inv_a), v_sg_scaled), v_127);
    __m256i v_tb = _mm256_add_epi32(_mm256_add_epi32(_mm256_mullo_epi32(v_db_lin, v_inv_a), v_sb_scaled), v_127);
    __m256i v_ta = _mm256_add_epi32(_mm256_mullo_epi32(v_da, v_inv_a), v_127);

    // Division vectorielle par 255
    __m256i v_out_r_lin = _mm256_min_epu32(c2_div255_epu32_avx2(v_tr), _mm256_set1_epi32(65535));
    __m256i v_out_g_lin = _mm256_min_epu32(c2_div255_epu32_avx2(v_tg), _mm256_set1_epi32(65535));
    __m256i v_out_b_lin = _mm256_min_epu32(c2_div255_epu32_avx2(v_tb), _mm256_set1_epi32(65535));
    __m256i v_out_a     = _mm256_min_epu32(_mm256_add_epi32(v_sa, c2_div255_epu32_avx2(v_ta)), _mm256_set1_epi32(255));

    // Indices pour la table de réencodage sRGB (lin >> 4)
    __m256i v_idx_r = _mm256_srli_epi32(v_out_r_lin, 4);
    __m256i v_idx_g = _mm256_srli_epi32(v_out_g_lin, 4);
    __m256i v_idx_b = _mm256_srli_epi32(v_out_b_lin, 4);

    // Gathers 8 voies sur la LUT 4096 (uint8 -> échelle 1 octet)
    __m256i v_out_r = _mm256_and_si256(_mm256_i32gather_epi32((const int*)ctx->lut_lin_to_srgb, v_idx_r, 1), v_mask8);
    __m256i v_out_g = _mm256_and_si256(_mm256_i32gather_epi32((const int*)ctx->lut_lin_to_srgb, v_idx_g, 1), v_mask8);
    __m256i v_out_b = _mm256_and_si256(_mm256_i32gather_epi32((const int*)ctx->lut_lin_to_srgb, v_idx_b, 1), v_mask8);

    // Reconditionnement RGBA 8 voies dans un registre 256-bit
    __m256i v_out = _mm256_or_si256(v_out_r,
                    _mm256_or_si256(_mm256_slli_epi32(v_out_g, 8),
                    _mm256_or_si256(_mm256_slli_epi32(v_out_b, 16),
                                    _mm256_slli_epi32(v_out_a, 24))));

    // Store vectoriel direct de 8 pixels
    _mm256_storeu_si256((__m256i *)out, v_out);
}

#endif
