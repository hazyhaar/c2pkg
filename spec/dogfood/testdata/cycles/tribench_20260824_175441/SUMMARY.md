# tribench report

- stamp: 2026-08-24T15:54:41Z
- host: linux/amd64
- go: go1.27rc3
- summary: 24/30 compared bit-exact vs C; sgoiter-match=24 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines |
|-----|------|---------|------|-----------|------------|-----------|
| fnv1a_64 | OK | OK | skip | true | false | 29 |
| crc32_ieee | OK | OK | skip | true | false | 27 |
| fast_xor | OK | OK | skip | true | false | 49 |
| siphash24 | OK | OK | skip | true | false | 150 |
| murmur3_x86_32 | OK | OK | skip | true | false | 62 |
| blake2b_compress | OK | OK | skip | false | false | 169 |
| chacha20_qr | OK | OK | skip | true | false | 65 |
| md5_transform | OK | OK | skip | true | false | 43 |
| poly1305_block5 | OK | OK | skip | true | false | 75 |
| poly1305_donna32 | OK | OK | skip | true | false | 83 |
| curve25519_donna64 | OK | OK | skip | true | false | 101 |
| yyjson_int | OK | OK | skip | true | false | 44 |
| cjson_core | OK | OK | skip | false | false | 63 |
| stbi_png_filter | OK | OK | skip | true | false | 108 |
| utf8proc_core | OK | FAIL | skip | false | false | 132 |
| fastlz_core | OK | OK | skip | true | false | 69 |
| murmur3_x64_128 | OK | OK | skip | false | false | 125 |
| tweetnacl_hsalsa | OK | OK | skip | true | false | 109 |
| base64_simd | OK | OK | skip | true | false | 42 |
| tweetnacl_dogfood | OK | FAIL | skip | false | false | 677 |
| simsimd_dot_f32 | OK | OK | skip | true | false | 24 |
| strlenspn_lab | OK | OK | skip | true | false | 24 |
| md5_transform_full | OK | OK | skip | false | false | 322 |
| xxhash64_core | OK | OK | skip | true | false | 91 |
| miniz_adler32 | OK | OK | skip | true | false | 66 |
| stbi_crc_dogfood | OK | OK | skip | true | false | 27 |
| yyjson_digit_dogfood | OK | OK | skip | true | false | 56 |
| utf8_iterate_dogfood | OK | OK | skip | true | false | 57 |
| quickjs_minmax_dogfood | OK | OK | skip | true | false | 26 |
| simsimd_l2sq_f32 | OK | OK | skip | true | false | 14 |
