#include <stdint.h>
#include <stddef.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

/* kernel symbols */
void md5_transform_block(uint32_t state[4], const uint8_t block[64]);

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
  { uint32_t st[4] = {0x67452301,0xefcdab89,0x98badcfe,0x10325476};
    uint8_t block[64]; memset(block,0,64); size_t n=0; if(n>64)n=64; if(n)memcpy(block,d0,n);
    md5_transform_block(st, block); int j; for(j=0;j<4;j++) hex_u32(st[j]); printf("\n"); }
  printf("pattern ");
  { uint32_t st[4] = {0x67452301,0xefcdab89,0x98badcfe,0x10325476};
    uint8_t block[64]; memset(block,0,64); size_t n=64; if(n>64)n=64; if(n)memcpy(block,d1,n);
    md5_transform_block(st, block); int j; for(j=0;j<4;j++) hex_u32(st[j]); printf("\n"); }
  return 0;
}
