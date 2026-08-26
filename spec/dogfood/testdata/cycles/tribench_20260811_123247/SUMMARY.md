# tribench report

- stamp: 2026-08-11T10:32:47Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | skip | true | false | 16 | 2150233 |
| crc32_ieee | OK | OK | skip | true | false | 21 | 1602061 |
| fast_xor | OK | OK | skip | true | false | 21 | 1192432 |
| siphash24 | OK | OK | skip | true | false | 210 | 1085618 |
| murmur3_x86_32 | OK | OK | skip | true | false | 62 | 1528199 |
| blake2b_compress | OK | OK | skip | true | false | 425 | 1019492 |
| chacha20_qr | OK | OK | skip | true | false | 38 | 1194865 |
| md5_transform | OK | OK | skip | true | false | 50 | 947507 |
| poly1305_block5 | OK | OK | skip | true | false | 60 | 1128354 |
| base64_simd | OK | OK | skip | true | false | 65 | 1188423 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 266 | 1110447 |
| libinjection_sqli | skip | OK | skip | true | false | 55 | 994298 |
