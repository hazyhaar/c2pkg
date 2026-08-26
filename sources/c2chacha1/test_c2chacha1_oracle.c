/* Oracle gcc -O2 pour c2chacha1 : vecteur RFC 8439 §2.4.2 (bloc 1) puis
 * vecteurs LCG (graine = taille) sur 0..64 octets. Imprime "n:hex". */
#include <stdio.h>
#include <string.h>
#include "c2chacha1.h"

static uint64_t lcg_state;
static void lcg_seed(uint64_t s) { lcg_state = s * 0x9E3779B97F4A7C15ULL + 1; }
static uint8_t lcg_byte(void) {
    lcg_state = lcg_state * 6364136223846793005ULL + 1442695040888963407ULL;
    return (uint8_t)(lcg_state >> 56);
}

int main(void) {
    /* RFC 8439 §2.4.2 : clé 00..1f, nonce 00 00 00 00 00 00 00 4a 00 00 00 00,
     * compteur 1, texte "Ladies and Gentlemen..." (premiers 64 octets). */
    uint8_t key[32], nonce[12], out[64];
    for (int i = 0; i < 32; i++) key[i] = (uint8_t)i;
    memset(nonce, 0, 12);
    nonce[7] = 0x4a;
    const char *msg = "Ladies and Gentlemen of the class of '99: If I could offer you only one tip for the future, sunscreen would be it.";
    c2chacha1_xor_block(out, (const uint8_t *)msg, 64, key, nonce, 1);
    printf("rfc:");
    for (int i = 0; i < 64; i++) printf("%02x", out[i]);
    printf("\n");

    static const size_t sizes[] = {0, 1, 15, 16, 17, 31, 32, 33, 47, 48, 63, 64};
    uint8_t in[64];
    for (size_t i = 0; i < sizeof sizes / sizeof sizes[0]; i++) {
        size_t n = sizes[i];
        lcg_seed((uint64_t)n);
        for (int k = 0; k < 32; k++) key[k] = lcg_byte();
        for (int k = 0; k < 12; k++) nonce[k] = lcg_byte();
        uint32_t ctr = (uint32_t)lcg_byte() | ((uint32_t)lcg_byte() << 8);
        for (size_t k = 0; k < n; k++) in[k] = lcg_byte();
        c2chacha1_xor_block(out, in, n, key, nonce, ctr);
        printf("%zu:", n);
        for (size_t k = 0; k < n; k++) printf("%02x", out[k]);
        printf("\n");
    }
    return 0;
}
