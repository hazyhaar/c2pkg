# tribench report

- stamp: 2026-08-11T09:27:44Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=11

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | OK | true | true | 21 | 1048960 |
| crc32_ieee | OK | OK | OK | true | true | 36 | 1468343 |
| fast_xor | OK | OK | OK | true | true | 36 | 1131517 |
| siphash24 | OK | OK | OK | true | true | 504 | 1103733 |
| murmur3_x86_32 | OK | OK | OK | true | true | 122 | 1048121 |
| blake2b_compress | OK | OK | OK | true | true | 866 | 939491 |
| chacha20_qr | OK | OK | OK | true | true | 89 | 960638 |
| md5_transform | OK | OK | OK | true | true | 143 | 983903 |
| poly1305_block5 | OK | OK | OK | true | true | 211 | 948819 |
| base64_simd | OK | OK | OK | true | true | 138 | 1209918 |
| tweetnacl_dogfood | OK | OK | OK | true | true | 447 | 976086 |
| libinjection_sqli | skip | OK | skip | true | false | 76 | 1008411 |
