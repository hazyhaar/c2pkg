# tribench report

- stamp: 2026-08-11T09:14:31Z
- host: linux/amd64
- go: go1.26.5
- summary: 8/12 libs all-match; sgoiter-match=8 ccgo-match=11

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | OK | true | true | 21 | 1050361 |
| crc32_ieee | OK | OK | OK | true | true | 36 | 1484820 |
| fast_xor | OK | OK | OK | true | true | 36 | 1056965 |
| siphash24 | OK | OK | OK | true | true | 504 | 994318 |
| murmur3_x86_32 | OK | OK | OK | true | true | 122 | 984426 |
| blake2b_compress | OK | OK | OK | false | true | 578 | 982512 |
| chacha20_qr | OK | OK | OK | false | true | 59 | 918640 |
| md5_transform | OK | OK | OK | true | true | 143 | 942398 |
| poly1305_block5 | OK | OK | OK | false | true | 205 | 918341 |
| base64_simd | OK | OK | OK | false | true | 138 | 1128256 |
| tweetnacl_dogfood | OK | OK | OK | true | true | 447 | 899912 |
| libinjection_sqli | skip | OK | skip | true | false | 76 | 981150 |
