# REVIEW dogfood — fnv1a_64 (20260810d)

## Métriques
```json
{
  "kernel": "fnv1a_64",
  "stamp": "20260810d",
  "src_c": "spec/c_sources/testdata/c_sources/fnv1a_64.c",
  "ccgo": "/home/cl-ment/go/bin/ccgo",
  "raw_lines": 504,
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
482:func fnv1a_64(tls *libc.TLS, data uintptr, len1 size_t) (r uint64_t) {
502:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt funcs ---
482:func fnv1a_64(data uintptr, len1 size_t) (r uint64_t) {
502:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt RotateLeft ---
(none)
--- opt remaining rotl/>> << patterns (sample) ---
(none)
```

## Hot path opt (première fonction non-__ccgo, extrait 80 lignes)
```go
func fnv1a_64(data uintptr, len1 size_t) (r uint64_t) {
	var h uint64_t
	var i size_t
	_, _ = h, i
	h = uint64(14695981039346656037)
	i = uint64(0)
	for {
		if !(i < len1) {
			break
		}
		h = h ^ uint64(**(**uint8_t)(__ccgo_up(data + uintptr(i))))
		h = uint64(h * libc.Uint64FromUint64(1099511628211))
		goto _1
	_1:
		;
		i = i + 1
	}
	return h
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
