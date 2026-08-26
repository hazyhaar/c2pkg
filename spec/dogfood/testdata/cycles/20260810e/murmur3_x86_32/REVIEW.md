# REVIEW dogfood — murmur3_x86_32 (20260810e)

## Métriques
```json
{
  "kernel": "murmur3_x86_32",
  "stamp": "20260810e",
  "src_c": "spec/c_sources/testdata/c_sources/murmur3_x86_32.c",
  "ccgo": "/home/cl-ment/go/bin/ccgo",
  "raw_lines": 544,
  "opt_lines": 541,
  "raw_rotl_calls": 4,
  "opt_rotl_calls": 0,
  "raw_bits_rotate": 0,
  "opt_bits_rotate": 3,
  "raw_tls_first_param_funcs": 3,
  "opt_tls_first_param_funcs": 0,
  "tls_params_elided": 3,
  "bits_rotate_gained": 3,
  "build_raw_ok": 1,
  "build_opt_ok": 1
}
```

## Diff structurel raw → opt (grep signatures / rotate / tls / unsafe)
```
--- raw funcs ---
481:func rotl32(tls *libc.TLS, x uint32_t, r int8_t) (r1 uint32_t) {
485:func fmix32(tls *libc.TLS, h uint32_t) (r uint32_t) {
494:func murmur3_x86_32(tls *libc.TLS, key uintptr, len1 size_t, seed uint32_t) (r uint32_t) {
542:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt funcs ---
482:func fmix32(h uint32_t) (r uint32_t) {
491:func murmur3_x86_32(key uintptr, len1 size_t, seed uint32_t) (r uint32_t) {
539:func __ccgo_up(n uintptr) unsafe.Pointer {
--- opt RotateLeft ---
509:		k1 = bits.RotateLeft32(k1, int(int8(15)))
512:		h1 = bits.RotateLeft32(h1, int(int8(13)))
531:		k11 = bits.RotateLeft32(k11, int(int8(15)))
--- opt remaining rotl/>> << patterns (sample) ---
(none)
```

## Hot path opt (première fonction non-__ccgo, extrait 80 lignes)
```go
func fmix32(h uint32_t) (r uint32_t) {
	h = h ^ h>>int32(16)
	h = h * uint32(0x85ebca6b)
	h = h ^ h>>int32(13)
	h = h * uint32(0xc2b2ae35)
	h = h ^ h>>int32(16)
	return h
}

func murmur3_x86_32(key uintptr, len1 size_t, seed uint32_t) (r uint32_t) {
	var blocks, data, tail uintptr
	var c1, c2, h1, k1, k11 uint32_t
	var i, nblocks int32
	_, _, _, _, _, _, _, _, _, _ = blocks, c1, c2, data, h1, i, k1, k11, nblocks, tail
	data = key
	nblocks = libc.Int32FromUint64(len1 / uint64(4))
	h1 = seed
	c1 = uint32(0xcc9e2d51)
	c2 = uint32(0x1b873593)
	blocks = data + uintptr(nblocks*int32(4))
	i = -nblocks
	for {
		if !(i != 0) {
			break
		}
		k1 = **(**uint32_t)(__ccgo_up(blocks + uintptr(i)*4))
		k1 = k1 * c1
		k1 = bits.RotateLeft32(k1, int(int8(15)))
		k1 = k1 * c2
		h1 = h1 ^ k1
		h1 = bits.RotateLeft32(h1, int(int8(13)))
		h1 = h1*uint32(5) + uint32(0xe6546b64)
		goto _1
	_1:
		;
		i = i + 1
	}
	tail = data + uintptr(nblocks*int32(4))
	k11 = uint32(0)
	switch len1 & uint64(3) {
	case uint64(3):
		k11 = k11 ^ uint32(**(**uint8_t)(__ccgo_up(tail + 2)))<<int32(16) /* fallthrough */
		fallthrough
	case uint64(2):
		k11 = k11 ^ uint32(**(**uint8_t)(__ccgo_up(tail + 1)))<<int32(8) /* fallthrough */
		fallthrough
	case uint64(1):
		k11 = k11 ^ uint32(**(**uint8_t)(__ccgo_up(tail)))
		k11 = k11 * c1
		k11 = bits.RotateLeft32(k11, int(int8(15)))
		k11 = k11 * c2
		h1 = h1 ^ k11
	}
	h1 = h1 ^ uint32(len1)
	return fmix32(h1)
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
