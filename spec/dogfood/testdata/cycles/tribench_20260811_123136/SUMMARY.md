# tribench report

- stamp: 2026-08-11T10:31:36Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | skip | true | false | 16 | 1057897 |
| crc32_ieee | OK | OK | skip | true | false | 21 | 1516622 |
| fast_xor | OK | OK | skip | true | false | 21 | 1045580 |
| siphash24 | OK | OK | skip | true | false | 210 | 1043923 |
| murmur3_x86_32 | OK | OK | skip | true | false | 62 | 1057966 |
| blake2b_compress | OK | OK | skip | true | false | 425 | 1164074 |
| chacha20_qr | OK | OK | skip | true | false | 38 | 969283 |
| md5_transform | OK | OK | skip | true | false | 50 | 960760 |
| poly1305_block5 | OK | OK | skip | true | false | 60 | 1076468 |
| base64_simd | OK | OK | skip | true | false | 65 | 1147176 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 266 | 1043772 |
| libinjection_sqli | skip | OK | skip | true | false | 55 | 1090911 |
