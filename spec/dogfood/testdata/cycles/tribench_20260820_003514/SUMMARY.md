# tribench report

- stamp: 2026-08-19T22:35:14Z
- host: linux/amd64
- go: go1.27rc3
- summary: 0/1 compared bit-exact vs C; sgoiter-match=0 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines |
|-----|------|---------|------|-----------|------------|-----------|
| xxhash64_core | FAIL | FAIL | skip | false | false | 89 |
