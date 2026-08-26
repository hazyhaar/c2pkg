# tribench report

- stamp: 2026-08-19T22:38:03Z
- host: linux/amd64
- go: go1.27rc3
- summary: 1/1 compared bit-exact vs C; sgoiter-match=1 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines |
|-----|------|---------|------|-----------|------------|-----------|
| miniz_adler32 | OK | OK | skip | true | false | 56 |
