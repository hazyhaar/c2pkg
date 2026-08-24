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

    // 2. Preuve mathématique d'exactitude bit-exacte de c2_div255_u32 sur tout le domaine [0, 65535*255]
    for (uint32_t t = 0; t <= 65535 * 255; t++) {
        uint32_t exact = (t + 127) / 255;
        uint32_t fast = c2_div255_u32(t);
        if (exact != fast) {
            fprintf(stderr, "DIV255 MISMATCH at t=%u : exact=%u fast=%u\n", t, exact, fast);
            assert(exact == fast);
        }
    }

    // 3. Test blend_cov : blanc opaque avec couverture 128
    uint32_t white_opaque = 0xFFFFFFFF;
    uint32_t blended_cov = c2_blend_pixel_cov_photometric(black, white_opaque, 128);
    assert((blended_cov & 0xFF) == 188);

    // 4. Test blend_solid_span : validation bit-exacte sur 16 longueurs {1, 2, 3, 7, 8, 9, 15, 16, 31, 32, 63, 64, 127, 128, 255, 1024}
    // avec offsets non alignés (dst[1:]) et couleurs variées (opaque, transparent, teinté)
    int lengths[] = {1, 2, 3, 7, 8, 9, 15, 16, 31, 32, 63, 64, 127, 128, 255, 1024};
    uint32_t test_colors[] = {
        0x80FFFFFF, // 50% blanc
        0xFF123456, // Opaque
        0x00123456, // Transparent
        0x9955AA33, // Teinté alpha 153
        0x33AABBCC  // Teinté alpha 51
    };

    for (int li = 0; li < 16; li++) {
        int n = lengths[li];
        for (int ci = 0; ci < 5; ci++) {
            uint32_t col = test_colors[ci];
            uint32_t *raw_buf = (uint32_t*)malloc((n + 2) * sizeof(uint32_t));

            // Test aligné et test décalé (offset +1)
            for (int offset = 0; offset <= 1; offset++) {
                uint32_t *dst_span = raw_buf + offset;
                for (int i = 0; i < n; i++) {
                    dst_span[i] = 0xFF000000 | ((i * 17) & 0xFF) | (((i * 31) & 0xFF) << 8) | (((i * 47) & 0xFF) << 16);
                }

                uint32_t *ref_span = (uint32_t*)malloc(n * sizeof(uint32_t));
                for (int i = 0; i < n; i++) {
                    ref_span[i] = c2_blend_photometric(dst_span[i], col);
                }

                c2_blend_solid_span_photometric(dst_span, col, n);

                for (int i = 0; i < n; i++) {
                    if (dst_span[i] != ref_span[i]) {
                        fprintf(stderr, "SPAN MISMATCH at len=%d col=0x%08X offset=%d idx=%d : span=0x%08X ref=0x%08X\n",
                                n, col, offset, i, dst_span[i], ref_span[i]);
                        assert(dst_span[i] == ref_span[i]);
                    }
                }
                free(ref_span);
            }
            free(raw_buf);
        }
    }

    printf("ORACLE C PHOTOMETRIC PASS: R=%d G=%d B=%d A=%d (50%% luminance exacte, 16M div255 validés, 16 longueurs bit-exactes)\n", r, g, b, a);
    return 0;
}
