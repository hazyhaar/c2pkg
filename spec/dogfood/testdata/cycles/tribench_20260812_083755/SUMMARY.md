# tribench report

- stamp: 2026-08-12T06:37:55Z
- host: linux/amd64
- go: go1.26.5
- summary: 4/4 compared bit-exact vs C; sgoiter-match=4 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines |
|-----|------|---------|------|-----------|------------|-----------|
| fnv1a_64 | OK | OK | skip | true | false | 28 |
| fast_xor | OK | OK | skip | true | false | 50 |
| murmur3_x86_32 | OK | OK | skip | true | false | 62 |
| blake2b_compress | OK | OK | skip | true | false | 813 |
