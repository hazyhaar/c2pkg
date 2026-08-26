# tribench report

- stamp: 2026-08-11T15:02:12Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | skip | true | false | 17 | 1105301 |
| crc32_ieee | OK | OK | skip | true | false | 22 | 1589642 |
| fast_xor | OK | OK | skip | true | false | 21 | 1221988 |
| siphash24 | OK | OK | skip | true | false | 173 | 1056593 |
| murmur3_x86_32 | OK | OK | skip | true | false | 59 | 1052245 |
| blake2b_compress | OK | OK | skip | true | false | 362 | 1037076 |
| chacha20_qr | OK | OK | skip | true | false | 26 | 979300 |
| md5_transform | OK | OK | skip | true | false | 50 | 980023 |
| poly1305_block5 | OK | OK | skip | true | false | 54 | 1000586 |
| base64_simd | OK | OK | skip | true | false | 57 | 992141 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 270 | 950668 |
| libinjection_sqli | skip | OK | skip | true | false | 56 | 1050596 |
