# tribench report

- stamp: 2026-08-11T20:47:07Z
- host: linux/amd64
- go: go1.26.5
- summary: 11/11 compared bit-exact vs C; sgoiter-match=11 ccgo-match=11
- 1 kernel(s) built with no C oracle: compiled and run, never compared

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines |
|-----|------|---------|------|-----------|------------|-----------|
| fnv1a_64 | OK | OK | OK | true | true | 17 |
| crc32_ieee | OK | OK | OK | true | true | 22 |
| fast_xor | OK | OK | OK | true | true | 21 |
| siphash24 | OK | OK | OK | true | true | 173 |
| murmur3_x86_32 | OK | OK | OK | true | true | 62 |
| blake2b_compress | OK | OK | OK | true | true | 112 |
| chacha20_qr | OK | OK | OK | true | true | 26 |
| md5_transform | OK | OK | OK | true | true | 50 |
| poly1305_block5 | OK | OK | OK | true | true | 54 |
| base64_simd | OK | OK | OK | true | true | 48 |
| tweetnacl_dogfood | OK | OK | OK | true | true | 282 |
| libinjection_sqli | skip | OK | skip | no oracle | false | 56 |
