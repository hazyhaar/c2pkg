#include <stdint.h>
#include <stddef.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

/* kernel symbols */
void poly1305_block5(uint32_t h[5], const uint32_t r[5], const uint32_t m[4]);

static void hex_u64(uint64_t v) { printf("%016llx", (unsigned long long)v); }
static void hex_u32(uint32_t v) { printf("%08x", v); }
static void hex_buf(const uint8_t *p, size_t n) {
  size_t i; for (i = 0; i < n; i++) printf("%02x", p[i]);
}

int main(void) {
  /* fixture zero */
  /* fixture pattern */
  printf("zero ");
  { uint32_t h[5]={0}, r[5]={1,2,3,4,5}, m[4]={9,8,7,6};
    poly1305_block5(h,r,m); int j; for(j=0;j<5;j++) hex_u32(h[j]); printf("\n"); }
  printf("pattern ");
  { uint32_t h[5]={0}, r[5]={1,2,3,4,5}, m[4]={9,8,7,6};
    poly1305_block5(h,r,m); int j; for(j=0;j<5;j++) hex_u32(h[j]); printf("\n"); }
  return 0;
}
