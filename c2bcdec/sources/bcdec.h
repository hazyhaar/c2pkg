/* SPDX-License-Identifier: MIT */
#ifndef BCDEC_H
#define BCDEC_H

#include <stdint.h>
#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

#define BCDEC_BC1_BLOCK_SIZE 8
#define BCDEC_BC3_BLOCK_SIZE 16
#define BCDEC_BC4_BLOCK_SIZE 8
#define BCDEC_BC5_BLOCK_SIZE 16

/* Décode un bloc 4x4 BC1 (DXT1) en 16 pixels RGBA (4 octets par pixel, format pitch RGBA) */
void bcdec_bc1(const void *compressedBlock, void *decompressedRGBA, int destinationPitch);

/* Décode un bloc 4x4 BC3 (DXT5) en 16 pixels RGBA */
void bcdec_bc3(const void *compressedBlock, void *decompressedRGBA, int destinationPitch);

/* Décode un bloc 4x4 BC4 (canal unique / R) */
void bcdec_bc4(const void *compressedBlock, void *decompressedR, int destinationPitch);

/* Décode un bloc 4x4 BC5 (deux canaux / RG / Normales) */
void bcdec_bc5(const void *compressedBlock, void *decompressedRG, int destinationPitch);

#ifdef __cplusplus
}
#endif

#endif /* BCDEC_H */
