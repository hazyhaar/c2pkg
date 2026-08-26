#include <stdint.h>
#include <stddef.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

/* kernel symbols */
void chacha20_quarter_round(uint32_t *a, uint32_t *b, uint32_t *c, uint32_t *d);

static void hex_u64(uint64_t v) { printf("%016llx", (unsigned long long)v); }
static void hex_u32(uint32_t v) { printf("%08x", v); }
static void hex_buf(const uint8_t *p, size_t n) {
  size_t i; for (i = 0; i < n; i++) printf("%02x", p[i]);
}

int main(void) {
  /* fixture zero */
  /* fixture pattern */
  printf("zero ");
  { uint32_t a=0x11111111,b=0x22222222,c=0x33333333,d=0x44444444;
    chacha20_quarter_round(&a,&b,&c,&d); hex_u32(a); hex_u32(b); hex_u32(c); hex_u32(d); printf("\n"); }
  printf("pattern ");
  { uint32_t a=0x11111111,b=0x22222222,c=0x33333333,d=0x44444444;
    chacha20_quarter_round(&a,&b,&c,&d); hex_u32(a); hex_u32(b); hex_u32(c); hex_u32(d); printf("\n"); }
  return 0;
}
