# tribench report

- stamp: 2026-08-19T22:17:19Z
- host: linux/amd64
- go: go1.27rc3
- summary: 23/23 compared bit-exact vs C; sgoiter-match=23 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines |
|-----|------|---------|------|-----------|------------|-----------|
| fnv1a_64 | OK | OK | skip | true | false | 27 |
| crc32_ieee | OK | OK | skip | true | false | 23 |
| fast_xor | OK | OK | skip | true | false | 50 |
| siphash24 | OK | OK | skip | true | false | 69 |
| murmur3_x86_32 | OK | OK | skip | true | false | 62 |
| blake2b_compress | OK | OK | skip | true | false | 144 |
| chacha20_qr | OK | OK | skip | true | false | 57 |
| md5_transform | OK | OK | skip | true | false | 44 |
| poly1305_block5 | OK | OK | skip | true | false | 55 |
| poly1305_donna32 | OK | OK | skip | true | false | 68 |
| curve25519_donna64 | OK | OK | skip | true | false | 89 |
| yyjson_int | OK | OK | skip | true | false | 43 |
| cjson_core | OK | OK | skip | true | false | 57 |
| stbi_png_filter | OK | OK | skip | true | false | 91 |
| utf8proc_core | OK | OK | skip | true | false | 95 |
| fastlz_core | OK | OK | skip | true | false | 60 |
| murmur3_x64_128 | OK | OK | skip | true | false | 123 |
| tweetnacl_hsalsa | OK | OK | skip | true | false | 107 |
| base64_simd | OK | OK | skip | true | false | 41 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 299 |
| simsimd_dot_f32 | OK | OK | skip | true | false | 29 |
| strlenspn_lab | OK | OK | skip | true | false | 21 |
| md5_transform_full | OK | OK | skip | true | false | 248 |
