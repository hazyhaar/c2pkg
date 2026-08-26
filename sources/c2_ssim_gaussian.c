#include "c2_ssim_gaussian.h"

static const int32_t c2_gauss_weights[11] = {
    67, 498, 2360, 7167, 13959, 17434, 13959, 7167, 2360, 498, 67
};

static int c2_clamp_coord(int coord, int max_val) {
    if (coord < 0) {
        return 0;
    }
    if (coord >= max_val) {
        return max_val - 1;
    }
    return coord;
}

void c2_gaussian_blur_1d_h(const uint8_t *src, uint16_t *tmp, int width, int height, int stride) {
    int y = 0;
    int x = 0;
    int k = 0;

    if (width <= 0 || height <= 0) {
        return;
    }

    for (y = 0; y < height; y++) {
        int row_src_off = y * stride;
        int row_tmp_off = y * width;

        for (x = 0; x < width; x++) {
            int64_t acc = 32768LL;
            for (k = -5; k <= 5; k++) {
                int cx = c2_clamp_coord(x + k, width);
                int32_t w = c2_gauss_weights[k + 5];
                uint8_t val = src[row_src_off + cx];
                acc += (int64_t)w * (int64_t)val;
            }
            tmp[row_tmp_off + x] = (uint16_t)(acc >> 16);
        }
    }
}

void c2_gaussian_blur_1d_v(const uint16_t *tmp, uint8_t *dst, int width, int height, int stride) {
    int y = 0;
    int x = 0;
    int k = 0;

    if (width <= 0 || height <= 0) {
        return;
    }

    for (y = 0; y < height; y++) {
        int row_dst_off = y * stride;

        for (x = 0; x < width; x++) {
            int64_t acc = 32768LL;
            for (k = -5; k <= 5; k++) {
                int cy = c2_clamp_coord(y + k, height);
                int32_t w = c2_gauss_weights[k + 5];
                uint16_t val = tmp[cy * width + x];
                acc += (int64_t)w * (int64_t)val;
            }
            int64_t out_val = acc >> 16;
            if (out_val > 255LL) {
                out_val = 255LL;
            }
            dst[row_dst_off + x] = (uint8_t)out_val;
        }
    }
}

void c2_gaussian_blur_2d(const uint8_t *src, uint8_t *dst, uint16_t *tmp, int width, int height, int stride) {
    c2_gaussian_blur_1d_h(src, tmp, width, height, stride);
    c2_gaussian_blur_1d_v(tmp, dst, width, height, stride);
}

void c2_gaussian_blur_1d_h_u32(const uint32_t *src, uint32_t *tmp, int width, int height, int stride) {
    int y = 0;
    int x = 0;
    int k = 0;

    if (width <= 0 || height <= 0) {
        return;
    }

    for (y = 0; y < height; y++) {
        int row_src_off = y * stride;
        int row_tmp_off = y * width;

        for (x = 0; x < width; x++) {
            int64_t acc = 32768LL;
            for (k = -5; k <= 5; k++) {
                int cx = c2_clamp_coord(x + k, width);
                int32_t w = c2_gauss_weights[k + 5];
                uint32_t val = src[row_src_off + cx];
                acc += (int64_t)w * (int64_t)val;
            }
            tmp[row_tmp_off + x] = (uint32_t)(acc >> 16);
        }
    }
}

void c2_gaussian_blur_1d_v_u32(const uint32_t *tmp, uint32_t *dst, int width, int height, int stride) {
    int y = 0;
    int x = 0;
    int k = 0;

    if (width <= 0 || height <= 0) {
        return;
    }

    for (y = 0; y < height; y++) {
        int row_dst_off = y * stride;

        for (x = 0; x < width; x++) {
            int64_t acc = 32768LL;
            for (k = -5; k <= 5; k++) {
                int cy = c2_clamp_coord(y + k, height);
                int32_t w = c2_gauss_weights[k + 5];
                uint32_t val = tmp[cy * width + x];
                acc += (int64_t)w * (int64_t)val;
            }
            dst[row_dst_off + x] = (uint32_t)(acc >> 16);
        }
    }
}

void c2_gaussian_blur_2d_u32(const uint32_t *src, uint32_t *dst, uint32_t *tmp, int width, int height, int stride) {
    c2_gaussian_blur_1d_h_u32(src, tmp, width, height, stride);
    c2_gaussian_blur_1d_v_u32(tmp, dst, width, height, stride);
}

int64_t c2_ssim_compute_milli(const uint8_t *img1, const uint8_t *img2, int width, int height, int stride) {
    int y = 0;
    int x = 0;
    int64_t total_ssim = 0;
    int64_t num_pixels = 0;

    /* Constantes standard SSIM : C1 = 6.5025, C2 = 58.5225 */
    /* En échelle x 100 : C1 = 650, C2 = 5852 */
    const int64_t c1 = 650LL;
    const int64_t c2 = 5852LL;

    if (width <= 0 || height <= 0 || !img1 || !img2) {
        return 0;
    }

    for (y = 0; y < height; y++) {
        for (x = 0; x < width; x++) {
            int64_t mu1_acc = 32768LL;
            int64_t mu2_acc = 32768LL;
            int64_t mu1_sq_acc = 32768LL;
            int64_t mu2_sq_acc = 32768LL;
            int64_t mu12_acc = 32768LL;
            int ky = 0;
            int kx = 0;

            for (ky = -5; ky <= 5; ky++) {
                int cy = c2_clamp_coord(y + ky, height);
                int32_t wy = c2_gauss_weights[ky + 5];

                for (kx = -5; kx <= 5; kx++) {
                    int cx = c2_clamp_coord(x + kx, width);
                    int32_t wx = c2_gauss_weights[kx + 5];
                    /* Poids 2D séparable combiné en virgule fixe */
                    int64_t w2d = ((int64_t)wy * (int64_t)wx + 32768LL) >> 16;

                    uint8_t p1 = img1[cy * stride + cx];
                    uint8_t p2 = img2[cy * stride + cx];

                    mu1_acc += w2d * (int64_t)p1;
                    mu2_acc += w2d * (int64_t)p2;
                    mu1_sq_acc += w2d * ((int64_t)p1 * (int64_t)p1);
                    mu2_sq_acc += w2d * ((int64_t)p2 * (int64_t)p2);
                    mu12_acc += w2d * ((int64_t)p1 * (int64_t)p2);
                }
            }

            int64_t mu1 = mu1_acc >> 16;
            int64_t mu2 = mu2_acc >> 16;
            int64_t mu1_sq = mu1_sq_acc >> 16;
            int64_t mu2_sq = mu2_sq_acc >> 16;
            int64_t mu12 = mu12_acc >> 16;

            int64_t sigma1_sq = mu1_sq - (mu1 * mu1);
            int64_t sigma2_sq = mu2_sq - (mu2 * mu2);
            int64_t sigma12 = mu12 - (mu1 * mu2);

            if (sigma1_sq < 0) sigma1_sq = 0;
            if (sigma2_sq < 0) sigma2_sq = 0;

            int64_t num1 = 2LL * mu1 * mu2 + c1;
            int64_t num2 = 2LL * sigma12 + c2;
            int64_t den1 = mu1 * mu1 + mu2 * mu2 + c1;
            int64_t den2 = sigma1_sq + sigma2_sq + c2;

            int64_t num = num1 * num2;
            int64_t den = den1 * den2;

            if (den != 0) {
                int64_t ssim_pix = (num * 1000000LL) / den;
                if (ssim_pix > 1000000LL) ssim_pix = 1000000LL;
                if (ssim_pix < -1000000LL) ssim_pix = -1000000LL;
                total_ssim += ssim_pix;
            }
            num_pixels++;
        }
    }

    if (num_pixels == 0) {
        return 0;
    }
    return total_ssim / num_pixels;
}

int64_t c2_ssim_compute_q16(const uint8_t *img1, const uint8_t *img2, int width, int height, int stride) {
    int64_t milli = c2_ssim_compute_milli(img1, img2, width, height, stride);
    return (milli * 65536LL + 500000LL) / 1000000LL;
}
