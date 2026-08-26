# tribench report

- stamp: 2026-08-11T23:09:50Z
- host: linux/amd64
- go: go1.26.5
- summary: 0/1 compared bit-exact vs C; sgoiter-match=0 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines |
|-----|------|---------|------|-----------|------------|-----------|
| strlenspn_lab | OK | FAIL | skip | false | false | 0 |
