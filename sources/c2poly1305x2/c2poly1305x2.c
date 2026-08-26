/* c2poly1305x2.c — Poly1305 (r, r²), membres de 26 bits, C99 portable. */

#include "c2poly1305x2.h"

static uint32_t load32_le(const uint8_t *p)
{
	return (uint32_t)p[0]
	     | ((uint32_t)p[1] << 8)
	     | ((uint32_t)p[2] << 16)
	     | ((uint32_t)p[3] << 24);
}

static void store32_le(uint8_t *p, uint32_t v)
{
	p[0] = (uint8_t)(v & 0xffu);
	p[1] = (uint8_t)((v >> 8) & 0xffu);
	p[2] = (uint8_t)((v >> 16) & 0xffu);
	p[3] = (uint8_t)((v >> 24) & 0xffu);
}

/* h = h * r  (produit 5×5, réduction partielle donna). */
static void poly_mul(uint32_t h[5], const uint32_t r[5])
{
	uint32_t r0 = r[0];
	uint32_t r1 = r[1];
	uint32_t r2 = r[2];
	uint32_t r3 = r[3];
	uint32_t r4 = r[4];
	uint32_t s1 = r1 * 5u;
	uint32_t s2 = r2 * 5u;
	uint32_t s3 = r3 * 5u;
	uint32_t s4 = r4 * 5u;
	uint32_t h0 = h[0];
	uint32_t h1 = h[1];
	uint32_t h2 = h[2];
	uint32_t h3 = h[3];
	uint32_t h4 = h[4];
	uint64_t d0 = (uint64_t)h0 * r0 + (uint64_t)h1 * s4 + (uint64_t)h2 * s3
	            + (uint64_t)h3 * s2 + (uint64_t)h4 * s1;
	uint64_t d1 = (uint64_t)h0 * r1 + (uint64_t)h1 * r0 + (uint64_t)h2 * s4
	            + (uint64_t)h3 * s3 + (uint64_t)h4 * s2;
	uint64_t d2 = (uint64_t)h0 * r2 + (uint64_t)h1 * r1 + (uint64_t)h2 * r0
	            + (uint64_t)h3 * s4 + (uint64_t)h4 * s3;
	uint64_t d3 = (uint64_t)h0 * r3 + (uint64_t)h1 * r2 + (uint64_t)h2 * r1
	            + (uint64_t)h3 * r0 + (uint64_t)h4 * s4;
	uint64_t d4 = (uint64_t)h0 * r4 + (uint64_t)h1 * r3 + (uint64_t)h2 * r2
	            + (uint64_t)h3 * r1 + (uint64_t)h4 * r0;
	uint64_t c;

	c = d0 >> 26;
	h0 = (uint32_t)d0 & 0x3ffffffu;
	d1 += c;
	c = d1 >> 26;
	h1 = (uint32_t)d1 & 0x3ffffffu;
	d2 += c;
	c = d2 >> 26;
	h2 = (uint32_t)d2 & 0x3ffffffu;
	d3 += c;
	c = d3 >> 26;
	h3 = (uint32_t)d3 & 0x3ffffffu;
	d4 += c;
	c = d4 >> 26;
	h4 = (uint32_t)d4 & 0x3ffffffu;
	h0 += (uint32_t)c * 5u;
	c = (uint64_t)h0 >> 26;
	h0 &= 0x3ffffffu;
	h1 += (uint32_t)c;

	h[0] = h0;
	h[1] = h1;
	h[2] = h2;
	h[3] = h3;
	h[4] = h4;
}

/* Ajoute le bloc de 16 octets (bit haut hibit, 0 ou 0x1000000) à h. */
static void poly_add_block(uint32_t h[5], const uint8_t *m, uint32_t hibit)
{
	uint32_t w0 = load32_le(m);
	uint32_t w1 = load32_le(m + 3);
	uint32_t w2 = load32_le(m + 6);
	uint32_t w3 = load32_le(m + 9);
	uint32_t w4 = load32_le(m + 12);
	h[0] += w0 & 0x3ffffffu;
	h[1] += (w1 >> 2) & 0x3ffffffu;
	h[2] += (w2 >> 4) & 0x3ffffffu;
	h[3] += (w3 >> 6) & 0x3ffffffu;
	h[4] += (w4 >> 8) | hibit;
}

/* h = h + t, une passe de retenue 26 bits. */
static void poly_add5(uint32_t h[5], const uint32_t t[5])
{
	uint64_t c;
	c = (uint64_t)h[0] + t[0];
	h[0] = (uint32_t)c & 0x3ffffffu;
	c = (c >> 26) + (uint64_t)h[1] + t[1];
	h[1] = (uint32_t)c & 0x3ffffffu;
	c = (c >> 26) + (uint64_t)h[2] + t[2];
	h[2] = (uint32_t)c & 0x3ffffffu;
	c = (c >> 26) + (uint64_t)h[3] + t[3];
	h[3] = (uint32_t)c & 0x3ffffffu;
	c = (c >> 26) + (uint64_t)h[4] + t[4];
	h[4] = (uint32_t)c;
}

/* Un bloc plein : h = (h + m) * r, bit haut posé. */
static void poly_block_r(crypto_poly1305x2_ctx *ctx, const uint8_t *m)
{
	poly_add_block(ctx->h, m, 0x1000000u);
	poly_mul(ctx->h, ctx->r);
}

/* Un bloc partiel final : h = (h + m) * r, bit haut absent. */
static void poly_block_r_final(crypto_poly1305x2_ctx *ctx, const uint8_t *m)
{
	poly_add_block(ctx->h, m, 0);
	poly_mul(ctx->h, ctx->r);
}

/* Deux blocs : h = (h + m0) * r² + m1 * r. */
static void poly_pair(crypto_poly1305x2_ctx *ctx, const uint8_t *m)
{
	uint32_t t[5];
	poly_add_block(ctx->h, m, 0x1000000u);
	poly_mul(ctx->h, ctx->r2);
	t[0] = 0;
	t[1] = 0;
	t[2] = 0;
	t[3] = 0;
	t[4] = 0;
	poly_add_block(t, m + 16, 0x1000000u);
	poly_mul(t, ctx->r);
	poly_add5(ctx->h, t);
}

void crypto_poly1305x2_init(crypto_poly1305x2_ctx *ctx, const uint8_t key[32])
{
	uint32_t t0 = load32_le(key);
	uint32_t t1 = load32_le(key + 4);
	uint32_t t2 = load32_le(key + 8);
	uint32_t t3 = load32_le(key + 12);
	size_t i;

	ctx->r[0] = t0 & 0x3ffffffu;
	ctx->r[1] = ((t0 >> 26) | (t1 << 6)) & 0x3ffff03u;
	ctx->r[2] = ((t1 >> 20) | (t2 << 12)) & 0x3ffc0ffu;
	ctx->r[3] = ((t2 >> 14) | (t3 << 18)) & 0x3f03fffu;
	ctx->r[4] = (t3 >> 8) & 0x00fffffu;

	ctx->r2[0] = ctx->r[0];
	ctx->r2[1] = ctx->r[1];
	ctx->r2[2] = ctx->r[2];
	ctx->r2[3] = ctx->r[3];
	ctx->r2[4] = ctx->r[4];
	poly_mul(ctx->r2, ctx->r);

	ctx->pad[0] = load32_le(key + 16);
	ctx->pad[1] = load32_le(key + 20);
	ctx->pad[2] = load32_le(key + 24);
	ctx->pad[3] = load32_le(key + 28);

	ctx->h[0] = 0;
	ctx->h[1] = 0;
	ctx->h[2] = 0;
	ctx->h[3] = 0;
	ctx->h[4] = 0;
	ctx->c_idx = 0;
	for (i = 0; i < 32; i++) {
		ctx->c[i] = 0;
	}
}

void crypto_poly1305x2_update(crypto_poly1305x2_ctx *ctx,
                              const uint8_t *message, size_t message_size)
{
	size_t i;
	size_t need;
	size_t pairs;

	if (message_size == 0) {
		return;
	}

	if (ctx->c_idx != 0) {
		need = 32 - ctx->c_idx;
		if (need > message_size) {
			need = message_size;
		}
		for (i = 0; i < need; i++) {
			ctx->c[ctx->c_idx] = *message;
			ctx->c_idx = ctx->c_idx + 1;
			message++;
			message_size--;
		}
		if (ctx->c_idx == 32) {
			poly_pair(ctx, ctx->c);
			ctx->c_idx = 0;
		}
	}

	pairs = message_size >> 5;
	for (i = 0; i < pairs; i++) {
		poly_pair(ctx, message);
		message += 32;
		message_size -= 32;
	}

	for (i = 0; i < message_size; i++) {
		ctx->c[ctx->c_idx] = message[i];
		ctx->c_idx = ctx->c_idx + 1;
	}
}

void crypto_poly1305x2_final(crypto_poly1305x2_ctx *ctx, uint8_t mac[16])
{
	size_t i;
	size_t leftover;
	uint32_t h0, h1, h2, h3, h4;
	uint32_t g0, g1, g2, g3, g4;
	uint32_t c;
	uint32_t mask;
	uint64_t f;

	if (ctx->c_idx >= 16) {
		poly_block_r(ctx, ctx->c);
		leftover = ctx->c_idx - 16;
		for (i = 0; i < leftover; i++) {
			ctx->c[i] = ctx->c[16 + i];
		}
		ctx->c_idx = leftover;
	}
	if (ctx->c_idx != 0) {
		ctx->c[ctx->c_idx] = 1;
		for (i = ctx->c_idx + 1; i < 16; i++) {
			ctx->c[i] = 0;
		}
		poly_block_r_final(ctx, ctx->c);
	}

	h0 = ctx->h[0];
	h1 = ctx->h[1];
	h2 = ctx->h[2];
	h3 = ctx->h[3];
	h4 = ctx->h[4];

	c = h1 >> 26;
	h1 = h1 & 0x3ffffffu;
	h2 = h2 + c;
	c = h2 >> 26;
	h2 = h2 & 0x3ffffffu;
	h3 = h3 + c;
	c = h3 >> 26;
	h3 = h3 & 0x3ffffffu;
	h4 = h4 + c;
	c = h4 >> 26;
	h4 = h4 & 0x3ffffffu;
	h0 = h0 + c * 5u;
	c = h0 >> 26;
	h0 = h0 & 0x3ffffffu;
	h1 = h1 + c;

	g0 = h0 + 5u;
	c = g0 >> 26;
	g0 = g0 & 0x3ffffffu;
	g1 = h1 + c;
	c = g1 >> 26;
	g1 = g1 & 0x3ffffffu;
	g2 = h2 + c;
	c = g2 >> 26;
	g2 = g2 & 0x3ffffffu;
	g3 = h3 + c;
	c = g3 >> 26;
	g3 = g3 & 0x3ffffffu;
	g4 = h4 + c - ((uint32_t)1 << 26);

	mask = (g4 >> 31) - 1u;
	g0 = g0 & mask;
	g1 = g1 & mask;
	g2 = g2 & mask;
	g3 = g3 & mask;
	g4 = g4 & mask;
	mask = ~mask;
	h0 = (h0 & mask) | g0;
	h1 = (h1 & mask) | g1;
	h2 = (h2 & mask) | g2;
	h3 = (h3 & mask) | g3;
	h4 = (h4 & mask) | g4;

	h0 = (h0 | (h1 << 26)) & 0xffffffffu;
	h1 = ((h1 >> 6) | (h2 << 20)) & 0xffffffffu;
	h2 = ((h2 >> 12) | (h3 << 14)) & 0xffffffffu;
	h3 = ((h3 >> 18) | (h4 << 8)) & 0xffffffffu;

	f = (uint64_t)h0 + ctx->pad[0];
	h0 = (uint32_t)f;
	f = (uint64_t)h1 + ctx->pad[1] + (f >> 32);
	h1 = (uint32_t)f;
	f = (uint64_t)h2 + ctx->pad[2] + (f >> 32);
	h2 = (uint32_t)f;
	f = (uint64_t)h3 + ctx->pad[3] + (f >> 32);
	h3 = (uint32_t)f;

	store32_le(mac, h0);
	store32_le(mac + 4, h1);
	store32_le(mac + 8, h2);
	store32_le(mac + 12, h3);

	for (i = 0; i < 32; i++) {
		ctx->c[i] = 0;
	}
	ctx->c_idx = 0;
	for (i = 0; i < 5; i++) {
		ctx->r[i] = 0;
		ctx->r2[i] = 0;
		ctx->h[i] = 0;
	}
	for (i = 0; i < 4; i++) {
		ctx->pad[i] = 0;
	}
}

void crypto_poly1305x2(uint8_t mac[16], const uint8_t *message,
                       size_t message_size, const uint8_t key[32])
{
	crypto_poly1305x2_ctx ctx;
	crypto_poly1305x2_init(&ctx, key);
	crypto_poly1305x2_update(&ctx, message, message_size);
	crypto_poly1305x2_final(&ctx, mac);
}
