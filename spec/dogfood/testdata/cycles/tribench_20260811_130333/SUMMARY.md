# tribench report

- stamp: 2026-08-11T11:03:33Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | skip | true | false | 17 | 1096566 |
| crc32_ieee | OK | OK | skip | true | false | 22 | 1591904 |
| fast_xor | OK | OK | skip | true | false | 21 | 1104387 |
| siphash24 | OK | OK | skip | true | false | 185 | 1063286 |
| murmur3_x86_32 | OK | OK | skip | true | false | 59 | 1023047 |
| blake2b_compress | OK | OK | skip | true | false | 423 | 1007737 |
| chacha20_qr | OK | OK | skip | true | false | 26 | 1007822 |
| md5_transform | OK | OK | skip | true | false | 50 | 918495 |
| poly1305_block5 | OK | OK | skip | true | false | 60 | 1007648 |
| base64_simd | OK | OK | skip | true | false | 57 | 977566 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 270 | 971734 |
| libinjection_sqli | skip | OK | skip | true | false | 56 | 988659 |
