# tribench report

- stamp: 2026-08-12T08:17:08Z
- host: linux/amd64
- go: go1.26.5
- summary: 2/3 compared bit-exact vs C; sgoiter-match=2 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines |
|-----|------|---------|------|-----------|------------|-----------|
| crc32_ieee | OK | OK | skip | true | false | 23 |
| blake2b_compress | OK | OK | skip | true | false | 143 |
| tweetnacl_dogfood | OK | FAIL | skip | false | false | 311 |
