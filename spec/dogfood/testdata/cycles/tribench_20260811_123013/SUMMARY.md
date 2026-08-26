# tribench report

- stamp: 2026-08-11T10:30:13Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | skip | true | false | 16 | 1033256 |
| crc32_ieee | OK | OK | skip | true | false | 21 | 1514925 |
| fast_xor | OK | OK | skip | true | false | 22 | 1101061 |
| siphash24 | OK | OK | skip | true | false | 210 | 1085294 |
| murmur3_x86_32 | OK | OK | skip | true | false | 62 | 1029658 |
| blake2b_compress | OK | OK | skip | true | false | 425 | 938125 |
| chacha20_qr | OK | OK | skip | true | false | 38 | 982241 |
| md5_transform | OK | OK | skip | true | false | 50 | 1020456 |
| poly1305_block5 | OK | OK | skip | true | false | 60 | 1040980 |
| base64_simd | OK | OK | skip | true | false | 65 | 971980 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 266 | 1215119 |
| libinjection_sqli | skip | OK | skip | true | false | 57 | 1079909 |
