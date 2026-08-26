# tribench report

- stamp: 2026-08-11T10:32:16Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | skip | true | false | 16 | 1066429 |
| crc32_ieee | OK | OK | skip | true | false | 21 | 1558394 |
| fast_xor | OK | OK | skip | true | false | 21 | 1121170 |
| siphash24 | OK | OK | skip | true | false | 210 | 1065769 |
| murmur3_x86_32 | OK | OK | skip | true | false | 62 | 1033673 |
| blake2b_compress | OK | OK | skip | true | false | 425 | 990850 |
| chacha20_qr | OK | OK | skip | true | false | 38 | 1027976 |
| md5_transform | OK | OK | skip | true | false | 50 | 1008484 |
| poly1305_block5 | OK | OK | skip | true | false | 60 | 1040026 |
| base64_simd | OK | OK | skip | true | false | 65 | 1043935 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 266 | 1147767 |
| libinjection_sqli | skip | OK | skip | true | false | 55 | 1154125 |
