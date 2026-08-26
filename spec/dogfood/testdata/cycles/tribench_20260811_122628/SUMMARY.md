# tribench report

- stamp: 2026-08-11T10:26:28Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | skip | true | false | 17 | 1086166 |
| crc32_ieee | OK | OK | skip | true | false | 22 | 1558106 |
| fast_xor | OK | OK | skip | true | false | 21 | 1122668 |
| siphash24 | OK | OK | skip | true | false | 274 | 1182998 |
| murmur3_x86_32 | OK | OK | skip | true | false | 59 | 1068383 |
| blake2b_compress | OK | OK | skip | true | false | 427 | 922360 |
| chacha20_qr | OK | OK | skip | true | false | 26 | 967150 |
| md5_transform | OK | OK | skip | true | false | 50 | 974158 |
| poly1305_block5 | OK | OK | skip | true | false | 73 | 921608 |
| base64_simd | OK | OK | skip | true | false | 57 | 1185418 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 270 | 987671 |
| libinjection_sqli | skip | OK | skip | true | false | 56 | 1082381 |
