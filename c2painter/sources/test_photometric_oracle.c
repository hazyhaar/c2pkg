#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <assert.h>
#include "c2_photometric.h"

int main() {
    // 1. Test nominal : 50% blanc sur noir
    uint32_t black = 0xFF000000;
    uint32_t white_half = 0x80FFFFFF; // Alpha 128, R=255, G=255, B=255
    uint32_t blended = c2_blend_photometric(black, white_half);

    uint8_t r = blended & 0xFF;
    uint8_t g = (blended >> 8) & 0xFF;
    uint8_t b = (blended >> 16) & 0xFF;
    uint8_t a = (blended >> 24) & 0xFF;

    assert(a == 255);
    // 50% d'énergie lumineuse dans l'espace sRGB correspond à 188 bit-exact
    assert(r == 188);
    assert(g == 188);
    assert(b == 188);

    // 3. Test blend_solid_span : validation bit-exacte sur toutes les longueurs (1, 7, 8, 15, 16, 64, 127)
    int lengths[] = {1, 7, 8, 15, 16, 64, 127};
    for (int li = 0; li < 7; li++) {
        int n = lengths[li];
        uint32_t *dst_span = (uint32_t*)malloc(n * sizeof(uint32_t));
        for (int i = 0; i < n; i++) {
            dst_span[i] = black;
        }
        c2_blend_solid_span_photometric(dst_span, white_half, n);
        for (int i = 0; i < n; i++) {
            assert((dst_span[i] & 0xFF) == 188);
            assert(((dst_span[i] >> 8) & 0xFF) == 188);
            assert(((dst_span[i] >> 16) & 0xFF) == 188);
            assert(((dst_span[i] >> 24) & 0xFF) == 255);
            // Vérification stricte contre le résultat scalaire
            assert(dst_span[i] == c2_blend_photometric(black, white_half));
        }
        free(dst_span);
    }

    printf("ORACLE C PHOTOMETRIC PASS: R=%d G=%d B=%d A=%d (50%% luminance exacte)\n", r, g, b, a);
    return 0;
}
