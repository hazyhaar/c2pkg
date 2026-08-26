# tribench report

- stamp: 2026-08-11T10:37:29Z
- host: linux/amd64
- go: go1.26.5
- summary: 12/12 libs all-match; sgoiter-match=12 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | skip | true | false | 16 | 1105055 |
| crc32_ieee | OK | OK | skip | true | false | 21 | 1493567 |
| fast_xor | OK | OK | skip | true | false | 21 | 1117512 |
| siphash24 | OK | OK | skip | true | false | 210 | 1033048 |
| murmur3_x86_32 | OK | OK | skip | true | false | 62 | 1044849 |
| blake2b_compress | OK | OK | skip | true | false | 425 | 1004104 |
| chacha20_qr | OK | OK | skip | true | false | 38 | 972232 |
| md5_transform | OK | OK | skip | true | false | 50 | 961573 |
| poly1305_block5 | OK | OK | skip | true | false | 60 | 1009028 |
| base64_simd | OK | OK | skip | true | false | 65 | 996769 |
| tweetnacl_dogfood | OK | OK | skip | true | false | 266 | 1007449 |
| libinjection_sqli | skip | OK | skip | true | false | 55 | 995038 |
