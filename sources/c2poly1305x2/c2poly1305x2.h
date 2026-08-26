/* c2poly1305x2 — Poly1305 scalaire à double accumulateur (r, r²).
 *
 * C99 portable, aucun intrinsèque. Deux blocs de 16 octets par itération :
 *   h = ((h + m0) * r² + m1 * r)  mod (2^130 − 5)
 * Le reste d'un bloc plein est traité avec r. Le bloc partiel final
 * porte l'octet de bourrage 0x01, sans le bit haut 2^128.
 */
#ifndef C2POLY1305X2_H
#define C2POLY1305X2_H

#include <stddef.h>
#include <stdint.h>

typedef struct {
	uint8_t  c[32];
	size_t   c_idx;
	uint32_t r[5];
	uint32_t r2[5];
	uint32_t pad[4];
	uint32_t h[5];
} crypto_poly1305x2_ctx;

void crypto_poly1305x2_init(crypto_poly1305x2_ctx *ctx, const uint8_t key[32]);
void crypto_poly1305x2_update(crypto_poly1305x2_ctx *ctx,
                              const uint8_t *message, size_t message_size);
void crypto_poly1305x2_final(crypto_poly1305x2_ctx *ctx, uint8_t mac[16]);
void crypto_poly1305x2(uint8_t mac[16], const uint8_t *message,
                       size_t message_size, const uint8_t key[32]);

#endif
