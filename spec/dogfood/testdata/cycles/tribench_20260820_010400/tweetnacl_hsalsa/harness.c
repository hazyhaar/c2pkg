#include <stdint.h>
#include <stddef.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

/* kernel symbols */
int crypto_core_hsalsa20(uint8_t *out, const uint8_t *in, const uint8_t *k, const uint8_t *c);

static void hex_u64(uint64_t v) { printf("%016llx", (unsigned long long)v); }
static void hex_u32(uint32_t v) { printf("%08x", v); }
static void hex_buf(const uint8_t *p, size_t n) {
  size_t i; for (i = 0; i < n; i++) printf("%02x", p[i]);
}

int main(void) {
  /* fixture zero */
  /* fixture pattern */
  printf("zero ");
  { uint8_t in[16]={1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16};
    uint8_t k[32]={0x10,0x20,0x30,0x40,0x50,0x60,0x70,0x80,0x90,0xa0,0xb0,0xc0,0xd0,0xe0,0xf0,0x01,
                   0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x09,0x0a,0x0b,0x0c,0x0d,0x0e,0x0f,0x11,0x22};
    uint8_t c[16]="expand 32-byte k";
    uint8_t out[32]={0};
    crypto_core_hsalsa20(out, in, k, c);
    hex_buf(out, 32); printf("\n"); }
  printf("pattern ");
  { uint8_t in[16]={1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16};
    uint8_t k[32]={0x10,0x20,0x30,0x40,0x50,0x60,0x70,0x80,0x90,0xa0,0xb0,0xc0,0xd0,0xe0,0xf0,0x01,
                   0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x09,0x0a,0x0b,0x0c,0x0d,0x0e,0x0f,0x11,0x22};
    uint8_t c[16]="expand 32-byte k";
    uint8_t out[32]={0};
    crypto_core_hsalsa20(out, in, k, c);
    hex_buf(out, 32); printf("\n"); }
  return 0;
}
