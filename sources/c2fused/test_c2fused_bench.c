/* Banc C : fusionné vs séquentiel (c2chacha8 + Poly 64 bits) à 8192 octets.
 * Cinq brutes, CLOCK_MONOTONIC, même harnais. */
#include <stdio.h>
#include <string.h>
#include <stdint.h>
#include <time.h>
#include "c2fused.h"
#include "c2chacha8.h"

#define N 8192
#define CHUNKS 16
#define ITER 80000

static uint64_t nsec_now(void) {
	struct timespec ts;
	clock_gettime(CLOCK_MONOTONIC, &ts);
	return (uint64_t)ts.tv_sec * 1000000000ull + (uint64_t)ts.tv_nsec;
}

static void run_seq(uint8_t *out, const uint8_t *in,
                    const uint8_t key[32], const uint8_t nonce[12],
                    const uint8_t polykey[32]) {
	c2fused_poly st;
	uint32_t ctr = 1;
	int i;
	c2fused_poly_init(&st, polykey);
	for (i = 0; i < CHUNKS; i++) {
		c2chacha8_xor_blocks(out + i * 512, in + i * 512, 512, key, nonce, ctr);
		c2fused_poly_blocks(&st, out + i * 512, 32);
		ctr += 8;
	}
}

static void run_fused(uint8_t *out, const uint8_t *in,
                      const uint8_t key[32], const uint8_t nonce[12],
                      const uint8_t polykey[32]) {
	c2fused_poly st;
	uint32_t ctr = 1;
	int i;
	c2fused_poly_init(&st, polykey);
	c2chacha8_xor_blocks(out, in, 512, key, nonce, ctr);
	ctr += 8;
	for (i = 1; i < CHUNKS; i++) {
		c2fused_seal_blocks(out + i * 512, in + i * 512, 512, key, nonce, ctr, &st,
		                    out + (i - 1) * 512);
		ctr += 8;
	}
	c2fused_poly_blocks(&st, out + (CHUNKS - 1) * 512, 32);
}

static double brute(void (*fn)(uint8_t *, const uint8_t *, const uint8_t *,
                               const uint8_t *, const uint8_t *),
                    uint8_t *out, const uint8_t *in,
                    const uint8_t key[32], const uint8_t nonce[12],
                    const uint8_t polykey[32]) {
	int i;
	uint64_t t0, t1;
	for (i = 0; i < 200; i++) {
		fn(out, in, key, nonce, polykey);
	}
	t0 = nsec_now();
	for (i = 0; i < ITER; i++) {
		fn(out, in, key, nonce, polykey);
	}
	t1 = nsec_now();
	return (double)(t1 - t0) / (double)ITER;
}

int main(void) {
	static uint8_t in[N], out_s[N], out_f[N];
	uint8_t key[32], nonce[12], polykey[32];
	int i;
	double seq[5], fused[5];

	for (i = 0; i < 32; i++) {
		key[i] = (uint8_t)i;
		polykey[i] = (uint8_t)(i * 3 + 1);
	}
	memset(nonce, 0, 12);
	nonce[0] = 1;
	for (i = 0; i < N; i++) {
		in[i] = (uint8_t)(i * 7);
	}

	run_seq(out_s, in, key, nonce, polykey);
	run_fused(out_f, in, key, nonce, polykey);
	if (memcmp(out_s, out_f, N) != 0) {
		fprintf(stderr, "FAIL bench parité chiffré\n");
		return 1;
	}

	printf("ns_per_op seq fused\n");
	for (i = 0; i < 5; i++) {
		seq[i] = brute(run_seq, out_s, in, key, nonce, polykey);
		fused[i] = brute(run_fused, out_f, in, key, nonce, polykey);
		printf("brute%d %.1f %.1f\n", i + 1, seq[i], fused[i]);
		fflush(stdout);
	}
	return 0;
}
