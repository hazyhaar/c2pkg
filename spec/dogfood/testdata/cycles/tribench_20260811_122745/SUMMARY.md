# tribench report

- stamp: 2026-08-11T10:27:45Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | skip | true | false | 17 | 1098122 |
| crc32_ieee | OK | OK | skip | true | false | 22 | 1536407 |
| fast_xor | OK | OK | skip | true | false | 21 | 1083496 |
| siphash24 | OK | OK | skip | true | false | 210 | 993312 |
| murmur3_x86_32 | OK | OK | skip | true | false | 59 | 1008327 |
| blake2b_compress | OK | OK | skip | true | false | 423 | 982609 |
| chacha20_qr | OK | OK | skip | true | false | 26 | 976787 |
| md5_transform | OK | OK | skip | true | false | 50 | 966696 |
| poly1305_block5 | OK | OK | skip | true | false | 60 | 979114 |
| base64_simd | OK | OK | skip | true | false | 57 | 944706 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 270 | 955507 |
| libinjection_sqli | skip | OK | skip | true | false | 56 | 1021349 |
