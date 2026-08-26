# tribench report

- stamp: 2026-08-12T20:08:08Z
- host: linux/amd64
- go: go1.26.5
- summary: 3/3 compared bit-exact vs C; sgoiter-match=3 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines |
|-----|------|---------|------|-----------|------------|-----------|
| fnv1a_64 | OK | OK | skip | true | false | 27 |
| murmur3_x86_32 | OK | OK | skip | true | false | 58 |
| base64_simd | OK | OK | skip | true | false | 49 |
