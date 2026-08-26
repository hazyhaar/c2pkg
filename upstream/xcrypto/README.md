# Reproduction Harness: x/crypto/chacha20poly1305 (RFC 8439)

This directory contains the reproduction harness evaluating the pure-Go `simd/archsimd` implementation against the official upstream `golang.org/x/crypto/chacha20poly1305` test suite and benchmark suite.

## Invariants & Design

1. **Untouched Upstream Tests:**  
   The upstream test suite (`TestVectors` with 102 test vectors, `TestRandom`, and `BenchmarkChacha20Poly1305`) is run without modifying a single line of the upstream test files (`cmp` verified against the module cache).
2. **Zero Allocation Profile:**  
   The `archsimd` graft achieves 0 B/op and 0 allocs/op across all message sizes (64 B, 1350 B, 8192 B).
3. **Reproducibility:**  
   The test matrix and `benchstat` comparisons can be executed in a single command.

---

## Directory Contents

| File | Description |
| :--- | :--- |
| **`chacha20poly1305_archsimd.go`** | Pure-Go AVX2 graft delegating the IETF 12-byte nonce path to `c2simd.AEADSubkeyLockDst` / `AEADSubkeyUnlockDst`. |
| **`harness.sh`** | Materializes `x/crypto@v0.54.0` from module cache, applies the archsimd build tag, and executes tests and benchmarks across the 3 configurations (`asm`, `purego`, `archsimd`). |
| **`results/`** | Raw benchmark outputs (`asm.txt`, `purego.txt`, `archsimd.txt`), `benchstat.md` comparison, stratified pipeline breakdown, and mutation metrics. |

---

## Usage

```bash
# Run unit tests across all 3 configurations (asm, purego, archsimd)
./harness.sh test

# Run benchmarks pinned to performance cores with benchstat output
PIN="taskset -c 0-15" ./harness.sh bench

# Run complete suite: tests, benchmarks, stratified pipeline and mutation breakdown
PIN="taskset -c 0-15" ./harness.sh all
```
