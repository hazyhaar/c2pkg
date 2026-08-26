# REVIEW dogfood — adv_stack_buf (20260810f)

## Métriques
```json
{
  "kernel": "adv_stack_buf",
  "stamp": "20260810f",
  "src_c": "spec/c_sources/testdata/c_sources/adv_stack_buf.c",
  "ccgo": "/home/cl-ment/go/bin/ccgo",
  "raw_lines": 526,
  "opt_lines": 526,
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
487:func adv_stack_workspace(tls *libc.TLS, in uintptr, n size_t) (r uint32_t) {
524:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt funcs ---
487:func adv_stack_workspace(in uintptr, n size_t) (r uint32_t) {
524:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt RotateLeft ---
(none)
--- opt remaining rotl/>> << patterns (sample) ---
(none)
```

## Hot path opt (première fonction non-__ccgo, extrait 80 lignes)
```go
func adv_stack_workspace(in uintptr, n size_t) (r uint32_t) {
	var h uint32_t
	var i size_t
	var ws [4096]uint8_t
	_, _, _ = h, i, ws
	h = uint32(0x811c9dc5)
	if n > uint64(4096) {
		n = uint64(4096)
	}
	i = uint64(0)
	for {
		if !(i < n) {
			break
		}
		ws[i] = libc.Uint8FromInt32(libc.Int32FromUint8(**(**uint8_t)(__ccgo_up(in + uintptr(i)))) ^ libc.Int32FromUint8(uint8(i)))
		goto _1
	_1:
		;
		i = i + 1
	}
	i = uint64(0)
	for {
		if !(i < n) {
			break
		}
		h = h ^ uint32(ws[i])
		h = h * uint32(0x01000193)
		goto _2
	_2:
		;
		i = i + 1
	}
	/* intentional leftover: touch end of buffer */
	h = h ^ uint32(ws[libc.Uint64FromInt64(4096)-uint64(1)])
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
