# tribench report

- stamp: 2026-08-11T09:46:55Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=11

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | OK | true | true | 21 | 1037269 |
| crc32_ieee | OK | OK | OK | true | true | 36 | 1519631 |
| fast_xor | OK | OK | OK | true | true | 36 | 1055895 |
| siphash24 | OK | OK | OK | true | true | 504 | 962785 |
| murmur3_x86_32 | OK | OK | OK | true | true | 122 | 986758 |
| blake2b_compress | OK | OK | OK | true | true | 866 | 972434 |
| chacha20_qr | OK | OK | OK | true | true | 89 | 937872 |
| md5_transform | OK | OK | OK | true | true | 143 | 954564 |
| poly1305_block5 | OK | OK | OK | true | true | 211 | 945642 |
| base64_simd | OK | OK | OK | true | true | 138 | 1162633 |
| tweetnacl_dogfood | OK | OK | OK | true | true | 447 | 1212027 |
| libinjection_sqli | skip | OK | skip | true | false | 76 | 1177260 |
