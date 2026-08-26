# tribench report

- stamp: 2026-08-11T10:29:38Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | skip | true | false | 16 | 1112422 |
| crc32_ieee | OK | OK | skip | true | false | 21 | 1394934 |
| fast_xor | OK | OK | skip | true | false | 22 | 1043853 |
| siphash24 | OK | OK | skip | true | false | 210 | 955829 |
| murmur3_x86_32 | OK | OK | skip | true | false | 62 | 950007 |
| blake2b_compress | OK | OK | skip | true | false | 425 | 947053 |
| chacha20_qr | OK | OK | skip | true | false | 38 | 897276 |
| md5_transform | OK | OK | skip | true | false | 50 | 938224 |
| poly1305_block5 | OK | OK | skip | true | false | 60 | 920304 |
| base64_simd | OK | OK | skip | true | false | 65 | 1047707 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 266 | 1083644 |
| libinjection_sqli | skip | OK | skip | true | false | 57 | 1009803 |
