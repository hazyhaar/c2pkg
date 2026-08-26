# REVIEW dogfood — adv_tls_depth (20260810f)

## Métriques
```json
{
  "kernel": "adv_tls_depth",
  "stamp": "20260810f",
  "src_c": "spec/c_sources/testdata/c_sources/adv_tls_depth.c",
  "ccgo": "/home/cl-ment/go/bin/ccgo",
  "raw_lines": 497,
  "opt_lines": 495,
  "raw_rotl_calls": 0,
  "opt_rotl_calls": 0,
  "raw_bits_rotate": 0,
  "opt_bits_rotate": 0,
  "raw_tls_first_param_funcs": 6,
  "opt_tls_first_param_funcs": 0,
  "tls_params_elided": 6,
  "bits_rotate_gained": 0,
  "build_raw_ok": 1,
  "build_opt_ok": 1
}
```

## Diff structurel raw → opt (grep signatures / rotate / tls / unsafe)
```
--- raw funcs ---
471:func leaf(tls *libc.TLS, x uint32_t) (r uint32_t) {
475:func m1(tls *libc.TLS, x uint32_t) (r uint32_t) {
479:func m2(tls *libc.TLS, x uint32_t) (r uint32_t) {
483:func m3(tls *libc.TLS, x uint32_t) (r uint32_t) {
487:func m4(tls *libc.TLS, x uint32_t) (r uint32_t) {
491:func adv_tls_depth(tls *libc.TLS, x uint32_t) (r uint32_t) {
495:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt funcs ---
469:func leaf(x uint32_t) (r uint32_t) {
473:func m1(x uint32_t) (r uint32_t) {
477:func m2(x uint32_t) (r uint32_t) {
481:func m3(x uint32_t) (r uint32_t) {
485:func m4(x uint32_t) (r uint32_t) {
489:func adv_tls_depth(x uint32_t) (r uint32_t) {
493:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt RotateLeft ---
(none)
--- opt remaining rotl/>> << patterns (sample) ---
(none)
```

## Hot path opt (première fonction non-__ccgo, extrait 80 lignes)
```go
func leaf(x uint32_t) (r uint32_t) {
	return x*uint32(0x9e3779b1) ^ x>>int32(16)
}

func m1(x uint32_t) (r uint32_t) {
	return leaf(x + uint32(1))
}

func m2(x uint32_t) (r uint32_t) {
	return m1(x ^ uint32(0x85ebca6b))
}

func m3(x uint32_t) (r uint32_t) {
	return m2(x + uint32(0xc2b2ae35))
}

func m4(x uint32_t) (r uint32_t) {
	return m3(x ^ x<<int32(13))
}

func adv_tls_depth(x uint32_t) (r uint32_t) {
	return m4(x)
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
