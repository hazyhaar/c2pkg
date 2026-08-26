#include <stdint.h>
#include <stddef.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

/* kernel symbols */
size_t yyjson_write_u32(char *buf, uint32_t val);

static void hex_u64(uint64_t v) { printf("%016llx", (unsigned long long)v); }
static void hex_u32(uint32_t v) { printf("%08x", v); }
static void hex_buf(const uint8_t *p, size_t n) {
  size_t i; for (i = 0; i < n; i++) printf("%02x", p[i]);
}

int main(void) {
  /* fixture zero */
  /* fixture single */
  /* fixture double */
  /* fixture medium */
  /* fixture large */
  /* fixture max */
  printf("zero ");
  { char buf[32]; size_t n = yyjson_write_u32(buf, 0U); printf("%.*s\n", (int)n, buf); }
  printf("single ");
  { char buf[32]; size_t n = yyjson_write_u32(buf, 7U); printf("%.*s\n", (int)n, buf); }
  printf("double ");
  { char buf[32]; size_t n = yyjson_write_u32(buf, 42U); printf("%.*s\n", (int)n, buf); }
  printf("medium ");
  { char buf[32]; size_t n = yyjson_write_u32(buf, 12345U); printf("%.*s\n", (int)n, buf); }
  printf("large ");
  { char buf[32]; size_t n = yyjson_write_u32(buf, 123456789U); printf("%.*s\n", (int)n, buf); }
  printf("max ");
  { char buf[32]; size_t n = yyjson_write_u32(buf, 4294967295U); printf("%.*s\n", (int)n, buf); }
  return 0;
}
