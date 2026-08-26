/* Oracle C gcc -O2 pour c2pkg/c2poly1305 : Poly1305 de monocypher 4.0.2.
 * Pour chaque taille de la liste, clé et message sont dérivés d'un LCG
 * déterministe (graine = taille) ; le MAC est imprimé en hexadécimal.
 * Le test Go rejoue le même LCG et compare octet à octet. */
#include <stdint.h>
#include <stdio.h>
#include <string.h>
#include "monocypher.h"

static uint64_t lcg_state;
static void lcg_seed(uint64_t s) { lcg_state = s * 0x9E3779B97F4A7C15ULL + 1; }
static uint8_t lcg_byte(void) {
    lcg_state = lcg_state * 6364136223846793005ULL + 1442695040888963407ULL;
    return (uint8_t)(lcg_state >> 56);
}

int main(void) {
    static const size_t sizes[] = {0, 1, 15, 16, 17, 31, 32, 33, 47, 48, 63, 64, 65, 96,
                                   127, 128, 129, 255, 256, 1000, 1023, 1024, 1025, 4096};
    static uint8_t msg[4096];
    uint8_t key[32], mac[16];
    for (size_t i = 0; i < sizeof sizes / sizeof sizes[0]; i++) {
        size_t n = sizes[i];
        lcg_seed((uint64_t)n);
        for (int k = 0; k < 32; k++) key[k] = lcg_byte();
        for (size_t k = 0; k < n; k++) msg[k] = lcg_byte();
        crypto_poly1305(mac, msg, n, key);
        printf("%zu:", n);
        for (int k = 0; k < 16; k++) printf("%02x", mac[k]);
        printf("\n");
    }
    return 0;
}
