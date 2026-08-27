/* SPDX-License-Identifier: Apache-2.0 OR MIT */
#include "c2raster_3d.h"
#include <math.h>
#include <string.h>

#define C2_SUBPIXEL_BITS 4
#define C2_SUBPIXEL_SCALE (1 << C2_SUBPIXEL_BITS)
#define C2_SUBPIXEL_MASK (C2_SUBPIXEL_SCALE - 1)

static inline int c2_clamp_int(int v, int min_val, int max_val) {
    if (v < min_val) return min_val;
    if (v > max_val) return max_val;
    return v;
}

static inline float c2_min3f(float a, float b, float c) {
    float m = a < b ? a : b;
    return m < c ? m : c;
}

static inline float c2_max3f(float a, float b, float c) {
    float m = a > b ? a : b;
    return m > c ? m : c;
}

static inline int64_t c2_edge_func_fixed(int64_t ax, int64_t ay, int64_t bx, int64_t by, int64_t px, int64_t py) {
    return (bx - ax) * (py - ay) - (by - ay) * (px - ax);
}

static inline bool c2_is_top_left_fixed(int64_t ax, int64_t ay, int64_t bx, int64_t by) {
    int64_t dx = bx - ax;
    int64_t dy = by - ay;
    return (dy < 0) || (dy == 0 && dx < 0);
}

void c2_clear_buffers(
    uint32_t *color_buffer,
    float *depth_buffer,
    int width,
    int height,
    int stride,
    uint32_t clear_color,
    float clear_depth
) {
    if (!color_buffer && !depth_buffer) return;
    for (int y = 0; y < height; y++) {
        int row = y * stride;
        if (color_buffer) {
            for (int x = 0; x < width; x++) {
                color_buffer[row + x] = clear_color;
            }
        }
        if (depth_buffer) {
            for (int x = 0; x < width; x++) {
                depth_buffer[row + x] = clear_depth;
            }
        }
    }
}

void c2_rasterize_triangle3d(
    const c2_triangle3d_t *tri,
    uint32_t *color_buffer,
    float *depth_buffer,
    int width,
    int height,
    int stride,
    int tile_x0,
    int tile_y0,
    int tile_x1,
    int tile_y1
) {
    if (!tri || width <= 0 || height <= 0 || stride < width) return;

    // Rejet des triangles derrière ou sur le plan singulier
    if (tri->v0.inv_w <= 0.0f || tri->v1.inv_w <= 0.0f || tri->v2.inv_w <= 0.0f) {
        return;
    }

    // Conversion en coordonnées virgule fixe sous-pixel 28.4
    int64_t x0 = (int64_t)roundf(tri->v0.x * (float)C2_SUBPIXEL_SCALE);
    int64_t y0 = (int64_t)roundf(tri->v0.y * (float)C2_SUBPIXEL_SCALE);
    int64_t x1 = (int64_t)roundf(tri->v1.x * (float)C2_SUBPIXEL_SCALE);
    int64_t y1 = (int64_t)roundf(tri->v1.y * (float)C2_SUBPIXEL_SCALE);
    int64_t x2 = (int64_t)roundf(tri->v2.x * (float)C2_SUBPIXEL_SCALE);
    int64_t y2 = (int64_t)roundf(tri->v2.y * (float)C2_SUBPIXEL_SCALE);

    // Calcul de l'aire signée exacte en sous-pixels
    int64_t area_fixed = c2_edge_func_fixed(x0, y0, x1, y1, x2, y2);
    if (area_fixed <= 0) {
        // Back-face ou triangle dégénéré
        return;
    }

    // Détermination de l'AABB en pixels
    float min_x = c2_min3f(tri->v0.x, tri->v1.x, tri->v2.x);
    float max_x = c2_max3f(tri->v0.x, tri->v1.x, tri->v2.x);
    float min_y = c2_min3f(tri->v0.y, tri->v1.y, tri->v2.y);
    float max_y = c2_max3f(tri->v0.y, tri->v1.y, tri->v2.y);

    int min_ix = c2_clamp_int((int)floorf(min_x), tile_x0, tile_x1 - 1);
    int max_ix = c2_clamp_int((int)ceilf(max_x), tile_x0, tile_x1 - 1);
    int min_iy = c2_clamp_int((int)floorf(min_y), tile_y0, tile_y1 - 1);
    int max_iy = c2_clamp_int((int)ceilf(max_y), tile_y0, tile_y1 - 1);

    if (min_ix > max_ix || min_iy > max_iy) return;

    double inv_area = 1.0 / (double)area_fixed;

    // Détermination des biais top-left stricts (0 si top-left, -1 LSB sinon)
    int64_t bias0 = c2_is_top_left_fixed(x1, y1, x2, y2) ? 0 : -1;
    int64_t bias1 = c2_is_top_left_fixed(x2, y2, x0, y0) ? 0 : -1;
    int64_t bias2 = c2_is_top_left_fixed(x0, y0, x1, y1) ? 0 : -1;

    // Pente incrémentale Pineda
    int64_t a01 = y0 - y1, b01 = x1 - x0;
    int64_t a12 = y1 - y2, b12 = x2 - x1;
    int64_t a20 = y2 - y0, b20 = x0 - x2;

    // Décomposition et pré-multiplication des attributs par inv_w
    uint32_t c0 = tri->v0.color;
    uint32_t c1 = tri->v1.color;
    uint32_t c2 = tri->v2.color;

    float r0 = (float)(c0 & 0xff) * tri->v0.inv_w;
    float g0 = (float)((c0 >> 8) & 0xff) * tri->v0.inv_w;
    float b0 = (float)((c0 >> 16) & 0xff) * tri->v0.inv_w;
    float a0 = (float)((c0 >> 24) & 0xff) * tri->v0.inv_w;

    float r1 = (float)(c1 & 0xff) * tri->v1.inv_w;
    float g1 = (float)((c1 >> 8) & 0xff) * tri->v1.inv_w;
    float b1 = (float)((c1 >> 16) & 0xff) * tri->v1.inv_w;
    float a1 = (float)((c1 >> 24) & 0xff) * tri->v1.inv_w;

    float r2 = (float)(c2 & 0xff) * tri->v2.inv_w;
    float g2 = (float)((c2 >> 8) & 0xff) * tri->v2.inv_w;
    float b2 = (float)((c2 >> 16) & 0xff) * tri->v2.inv_w;
    float a2 = (float)((c2 >> 24) & 0xff) * tri->v2.inv_w;

    // Z fenêtre affine : z0, z1, z2 sans re-hyperbolisation
    float z0 = tri->v0.z;
    float z1 = tri->v1.z;
    float z2 = tri->v2.z;

    for (int y = min_iy; y <= max_iy; y++) {
        int64_t py = ((int64_t)y << C2_SUBPIXEL_BITS) + (C2_SUBPIXEL_SCALE / 2);
        int row = y * stride;

        for (int x = min_ix; x <= max_ix; x++) {
            int64_t px = ((int64_t)x << C2_SUBPIXEL_BITS) + (C2_SUBPIXEL_SCALE / 2);

            // Fonctions d'arêtes brutes pour calcul barycentrique exact
            int64_t w0 = c2_edge_func_fixed(x1, y1, x2, y2, px, py);
            int64_t w1 = c2_edge_func_fixed(x2, y2, x0, y0, px, py);
            int64_t w2 = c2_edge_func_fixed(x0, y0, x1, y1, px, py);

            // Test de couverture avec règle top-left stricte
            if ((w0 + bias0 >= 0) && (w1 + bias1 >= 0) && (w2 + bias2 >= 0)) {
                double l0 = (double)w0 * inv_area;
                double l1 = (double)w1 * inv_area;
                double l2 = (double)w2 * inv_area;

                float fl0 = (float)l0;
                float fl1 = (float)l1;
                float fl2 = (float)l2;

                // 1. Z-Buffer affine direct en espace fenêtre
                float z_interp = fl0 * z0 + fl1 * z1 + fl2 * z2;

                int idx = row + x;
                if (depth_buffer) {
                    if (z_interp >= depth_buffer[idx]) {
                        continue;
                    }
                    depth_buffer[idx] = z_interp;
                }

                // 2. Interpolation des attributs avec correction de perspective
                float inv_w = fl0 * tri->v0.inv_w + fl1 * tri->v1.inv_w + fl2 * tri->v2.inv_w;
                if (inv_w <= 0.0f) continue;
                float w_interp = 1.0f / inv_w;

                if (color_buffer) {
                    float r = (fl0 * r0 + fl1 * r1 + fl2 * r2) * w_interp;
                    float g = (fl0 * g0 + fl1 * g1 + fl2 * g2) * w_interp;
                    float b = (fl0 * b0 + fl1 * b1 + fl2 * b2) * w_interp;
                    float a = (fl0 * a0 + fl1 * a1 + fl2 * a2) * w_interp;

                    uint8_t ur = (uint8_t)(r < 0.0f ? 0 : (r > 255.0f ? 255 : (int)r));
                    uint8_t ug = (uint8_t)(g < 0.0f ? 0 : (g > 255.0f ? 255 : (int)g));
                    uint8_t ub = (uint8_t)(b < 0.0f ? 0 : (b > 255.0f ? 255 : (int)b));
                    uint8_t ua = (uint8_t)(a < 0.0f ? 0 : (a > 255.0f ? 255 : (int)a));

                    color_buffer[idx] = (uint32_t)ur | ((uint32_t)ug << 8) | ((uint32_t)ub << 16) | ((uint32_t)ua << 24);
                }
            }
        }
    }
}
