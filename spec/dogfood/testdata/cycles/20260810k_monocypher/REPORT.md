# monocypher sgoiter — 20260810k Exit 2 + parity ccgo

## Résultats
| Oracle | Résultat |
|--------|----------|
| go build emit | OK ~9440 L |
| Self-KAT 36B / 1KB / chacha128 | PASS (`sgoiter_out/`) |
| **vs ccgo** secretstream55/internal/monocypher | **PASS bit-exact** CT+MAC |
| ci_check + MultiBlock 1KB | OK |
| Package dual | `pkg/secretstream55/internal/monocypher_sgoiter` |

## Vecteurs
`sgoiter_out/vectors_ccgo_parity.json`

## Fix multi-bloc
`ptr += N` → `ptr = ptr[N:]`
