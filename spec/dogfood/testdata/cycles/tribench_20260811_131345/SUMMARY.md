# tribench report

- stamp: 2026-08-11T11:13:45Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | skip | true | false | 17 | 1067255 |
| crc32_ieee | OK | OK | skip | true | false | 22 | 1498789 |
| fast_xor | OK | OK | skip | true | false | 21 | 1008183 |
| siphash24 | OK | OK | skip | true | false | 181 | 989911 |
| murmur3_x86_32 | OK | OK | skip | true | false | 59 | 940422 |
| blake2b_compress | OK | OK | skip | true | false | 423 | 977904 |
| chacha20_qr | OK | OK | skip | true | false | 26 | 1041003 |
| md5_transform | OK | OK | skip | true | false | 50 | 1061829 |
| poly1305_block5 | OK | OK | skip | true | false | 60 | 1047130 |
| base64_simd | OK | OK | skip | true | false | 57 | 1047969 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 270 | 1046068 |
| libinjection_sqli | skip | OK | skip | true | false | 56 | 949624 |
