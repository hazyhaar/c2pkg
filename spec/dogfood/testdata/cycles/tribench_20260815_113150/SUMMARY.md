# tribench report

- stamp: 2026-08-15T09:31:50Z
- host: linux/amd64
- go: go1.26.5
- summary: 23/23 compared bit-exact vs C; sgoiter-match=23 ccgo-match=11
- 1 kernel(s) built with no C oracle: compiled and run, never compared

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines |
|-----|------|---------|------|-----------|------------|-----------|
| fnv1a_64 | OK | OK | OK | true | true | 27 |
| crc32_ieee | OK | OK | OK | true | true | 23 |
| fast_xor | OK | OK | FAIL | true | false | 50 |
| siphash24 | OK | OK | OK | true | true | 69 |
| murmur3_x86_32 | OK | OK | OK | true | true | 62 |
| blake2b_compress | OK | OK | OK | true | true | 143 |
| chacha20_qr | OK | OK | OK | true | true | 57 |
| md5_transform | OK | OK | OK | true | true | 44 |
| poly1305_block5 | OK | OK | OK | true | true | 55 |
| poly1305_donna32 | OK | OK | FAIL | true | false | 68 |
| curve25519_donna64 | OK | OK | FAIL | true | false | 89 |
| yyjson_int | OK | OK | FAIL | true | false | 43 |
| cjson_core | OK | OK | FAIL | true | false | 57 |
| stbi_png_filter | OK | OK | FAIL | true | false | 96 |
| utf8proc_core | OK | OK | FAIL | true | false | 95 |
| fastlz_core | OK | OK | FAIL | true | false | 60 |
| murmur3_x64_128 | OK | OK | FAIL | true | false | 123 |
| tweetnacl_hsalsa | OK | OK | FAIL | true | false | 120 |
| base64_simd | OK | OK | OK | true | true | 41 |
| tweetnacl_dogfood | OK | OK | OK | true | true | 332 |
| simsimd_dot_f32 | OK | OK | FAIL | true | false | 29 |
| libinjection_sqli | skip | FAIL | skip | no oracle | false | 65 |
| strlenspn_lab | OK | OK | FAIL | true | false | 21 |
| md5_transform_full | OK | OK | OK | true | true | 248 |
