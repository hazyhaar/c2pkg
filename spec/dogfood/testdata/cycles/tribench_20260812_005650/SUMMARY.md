# tribench report

- stamp: 2026-08-11T22:56:50Z
- host: linux/amd64
- go: go1.26.5
- summary: 11/11 compared bit-exact vs C; sgoiter-match=11 ccgo-match=0
- 1 kernel(s) built with no C oracle: compiled and run, never compared

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines |
|-----|------|---------|------|-----------|------------|-----------|
| fnv1a_64 | OK | OK | skip | true | false | 23 |
| crc32_ieee | OK | OK | skip | true | false | 23 |
| fast_xor | OK | OK | skip | true | false | 30 |
| siphash24 | OK | OK | skip | true | false | 69 |
| murmur3_x86_32 | OK | OK | skip | true | false | 62 |
| blake2b_compress | OK | OK | skip | true | false | 1005 |
| chacha20_qr | OK | OK | skip | true | false | 27 |
| md5_transform | OK | OK | skip | true | false | 44 |
| poly1305_block5 | OK | OK | skip | true | false | 55 |
| base64_simd | OK | OK | skip | true | false | 33 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 311 |
| libinjection_sqli | skip | OK | skip | no oracle | false | 57 |
