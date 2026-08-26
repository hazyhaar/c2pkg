# tribench report

- stamp: 2026-08-11T10:33:24Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | skip | true | false | 16 | 1150405 |
| crc32_ieee | OK | OK | skip | true | false | 21 | 1614229 |
| fast_xor | OK | OK | skip | true | false | 21 | 1145490 |
| siphash24 | OK | OK | skip | true | false | 210 | 1126812 |
| murmur3_x86_32 | OK | OK | skip | true | false | 62 | 1045750 |
| blake2b_compress | OK | OK | skip | true | false | 425 | 1013098 |
| chacha20_qr | OK | OK | skip | true | false | 38 | 1037244 |
| md5_transform | OK | OK | skip | true | false | 50 | 1061922 |
| poly1305_block5 | OK | OK | skip | true | false | 60 | 1478849 |
| base64_simd | OK | OK | skip | true | false | 65 | 1124654 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 266 | 1063962 |
| libinjection_sqli | skip | OK | skip | true | false | 55 | 1108485 |
