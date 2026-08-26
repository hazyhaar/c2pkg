# tribench report

- stamp: 2026-08-11T10:27:05Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | skip | true | false | 17 | 1013014 |
| crc32_ieee | OK | OK | skip | true | false | 22 | 1476501 |
| fast_xor | OK | OK | skip | true | false | 21 | 1091338 |
| siphash24 | OK | OK | skip | true | false | 210 | 998307 |
| murmur3_x86_32 | OK | OK | skip | true | false | 59 | 971888 |
| blake2b_compress | OK | OK | skip | true | false | 423 | 989182 |
| chacha20_qr | OK | OK | skip | true | false | 26 | 966780 |
| md5_transform | OK | OK | skip | true | false | 50 | 935210 |
| poly1305_block5 | OK | OK | skip | true | false | 60 | 928915 |
| base64_simd | OK | OK | skip | true | false | 57 | 1119525 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 270 | 917794 |
| libinjection_sqli | skip | OK | skip | true | false | 56 | 949230 |
