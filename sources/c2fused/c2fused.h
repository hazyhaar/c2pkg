/* c2fused — ChaCha20 8 blocs et Poly1305 scalaire 64 bits entrelacés.
 *
 * c2fused_seal_blocks chiffre n ≤ 512 octets comme c2chacha8 et, si prev
 * n'est pas NULL, absorbe les 32 blocs Poly de 16 octets de prev (le chiffré
 * de la tranche précédente) après les quarts de tour désignés par
 * c2fused_schedule.h (grille ARCHTIME). Prologue : prev = NULL. Épilogue :
 * c2fused_poly_blocks sur la dernière tranche, sans keystream.
 */
#ifndef C2FUSED_H
#define C2FUSED_H

#include <stddef.h>
#include <stdint.h>

typedef struct {
	uint64_t h[3];
	uint64_t r[2];
	uint64_t pad[2];
	uint8_t  leftover[16];
	size_t   leftover_len;
} c2fused_poly;

void c2fused_poly_init(c2fused_poly *st, const uint8_t key[32]);
void c2fused_poly_blocks(c2fused_poly *st, const uint8_t *m, size_t nblocks);
void c2fused_poly_update(c2fused_poly *st, const uint8_t *m, size_t n);
void c2fused_poly_final(c2fused_poly *st, uint8_t mac[16]);

size_t c2fused_seal_blocks(uint8_t *out, const uint8_t *in, size_t n,
                           const uint8_t key[32], const uint8_t nonce[12],
                           uint32_t counter, c2fused_poly *st,
                           const uint8_t *prev);

#endif
