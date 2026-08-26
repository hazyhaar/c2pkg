/* Oracle gcc -O2 -mavx2 pour c2fused : RFC 8439 §2.5.2 (Poly), §2.8.2 (AEAD),
 * parité chiffré vs c2chacha8, 4000 tirages fusionné vs séquentiel. */
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "c2fused.h"
#include "c2chacha8.h"

static uint64_t lcg_state;
static void lcg_seed(uint64_t s) { lcg_state = s * 0x9E3779B97F4A7C15ULL + 1; }
static uint8_t lcg_byte(void) {
	lcg_state = lcg_state * 6364136223846793005ULL + 1442695040888963407ULL;
	return (uint8_t)(lcg_state >> 56);
}

static void store64_le(uint8_t *p, uint64_t v) {
	int i;
	for (i = 0; i < 8; i++) {
		p[i] = (uint8_t)(v >> (8 * i));
	}
}

static void poly_pad(c2fused_poly *st, size_t n) {
	size_t rem = n & 15;
	if (rem != 0) {
		uint8_t z[16];
		memset(z, 0, 16);
		c2fused_poly_update(st, z, 16 - rem);
	}
}

static void aead_seq(uint8_t *out, uint8_t tag[16],
                     const uint8_t *in, size_t n,
                     const uint8_t *ad, size_t adn,
                     const uint8_t key[32], const uint8_t nonce[12]) {
	uint8_t polykey[64];
	uint8_t zeros[64];
	c2fused_poly st;
	uint32_t ctr;
	size_t off;
	uint8_t lens[16];

	memset(zeros, 0, 64);
	c2chacha8_xor_blocks(polykey, zeros, 64, key, nonce, 0);
	c2fused_poly_init(&st, polykey);
	c2fused_poly_update(&st, ad, adn);
	poly_pad(&st, adn);

	ctr = 1;
	for (off = 0; off < n; off += 512) {
		size_t chunk = n - off;
		if (chunk > 512) {
			chunk = 512;
		}
		c2chacha8_xor_blocks(out + off, in + off, chunk, key, nonce, ctr);
		ctr += 8;
		c2fused_poly_update(&st, out + off, chunk);
	}
	poly_pad(&st, n);
	memset(lens, 0, 16);
	store64_le(lens, adn);
	store64_le(lens + 8, n);
	c2fused_poly_update(&st, lens, 16);
	c2fused_poly_final(&st, tag);
}

static void aead_fused(uint8_t *out, uint8_t tag[16],
                       const uint8_t *in, size_t n,
                       const uint8_t *ad, size_t adn,
                       const uint8_t key[32], const uint8_t nonce[12]) {
	uint8_t polykey[64];
	uint8_t zeros[64];
	c2fused_poly st;
	uint32_t ctr;
	size_t off;
	uint8_t lens[16];
	const uint8_t *prev;

	memset(zeros, 0, 64);
	c2chacha8_xor_blocks(polykey, zeros, 64, key, nonce, 0);
	c2fused_poly_init(&st, polykey);
	c2fused_poly_update(&st, ad, adn);
	poly_pad(&st, adn);

	ctr = 1;
	prev = NULL;
	for (off = 0; off < n; off += 512) {
		size_t chunk = n - off;
		if (chunk > 512) {
			chunk = 512;
		}
		if (prev != NULL && chunk == 512) {
			c2fused_seal_blocks(out + off, in + off, chunk, key, nonce, ctr, &st, prev);
		} else {
			if (prev != NULL) {
				c2fused_poly_blocks(&st, prev, 32);
			}
			c2chacha8_xor_blocks(out + off, in + off, chunk, key, nonce, ctr);
		}
		prev = out + off;
		if (chunk != 512) {
			c2fused_poly_update(&st, prev, chunk);
			prev = NULL;
		}
		ctr += 8;
	}
	if (prev != NULL) {
		c2fused_poly_blocks(&st, prev, 32);
	}
	poly_pad(&st, n);
	memset(lens, 0, 16);
	store64_le(lens, adn);
	store64_le(lens + 8, n);
	c2fused_poly_update(&st, lens, 16);
	c2fused_poly_final(&st, tag);
}

static int hex_eq(const uint8_t *got, const char *want, size_t n) {
	size_t i;
	for (i = 0; i < n; i++) {
		unsigned v;
		if (sscanf(want + 2 * i, "%2x", &v) != 1) {
			return 0;
		}
		if (got[i] != (uint8_t)v) {
			return 0;
		}
	}
	return 1;
}

int main(void) {
	/* RFC 8439 §2.5.2 */
	{
		static const uint8_t key[32] = {
			0x85, 0xd6, 0xbe, 0x78, 0x57, 0x55, 0x6d, 0x33,
			0x7f, 0x44, 0x52, 0xfe, 0x42, 0xd5, 0x06, 0xa8,
			0x01, 0x03, 0x80, 0x8a, 0xfb, 0x0d, 0xb2, 0xfd,
			0x4a, 0xbf, 0xf6, 0xaf, 0x41, 0x49, 0xf5, 0x1b};
		const char *msg = "Cryptographic Forum Research Group";
		uint8_t mac[16];
		c2fused_poly st;
		c2fused_poly_init(&st, key);
		c2fused_poly_update(&st, (const uint8_t *)msg, 34);
		c2fused_poly_final(&st, mac);
		printf("poly:");
		for (int i = 0; i < 16; i++) printf("%02x", mac[i]);
		printf("\n");
		if (!hex_eq(mac, "a8061dc1305136c6c22b8baf0c0127a9", 16)) {
			fprintf(stderr, "FAIL RFC 8439 §2.5.2\n");
			return 1;
		}
	}

	/* RFC 8439 §2.8.2 */
	{
		uint8_t key[32], nonce[12], ad[12], out[114], tag[16];
		const char *msg = "Ladies and Gentlemen of the class of '99: If I could offer you only one tip for the future, sunscreen would be it.";
		int i;
		for (i = 0; i < 32; i++) key[i] = (uint8_t)(0x80 + i);
		memset(nonce, 0, 12);
		nonce[0] = 0x07;
		nonce[4] = 0x40;
		nonce[5] = 0x41;
		nonce[6] = 0x42;
		nonce[7] = 0x43;
		nonce[8] = 0x44;
		nonce[9] = 0x45;
		nonce[10] = 0x46;
		nonce[11] = 0x47;
		ad[0] = 0x50;
		ad[1] = 0x51;
		ad[2] = 0x52;
		ad[3] = 0x53;
		ad[4] = 0xc0;
		ad[5] = 0xc1;
		ad[6] = 0xc2;
		ad[7] = 0xc3;
		ad[8] = 0xc4;
		ad[9] = 0xc5;
		ad[10] = 0xc6;
		ad[11] = 0xc7;
		aead_fused(out, tag, (const uint8_t *)msg, 114, ad, 12, key, nonce);
		printf("rfc:");
		for (i = 0; i < 114; i++) printf("%02x", out[i]);
		printf("\nrfctag:");
		for (i = 0; i < 16; i++) printf("%02x", tag[i]);
		printf("\n");
		{
			uint8_t out_s[114], tag_s[16];
			aead_seq(out_s, tag_s, (const uint8_t *)msg, 114, ad, 12, key, nonce);
			if (memcmp(out, out_s, 114) != 0 || memcmp(tag, tag_s, 16) != 0) {
				fprintf(stderr, "FAIL RFC 8439 §2.8.2 fused vs seq\n");
				return 1;
			}
		}
		if (!hex_eq(tag, "1ae10b594f09e26a7e902ecbd0600691", 16)) {
			fprintf(stderr, "FAIL RFC 8439 §2.8.2 tag\n");
			return 1;
		}
	}

	/* 4000 tirages : fusionné == séquentiel, chiffré == c2chacha8. */
	{
		enum { MAXN = 8192 };
		uint8_t *in = malloc(MAXN);
		uint8_t *out_f = malloc(MAXN);
		uint8_t *out_s = malloc(MAXN);
		uint8_t *out_c = malloc(MAXN);
		uint8_t key[32], nonce[12], ad[13], tag_f[16], tag_s[16];
		int draw;
		if (!in || !out_f || !out_s || !out_c) {
			fprintf(stderr, "malloc\n");
			return 1;
		}
		for (draw = 0; draw < 4000; draw++) {
			size_t n, adn, k;
			uint32_t ctr;
			lcg_seed((uint64_t)draw + 1);
			n = (size_t)(1 + (lcg_byte() % 16)) * 512;
			if (n > MAXN) {
				n = MAXN;
			}
			adn = (size_t)(lcg_byte() % 14);
			for (k = 0; k < 32; k++) key[k] = lcg_byte();
			for (k = 0; k < 12; k++) nonce[k] = lcg_byte();
			for (k = 0; k < adn; k++) ad[k] = lcg_byte();
			for (k = 0; k < n; k++) in[k] = lcg_byte();

			aead_fused(out_f, tag_f, in, n, ad, adn, key, nonce);
			aead_seq(out_s, tag_s, in, n, ad, adn, key, nonce);
			if (memcmp(out_f, out_s, n) != 0 || memcmp(tag_f, tag_s, 16) != 0) {
				fprintf(stderr, "FAIL draw %d n=%zu fused vs seq\n", draw, n);
				return 1;
			}
			ctr = 1;
			for (k = 0; k < n; k += 512) {
				size_t chunk = n - k;
				if (chunk > 512) {
					chunk = 512;
				}
				c2chacha8_xor_blocks(out_c + k, in + k, chunk, key, nonce, ctr);
				ctr += 8;
			}
			if (memcmp(out_f, out_c, n) != 0) {
				fprintf(stderr, "FAIL draw %d n=%zu vs c2chacha8\n", draw, n);
				return 1;
			}
		}
		free(in);
		free(out_f);
		free(out_s);
		free(out_c);
		printf("draws:4000 ok\n");
	}
	return 0;
}
