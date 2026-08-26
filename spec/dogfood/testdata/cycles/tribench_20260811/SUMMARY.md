# tribench report

- stamp: 2026-08-11T09:06:04Z
- host: linux/amd64
- go: go1.26.5
- summary: 8/12 libs all-match; sgoiter-match=8 ccgo-match=11

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | OK | true | true | 21 | 1008786 |
| crc32_ieee | OK | OK | OK | true | true | 36 | 2052653 |
| fast_xor | OK | OK | OK | true | true | 36 | 1037100 |
| siphash24 | OK | OK | OK | true | true | 466 | 959703 |
| murmur3_x86_32 | OK | OK | OK | true | true | 122 | 959943 |
| blake2b_compress | OK | OK | OK | false | true | 586 | 893162 |
| chacha20_qr | OK | OK | OK | false | true | 67 | 968629 |
| md5_transform | OK | OK | OK | true | true | 151 | 961531 |
| poly1305_block5 | OK | OK | OK | false | true | 205 | 953185 |
| base64_simd | OK | OK | OK | false | true | 118 | 1018559 |
| tweetnacl_dogfood | OK | OK | OK | true | true | 447 | 927497 |
| libinjection_sqli | skip | OK | skip | true | false | 76 | 971460 |
