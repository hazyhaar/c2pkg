# REVIEW dogfood — adv_pointer_alias (20260810f)

## Métriques
```json
{
  "kernel": "adv_pointer_alias",
  "stamp": "20260810f",
  "src_c": "spec/c_sources/testdata/c_sources/adv_pointer_alias.c",
  "ccgo": "/home/cl-ment/go/bin/ccgo",
  "raw_lines": 515,
  "opt_lines": 515,
  "raw_rotl_calls": 0,
  "opt_rotl_calls": 0,
  "raw_bits_rotate": 0,
  "opt_bits_rotate": 0,
  "raw_tls_first_param_funcs": 3,
  "opt_tls_first_param_funcs": 0,
  "tls_params_elided": 3,
  "bits_rotate_gained": 0,
  "build_raw_ok": 1,
  "build_opt_ok": 1
}
```

## Diff structurel raw → opt (grep signatures / rotate / tls / unsafe)
```
--- raw funcs ---
483:func adv_alias_load32(tls *libc.TLS, p uintptr) (r uint32_t) {
488:func adv_alias_store32(tls *libc.TLS, p uintptr, v uint32_t) {
495:func adv_overlap_xor(tls *libc.TLS, dst uintptr, src uintptr, n size_t) {
513:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt funcs ---
483:func adv_alias_load32(p uintptr) (r uint32_t) {
488:func adv_alias_store32(p uintptr, v uint32_t) {
495:func adv_overlap_xor(dst uintptr, src uintptr, n size_t) {
513:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt RotateLeft ---
(none)
--- opt remaining rotl/>> << patterns (sample) ---
(none)
```

## Hot path opt (première fonction non-__ccgo, extrait 80 lignes)
```go
func adv_alias_load32(p uintptr) (r uint32_t) {
	/* classic unaligned pun — UB in C strict, common in codecs */
	return **(**uint32_t)(__ccgo_up(p))
}

func adv_alias_store32(p uintptr, v uint32_t) {
	**(**uint32_t)(__ccgo_up(p)) = v
}

// C documentation
//
//	/* overlapping memcpy-like with pointer arithmetic */
func adv_overlap_xor(dst uintptr, src uintptr, n size_t) {
	var i size_t
	var v2 uintptr
	_, _ = i, v2
	i = uint64(0)
	for {
		if !(i < n) {
			break
		}
		v2 = dst + uintptr(i)
		*(*uint8_t)(unsafe.Pointer(v2)) = uint8_t(int32(*(*uint8_t)(unsafe.Pointer(v2))) ^ libc.Int32FromUint8(**(**uint8_t)(__ccgo_up(src + uintptr(i)))))
		goto _1
	_1:
		;
		i = i + 1
	}
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
