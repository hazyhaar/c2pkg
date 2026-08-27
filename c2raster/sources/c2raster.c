#include "c2raster.h"

static float c2_sqrtf(float x) {
    if (x <= 0.0f) {
        return 0.0f;
    }
    float y = x;
    if (y < 1.0f) {
        y = 1.0f;
    }
    for (int i = 0; i < 8; i++) {
        y = 0.5f * (y + x / y);
    }
    return y;
}

static int64_t c2_isqrt64(int64_t val) {
    if (val <= 0) {
        return 0;
    }
    int64_t res = 0;
    int64_t bit = (int64_t)1 << 62;
    while (bit > val) {
        bit = bit >> 2;
    }
    while (bit != 0) {
        if (val >= res + bit) {
            val = val - (res + bit);
            res = (res >> 1) + bit;
        } else {
            res = res >> 1;
        }
        bit = bit >> 2;
    }
    return res;
}

static uint32_t c2_blend_pixel(uint32_t dst, uint32_t src) {
    uint32_t sa = (src >> 24) & 0xFF;
    if (sa == 255) {
        return src;
    }
    if (sa == 0) {
        return dst;
    }
    uint32_t sr = src & 0xFF;
    uint32_t sg = (src >> 8) & 0xFF;
    uint32_t sb = (src >> 16) & 0xFF;

    uint32_t dr = dst & 0xFF;
    uint32_t dg = (dst >> 8) & 0xFF;
    uint32_t db = (dst >> 16) & 0xFF;
    uint32_t da = (dst >> 24) & 0xFF;

    uint32_t inv_a = 255 - sa;
    uint32_t out_r = sr + (dr * inv_a + 127) / 255;
    uint32_t out_g = sg + (dg * inv_a + 127) / 255;
    uint32_t out_b = sb + (db * inv_a + 127) / 255;
    uint32_t out_a = sa + (da * inv_a + 127) / 255;

    if (out_r > 255) out_r = 255;
    if (out_g > 255) out_g = 255;
    if (out_b > 255) out_b = 255;
    if (out_a > 255) out_a = 255;

    return out_r | (out_g << 8) | (out_b << 16) | (out_a << 24);
}

static uint32_t c2_blend_pixel_cov(uint32_t dst, uint32_t src, uint32_t cov) {
    if (cov == 0) {
        return dst;
    }
    if (cov >= 255) {
        return c2_blend_pixel(dst, src);
    }
    uint32_t sr = ((src & 0xFF) * cov + 127) / 255;
    uint32_t sg = (((src >> 8) & 0xFF) * cov + 127) / 255;
    uint32_t sb = (((src >> 16) & 0xFF) * cov + 127) / 255;
    uint32_t sa = (((src >> 24) & 0xFF) * cov + 127) / 255;
    uint32_t scaled_src = sr | (sg << 8) | (sb << 16) | (sa << 24);
    return c2_blend_pixel(dst, scaled_src);
}

static void fill_solid_rect(uint32_t *framebuffer, int stride, int x, int y, int w, int h, uint32_t color) {
    if (w <= 0 || h <= 0) {
        return;
    }
    int x0 = x;
    if (x0 < 0) {
        x0 = 0;
    }
    int y0 = y;
    if (y0 < 0) {
        y0 = 0;
    }
    int x1 = x + w;
    int y1 = y + h;
    if (x1 > stride) {
        x1 = stride;
    }

    if (x0 >= x1 || y0 >= y1) {
        return;
    }

    uint32_t sa = (color >> 24) & 0xFF;
    if (sa == 255) {
        for (int cy = y0; cy < y1; cy++) {
            int row = cy * stride;
            for (int cx = x0; cx < x1; cx++) {
                framebuffer[row + cx] = color;
            }
        }
    } else if (sa > 0) {
        for (int cy = y0; cy < y1; cy++) {
            int row = cy * stride;
            for (int cx = x0; cx < x1; cx++) {
                framebuffer[row + cx] = c2_blend_pixel(framebuffer[row + cx], color);
            }
        }
    }
}

void rasterize_bezier_quad(float x0, float y0, float x1, float y1, float x2, float y2, uint8_t *coverage, int w, int h) {
    if (w <= 0 || h <= 0 || coverage == 0) {
        return;
    }
    if (x0 != x0 || y0 != y0 || x1 != x1 || y1 != y1 || x2 != x2 || y2 != y2) {
        return;
    }

    float min_x = x0;
    if (x1 < min_x) min_x = x1;
    if (x2 < min_x) min_x = x2;

    float max_x = x0;
    if (x1 > max_x) max_x = x1;
    if (x2 > max_x) max_x = x2;

    float min_y = y0;
    if (y1 < min_y) min_y = y1;
    if (y2 < min_y) min_y = y2;

    float max_y = y0;
    if (y1 > max_y) max_y = y1;
    if (y2 > max_y) max_y = y2;

    int imin_x = (int)min_x - 2;
    int imax_x = (int)max_x + 3;
    int imin_y = (int)min_y - 2;
    int imax_y = (int)max_y + 3;

    if (imin_x < 0) imin_x = 0;
    if (imin_y < 0) imin_y = 0;
    if (imax_x > w) imax_x = w;
    if (imax_y > h) imax_y = h;

    if (imin_x >= imax_x || imin_y >= imax_y) {
        return;
    }

    float pts_x[33];
    float pts_y[33];
    for (int i = 0; i <= 32; i++) {
        float t = (float)i / 32.0f;
        float inv_t = 1.0f - t;
        float b0 = inv_t * inv_t;
        float b1 = 2.0f * inv_t * t;
        float b2 = t * t;
        pts_x[i] = b0 * x0 + b1 * x1 + b2 * x2;
        pts_y[i] = b0 * y0 + b1 * y1 + b2 * y2;
    }

    for (int iy = imin_y; iy < imax_y; iy++) {
        float cy = (float)iy + 0.5f;
        int row = iy * w;
        for (int ix = imin_x; ix < imax_x; ix++) {
            float cx = (float)ix + 0.5f;
            float min_dist_sq = 1000000.0f;

            for (int k = 0; k < 32; k++) {
                float ax = pts_x[k];
                float ay = pts_y[k];
                float bx = pts_x[k + 1];
                float by = pts_y[k + 1];
                float dx = bx - ax;
                float dy = by - ay;
                float seg_len_sq = dx * dx + dy * dy;
                float vx = cx - ax;
                float vy = cy - ay;
                float dist_sq;

                if (seg_len_sq > 0.000001f) {
                    float t_seg = (vx * dx + vy * dy) / seg_len_sq;
                    if (t_seg < 0.0f) {
                        t_seg = 0.0f;
                    } else if (t_seg > 1.0f) {
                        t_seg = 1.0f;
                    }
                    float qx = ax + t_seg * dx;
                    float qy = ay + t_seg * dy;
                    float ex = cx - qx;
                    float ey = cy - qy;
                    dist_sq = ex * ex + ey * ey;
                } else {
                    dist_sq = vx * vx + vy * vy;
                }

                if (dist_sq < min_dist_sq) {
                    min_dist_sq = dist_sq;
                }
            }

            if (min_dist_sq < 2.25f) {
                float dist = c2_sqrtf(min_dist_sq);
                int cov = 0;
                if (dist <= 0.5f) {
                    cov = 255;
                } else if (dist < 1.5f) {
                    cov = (int)((1.5f - dist) * 255.0f + 0.5f);
                    if (cov > 255) cov = 255;
                    if (cov < 0) cov = 0;
                }
                int idx = row + ix;
                if (cov > (int)coverage[idx]) {
                    coverage[idx] = (uint8_t)cov;
                }
            }
        }
    }
}

void rasterize_rounded_rect(uint32_t *framebuffer, int stride, int x, int y, int w, int h, int radius, uint32_t color) {
    if (w <= 0 || h <= 0 || stride <= 0 || framebuffer == 0) {
        return;
    }
    if (radius <= 0) {
        fill_solid_rect(framebuffer, stride, x, y, w, h, color);
        return;
    }
    if (radius * 2 > w) {
        radius = w / 2;
    }
    if (radius * 2 > h) {
        radius = h / 2;
    }

    fill_solid_rect(framebuffer, stride, x + radius, y, w - 2 * radius, h, color);
    fill_solid_rect(framebuffer, stride, x, y + radius, radius, h - 2 * radius, color);
    fill_solid_rect(framebuffer, stride, x + w - radius, y + radius, radius, h - 2 * radius, color);

    int r = radius;
    int r_fp = r << 8;

    int corner_centers_x[4];
    int corner_centers_y[4];
    int corner_min_x[4];
    int corner_min_y[4];
    int corner_max_x[4];
    int corner_max_y[4];

    /* Top-left */
    corner_centers_x[0] = x + r;
    corner_centers_y[0] = y + r;
    corner_min_x[0] = x;
    corner_min_y[0] = y;
    corner_max_x[0] = x + r;
    corner_max_y[0] = y + r;

    /* Top-right */
    corner_centers_x[1] = x + w - r;
    corner_centers_y[1] = y + r;
    corner_min_x[1] = x + w - r;
    corner_min_y[1] = y;
    corner_max_x[1] = x + w;
    corner_max_y[1] = y + r;

    /* Bottom-left */
    corner_centers_x[2] = x + r;
    corner_centers_y[2] = y + h - r;
    corner_min_x[2] = x;
    corner_min_y[2] = y + h - r;
    corner_max_x[2] = x + r;
    corner_max_y[2] = y + h;

    /* Bottom-right */
    corner_centers_x[3] = x + w - r;
    corner_centers_y[3] = y + h - r;
    corner_min_x[3] = x + w - r;
    corner_min_y[3] = y + h - r;
    corner_max_x[3] = x + w;
    corner_max_y[3] = y + h;

    int64_t inner_r = (int64_t)r - 1;
    int64_t inner_sq = 0;
    if (inner_r > 0) {
        inner_sq = inner_r * inner_r;
    }
    int64_t outer_sq = (int64_t)(r + 2) * (r + 2);

    for (int c = 0; c < 4; c++) {
        int ccx = corner_centers_x[c];
        int ccy = corner_centers_y[c];
        int min_x = corner_min_x[c];
        int min_y = corner_min_y[c];
        int max_x = corner_max_x[c];
        int max_y = corner_max_y[c];

        if (min_x < 0) min_x = 0;
        if (min_y < 0) min_y = 0;
        if (max_x > stride) max_x = stride;

        if (min_x >= max_x || min_y >= max_y) {
            continue;
        }

        for (int py = min_y; py < max_y; py++) {
            int row = py * stride;
            int64_t dy;
            if (py < ccy) {
                dy = (int64_t)(ccy - py);
            } else {
                dy = (int64_t)(py - ccy + 1);
            }
            int64_t dy2 = dy * dy;

            for (int px = min_x; px < max_x; px++) {
                int64_t dx;
                if (px < ccx) {
                    dx = (int64_t)(ccx - px);
                } else {
                    dx = (int64_t)(px - ccx + 1);
                }
                int64_t dist_sq = dx * dx + dy2;

                if (dist_sq <= inner_sq) {
                    framebuffer[row + px] = c2_blend_pixel(framebuffer[row + px], color);
                } else if (dist_sq < outer_sq) {
                    int dist_fp = (int)c2_isqrt64(dist_sq << 16);
                    int diff = r_fp + 128 - dist_fp;
                    if (diff >= 255) {
                        framebuffer[row + px] = c2_blend_pixel(framebuffer[row + px], color);
                    } else if (diff > 0) {
                        framebuffer[row + px] = c2_blend_pixel_cov(framebuffer[row + px], color, (uint32_t)diff);
                    }
                }
            }
        }
    }
}
