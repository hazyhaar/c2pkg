# tribench report

- stamp: 2026-08-11T10:29:01Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | skip | true | false | 17 | 1080369 |
| crc32_ieee | OK | OK | skip | true | false | 22 | 1527385 |
| fast_xor | OK | OK | skip | true | false | 21 | 1125728 |
| siphash24 | OK | OK | skip | true | false | 210 | 1080148 |
| murmur3_x86_32 | OK | OK | skip | true | false | 59 | 1021451 |
| blake2b_compress | OK | OK | skip | true | false | 423 | 1022102 |
| chacha20_qr | OK | OK | skip | true | false | 26 | 908852 |
| md5_transform | OK | OK | skip | true | false | 50 | 973155 |
| poly1305_block5 | OK | OK | skip | true | false | 60 | 979919 |
| base64_simd | OK | OK | skip | true | false | 57 | 1012539 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 270 | 1066545 |
| libinjection_sqli | skip | OK | skip | true | false | 56 | 1037091 |
