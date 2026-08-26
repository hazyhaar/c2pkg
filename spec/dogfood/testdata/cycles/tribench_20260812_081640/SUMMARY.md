# tribench report

- stamp: 2026-08-12T06:16:40Z
- host: linux/amd64
- go: go1.26.5
- summary: 13/13 compared bit-exact vs C; sgoiter-match=13 ccgo-match=0
- 1 kernel(s) built with no C oracle: compiled and run, never compared

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines |
|-----|------|---------|------|-----------|------------|-----------|
| fnv1a_64 | OK | OK | skip | true | false | 28 |
| crc32_ieee | OK | OK | skip | true | false | 23 |
| fast_xor | OK | OK | skip | true | false | 25 |
| siphash24 | OK | OK | skip | true | false | 69 |
| murmur3_x86_32 | OK | OK | skip | true | false | 62 |
| blake2b_compress | OK | OK | skip | true | false | 813 |
| chacha20_qr | OK | OK | skip | true | false | 27 |
| md5_transform | OK | OK | skip | true | false | 44 |
| poly1305_block5 | OK | OK | skip | true | false | 55 |
| base64_simd | OK | OK | skip | true | false | 37 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 311 |
| libinjection_sqli | skip | FAIL | skip | no oracle | false | 57 |
| strlenspn_lab | OK | OK | skip | true | false | 21 |
| md5_transform_full | OK | OK | skip | true | false | 248 |
