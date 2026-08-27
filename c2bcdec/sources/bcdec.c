/* SPDX-License-Identifier: MIT */
#include "bcdec.h"
#include <string.h>

static inline uint8_t bcdec_expand_5(uint8_t v) {
    return (v << 3) | (v >> 2);
}

static inline uint8_t bcdec_expand_6(uint8_t v) {
    return (v << 2) | (v >> 4);
}

static inline void bcdec_decode_rgb565(uint16_t c, uint8_t *rgb) {
    rgb[0] = bcdec_expand_5((c >> 11) & 0x1F);
    rgb[1] = bcdec_expand_6((c >> 5) & 0x3F);
    rgb[2] = bcdec_expand_5(c & 0x1F);
}

void bcdec_bc1(const void *compressedBlock, void *decompressedRGBA, int destinationPitch) {
    const uint8_t *src = (const uint8_t *)compressedBlock;
    uint8_t *dst = (uint8_t *)decompressedRGBA;

    uint16_t c0 = (uint16_t)(src[0] | (src[1] << 8));
    uint16_t c1 = (uint16_t)(src[2] | (src[3] << 8));
    uint32_t indices = (uint32_t)(src[4] | (src[5] << 8) | (src[6] << 16) | (src[7] << 24));

    uint8_t pal[4][4];
    bcdec_decode_rgb565(c0, pal[0]);
    pal[0][3] = 0xFF;
    bcdec_decode_rgb565(c1, pal[1]);
    pal[1][3] = 0xFF;

    if (c0 > c1) {
        pal[2][0] = (uint8_t)((2 * pal[0][0] + pal[1][0] + 1) / 3);
        pal[2][1] = (uint8_t)((2 * pal[0][1] + pal[1][1] + 1) / 3);
        pal[2][2] = (uint8_t)((2 * pal[0][2] + pal[1][2] + 1) / 3);
        pal[2][3] = 0xFF;

        pal[3][0] = (uint8_t)((pal[0][0] + 2 * pal[1][0] + 1) / 3);
        pal[3][1] = (uint8_t)((pal[0][1] + 2 * pal[1][1] + 1) / 3);
        pal[3][2] = (uint8_t)((pal[0][2] + 2 * pal[1][2] + 1) / 3);
        pal[3][3] = 0xFF;
    } else {
        pal[2][0] = (uint8_t)((pal[0][0] + pal[1][0]) / 2);
        pal[2][1] = (uint8_t)((pal[0][1] + pal[1][1]) / 2);
        pal[2][2] = (uint8_t)((pal[0][2] + pal[1][2]) / 2);
        pal[2][3] = 0xFF;

        pal[3][0] = 0;
        pal[3][1] = 0;
        pal[3][2] = 0;
        pal[3][3] = 0;
    }

    for (int y = 0; y < 4; y++) {
        uint8_t *row = dst + y * destinationPitch;
        for (int x = 0; x < 4; x++) {
            uint8_t idx = (indices >> ((y * 4 + x) * 2)) & 0x03;
            row[x * 4 + 0] = pal[idx][0];
            row[x * 4 + 1] = pal[idx][1];
            row[x * 4 + 2] = pal[idx][2];
            row[x * 4 + 3] = pal[idx][3];
        }
    }
}

void bcdec_bc3(const void *compressedBlock, void *decompressedRGBA, int destinationPitch) {
    const uint8_t *src = (const uint8_t *)compressedBlock;
    uint8_t *dst = (uint8_t *)decompressedRGBA;

    // Décodage de la partie couleur BC1 (octets 8 à 15)
    bcdec_bc1(src + 8, dst, destinationPitch);

    // Décodage de la table d'alpha (octets 0 à 7)
    uint8_t a0 = src[0];
    uint8_t a1 = src[1];
    uint64_t aIndices = 0;
    for (int i = 0; i < 6; i++) {
        aIndices |= ((uint64_t)src[2 + i]) << (i * 8);
    }

    uint8_t aPal[8];
    aPal[0] = a0;
    aPal[1] = a1;

    if (a0 > a1) {
        for (int i = 1; i <= 6; i++) {
            aPal[1 + i] = (uint8_t)(((7 - i) * a0 + i * a1 + 3) / 7);
        }
    } else {
        for (int i = 1; i <= 4; i++) {
            aPal[1 + i] = (uint8_t)(((5 - i) * a0 + i * a1 + 2) / 5);
        }
        aPal[6] = 0x00;
        aPal[7] = 0xFF;
    }

    for (int y = 0; y < 4; y++) {
        uint8_t *row = dst + y * destinationPitch;
        for (int x = 0; x < 4; x++) {
            uint8_t aIdx = (aIndices >> ((y * 4 + x) * 3)) & 0x07;
            row[x * 4 + 3] = aPal[aIdx];
        }
    }
}

void bcdec_bc4(const void *compressedBlock, void *decompressedR, int destinationPitch) {
    const uint8_t *src = (const uint8_t *)compressedBlock;
    uint8_t *dst = (uint8_t *)decompressedR;

    uint8_t r0 = src[0];
    uint8_t r1 = src[1];
    uint64_t rIndices = 0;
    for (int i = 0; i < 6; i++) {
        rIndices |= ((uint64_t)src[2 + i]) << (i * 8);
    }

    uint8_t rPal[8];
    rPal[0] = r0;
    rPal[1] = r1;

    if (r0 > r1) {
        for (int i = 1; i <= 6; i++) {
            rPal[1 + i] = (uint8_t)(((7 - i) * r0 + i * r1 + 3) / 7);
        }
    } else {
        for (int i = 1; i <= 4; i++) {
            rPal[1 + i] = (uint8_t)(((5 - i) * r0 + i * r1 + 2) / 5);
        }
        rPal[6] = 0x00;
        rPal[7] = 0xFF;
    }

    for (int y = 0; y < 4; y++) {
        uint8_t *row = dst + y * destinationPitch;
        for (int x = 0; x < 4; x++) {
            uint8_t rIdx = (rIndices >> ((y * 4 + x) * 3)) & 0x07;
            row[x] = rPal[rIdx];
        }
    }
}

void bcdec_bc5(const void *compressedBlock, void *decompressedRG, int destinationPitch) {
    const uint8_t *src = (const uint8_t *)compressedBlock;
    uint8_t *dst = (uint8_t *)decompressedRG;

    uint8_t tempR[16];
    uint8_t tempG[16];

    bcdec_bc4(src, tempR, 4);
    bcdec_bc4(src + 8, tempG, 4);

    for (int y = 0; y < 4; y++) {
        uint8_t *row = dst + y * destinationPitch;
        for (int x = 0; x < 4; x++) {
            row[x * 2 + 0] = tempR[y * 4 + x];
            row[x * 2 + 1] = tempG[y * 4 + x];
        }
    }
}
