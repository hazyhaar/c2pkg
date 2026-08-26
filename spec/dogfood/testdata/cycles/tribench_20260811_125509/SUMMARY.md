# tribench report

- stamp: 2026-08-11T10:55:09Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | skip | true | false | 17 | 1092314 |
| crc32_ieee | OK | OK | skip | true | false | 22 | 1672673 |
| fast_xor | OK | OK | skip | true | false | 21 | 1121419 |
| siphash24 | OK | OK | skip | true | false | 210 | 1118222 |
| murmur3_x86_32 | OK | OK | skip | true | false | 59 | 1059068 |
| blake2b_compress | OK | OK | skip | true | false | 423 | 950299 |
| chacha20_qr | OK | OK | skip | true | false | 26 | 1022520 |
| md5_transform | OK | OK | skip | true | false | 50 | 991972 |
| poly1305_block5 | OK | OK | skip | true | false | 60 | 1068509 |
| base64_simd | OK | OK | skip | true | false | 57 | 1129048 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 270 | 965773 |
| libinjection_sqli | skip | OK | skip | true | false | 56 | 1017568 |
