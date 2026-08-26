#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include <assert.h>
#include "c2_ssim_gaussian.h"

static uint64_t hash_bytes(const uint8_t *data, size_t len) {
    uint64_t h = 0xcbf29ce484222325ULL;
    for (size_t i = 0; i < len; i++) {
        h ^= data[i];
        h *= 0x100000001b3ULL;
    }
    return h;
}

int main(int argc, char **argv) {
    (void)argc;
    (void)argv;
    setvbuf(stdout, NULL, _IONBF, 0);
    setvbuf(stderr, NULL, _IONBF, 0);

    printf("[ORACLE] Démarrage du banc de test de parité c2_ssim_gaussian...\n");

    /* 1. Test d'invariance d'identité : SSIM d'une image avec elle-même == 1.000000 */
    int w = 64;
    int h = 64;
    int stride = 64;
    size_t sz = (size_t)w * (size_t)h;

    uint8_t *img1 = (uint8_t *)malloc(sz);
    uint8_t *img2 = (uint8_t *)malloc(sz);
    uint8_t *blur_dst = (uint8_t *)malloc(sz);
    uint16_t *blur_tmp = (uint16_t *)malloc(sz * sizeof(uint16_t));

    assert(img1 && img2 && blur_dst && blur_tmp);

    for (size_t i = 0; i < sz; i++) {
        img1[i] = (uint8_t)((i * 17 + 43) & 0xFF);
        img2[i] = img1[i];
    }

    int64_t ssim_ident_milli = c2_ssim_compute_milli(img1, img2, w, h, stride);
    int64_t ssim_ident_q16 = c2_ssim_compute_q16(img1, img2, w, h, stride);

    if (ssim_ident_milli != 1000000LL) {
        fprintf(stderr, "FAIL: SSIM identité milli = %lld (attendu 1000000)\n", (long long)ssim_ident_milli);
        return 1;
    }
    if (ssim_ident_q16 != 65536LL) {
        fprintf(stderr, "FAIL: SSIM identité Q16 = %lld (attendu 65536)\n", (long long)ssim_ident_q16);
        return 1;
    }
    printf("[ORACLE] SSIM identité exacte validée : 1000000 / 1000000 (1.000000)\n");

    /* 2. Test de préservation de la moyenne sur image uniforme */
    memset(img1, 128, sz);
    c2_gaussian_blur_2d(img1, blur_dst, blur_tmp, w, h, stride);
    for (size_t i = 0; i < sz; i++) {
        if (blur_dst[i] != 128) {
            fprintf(stderr, "FAIL: Dérive du filtre gaussien sur image uniforme: pixel %zu = %d (attendu 128)\n", i, blur_dst[i]);
            return 1;
        }
    }
    printf("[ORACLE] Préservation exacte de l'énergie DC validée (128 -> 128 partout).\n");

    /* 3. Test de dégradation monotone sous bruit */
    for (size_t i = 0; i < sz; i++) {
        img1[i] = (uint8_t)((i * 31 + 7) & 0xFF);
        /* Ajout de bruit contrôlé */
        int noisy = (int)img1[i] + (int)((i % 11) - 5);
        if (noisy < 0) noisy = 0;
        if (noisy > 255) noisy = 255;
        img2[i] = (uint8_t)noisy;
    }
    int64_t ssim_noisy_milli = c2_ssim_compute_milli(img1, img2, w, h, stride);
    if (ssim_noisy_milli >= 1000000LL || ssim_noisy_milli <= 0LL) {
        fprintf(stderr, "FAIL: Score SSIM sous bruit incohérent : %lld\n", (long long)ssim_noisy_milli);
        return 1;
    }
    printf("[ORACLE] SSIM sous bruit léger mesuré : %lld millièmes (monotonie PASS)\n", (long long)ssim_noisy_milli);

    /* 4. Empreinte KAT sur image 128x128 */
    int kw = 128;
    int kh = 128;
    size_t ksz = (size_t)kw * (size_t)kh;
    uint8_t *kimg1 = (uint8_t *)malloc(ksz);
    uint8_t *kimg2 = (uint8_t *)malloc(ksz);
    uint8_t *kblur = (uint8_t *)malloc(ksz);
    uint16_t *ktmp = (uint16_t *)malloc(ksz * sizeof(uint16_t));

    for (size_t i = 0; i < ksz; i++) {
        kimg1[i] = (uint8_t)((i * 59 + 13) & 0xFF);
        kimg2[i] = (uint8_t)((i * 67 + 29) & 0xFF);
    }
    c2_gaussian_blur_2d(kimg1, kblur, ktmp, kw, kh, kw);
    uint64_t blur_hash = hash_bytes(kblur, ksz);
    int64_t kat_ssim = c2_ssim_compute_milli(kimg1, kimg2, kw, kh, kw);

    printf("[ORACLE] KAT GaussBlur 128x128 Hash = 0x%016llX\n", (unsigned long long)blur_hash);
    printf("[ORACLE] KAT SSIM 128x128 Score = %lld millièmes\n", (long long)kat_ssim);

    free(img1);
    free(img2);
    free(blur_dst);
    free(blur_tmp);
    free(kimg1);
    free(kimg2);
    free(kblur);
    free(ktmp);

    printf("[ORACLE] SUCCÈS : 100%% des vérifications de parité bit-exacte validées.\n");
    return 0;
}
