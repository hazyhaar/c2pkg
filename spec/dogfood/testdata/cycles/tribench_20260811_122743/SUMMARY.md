# tribench report

- stamp: 2026-08-11T10:27:43Z
- host: linux/amd64
- go: go1.26.5
- summary: 1/1 libs all-match; sgoiter-match=1 ccgo-match=0

| lib | c_O2 | sgoiter | ccgo | sgo match | ccgo match | sgo lines | sgo ns/op |
|-----|------|---------|------|-----------|------------|-----------|----------|
| base64_simd | OK | OK | skip | true | false | 57 | 911301 |
