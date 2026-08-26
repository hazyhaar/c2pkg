# tribench report

- stamp: 2026-08-12T06:52:14Z
- host: linux/amd64
- go: go1.26.5
- summary: 1/1 compared bit-exact vs C; sgoiter-match=1 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines |
|-----|------|---------|------|-----------|------------|-----------|
| blake2b_compress | OK | OK | skip | true | false | 6921 |
