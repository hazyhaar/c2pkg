/* SPDX-License-Identifier: Apache-2.0 OR MIT */
#ifndef C2RASTER_3D_H
#define C2RASTER_3D_H

#include <stdint.h>
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    float x, y, z, inv_w;
    float u, v;
    float nx, ny, nz;
    uint32_t color;
} c2_vertex3d_t;

typedef struct {
    c2_vertex3d_t v0;
    c2_vertex3d_t v1;
    c2_vertex3d_t v2;
} c2_triangle3d_t;

/*
 * c2_rasterize_triangle3d rastérise un triangle 3D avec interpolation barycentrique
 * sub-pixel perspective-correcte et test de profondeur Z-Buffer 32-bit flottant.
 */
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
);

/*
 * c2_clear_buffers réinitialise les tampons couleur et profondeur (Z = 1.0f).
 */
void c2_clear_buffers(
    uint32_t *color_buffer,
    float *depth_buffer,
    int width,
    int height,
    int stride,
    uint32_t clear_color,
    float clear_depth
);

#ifdef __cplusplus
}
#endif

#endif /* C2RASTER_3D_H */
