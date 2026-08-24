//go:build amd64

#include "textflag.h"

// func c2archsimd_blend_solid_photometric8_avx2_asm(in *uint32, ctx *C2archsimd_solid_blend_ctx_t, out *uint32)
TEXT ·c2archsimd_blend_solid_photometric8_avx2_asm(SB), NOSPLIT, $0-24
	MOVQ in+0(FP), DI
	MOVQ ctx+8(FP), DX
	MOVQ out+16(FP), SI

	// 1. Charger 8 pixels (256 bits)
	VMOVDQU (DI), Y0

	// 2. Extraire R, G, B, A
	MOVQ $·maskFF_data(SB), AX
	VMOVDQU (AX), Y15
	VPAND Y15, Y0, Y1           // Y1 = dr [8x u32]
	VPSRLD $8, Y0, Y2
	VPAND Y15, Y2, Y2           // Y2 = dg [8x u32]
	VPSRLD $16, Y0, Y3
	VPAND Y15, Y3, Y3           // Y3 = db [8x u32]
	VPSRLD $24, Y0, Y4           // Y4 = da [8x u32]

	// 3. Gather lut_srgb_to_lin (uint16)
	MOVQ 24(DX), AX              // AX = ptr lut_srgb_to_lin
	MOVQ $·maskFFFF_data(SB), BX
	VMOVDQU (BX), Y14

	VPCMPEQD Y5, Y5, Y5
	VPGATHERDD Y5, (AX)(Y1*2), Y6
	VPAND Y14, Y6, Y6            // Y6 = dr_lin

	VPCMPEQD Y5, Y5, Y5
	VPGATHERDD Y5, (AX)(Y2*2), Y7
	VPAND Y14, Y7, Y7            // Y7 = dg_lin

	VPCMPEQD Y5, Y5, Y5
	VPGATHERDD Y5, (AX)(Y3*2), Y8
	VPAND Y14, Y8, Y8            // Y8 = db_lin

	// 4. Charger coefficients de contexte
	VPBROADCASTD 4(DX), Y9       // Y9 = inv_a
	VPBROADCASTD 8(DX), Y10      // Y10 = sr_scaled
	VPBROADCASTD 12(DX), Y11     // Y11 = sg_scaled
	VPBROADCASTD 16(DX), Y12     // Y12 = sb_scaled
	VPBROADCASTD 0(DX), Y13      // Y13 = sa

	MOVQ $·c127_data(SB), BX
	VMOVDQU (BX), Y0             // Y0 = 127

	// tr = dr_lin * inv_a + sr_scaled + 127
	VPMULLD Y9, Y6, Y6
	VPADDD Y10, Y6, Y6
	VPADDD Y0, Y6, Y6

	// tg = dg_lin * inv_a + sg_scaled + 127
	VPMULLD Y9, Y7, Y7
	VPADDD Y11, Y7, Y7
	VPADDD Y0, Y7, Y7

	// tb = db_lin * inv_a + sb_scaled + 127
	VPMULLD Y9, Y8, Y8
	VPADDD Y12, Y8, Y8
	VPADDD Y0, Y8, Y8

	// ta = da * inv_a + 127
	VPMULLD Y9, Y4, Y4
	VPADDD Y0, Y4, Y4

	// 5. Division exacte vectorielle par 255 (c2_div255_epu32_avx2)
	MOVQ $·cMagic_data(SB), BX
	VMOVDQU (BX), Y10            // Y10 = magic 33686019

	// Division Y6 (tr) -> Y6
	VPMULUDQ Y10, Y6, Y0
	VPSRLQ $33, Y0, Y0
	VPSRLQ $32, Y6, Y1
	VPMULUDQ Y10, Y1, Y1
	VPSRLQ $33, Y1, Y1
	VPSLLQ $32, Y1, Y1
	VPOR Y1, Y0, Y6

	// Division Y7 (tg) -> Y7
	VPMULUDQ Y10, Y7, Y0
	VPSRLQ $33, Y0, Y0
	VPSRLQ $32, Y7, Y1
	VPMULUDQ Y10, Y1, Y1
	VPSRLQ $33, Y1, Y1
	VPSLLQ $32, Y1, Y1
	VPOR Y1, Y0, Y7

	// Division Y8 (tb) -> Y8
	VPMULUDQ Y10, Y8, Y0
	VPSRLQ $33, Y0, Y0
	VPSRLQ $32, Y8, Y1
	VPMULUDQ Y10, Y1, Y1
	VPSRLQ $33, Y1, Y1
	VPSLLQ $32, Y1, Y1
	VPOR Y1, Y0, Y8

	// Division Y4 (ta) -> Y4 (+ sa, clamp 255)
	VPMULUDQ Y10, Y4, Y0
	VPSRLQ $33, Y0, Y0
	VPSRLQ $32, Y4, Y1
	VPMULUDQ Y10, Y1, Y1
	VPSRLQ $33, Y1, Y1
	VPSLLQ $32, Y1, Y1
	VPOR Y1, Y0, Y4
	VPADDD Y13, Y4, Y4
	VPMINUD Y15, Y4, Y4          // clamp out_a <= 255

	// 6. Indices sRGB (>> 4 et clamp <= 4095)
	VPSRLD $4, Y6, Y6
	VPSRLD $4, Y7, Y7
	VPSRLD $4, Y8, Y8

	MOVQ $·c4095_data(SB), BX
	VMOVDQU (BX), Y0
	VPMINUD Y0, Y6, Y6
	VPMINUD Y0, Y7, Y7
	VPMINUD Y0, Y8, Y8

	// 7. Gather lut_lin_to_srgb (uint8)
	MOVQ 48(DX), BX              // BX = ptr lut_lin_to_srgb

	VPCMPEQD Y5, Y5, Y5
	VPGATHERDD Y5, (BX)(Y6*1), Y0
	VPAND Y15, Y0, Y6

	VPCMPEQD Y5, Y5, Y5
	VPGATHERDD Y5, (BX)(Y7*1), Y0
	VPAND Y15, Y0, Y7

	VPCMPEQD Y5, Y5, Y5
	VPGATHERDD Y5, (BX)(Y8*1), Y0
	VPAND Y15, Y0, Y8

	// 8. Recombinaison RGBA : Y6 | (Y7 << 8) | (Y8 << 16) | (Y4 << 24)
	VPSLLD $8, Y7, Y7
	VPOR Y7, Y6, Y6
	VPSLLD $16, Y8, Y8
	VPOR Y8, Y6, Y6
	VPSLLD $24, Y4, Y4
	VPOR Y4, Y6, Y6

	// 9. Écriture mémoire 8 pixels
	VMOVDQU Y6, (SI)
	VZEROUPPER
	RET

// func c2archsimd_blend_solid_span_photometric_avx2_asm(dst *uint32, ctx *C2archsimd_solid_blend_ctx_t, count int)
TEXT ·c2archsimd_blend_solid_span_photometric_avx2_asm(SB), NOSPLIT, $0-24
	MOVQ dst+0(FP), DI
	MOVQ ctx+8(FP), DX
	MOVQ count+16(FP), CX

	CMPQ CX, $8
	JL done

	MOVQ $·maskFF_data(SB), AX
	VMOVDQU (AX), Y15
	MOVQ $·maskFFFF_data(SB), AX
	VMOVDQU (AX), Y14
	MOVQ $·cMagic_data(SB), AX
	VMOVDQU (AX), Y13
	MOVQ $·c127_data(SB), AX
	VMOVDQU (AX), Y12
	MOVQ $·c4095_data(SB), AX
	VMOVDQU (AX), Y11

	MOVQ 24(DX), R8              // R8 = ptr lut_srgb_to_lin
	MOVQ 48(DX), R9              // R9 = ptr lut_lin_to_srgb

	VPBROADCASTD 4(DX), Y10      // Y10 = inv_a
	VPBROADCASTD 8(DX), Y9       // Y9 = sr_scaled
	VPBROADCASTD 12(DX), Y8      // Y8 = sg_scaled
	VPBROADCASTD 16(DX), Y7      // Y7 = sb_scaled
	VPBROADCASTD 0(DX), Y6       // Y6 = sa

span_loop:
	VMOVDQU (DI), Y0

	VPAND Y15, Y0, Y1           // dr
	VPSRLD $8, Y0, Y2
	VPAND Y15, Y2, Y2           // dg
	VPSRLD $16, Y0, Y3
	VPAND Y15, Y3, Y3           // db
	VPSRLD $24, Y0, Y4           // da

	VPCMPEQD Y5, Y5, Y5
	VPGATHERDD Y5, (R8)(Y1*2), Y0
	VPAND Y14, Y0, Y1

	VPCMPEQD Y5, Y5, Y5
	VPGATHERDD Y5, (R8)(Y2*2), Y0
	VPAND Y14, Y0, Y2

	VPCMPEQD Y5, Y5, Y5
	VPGATHERDD Y5, (R8)(Y3*2), Y0
	VPAND Y14, Y0, Y3

	VPMULLD Y10, Y1, Y1
	VPADDD Y9, Y1, Y1
	VPADDD Y12, Y1, Y1

	VPMULLD Y10, Y2, Y2
	VPADDD Y8, Y2, Y2
	VPADDD Y12, Y2, Y2

	VPMULLD Y10, Y3, Y3
	VPADDD Y7, Y3, Y3
	VPADDD Y12, Y3, Y3

	VPMULLD Y10, Y4, Y4
	VPADDD Y12, Y4, Y4

	// Div255 Y1
	VPMULUDQ Y13, Y1, Y0
	VPSRLQ $33, Y0, Y0
	VPSRLQ $32, Y1, Y5
	VPMULUDQ Y13, Y5, Y5
	VPSRLQ $33, Y5, Y5
	VPSLLQ $32, Y5, Y5
	VPOR Y5, Y0, Y1

	// Div255 Y2
	VPMULUDQ Y13, Y2, Y0
	VPSRLQ $33, Y0, Y0
	VPSRLQ $32, Y2, Y5
	VPMULUDQ Y13, Y5, Y5
	VPSRLQ $33, Y5, Y5
	VPSLLQ $32, Y5, Y5
	VPOR Y5, Y0, Y2

	// Div255 Y3
	VPMULUDQ Y13, Y3, Y0
	VPSRLQ $33, Y0, Y0
	VPSRLQ $32, Y3, Y5
	VPMULUDQ Y13, Y5, Y5
	VPSRLQ $33, Y5, Y5
	VPSLLQ $32, Y5, Y5
	VPOR Y5, Y0, Y3

	// Div255 Y4
	VPMULUDQ Y13, Y4, Y0
	VPSRLQ $33, Y0, Y0
	VPSRLQ $32, Y4, Y5
	VPMULUDQ Y13, Y5, Y5
	VPSRLQ $33, Y5, Y5
	VPSLLQ $32, Y5, Y5
	VPOR Y5, Y0, Y4
	VPADDD Y6, Y4, Y4
	VPMINUD Y15, Y4, Y4

	// Clamp + sRGB gather
	VPSRLD $4, Y1, Y1
	VPSRLD $4, Y2, Y2
	VPSRLD $4, Y3, Y3
	VPMINUD Y11, Y1, Y1
	VPMINUD Y11, Y2, Y2
	VPMINUD Y11, Y3, Y3

	VPCMPEQD Y5, Y5, Y5
	VPGATHERDD Y5, (R9)(Y1*1), Y0
	VPAND Y15, Y0, Y1

	VPCMPEQD Y5, Y5, Y5
	VPGATHERDD Y5, (R9)(Y2*1), Y0
	VPAND Y15, Y0, Y2

	VPCMPEQD Y5, Y5, Y5
	VPGATHERDD Y5, (R9)(Y3*1), Y0
	VPAND Y15, Y0, Y3

	VPSLLD $8, Y2, Y2
	VPOR Y2, Y1, Y1
	VPSLLD $16, Y3, Y3
	VPOR Y3, Y1, Y1
	VPSLLD $24, Y4, Y4
	VPOR Y4, Y1, Y1

	VMOVDQU Y1, (DI)

	ADDQ $32, DI
	SUBQ $8, CX
	CMPQ CX, $8
	JGE span_loop

done:
	VZEROUPPER
	RET

GLOBL ·maskFF_data(SB), RODATA, $32
DATA ·maskFF_data+0(SB)/8, $0x000000FF000000FF
DATA ·maskFF_data+8(SB)/8, $0x000000FF000000FF
DATA ·maskFF_data+16(SB)/8, $0x000000FF000000FF
DATA ·maskFF_data+24(SB)/8, $0x000000FF000000FF

GLOBL ·maskFFFF_data(SB), RODATA, $32
DATA ·maskFFFF_data+0(SB)/8, $0x0000FFFF0000FFFF
DATA ·maskFFFF_data+8(SB)/8, $0x0000FFFF0000FFFF
DATA ·maskFFFF_data+16(SB)/8, $0x0000FFFF0000FFFF
DATA ·maskFFFF_data+24(SB)/8, $0x0000FFFF0000FFFF

GLOBL ·c127_data(SB), RODATA, $32
DATA ·c127_data+0(SB)/8, $0x0000007F0000007F
DATA ·c127_data+8(SB)/8, $0x0000007F0000007F
DATA ·c127_data+16(SB)/8, $0x0000007F0000007F
DATA ·c127_data+24(SB)/8, $0x0000007F0000007F

GLOBL ·cMagic_data(SB), RODATA, $32
DATA ·cMagic_data+0(SB)/8, $0x0202020302020203
DATA ·cMagic_data+8(SB)/8, $0x0202020302020203
DATA ·cMagic_data+16(SB)/8, $0x0202020302020203
DATA ·cMagic_data+24(SB)/8, $0x0202020302020203

GLOBL ·c4095_data(SB), RODATA, $32
DATA ·c4095_data+0(SB)/8, $0x00000FFF00000FFF
DATA ·c4095_data+8(SB)/8, $0x00000FFF00000FFF
DATA ·c4095_data+16(SB)/8, $0x00000FFF00000FFF
DATA ·c4095_data+24(SB)/8, $0x00000FFF00000FFF
