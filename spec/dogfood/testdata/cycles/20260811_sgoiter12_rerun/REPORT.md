# sgoiter 12-lib — 20260811 **12/12 GREEN**

Post-monocypher AEAD engine + fixes régressions.

## Score
| | j (e5c92b8) | mid-regression | **final** |
|--|--:|--:|--:|
| build OK | 11/12 | 5/12 | **12/12** |

## Fixes (moteur)

| Fix | Fichiers | Libs |
|-----|----------|------|
| `(TYPE)-x` pas binop si TYPE | front findOp | crc32 |
| `((T*)p)[i]` indexable base | front isPtrCastBase | siphash, fast_xor, md5, blake |
| `ensureOffSlot` → f.Body | front | siphash for in!=end |
| uniqueExportName Sigma0/sigma0 | emit | tweetnacl |
| strchr builtin stub | emit | libinjection |
| uint8_t arr[N]={…} globals | front struct.go | blake2b flat sigma |
| `:=` + `_ = tmp` write-only | emit writeAssign | siphash unused |
| tweetnacl `m-=n` rewrite | src.c dogfood | tweetnacl hash |

## blake2b
Source dogfood aplati `sigma[12*16]` (pas de vrai 2D sgoiter encore). Finding 2D reste backlog langage.

## CI
`ci_check.sh` OK (labs + monocypher 1KB + sgoiter_out + parity).
