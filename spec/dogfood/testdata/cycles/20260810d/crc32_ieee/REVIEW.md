# REVIEW dogfood — crc32_ieee (20260810d)

## Métriques
```json
{
  "kernel": "crc32_ieee",
  "stamp": "20260810d",
  "src_c": "spec/c_sources/testdata/c_sources/crc32_ieee.c",
  "ccgo": "/home/cl-ment/go/bin/ccgo",
  "raw_lines": 516,
  "opt_lines": 516,
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
482:func crc32_ieee(tls *libc.TLS, data uintptr, len1 size_t) (r uint32_t) {
514:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt funcs ---
482:func crc32_ieee(data uintptr, len1 size_t) (r uint32_t) {
514:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt RotateLeft ---
(none)
--- opt remaining rotl/>> << patterns (sample) ---
(none)
```

## Hot path opt (première fonction non-__ccgo, extrait 80 lignes)
```go
func crc32_ieee(data uintptr, len1 size_t) (r uint32_t) {
	var b int32
	var crc, mask uint32_t
	var i size_t
	_, _, _, _ = b, crc, i, mask
	crc = uint32(0xFFFFFFFF)
	i = uint64(0)
	for {
		if !(i < len1) {
			break
		}
		crc = crc ^ uint32(**(**uint8_t)(__ccgo_up(data + uintptr(i))))
		b = 0
		for {
			if !(b < int32(8)) {
				break
			}
			mask = libc.Uint32FromInt32(-int32(crc & uint32(1)))
			crc = crc>>int32(1) ^ uint32(0xEDB88320)&mask
			goto _2
		_2:
			;
			b = b + 1
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return ^crc
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
