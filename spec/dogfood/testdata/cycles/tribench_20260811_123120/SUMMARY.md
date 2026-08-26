# tribench report

- stamp: 2026-08-11T10:31:20Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | skip | true | false | 16 | 1074655 |
| crc32_ieee | OK | OK | skip | true | false | 21 | 1517222 |
| fast_xor | OK | OK | skip | true | false | 21 | 1070316 |
| siphash24 | OK | OK | skip | true | false | 210 | 956576 |
| murmur3_x86_32 | OK | OK | skip | true | false | 62 | 987087 |
| blake2b_compress | OK | OK | skip | true | false | 425 | 936794 |
| chacha20_qr | OK | OK | skip | true | false | 38 | 1201751 |
| md5_transform | OK | OK | skip | true | false | 50 | 964585 |
| poly1305_block5 | OK | OK | skip | true | false | 60 | 998615 |
| base64_simd | OK | OK | skip | true | false | 65 | 985219 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 266 | 930983 |
| libinjection_sqli | skip | OK | skip | true | false | 55 | 946032 |
