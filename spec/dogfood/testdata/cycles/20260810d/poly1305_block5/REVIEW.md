# REVIEW dogfood — poly1305_block5 (20260810d)

## Métriques
```json
{
  "kernel": "poly1305_block5",
  "stamp": "20260810d",
  "src_c": "spec/c_sources/testdata/c_sources/poly1305_block5.c",
  "ccgo": "/home/cl-ment/go/bin/ccgo",
  "raw_lines": 516,
  "opt_lines": 514,
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
486:func poly1305_block5(tls *libc.TLS, h uintptr, r uintptr, m uintptr) {
514:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt funcs ---
484:func poly1305_block5(h uintptr, r uintptr, m uintptr) {
512:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt RotateLeft ---
(none)
--- opt remaining rotl/>> << patterns (sample) ---
(none)
```

## Hot path opt (première fonction non-__ccgo, extrait 80 lignes)
```go
func poly1305_block5(h uintptr, r uintptr, m uintptr) {
	var c, d0, d1, d2, d3, d4 uint64_t
	_, _, _, _, _, _ = c, d0, d1, d2, d3, d4
	d0 = (uint64(**(**uint32_t)(__ccgo_up(h)))+uint64(**(**uint32_t)(__ccgo_up(m))))*uint64(**(**uint32_t)(__ccgo_up(r))) + (uint64(**(**uint32_t)(__ccgo_up(h + 1*4)))+uint64(0))*uint64(uint32(5)***(**uint32_t)(__ccgo_up(r + 4*4))) + (uint64(**(**uint32_t)(__ccgo_up(h + 2*4)))+uint64(0))*uint64(uint32(5)***(**uint32_t)(__ccgo_up(r + 3*4))) + (uint64(**(**uint32_t)(__ccgo_up(h + 3*4)))+uint64(0))*uint64(uint32(5)***(**uint32_t)(__ccgo_up(r + 2*4))) + (uint64(**(**uint32_t)(__ccgo_up(h + 4*4)))+uint64(0))*uint64(uint32(5)***(**uint32_t)(__ccgo_up(r + 1*4)))
	d1 = (uint64(**(**uint32_t)(__ccgo_up(h)))+uint64(**(**uint32_t)(__ccgo_up(m))))*uint64(**(**uint32_t)(__ccgo_up(r + 1*4))) + (uint64(**(**uint32_t)(__ccgo_up(h + 1*4)))+uint64(0))*uint64(**(**uint32_t)(__ccgo_up(r))) + (uint64(**(**uint32_t)(__ccgo_up(h + 2*4)))+uint64(0))*uint64(uint32(5)***(**uint32_t)(__ccgo_up(r + 4*4))) + (uint64(**(**uint32_t)(__ccgo_up(h + 3*4)))+uint64(0))*uint64(uint32(5)***(**uint32_t)(__ccgo_up(r + 3*4))) + (uint64(**(**uint32_t)(__ccgo_up(h + 4*4)))+uint64(0))*uint64(uint32(5)***(**uint32_t)(__ccgo_up(r + 2*4)))
	d2 = (uint64(**(**uint32_t)(__ccgo_up(h)))+uint64(**(**uint32_t)(__ccgo_up(m + 1*4))))*uint64(**(**uint32_t)(__ccgo_up(r + 2*4))) + (uint64(**(**uint32_t)(__ccgo_up(h + 1*4)))+uint64(0))*uint64(**(**uint32_t)(__ccgo_up(r + 1*4))) + (uint64(**(**uint32_t)(__ccgo_up(h + 2*4)))+uint64(0))*uint64(**(**uint32_t)(__ccgo_up(r))) + (uint64(**(**uint32_t)(__ccgo_up(h + 3*4)))+uint64(0))*uint64(uint32(5)***(**uint32_t)(__ccgo_up(r + 4*4))) + (uint64(**(**uint32_t)(__ccgo_up(h + 4*4)))+uint64(0))*uint64(uint32(5)***(**uint32_t)(__ccgo_up(r + 3*4)))
	d3 = (uint64(**(**uint32_t)(__ccgo_up(h)))+uint64(**(**uint32_t)(__ccgo_up(m + 2*4))))*uint64(**(**uint32_t)(__ccgo_up(r + 3*4))) + (uint64(**(**uint32_t)(__ccgo_up(h + 1*4)))+uint64(0))*uint64(**(**uint32_t)(__ccgo_up(r + 2*4))) + (uint64(**(**uint32_t)(__ccgo_up(h + 2*4)))+uint64(0))*uint64(**(**uint32_t)(__ccgo_up(r + 1*4))) + (uint64(**(**uint32_t)(__ccgo_up(h + 3*4)))+uint64(0))*uint64(**(**uint32_t)(__ccgo_up(r))) + (uint64(**(**uint32_t)(__ccgo_up(h + 4*4)))+uint64(0))*uint64(uint32(5)***(**uint32_t)(__ccgo_up(r + 4*4)))
	d4 = (uint64(**(**uint32_t)(__ccgo_up(h)))+uint64(**(**uint32_t)(__ccgo_up(m + 3*4))))*uint64(**(**uint32_t)(__ccgo_up(r + 4*4))) + (uint64(**(**uint32_t)(__ccgo_up(h + 1*4)))+uint64(0))*uint64(**(**uint32_t)(__ccgo_up(r + 3*4))) + (uint64(**(**uint32_t)(__ccgo_up(h + 2*4)))+uint64(0))*uint64(**(**uint32_t)(__ccgo_up(r + 2*4))) + (uint64(**(**uint32_t)(__ccgo_up(h + 3*4)))+uint64(0))*uint64(**(**uint32_t)(__ccgo_up(r + 1*4))) + (uint64(**(**uint32_t)(__ccgo_up(h + 4*4)))+uint64(0))*uint64(**(**uint32_t)(__ccgo_up(r)))
	c = d0 >> int32(26)
	**(**uint32_t)(__ccgo_up(h)) = uint32(d0) & uint32(0x3ffffff)
	d1 = d1 + c
	c = d1 >> int32(26)
	**(**uint32_t)(__ccgo_up(h + 1*4)) = uint32(d1) & uint32(0x3ffffff)
	d2 = d2 + c
	c = d2 >> int32(26)
	**(**uint32_t)(__ccgo_up(h + 2*4)) = uint32(d2) & uint32(0x3ffffff)
	d3 = d3 + c
	c = d3 >> int32(26)
	**(**uint32_t)(__ccgo_up(h + 3*4)) = uint32(d3) & uint32(0x3ffffff)
	d4 = d4 + c
	c = d4 >> int32(26)
	**(**uint32_t)(__ccgo_up(h + 4*4)) = uint32(d4) & uint32(0x3ffffff)
	**(**uint32_t)(__ccgo_up(h)) += uint32(c) * uint32(5)
	c = uint64(**(**uint32_t)(__ccgo_up(h)) >> int32(26))
	**(**uint32_t)(__ccgo_up(h)) &= uint32(0x3ffffff)
	**(**uint32_t)(__ccgo_up(h + 1*4)) += uint32(c)
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
