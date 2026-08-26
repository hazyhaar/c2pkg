#include <stdint.h>
#include <stddef.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

/* kernel symbols */
void blake2b_compress_block(uint64_t h[8], const uint8_t block[128], uint64_t t0, uint64_t t1, uint64_t f0, uint64_t f1);

static void hex_u64(uint64_t v) { printf("%016llx", (unsigned long long)v); }
static void hex_u32(uint32_t v) { printf("%08x", v); }
static void hex_buf(const uint8_t *p, size_t n) {
  size_t i; for (i = 0; i < n; i++) printf("%02x", p[i]);
}

int main(void) {
  /* fixture zero */
  static const uint8_t d0[1] = {0};
  /* fixture pattern */
  static const uint8_t d1[64] = {
    0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,
    0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,
    0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,
    0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5,0xa5};
  printf("zero ");
  { uint64_t h[8] = {0}; uint8_t block[128]; memset(block, 0, 128);
    size_t n = 0; if (n > 128) n = 128; if (n) memcpy(block, d0, n);
    blake2b_compress_block(h, block, 0, 0, 0xffffffffffffffffULL, 0xffffffffffffffffULL);
    int j; for (j=0;j<8;j++) hex_u64(h[j]); printf("\n"); }
  printf("pattern ");
  { uint64_t h[8] = {0}; uint8_t block[128]; memset(block, 0, 128);
    size_t n = 64; if (n > 128) n = 128; if (n) memcpy(block, d1, n);
    blake2b_compress_block(h, block, 0, 0, 0xffffffffffffffffULL, 0xffffffffffffffffULL);
    int j; for (j=0;j<8;j++) hex_u64(h[j]); printf("\n"); }
  return 0;
}
