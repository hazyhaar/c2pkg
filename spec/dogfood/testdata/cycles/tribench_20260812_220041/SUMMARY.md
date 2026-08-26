# tribench report

- stamp: 2026-08-12T20:00:41Z
- host: linux/amd64
- go: go1.26.5
- summary: 3/4 compared bit-exact vs C; sgoiter-match=3 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines |
|-----|------|---------|------|-----------|------------|-----------|
| fnv1a_64 | OK | FAIL | skip | false | false | 28 |
| crc32_ieee | OK | OK | skip | true | false | 23 |
| blake2b_compress | OK | OK | skip | true | false | 143 |
| base64_simd | OK | OK | skip | true | false | 41 |
