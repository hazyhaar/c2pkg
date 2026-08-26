/* Oracle C gcc -O2 pour c2pkg/c2poly1305x2.
 * Ligne RFC 8439 §2.5.2 puis 24 tailles LCG (même graine que
 * sources/test_poly1305_oracle.c). */
#include <stdint.h>
#include <stdio.h>
#include <string.h>
#include "c2poly1305x2.h"

static uint64_t lcg_state;
static void lcg_seed(uint64_t s) { lcg_state = s * 0x9E3779B97F4A7C15ULL + 1; }
static uint8_t lcg_byte(void)
{
	lcg_state = lcg_state * 6364136223846793005ULL + 1442695040888963407ULL;
	return (uint8_t)(lcg_state >> 56);
}

int main(void)
{
	static const uint8_t rfc_key[32] = {
		0x85, 0xd6, 0xbe, 0x78, 0x57, 0x55, 0x6d, 0x33,
		0x7f, 0x44, 0x52, 0xfe, 0x42, 0xd5, 0x06, 0xa8,
		0x01, 0x03, 0x80, 0x8a, 0xfb, 0x0d, 0xb2, 0xfd,
		0x4a, 0xbf, 0xf6, 0xaf, 0x41, 0x49, 0xf5, 0x1b
	};
	static const char *rfc_msg = "Cryptographic Forum Research Group";
	uint8_t mac[16];
	int k;

	crypto_poly1305x2(mac, (const uint8_t *)rfc_msg, 34, rfc_key);
	printf("rfc:");
	for (k = 0; k < 16; k++) {
		printf("%02x", mac[k]);
	}
	printf("\n");

	{
		static const size_t sizes[] = {
			0, 1, 15, 16, 17, 31, 32, 33, 47, 48, 63, 64, 65, 96,
			127, 128, 129, 255, 256, 1000, 1023, 1024, 1025, 4096
		};
		static uint8_t msg[4096];
		uint8_t key[32];
		size_t i;
		for (i = 0; i < sizeof sizes / sizeof sizes[0]; i++) {
			size_t n = sizes[i];
			size_t j;
			lcg_seed((uint64_t)n);
			for (k = 0; k < 32; k++) {
				key[k] = lcg_byte();
			}
			for (j = 0; j < n; j++) {
				msg[j] = lcg_byte();
			}
			crypto_poly1305x2(mac, msg, n, key);
			printf("%zu:", n);
			for (k = 0; k < 16; k++) {
				printf("%02x", mac[k]);
			}
			printf("\n");
		}
	}
	return 0;
}
