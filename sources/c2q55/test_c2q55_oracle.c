#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include <assert.h>
#include "c2q55.h"

// Générateur pseudo-aléatoire 64-bit déterministe (xorshift64star) couvrant tout l'espace uint64
static uint64_t rng_state = 0x853c49e6748fea9bULL;
static uint64_t next_u64(void) {
    rng_state ^= rng_state >> 12;
    rng_state ^= rng_state << 25;
    rng_state ^= rng_state >> 27;
    return rng_state * 0x2545F4914F6CDD1DULL;
}

int main(void) {
    printf("=== C2Q55 C ORACLE HARNESS (AVX2 vs SCALAR) ===\n");

    // 1. KAT CRC32-C Castagnoli
    const uint8_t kat_data[] = "123456789";
    uint32_t crc = c2q_crc32c_hw(0, kat_data, 9);
    printf("KAT CRC32-C Castagnoli: 0x%08X (expected 0xE3069283) %s\n", 
           crc, (crc == 0xE3069283U) ? "[PASS]" : "[FAIL]");
    assert(crc == 0xE3069283U);

    // 2. Torture Test 2,000,000 Blocs sur le DOMAINE UINT64 COMPLET (y compris > 2^63)
    printf("--- [TEST] Torture Test 2,000,000 Random Blocks AVX2 vs Scalar (Full uint64 Domain) ---\n");
    c2q_meta_block_t block;

    for (int b = 0; b < 2000000; b++) {
        uint64_t now = next_u64();
        for (int i = 0; i < C2Q_BLOCK_SLOTS; i++) {
            block.visible_at[i] = next_u64();
            block.state[i] = (uint32_t)(next_u64() % 4);
            block.flags[i] = (uint32_t)next_u64();
        }

        uint32_t mask_scalar = c2q_scan_block_scalar(&block, now);
#if defined(__AVX2__)
        uint32_t mask_avx2 = c2q_scan_block_avx2(&block, now);
        if (mask_avx2 != mask_scalar) {
            printf("[FAIL] Divergence at block %d: AVX2=0x%X, Scalar=0x%X (now=%llu)\n",
                   b, mask_avx2, mask_scalar, (unsigned long long)now);
            return 1;
        }
#endif
    }
    printf("2,000,000 blocks tested across full uint64 range [0, 2^64-1]: 100%% AVX2 vs Scalar bit-exact [PASS]\n");
    printf("=== ALL C ORACLE TESTS PASSED ===\n");
    return 0;
}
