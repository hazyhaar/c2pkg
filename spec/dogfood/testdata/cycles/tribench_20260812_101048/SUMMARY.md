# tribench report

- stamp: 2026-08-12T08:10:48Z
- host: linux/amd64
- go: go1.26.5
- summary: 5/6 compared bit-exact vs C; sgoiter-match=5 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines |
|-----|------|---------|------|-----------|------------|-----------|
| fnv1a_64 | OK | OK | skip | true | false | 28 |
| crc32_ieee | OK | OK | skip | false | false | 17 |
| fast_xor | OK | OK | skip | true | false | 50 |
| murmur3_x86_32 | OK | OK | skip | true | false | 62 |
| blake2b_compress | OK | OK | skip | true | false | 159 |
| chacha20_qr | OK | OK | skip | true | false | 31 |
