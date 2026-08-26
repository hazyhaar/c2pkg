#include <immintrin.h>
#include "c2fused.h"
#include "c2fused_schedule.h"

typedef unsigned __int128 u128;

static uint32_t load32_le(const uint8_t *p) {
	return (uint32_t)p[0]
		| ((uint32_t)p[1] << 8)
		| ((uint32_t)p[2] << 16)
		| ((uint32_t)p[3] << 24);
}

static uint64_t load64_le(const uint8_t *p) {
	return (uint64_t)p[0]
		| ((uint64_t)p[1] << 8)
		| ((uint64_t)p[2] << 16)
		| ((uint64_t)p[3] << 24)
		| ((uint64_t)p[4] << 32)
		| ((uint64_t)p[5] << 40)
		| ((uint64_t)p[6] << 48)
		| ((uint64_t)p[7] << 56);
}

static void store64_le(uint8_t *p, uint64_t v) {
	p[0] = (uint8_t)v;
	p[1] = (uint8_t)(v >> 8);
	p[2] = (uint8_t)(v >> 16);
	p[3] = (uint8_t)(v >> 24);
	p[4] = (uint8_t)(v >> 32);
	p[5] = (uint8_t)(v >> 40);
	p[6] = (uint8_t)(v >> 48);
	p[7] = (uint8_t)(v >> 56);
}

/* Un bloc de 16 octets : h = (h + m + hibit·2^128) * r  mod (2^130−5).
 * hibit = 1 pour un bloc plein, 0 pour le bloc partiel final (octet 0x01 déjà
 * posé dans m). Trois membres 64 bits, six mulq via __uint128_t. */
__attribute__((always_inline)) static inline void
c2fused_poly_step(uint64_t *h0, uint64_t *h1, uint64_t *h2,
                  uint64_t r0, uint64_t r1, const uint8_t *m, uint64_t hibit) {
	uint64_t t0 = load64_le(m);
	uint64_t t1 = load64_le(m + 8);

	u128 s = (u128)(*h0) + t0;
	uint64_t n0 = (uint64_t)s;
	s = (u128)(*h1) + t1 + (s >> 64);
	uint64_t n1 = (uint64_t)s;
	uint64_t n2 = *h2 + (uint64_t)(s >> 64) + hibit;

	u128 p00 = (u128)n0 * r0;
	u128 p01 = (u128)n0 * r1;
	u128 p10 = (u128)n1 * r0;
	u128 p11 = (u128)n1 * r1;
	u128 p20 = (u128)n2 * r0;
	u128 p21 = (u128)n2 * r1;

	u128 acc = (p00 >> 64) + p01 + p10;
	uint64_t z0 = (uint64_t)p00;
	uint64_t z1 = (uint64_t)acc;
	acc = (acc >> 64) + p11 + p20;
	uint64_t z2 = (uint64_t)acc;
	acc = (acc >> 64) + p21;
	uint64_t z3 = (uint64_t)acc;
	uint64_t z4 = (uint64_t)(acc >> 64);

	u128 five_qlo = (u128)5 * (((u128)(z2 >> 2)) | ((u128)(z3 & 3) << 62));
	uint64_t qhi = (z3 >> 2) + (z4 << 62);
	u128 five_qhi = (u128)5 * qhi;

	u128 sum = ((u128)z0 | ((u128)z1 << 64)) + five_qlo;
	uint64_t s0 = (uint64_t)sum;
	u128 sum_hi = (sum >> 64) + five_qhi;
	*h0 = s0;
	*h1 = (uint64_t)sum_hi;
	*h2 = (uint64_t)(sum_hi >> 64) + (z2 & 3);
}

void c2fused_poly_init(c2fused_poly *st, const uint8_t key[32]) {
	st->r[0] = load64_le(key) & 0x0FFFFFFC0FFFFFFFULL;
	st->r[1] = load64_le(key + 8) & 0x0FFFFFFC0FFFFFFCULL;
	st->pad[0] = load64_le(key + 16);
	st->pad[1] = load64_le(key + 24);
	st->h[0] = 0;
	st->h[1] = 0;
	st->h[2] = 0;
	st->leftover_len = 0;
}

void c2fused_poly_blocks(c2fused_poly *st, const uint8_t *m, size_t nblocks) {
	uint64_t h0 = st->h[0];
	uint64_t h1 = st->h[1];
	uint64_t h2 = st->h[2];
	uint64_t r0 = st->r[0];
	uint64_t r1 = st->r[1];
	size_t i;
	for (i = 0; i < nblocks; i++) {
		c2fused_poly_step(&h0, &h1, &h2, r0, r1, m + i * 16, 1);
	}
	st->h[0] = h0;
	st->h[1] = h1;
	st->h[2] = h2;
}

void c2fused_poly_update(c2fused_poly *st, const uint8_t *m, size_t n) {
	size_t i;
	if (n == 0) {
		return;
	}
	if (st->leftover_len != 0) {
		size_t need = 16 - st->leftover_len;
		if (need > n) {
			need = n;
		}
		for (i = 0; i < need; i++) {
			st->leftover[st->leftover_len + i] = m[i];
		}
		st->leftover_len += need;
		m += need;
		n -= need;
		if (st->leftover_len == 16) {
			c2fused_poly_blocks(st, st->leftover, 1);
			st->leftover_len = 0;
		}
	}
	if (n >= 16) {
		size_t nb = n / 16;
		c2fused_poly_blocks(st, m, nb);
		m += nb * 16;
		n -= nb * 16;
	}
	for (i = 0; i < n; i++) {
		st->leftover[st->leftover_len + i] = m[i];
	}
	st->leftover_len += n;
}

void c2fused_poly_final(c2fused_poly *st, uint8_t mac[16]) {
	uint64_t h0 = st->h[0];
	uint64_t h1 = st->h[1];
	uint64_t h2 = st->h[2];
	size_t i;

	if (st->leftover_len != 0) {
		st->leftover[st->leftover_len] = 1;
		for (i = st->leftover_len + 1; i < 16; i++) {
			st->leftover[i] = 0;
		}
		c2fused_poly_step(&h0, &h1, &h2, st->r[0], st->r[1], st->leftover, 0);
	}

	{
		u128 t = (u128)h0 + 5;
		uint64_t g0 = (uint64_t)t;
		t = (u128)h1 + (t >> 64);
		uint64_t g1 = (uint64_t)t;
		uint64_t g2 = h2 + (uint64_t)(t >> 64);
		uint64_t mask = 0 - (uint64_t)(g2 >> 2);
		h0 = (h0 & ~mask) | (g0 & mask);
		h1 = (h1 & ~mask) | (g1 & mask);
	}

	{
		u128 f = (u128)h0 + st->pad[0];
		store64_le(mac, (uint64_t)f);
		f = (u128)h1 + st->pad[1] + (f >> 64);
		store64_le(mac + 8, (uint64_t)f);
	}

	st->h[0] = 0;
	st->h[1] = 0;
	st->h[2] = 0;
	st->r[0] = 0;
	st->r[1] = 0;
	st->pad[0] = 0;
	st->pad[1] = 0;
	st->leftover_len = 0;
	for (i = 0; i < 16; i++) {
		st->leftover[i] = 0;
	}
}

static const uint8_t c2fused_rot16_mask[32] = {
	2, 3, 0, 1, 6, 7, 4, 5, 10, 11, 8, 9, 14, 15, 12, 13,
	2, 3, 0, 1, 6, 7, 4, 5, 10, 11, 8, 9, 14, 15, 12, 13};
static const uint8_t c2fused_rot8_mask[32] = {
	3, 0, 1, 2, 7, 4, 5, 6, 11, 8, 9, 10, 15, 12, 13, 14,
	3, 0, 1, 2, 7, 4, 5, 6, 11, 8, 9, 10, 15, 12, 13, 14};
static const uint8_t c2fused_sigma[16] = {0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x20, 0x33,
                                          0x32, 0x2d, 0x62, 0x79, 0x74, 0x65, 0x20, 0x6b};
static const uint32_t c2fused_lane_ctr[8] = {0, 1, 2, 3, 4, 5, 6, 7};

#define C2FUSED_QR(a, b, c, d) do { \
	(a) = _mm256_add_epi32((a), (b)); \
	(d) = _mm256_xor_si256((d), (a)); \
	(d) = _mm256_shuffle_epi8((d), rot16); \
	(c) = _mm256_add_epi32((c), (d)); \
	(b) = _mm256_xor_si256((b), (c)); \
	(b) = _mm256_or_si256(_mm256_slli_epi32((b), 12), _mm256_srli_epi32((b), 20)); \
	(a) = _mm256_add_epi32((a), (b)); \
	(d) = _mm256_xor_si256((d), (a)); \
	(d) = _mm256_shuffle_epi8((d), rot8); \
	(c) = _mm256_add_epi32((c), (d)); \
	(b) = _mm256_xor_si256((b), (c)); \
	(b) = _mm256_or_si256(_mm256_slli_epi32((b), 7), _mm256_srli_epi32((b), 25)); \
} while (0)

#define C2FUSED_MAYBE_POLY(n) do { \
	const int _s = (int)c2fused_poly_slot[(n)]; \
	if (_s >= 0) { \
		c2fused_poly_step(&h0, &h1, &h2, r0, r1, prev + 16 * (size_t)_s, 1); \
	} \
} while (0)

#define C2FUSED_DOUBLE_ROUND(B) do { \
	C2FUSED_QR(x0, x4, x8, x12); C2FUSED_MAYBE_POLY((B) + 0); \
	C2FUSED_QR(x1, x5, x9, x13); C2FUSED_MAYBE_POLY((B) + 1); \
	C2FUSED_QR(x2, x6, x10, x14); C2FUSED_MAYBE_POLY((B) + 2); \
	C2FUSED_QR(x3, x7, x11, x15); C2FUSED_MAYBE_POLY((B) + 3); \
	C2FUSED_QR(x0, x5, x10, x15); C2FUSED_MAYBE_POLY((B) + 4); \
	C2FUSED_QR(x1, x6, x11, x12); C2FUSED_MAYBE_POLY((B) + 5); \
	C2FUSED_QR(x2, x7, x8, x13); C2FUSED_MAYBE_POLY((B) + 6); \
	C2FUSED_QR(x3, x4, x9, x14); C2FUSED_MAYBE_POLY((B) + 7); \
} while (0)

size_t c2fused_seal_blocks(uint8_t *out, const uint8_t *in, size_t n,
                           const uint8_t key[32], const uint8_t nonce[12],
                           uint32_t counter, c2fused_poly *st,
                           const uint8_t *prev) {
	if (n > 512) {
		n = 512;
	}
	if (n == 0) {
		return 0;
	}

	__m256i rot16 = _mm256_loadu_si256((const __m256i *)c2fused_rot16_mask);
	__m256i rot8 = _mm256_loadu_si256((const __m256i *)c2fused_rot8_mask);

	__m256i x0 = _mm256_set1_epi32((int)load32_le(c2fused_sigma + 0));
	__m256i x1 = _mm256_set1_epi32((int)load32_le(c2fused_sigma + 4));
	__m256i x2 = _mm256_set1_epi32((int)load32_le(c2fused_sigma + 8));
	__m256i x3 = _mm256_set1_epi32((int)load32_le(c2fused_sigma + 12));
	__m256i x4 = _mm256_set1_epi32((int)load32_le(key + 0));
	__m256i x5 = _mm256_set1_epi32((int)load32_le(key + 4));
	__m256i x6 = _mm256_set1_epi32((int)load32_le(key + 8));
	__m256i x7 = _mm256_set1_epi32((int)load32_le(key + 12));
	__m256i x8 = _mm256_set1_epi32((int)load32_le(key + 16));
	__m256i x9 = _mm256_set1_epi32((int)load32_le(key + 20));
	__m256i x10 = _mm256_set1_epi32((int)load32_le(key + 24));
	__m256i x11 = _mm256_set1_epi32((int)load32_le(key + 28));
	__m256i x12 = _mm256_add_epi32(_mm256_set1_epi32((int)counter),
	                               _mm256_loadu_si256((const __m256i *)c2fused_lane_ctr));
	__m256i x13 = _mm256_set1_epi32((int)load32_le(nonce + 0));
	__m256i x14 = _mm256_set1_epi32((int)load32_le(nonce + 4));
	__m256i x15 = _mm256_set1_epi32((int)load32_le(nonce + 8));

	uint64_t h0 = st->h[0];
	uint64_t h1 = st->h[1];
	uint64_t h2 = st->h[2];
	uint64_t r0 = st->r[0];
	uint64_t r1 = st->r[1];

	C2FUSED_DOUBLE_ROUND(0);
	C2FUSED_DOUBLE_ROUND(8);
	C2FUSED_DOUBLE_ROUND(16);
	C2FUSED_DOUBLE_ROUND(24);
	C2FUSED_DOUBLE_ROUND(32);
	C2FUSED_DOUBLE_ROUND(40);
	C2FUSED_DOUBLE_ROUND(48);
	C2FUSED_DOUBLE_ROUND(56);
	C2FUSED_DOUBLE_ROUND(64);
	C2FUSED_DOUBLE_ROUND(72);

	st->h[0] = h0;
	st->h[1] = h1;
	st->h[2] = h2;

	x0 = _mm256_add_epi32(x0, _mm256_set1_epi32((int)load32_le(c2fused_sigma + 0)));
	x1 = _mm256_add_epi32(x1, _mm256_set1_epi32((int)load32_le(c2fused_sigma + 4)));
	x2 = _mm256_add_epi32(x2, _mm256_set1_epi32((int)load32_le(c2fused_sigma + 8)));
	x3 = _mm256_add_epi32(x3, _mm256_set1_epi32((int)load32_le(c2fused_sigma + 12)));
	x4 = _mm256_add_epi32(x4, _mm256_set1_epi32((int)load32_le(key + 0)));
	x5 = _mm256_add_epi32(x5, _mm256_set1_epi32((int)load32_le(key + 4)));
	x6 = _mm256_add_epi32(x6, _mm256_set1_epi32((int)load32_le(key + 8)));
	x7 = _mm256_add_epi32(x7, _mm256_set1_epi32((int)load32_le(key + 12)));
	x8 = _mm256_add_epi32(x8, _mm256_set1_epi32((int)load32_le(key + 16)));
	x9 = _mm256_add_epi32(x9, _mm256_set1_epi32((int)load32_le(key + 20)));
	x10 = _mm256_add_epi32(x10, _mm256_set1_epi32((int)load32_le(key + 24)));
	x11 = _mm256_add_epi32(x11, _mm256_set1_epi32((int)load32_le(key + 28)));
	x12 = _mm256_add_epi32(x12, _mm256_add_epi32(_mm256_set1_epi32((int)counter),
	                                            _mm256_loadu_si256((const __m256i *)c2fused_lane_ctr)));
	x13 = _mm256_add_epi32(x13, _mm256_set1_epi32((int)load32_le(nonce + 0)));
	x14 = _mm256_add_epi32(x14, _mm256_set1_epi32((int)load32_le(nonce + 4)));
	x15 = _mm256_add_epi32(x15, _mm256_set1_epi32((int)load32_le(nonce + 8)));

	__m256i t0 = _mm256_unpacklo_epi32(x0, x1);
	__m256i t1 = _mm256_unpackhi_epi32(x0, x1);
	__m256i t2 = _mm256_unpacklo_epi32(x2, x3);
	__m256i t3 = _mm256_unpackhi_epi32(x2, x3);
	__m256i t4 = _mm256_unpacklo_epi32(x4, x5);
	__m256i t5 = _mm256_unpackhi_epi32(x4, x5);
	__m256i t6 = _mm256_unpacklo_epi32(x6, x7);
	__m256i t7 = _mm256_unpackhi_epi32(x6, x7);
	__m256i u0 = _mm256_unpacklo_epi64(t0, t2);
	__m256i u1 = _mm256_unpackhi_epi64(t0, t2);
	__m256i u2 = _mm256_unpacklo_epi64(t1, t3);
	__m256i u3 = _mm256_unpackhi_epi64(t1, t3);
	__m256i u4 = _mm256_unpacklo_epi64(t4, t6);
	__m256i u5 = _mm256_unpackhi_epi64(t4, t6);
	__m256i u6 = _mm256_unpacklo_epi64(t5, t7);
	__m256i u7 = _mm256_unpackhi_epi64(t5, t7);
	__m256i blo0 = _mm256_permute2x128_si256(u0, u4, 0x20);
	__m256i blo1 = _mm256_permute2x128_si256(u1, u5, 0x20);
	__m256i blo2 = _mm256_permute2x128_si256(u2, u6, 0x20);
	__m256i blo3 = _mm256_permute2x128_si256(u3, u7, 0x20);
	__m256i blo4 = _mm256_permute2x128_si256(u0, u4, 0x31);
	__m256i blo5 = _mm256_permute2x128_si256(u1, u5, 0x31);
	__m256i blo6 = _mm256_permute2x128_si256(u2, u6, 0x31);
	__m256i blo7 = _mm256_permute2x128_si256(u3, u7, 0x31);

	t0 = _mm256_unpacklo_epi32(x8, x9);
	t1 = _mm256_unpackhi_epi32(x8, x9);
	t2 = _mm256_unpacklo_epi32(x10, x11);
	t3 = _mm256_unpackhi_epi32(x10, x11);
	t4 = _mm256_unpacklo_epi32(x12, x13);
	t5 = _mm256_unpackhi_epi32(x12, x13);
	t6 = _mm256_unpacklo_epi32(x14, x15);
	t7 = _mm256_unpackhi_epi32(x14, x15);
	u0 = _mm256_unpacklo_epi64(t0, t2);
	u1 = _mm256_unpackhi_epi64(t0, t2);
	u2 = _mm256_unpacklo_epi64(t1, t3);
	u3 = _mm256_unpackhi_epi64(t1, t3);
	u4 = _mm256_unpacklo_epi64(t4, t6);
	u5 = _mm256_unpackhi_epi64(t4, t6);
	u6 = _mm256_unpacklo_epi64(t5, t7);
	u7 = _mm256_unpackhi_epi64(t5, t7);
	__m256i bhi0 = _mm256_permute2x128_si256(u0, u4, 0x20);
	__m256i bhi1 = _mm256_permute2x128_si256(u1, u5, 0x20);
	__m256i bhi2 = _mm256_permute2x128_si256(u2, u6, 0x20);
	__m256i bhi3 = _mm256_permute2x128_si256(u3, u7, 0x20);
	__m256i bhi4 = _mm256_permute2x128_si256(u0, u4, 0x31);
	__m256i bhi5 = _mm256_permute2x128_si256(u1, u5, 0x31);
	__m256i bhi6 = _mm256_permute2x128_si256(u2, u6, 0x31);
	__m256i bhi7 = _mm256_permute2x128_si256(u3, u7, 0x31);

	size_t nfull = n / 64;
	size_t rem = n % 64;

	if (nfull >= 1) {
		_mm256_storeu_si256((__m256i *)(out + 0), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 0)), blo0));
		_mm256_storeu_si256((__m256i *)(out + 32), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 32)), bhi0));
	}
	if (nfull >= 2) {
		_mm256_storeu_si256((__m256i *)(out + 64), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 64)), blo1));
		_mm256_storeu_si256((__m256i *)(out + 96), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 96)), bhi1));
	}
	if (nfull >= 3) {
		_mm256_storeu_si256((__m256i *)(out + 128), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 128)), blo2));
		_mm256_storeu_si256((__m256i *)(out + 160), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 160)), bhi2));
	}
	if (nfull >= 4) {
		_mm256_storeu_si256((__m256i *)(out + 192), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 192)), blo3));
		_mm256_storeu_si256((__m256i *)(out + 224), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 224)), bhi3));
	}
	if (nfull >= 5) {
		_mm256_storeu_si256((__m256i *)(out + 256), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 256)), blo4));
		_mm256_storeu_si256((__m256i *)(out + 288), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 288)), bhi4));
	}
	if (nfull >= 6) {
		_mm256_storeu_si256((__m256i *)(out + 320), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 320)), blo5));
		_mm256_storeu_si256((__m256i *)(out + 352), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 352)), bhi5));
	}
	if (nfull >= 7) {
		_mm256_storeu_si256((__m256i *)(out + 384), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 384)), blo6));
		_mm256_storeu_si256((__m256i *)(out + 416), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 416)), bhi6));
	}
	if (nfull >= 8) {
		_mm256_storeu_si256((__m256i *)(out + 448), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 448)), blo7));
		_mm256_storeu_si256((__m256i *)(out + 480), _mm256_xor_si256(_mm256_loadu_si256((const __m256i *)(in + 480)), bhi7));
	}

	if (rem != 0) {
		uint8_t ks[64];
		switch (nfull) {
		case 0:
			_mm256_storeu_si256((__m256i *)(ks + 0), blo0);
			_mm256_storeu_si256((__m256i *)(ks + 32), bhi0);
			break;
		case 1:
			_mm256_storeu_si256((__m256i *)(ks + 0), blo1);
			_mm256_storeu_si256((__m256i *)(ks + 32), bhi1);
			break;
		case 2:
			_mm256_storeu_si256((__m256i *)(ks + 0), blo2);
			_mm256_storeu_si256((__m256i *)(ks + 32), bhi2);
			break;
		case 3:
			_mm256_storeu_si256((__m256i *)(ks + 0), blo3);
			_mm256_storeu_si256((__m256i *)(ks + 32), bhi3);
			break;
		case 4:
			_mm256_storeu_si256((__m256i *)(ks + 0), blo4);
			_mm256_storeu_si256((__m256i *)(ks + 32), bhi4);
			break;
		case 5:
			_mm256_storeu_si256((__m256i *)(ks + 0), blo5);
			_mm256_storeu_si256((__m256i *)(ks + 32), bhi5);
			break;
		case 6:
			_mm256_storeu_si256((__m256i *)(ks + 0), blo6);
			_mm256_storeu_si256((__m256i *)(ks + 32), bhi6);
			break;
		default:
			_mm256_storeu_si256((__m256i *)(ks + 0), blo7);
			_mm256_storeu_si256((__m256i *)(ks + 32), bhi7);
			break;
		}
		size_t base = nfull * 64;
		size_t k;
		for (k = 0; k < rem; k++) {
			out[base + k] = in[base + k] ^ ks[k];
		}
	}
	return n;
}
