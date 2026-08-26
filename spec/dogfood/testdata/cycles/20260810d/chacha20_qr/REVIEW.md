# REVIEW dogfood — chacha20_qr (20260810d)

## Métriques
```json
{
  "kernel": "chacha20_qr",
  "stamp": "20260810d",
  "src_c": "spec/c_sources/testdata/c_sources/chacha20_qr.c",
  "ccgo": "/home/cl-ment/go/bin/ccgo",
  "raw_lines": 501,
  "opt_lines": 500,
  "raw_rotl_calls": 0,
  "opt_rotl_calls": 0,
  "raw_bits_rotate": 0,
  "opt_bits_rotate": 4,
  "raw_tls_first_param_funcs": 2,
  "opt_tls_first_param_funcs": 0,
  "tls_params_elided": 2,
  "bits_rotate_gained": 4,
  "build_raw_ok": 1,
  "build_opt_ok": 1
}
```

## Diff structurel raw → opt (grep signatures / rotate / tls / unsafe)
```
--- raw funcs ---
470:func chacha20_quarter_round(tls *libc.TLS, a uintptr, b uintptr, c uintptr, d uintptr) {
488:func chacha20_double_round(tls *libc.TLS, x uintptr) {
499:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt funcs ---
469:func chacha20_quarter_round(a uintptr, b uintptr, c uintptr, d uintptr) {
487:func chacha20_double_round(x uintptr) {
498:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt RotateLeft ---
472:	**(**uint32_t)(__ccgo_up(d)) = bits.RotateLeft32(**(**uint32_t)(__ccgo_up(d)), int(16))
475:	**(**uint32_t)(__ccgo_up(b)) = bits.RotateLeft32(**(**uint32_t)(__ccgo_up(b)), int(12))
478:	**(**uint32_t)(__ccgo_up(d)) = bits.RotateLeft32(**(**uint32_t)(__ccgo_up(d)), int(8))
481:	**(**uint32_t)(__ccgo_up(b)) = bits.RotateLeft32(**(**uint32_t)(__ccgo_up(b)), int(7))
--- opt remaining rotl/>> << patterns (sample) ---
(none)
```

## Hot path opt (première fonction non-__ccgo, extrait 80 lignes)
```go
func chacha20_quarter_round(a uintptr, b uintptr, c uintptr, d uintptr) {
	**(**uint32_t)(__ccgo_up(a)) += **(**uint32_t)(__ccgo_up(b))
	**(**uint32_t)(__ccgo_up(d)) ^= **(**uint32_t)(__ccgo_up(a))
	**(**uint32_t)(__ccgo_up(d)) = bits.RotateLeft32(**(**uint32_t)(__ccgo_up(d)), int(16))
	**(**uint32_t)(__ccgo_up(c)) += **(**uint32_t)(__ccgo_up(d))
	**(**uint32_t)(__ccgo_up(b)) ^= **(**uint32_t)(__ccgo_up(c))
	**(**uint32_t)(__ccgo_up(b)) = bits.RotateLeft32(**(**uint32_t)(__ccgo_up(b)), int(12))
	**(**uint32_t)(__ccgo_up(a)) += **(**uint32_t)(__ccgo_up(b))
	**(**uint32_t)(__ccgo_up(d)) ^= **(**uint32_t)(__ccgo_up(a))
	**(**uint32_t)(__ccgo_up(d)) = bits.RotateLeft32(**(**uint32_t)(__ccgo_up(d)), int(8))
	**(**uint32_t)(__ccgo_up(c)) += **(**uint32_t)(__ccgo_up(d))
	**(**uint32_t)(__ccgo_up(b)) ^= **(**uint32_t)(__ccgo_up(c))
	**(**uint32_t)(__ccgo_up(b)) = bits.RotateLeft32(**(**uint32_t)(__ccgo_up(b)), int(7))
}

// C documentation
//
//	/* Un double-round sur 16 words (in-place), 2 colonnes + 2 diagonales. */
func chacha20_double_round(x uintptr) {
	chacha20_quarter_round(x, x+4*4, x+8*4, x+12*4)
	chacha20_quarter_round(x+1*4, x+5*4, x+9*4, x+13*4)
	chacha20_quarter_round(x+2*4, x+6*4, x+10*4, x+14*4)
	chacha20_quarter_round(x+3*4, x+7*4, x+11*4, x+15*4)
	chacha20_quarter_round(x, x+5*4, x+10*4, x+15*4)
	chacha20_quarter_round(x+1*4, x+6*4, x+11*4, x+12*4)
	chacha20_quarter_round(x+2*4, x+7*4, x+8*4, x+13*4)
	chacha20_quarter_round(x+3*4, x+4*4, x+9*4, x+14*4)
}

```

## Checklist relecture agent (obligatoire)
- [ ] Sémantique rotate : ROL/ROR → RotateLeft correct (W-N)
- [ ] ABI tls : exportés gardent tls ; unexported T0 strip OK
- [ ] Pas de Go pointer passé en uintptr (callers futurs)
- [ ] Imports morts absents
- [ ] Motifs ratés documentés (finding proposed si récurrent)
- [ ] build_ok raw+opt

## Verdict agent
_À remplir après Read de opt.go / raw.go._

## Verdict agent (rempli) 20260810d

- Rotate : voir metrics bits_rotate_gained
- ABI tls : point fixe T0 OK sur unexported
- __ccgo_up : présent, goulot documenté F-20260810-ccgo-up-goulot
- build_ok : oui
- Pièges ccgo : F-20260810-ccgo-pitfalls-research cité
- Status : **accepté dogfood** (pas de régression KAT sur packages commités)
