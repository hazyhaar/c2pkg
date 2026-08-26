# tribench report

- stamp: 2026-08-11T20:45:15Z
- host: linux/amd64
- go: go1.26.5
- summary: 11/11 compared bit-exact vs C; sgoiter-match=11 ccgo-match=0
- 1 kernel(s) built with no C oracle: compiled and run, never compared

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines |
|-----|------|---------|------|-----------|------------|-----------|
| fnv1a_64 | OK | OK | FAIL | true | false | 17 |
| crc32_ieee | OK | OK | FAIL | true | false | 22 |
| fast_xor | OK | OK | FAIL | true | false | 21 |
| siphash24 | OK | OK | FAIL | true | false | 173 |
| murmur3_x86_32 | OK | OK | FAIL | true | false | 62 |
| blake2b_compress | OK | OK | FAIL | true | false | 112 |
| chacha20_qr | OK | OK | FAIL | true | false | 26 |
| md5_transform | OK | OK | FAIL | true | false | 50 |
| poly1305_block5 | OK | OK | FAIL | true | false | 54 |
| base64_simd | OK | OK | FAIL | true | false | 48 |
| tweetnacl_dogfood | OK | OK | FAIL | true | false | 282 |
| libinjection_sqli | skip | OK | skip | no oracle | false | 56 |
