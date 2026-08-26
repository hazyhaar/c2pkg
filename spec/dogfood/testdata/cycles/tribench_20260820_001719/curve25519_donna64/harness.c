#include <stdint.h>
#include <stddef.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

/* kernel symbols */
void curve25519_f51_mul121666(uint64_t out[5], const uint64_t in[5]);

static void hex_u64(uint64_t v) { printf("%016llx", (unsigned long long)v); }
static void hex_u32(uint32_t v) { printf("%08x", v); }
static void hex_buf(const uint8_t *p, size_t n) {
  size_t i; for (i = 0; i < n; i++) printf("%02x", p[i]);
}

int main(void) {
  /* fixture zero */
  /* fixture pattern */
  printf("zero ");
  { uint64_t in[5]={0x123456789ULL, 0x23456789aULL, 0x3456789abULL, 0x456789abcULL, 0x56789abcdULL};
    uint64_t out[5]={0};
    curve25519_f51_mul121666(out, in);
    int j; for(j=0;j<5;j++) hex_u64(out[j]); printf("\n"); }
  printf("pattern ");
  { uint64_t in[5]={0x123456789ULL, 0x23456789aULL, 0x3456789abULL, 0x456789abcULL, 0x56789abcdULL};
    uint64_t out[5]={0};
    curve25519_f51_mul121666(out, in);
    int j; for(j=0;j<5;j++) hex_u64(out[j]); printf("\n"); }
  return 0;
}
