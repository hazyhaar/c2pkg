# tribench report

- stamp: 2026-08-11T09:02:38Z
- host: linux/amd64
- go: go1.26.5
- summary: 0/3 libs all-match; sgoiter-match=3 ccgo-match=3

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| fnv1a_64 | OK | OK | OK | true | true | 21 | 1112860 |
| crc32_ieee | OK | OK | OK | true | true | 36 | 1443918 |
| fast_xor | OK | OK | OK | true | true | 36 | 1105952 |
