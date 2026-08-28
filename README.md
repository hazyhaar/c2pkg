# c2pkg — C2SIMD Sovereign Transpiled Packages (Pure Go / CGO=0)

`c2pkg` provides high-performance, pure-Go algorithmic packages mechanically transpiled from C99/C23 reference implementations by the **`sgoiter`** compiler (Go 1.27+ with `GOEXPERIMENT=simd`).

All packages provide strict invariants:
- **Zero CGO / 100% Pure Go:** No dynamic C runtime dependencies.
- **Bit-Exact Parity:** Verified against C99 oracles compiled with `gcc -O2` under ASan/UBSan.
- **Zero Heap Allocations:** On hot compute paths.
- **Hardware Acceleration:** Native SIMD intrinsics (AVX2, AVX-512, ARM NEON).

---

## Package Catalog

| Package | Upstream C Source | Description & Invariants |
| :--- | :--- | :--- |
| **`c2archtsim`** | `c2archsimd.c` | ARCHTIME `.rodata` tables, LUT16/LUT256 shuffles, vectorized hex codec. |
| **`c2chacha1-8`** | `monocypher.c` / RFC 8439 | Constant-time ChaCha stream ciphers (1, 2, 4, 8 rounds) with 0-alloc state. |
| **`c2poly1305`** | `poly1305.c` | RFC 8439 Poly1305 authenticator with 128-bit fused carry arithmetic. |
| **`c2poly1305x8`**| `poly1305.c` | 8-way parallel Poly1305 pipelined vector implementation. |
| **`c2blake2b`** | `blake2b.c` | High-throughput BLAKE2b cryptographic hashing with SIMD round unrolling. |
| **`blake3archtsim`**| `blake3.c` | Parallel tree-hashing BLAKE3 implementation with AVX2 backend. |
| **`c2lz4`** | `lz4.c` | High-speed compression / decompression with zero heap churn. |
| **`c2ryu`** | `ryu.c` | Ultra-fast float-to-string conversion with constant-table lookup. |
| **`c2json`** | `simdjson.c` | Vectorized structural JSON parser with stage 1/stage 2 bitmask sweeps. |
| **`c2pcre`** | `pcre.c` | Regular expression matching engine with DFA fast-path. |
| **`c2ordered_swiss`**| `abseil/swiss_table`| Cache-friendly ordered hash map with SIMD metadata group probing. |
| **`c2uuidv7`** | `uuidv7.c` | Monotonic UUIDv7 timestamp generator (0 allocs). |
| **`c2vtparser`** | `vtparser.c` | State-machine ANSI/VT500 escape sequence parser (Paul Flo Williams tables). |
| **`c2myers`** | `myers.c` | Vectorized Myers O(ND) diffing algorithm for text buffers. |
| **`c2swizzle`** | `swizzle.c` | RGBA <-> BGRA pixel transposition via `vpshufb` (>= 36 GB/s). |
| **`c2ssim`** | `ssim.c` | Separable 11x11 Gaussian blur and structural similarity index calculation. |

---

## Probes & Research

- **[`probes/abi_sroa_bench`](probes/abi_sroa_bench/README.md):** Isolated micro-benchmarks comparing parameter passing overhead (`[2]uint64` array stack spill vs `struct{Lo, Hi uint64}` register ABI in Go 1.27).

---

## Validation & Verification

All packages are validated under race detection:

```bash
GOEXPERIMENT=simd go test -v -race ./...
```

---

## License

Apache-2.0 OR MIT (with respective upstream attribution in `COPYRIGHT.md`).
