/* Oracle gcc -O2 -mavx2 pour c2chacha4 : vecteur RFC 8439 §2.4.2 (114 octets)
 * puis vecteurs LCG (graine = taille) sur 0..512 octets. Imprime "n:hex". */
#include <stdio.h>
#include <string.h>
#include "c2chacha4.h"

static uint64_t lcg_state;
static void lcg_seed(uint64_t s) { lcg_state = s * 0x9E3779B97F4A7C15ULL + 1; }
static uint8_t lcg_byte(void) {
    lcg_state = lcg_state * 6364136223846793005ULL + 1442695040888963407ULL;
    return (uint8_t)(lcg_state >> 56);
}

int main(void) {
    /* RFC 8439 §2.4.2 : clé 00..1f, nonce 00 00 00 00 00 00 00 4a 00 00 00 00,
     * compteur 1, texte "Ladies and Gentlemen..." (114 octets). */
    uint8_t key[32], nonce[12], out[512];
    for (int i = 0; i < 32; i++) key[i] = (uint8_t)i;
    memset(nonce, 0, 12);
    nonce[7] = 0x4a;
    const char *msg = "Ladies and Gentlemen of the class of '99: If I could offer you only one tip for the future, sunscreen would be it.";
    c2chacha4_xor_blocks(out, (const uint8_t *)msg, 114, key, nonce, 1);
    printf("rfc:");
    for (int i = 0; i < 114; i++) printf("%02x", out[i]);
    printf("\n");

    static const size_t sizes[] = {0, 1, 63, 64, 65, 127, 128, 129, 191, 192,
                                   255, 256, 320, 384, 447, 448, 511, 512};
    uint8_t in[512];
    for (size_t i = 0; i < sizeof sizes / sizeof sizes[0]; i++) {
        size_t n = sizes[i];
        lcg_seed((uint64_t)n);
        for (int k = 0; k < 32; k++) key[k] = lcg_byte();
        for (int k = 0; k < 12; k++) nonce[k] = lcg_byte();
        uint32_t ctr = (uint32_t)lcg_byte() | ((uint32_t)lcg_byte() << 8);
        for (size_t k = 0; k < n; k++) in[k] = lcg_byte();
        c2chacha4_xor_blocks(out, in, n, key, nonce, ctr);
        printf("%zu:", n);
        for (size_t k = 0; k < n; k++) printf("%02x", out[k]);
        printf("\n");
    }

    /* Débordement du compteur par voie (counter + [0..7] modulo 2^32) :
     * x/crypto refuse ces compteurs, seul le C peut servir d'oracle.
     * Lignes "wrap:<ctr hex>:<hex>" sur 512 octets, graine LCG = ctr. */
    static const uint32_t wraps[] = {0xFFFFFFF8u, 0xFFFFFFFCu, 0xFFFFFFFFu};
    for (size_t i = 0; i < sizeof wraps / sizeof wraps[0]; i++) {
        uint32_t ctr = wraps[i];
        lcg_seed((uint64_t)ctr);
        for (int k = 0; k < 32; k++) key[k] = lcg_byte();
        for (int k = 0; k < 12; k++) nonce[k] = lcg_byte();
        for (size_t k = 0; k < 512; k++) in[k] = lcg_byte();
        c2chacha4_xor_blocks(out, in, 512, key, nonce, ctr);
        printf("wrap:%08x:", ctr);
        for (size_t k = 0; k < 512; k++) printf("%02x", out[k]);
        printf("\n");
    }
    return 0;
}
