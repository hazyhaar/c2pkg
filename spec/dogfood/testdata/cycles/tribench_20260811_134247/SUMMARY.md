# tribench report

- stamp: 2026-08-11T11:42:47Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | skip | true | false | 17 | 0 |
| crc32_ieee | OK | OK | skip | true | false | 22 | 0 |
| fast_xor | OK | OK | skip | true | false | 21 | 0 |
| siphash24 | OK | OK | skip | true | false | 173 | 0 |
| murmur3_x86_32 | OK | OK | skip | true | false | 59 | 0 |
| blake2b_compress | OK | OK | skip | true | false | 362 | 0 |
| chacha20_qr | OK | OK | skip | true | false | 26 | 0 |
| md5_transform | OK | OK | skip | true | false | 50 | 0 |
| poly1305_block5 | OK | OK | skip | true | false | 54 | 0 |
| base64_simd | OK | OK | skip | true | false | 57 | 0 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 270 | 0 |
| libinjection_sqli | skip | OK | skip | true | false | 56 | 0 |
