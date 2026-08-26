# sgoiter 12-lib harvest — 2026-08-10 (final)

## ChaCha / secretstream

`secretstream55` production crypto already uses `x/crypto/chacha20poly1305` + c2simd SIMD
engine (`490ef3d`, `213fc36`). The dogfood kernel `chacha20_qr.c` is **lab transpile only**,
not a rewrite of gosecretstream.

## Scoreboard

| Iter | build OK | funcs |
|------|--------:|------:|
| 1 baseline v0.6 | 3/12 | 36 |
| 2–3 macros/emit | 7/12 | 40 |
| 4 globals/ptr-off/j++/char/shl | **11/12** | **43** |

## Per-lib

| lib | build | notes |
|-----|-------|-------|
| base64_simd | OK | global string table + `j++` + `'='` |
| blake2b_compress | FAIL | 2D `sigma[r][i]` + G macro |
| chacha20_qr | OK | lab only |
| crc32_ieee | OK | |
| fast_xor | OK | cast word store LE |
| fnv1a_64 | OK | ULL imm |
| libinjection_sqli | OK | partial harvest |
| md5_transform | OK | `(a)+=` + cast LE load |
| murmur3_x86_32 | OK | KAT gcc |
| poly1305_block5 | OK | |
| siphash24 | OK | ptr offset cursor |
| tweetnacl_dogfood | OK | 31 funcs + stubs |

## Remaining
blake2b 2D const tables.
