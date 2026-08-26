#include "c2_swizzle_simd.h"

uint32_t c2_swizzle_pixel(uint32_t pixel) {
    return (pixel & 0xFF00FF00U) | ((pixel & 0x000000FFU) << 16) | ((pixel & 0x00FF0000U) >> 16);
}

uint64_t c2_swizzle_pair(uint64_t pair) {
    return (pair & 0xFF00FF00FF00FF00ULL) | ((pair & 0x000000FF000000FFULL) << 16) | ((pair & 0x00FF000000FF0000ULL) >> 16);
}

void c2_swizzle_block4(const uint8_t in[16], uint8_t out[16]) {
    out[0]  = in[2];  out[1]  = in[1];  out[2]  = in[0];  out[3]  = in[3];
    out[4]  = in[6];  out[5]  = in[5];  out[6]  = in[4];  out[7]  = in[7];
    out[8]  = in[10]; out[9]  = in[9];  out[10] = in[8];  out[11] = in[11];
    out[12] = in[14]; out[13] = in[13]; out[14] = in[12]; out[15] = in[15];
}

void c2_swizzle_block8(const uint8_t in[32], uint8_t out[32]) {
    c2_swizzle_block4(&in[0], &out[0]);
    c2_swizzle_block4(&in[16], &out[16]);
}

void c2_swizzle_block16(const uint8_t in[64], uint8_t out[64]) {
    c2_swizzle_block4(&in[0], &out[0]);
    c2_swizzle_block4(&in[16], &out[16]);
    c2_swizzle_block4(&in[32], &out[32]);
    c2_swizzle_block4(&in[48], &out[48]);
}

void c2_swizzle_rgba_bgra(const uint8_t *src, uint8_t *dst, size_t num_pixels) {
    size_t i = 0;
    size_t blocks16 = num_pixels >> 4;
    size_t limit16 = blocks16 << 4;

    for (i = 0; i < limit16; i += 16) {
        size_t off = i * 4;
        c2_swizzle_block16(&src[off], &dst[off]);
    }

    for (i = limit16; i < num_pixels; i++) {
        size_t off = i * 4;
        dst[off + 0] = src[off + 2];
        dst[off + 1] = src[off + 1];
        dst[off + 2] = src[off + 0];
        dst[off + 3] = src[off + 3];
    }
}

void c2_swizzle_rgba_bgra_inplace(uint8_t *data, size_t num_pixels) {
    size_t i = 0;
    size_t blocks8 = num_pixels >> 3;
    size_t limit8 = blocks8 << 3;

    for (i = 0; i < limit8; i += 8) {
        size_t off0 = i * 4;
        uint8_t t0 = data[off0 + 0];
        data[off0 + 0] = data[off0 + 2];
        data[off0 + 2] = t0;

        size_t off1 = off0 + 4;
        uint8_t t1 = data[off1 + 0];
        data[off1 + 0] = data[off1 + 2];
        data[off1 + 2] = t1;

        size_t off2 = off0 + 8;
        uint8_t t2 = data[off2 + 0];
        data[off2 + 0] = data[off2 + 2];
        data[off2 + 2] = t2;

        size_t off3 = off0 + 12;
        uint8_t t3 = data[off3 + 0];
        data[off3 + 0] = data[off3 + 2];
        data[off3 + 2] = t3;

        size_t off4 = off0 + 16;
        uint8_t t4 = data[off4 + 0];
        data[off4 + 0] = data[off4 + 2];
        data[off4 + 2] = t4;

        size_t off5 = off0 + 20;
        uint8_t t5 = data[off5 + 0];
        data[off5 + 0] = data[off5 + 2];
        data[off5 + 2] = t5;

        size_t off6 = off0 + 24;
        uint8_t t6 = data[off6 + 0];
        data[off6 + 0] = data[off6 + 2];
        data[off6 + 2] = t6;

        size_t off7 = off0 + 28;
        uint8_t t7 = data[off7 + 0];
        data[off7 + 0] = data[off7 + 2];
        data[off7 + 2] = t7;
    }

    for (i = limit8; i < num_pixels; i++) {
        size_t off = i * 4;
        uint8_t tmp = data[off + 0];
        data[off + 0] = data[off + 2];
        data[off + 2] = tmp;
    }
}

void c2_swizzle_argb_to_rgba(const uint8_t *src, uint8_t *dst, size_t num_pixels) {
    size_t i = 0;
    for (i = 0; i < num_pixels; i++) {
        size_t off = i * 4;
        uint8_t a = src[off + 0];
        uint8_t r = src[off + 1];
        uint8_t g = src[off + 2];
        uint8_t b = src[off + 3];

        dst[off + 0] = r;
        dst[off + 1] = g;
        dst[off + 2] = b;
        dst[off + 3] = a;
    }
}

void c2_swizzle_rgba_to_argb(const uint8_t *src, uint8_t *dst, size_t num_pixels) {
    size_t i = 0;
    for (i = 0; i < num_pixels; i++) {
        size_t off = i * 4;
        uint8_t r = src[off + 0];
        uint8_t g = src[off + 1];
        uint8_t b = src[off + 2];
        uint8_t a = src[off + 3];

        dst[off + 0] = a;
        dst[off + 1] = r;
        dst[off + 2] = g;
        dst[off + 3] = b;
    }
}

void c2_swizzle_argb_to_bgra(const uint8_t *src, uint8_t *dst, size_t num_pixels) {
    size_t i = 0;
    for (i = 0; i < num_pixels; i++) {
        size_t off = i * 4;
        uint8_t a = src[off + 0];
        uint8_t r = src[off + 1];
        uint8_t g = src[off + 2];
        uint8_t b = src[off + 3];

        dst[off + 0] = b;
        dst[off + 1] = g;
        dst[off + 2] = r;
        dst[off + 3] = a;
    }
}

void c2_swizzle_rgba_bgra_stride(const uint8_t *src, uint8_t *dst, int width, int height, int src_stride, int dst_stride) {
    int y = 0;
    if (width <= 0 || height <= 0) {
        return;
    }
    for (y = 0; y < height; y++) {
        size_t src_row_off = (size_t)y * (size_t)src_stride;
        size_t dst_row_off = (size_t)y * (size_t)dst_stride;
        c2_swizzle_rgba_bgra(&src[src_row_off], &dst[dst_row_off], (size_t)width);
    }
}
