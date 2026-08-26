# tribench report

- stamp: 2026-08-11T09:38:07Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=11

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | OK | true | true | 21 | 1082818 |
| crc32_ieee | OK | OK | OK | true | true | 36 | 1441864 |
| fast_xor | OK | OK | OK | true | true | 36 | 1017383 |
| siphash24 | OK | OK | OK | true | true | 504 | 984442 |
| murmur3_x86_32 | OK | OK | OK | true | true | 122 | 1022105 |
| blake2b_compress | OK | OK | OK | true | true | 866 | 956828 |
| chacha20_qr | OK | OK | OK | true | true | 89 | 946018 |
| md5_transform | OK | OK | OK | true | true | 143 | 942440 |
| poly1305_block5 | OK | OK | OK | true | true | 211 | 939811 |
| base64_simd | OK | OK | OK | true | true | 138 | 1161056 |
| tweetnacl_dogfood | OK | OK | OK | true | true | 447 | 984105 |
| libinjection_sqli | skip | OK | skip | true | false | 76 | 952218 |
