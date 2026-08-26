# REVIEW dogfood — adv_computed_goto (20260810f)

## Métriques
```json
{
  "kernel": "adv_computed_goto",
  "stamp": "20260810f",
  "src_c": "spec/c_sources/testdata/c_sources/adv_computed_goto.c",
  "ccgo": "/home/cl-ment/go/bin/ccgo",
  "raw_lines": 506,
  "opt_lines": 504,
  "raw_rotl_calls": 0,
  "opt_rotl_calls": 0,
  "raw_bits_rotate": 0,
  "opt_bits_rotate": 0,
  "raw_tls_first_param_funcs": 1,
  "opt_tls_first_param_funcs": 0,
  "tls_params_elided": 1,
  "bits_rotate_gained": 0,
  "build_raw_ok": 1,
  "build_opt_ok": 1
}
```

## Diff structurel raw → opt (grep signatures / rotate / tls / unsafe)
```
--- raw funcs ---
482:func adv_dispatch(tls *libc.TLS, op uint32_t, a uint32_t, b uint32_t) (r uint32_t) {
504:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt funcs ---
480:func adv_dispatch(op uint32_t, a uint32_t, b uint32_t) (r uint32_t) {
502:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt RotateLeft ---
(none)
--- opt remaining rotl/>> << patterns (sample) ---
493:		return a<<(b&uint32(31)) | a>>(uint32(32)-b&uint32(31)) /* ROL var */
```

## Hot path opt (première fonction non-__ccgo, extrait 80 lignes)
```go
func adv_dispatch(op uint32_t, a uint32_t, b uint32_t) (r uint32_t) {
	switch op & uint32(7) {
	case uint32(0):
		return a + b
	case uint32(1):
		return a - b
	case uint32(2):
		return a ^ b
	case uint32(3):
		return a | b
	case uint32(4):
		return a & b
	case uint32(5):
		return a<<(b&uint32(31)) | a>>(uint32(32)-b&uint32(31)) /* ROL var */
	case uint32(6):
		return a * (b | uint32(1))
	default:
		return a ^ b*uint32(0x9e3779b1)
	}
	return r
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
