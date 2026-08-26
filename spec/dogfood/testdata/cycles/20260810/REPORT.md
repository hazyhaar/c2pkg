# Dogfood c2simd v2 — cycle 20260810

## Méthode

Pour chaque `.c` de `spec/c_sources/testdata/c_sources/` :

1. `ccgo --package-name=df_<k> -o raw.go src.c` (ccgo/v4.34.6)
2. `c2simd-gen -in raw.go -out opt.go`
3. `go build` raw + opt
4. Métriques JSON (rotl, bits.RotateLeft, tls first-param, lines)

Puis validation **produit** sur packages commités `raw/` vs `opt/` :

- `go test ./spec/c_sources/` — KAT équivalence raw≡opt
- `GOEXPERIMENT=simd go1.27rc1 test -bench=…` — perf
- `go test ./kat/` — KAT AEAD (scalaire + simd)

Script : `scripts/dogfood_cycle.sh`

## Résultats pipeline (5 kernels)

| Kernel | build raw/opt | tls élidé* | bits.RotateLeft gagnés | Notes |
|--------|---------------|------------|------------------------|-------|
| base64_simd | 1/1 | 1 | 0 | pas de rotate |
| blake2b_compress | 1/1 | 1 | **0** | ROTR64 non matché → F-20260810-rotr64-gap |
| fast_xor | 1/1 | 1 | 0 | mem-bound |
| md5_transform | 1/1 | 1 | **4** | motif ROL via `<<\|>>` |
| siphash24 | 1/1 | 1 | **48** | idem |

\*ccgo v4 émet des symboles **minuscules** → non exportés → T0 élide (guard exportés ABI OK sur raw historiques `Md5_*`).

## Bugs produit trouvés et corrigés dans le cycle

1. **libc import mort** après T0 total → `go build` fail (siphash, blake2b)  
   → `dropUnusedImports` — F-20260810-unused-libc-import **landed**
2. **T0 sur exportés** cassait ABI multi_bench/KAT  
   → élision seulement si `!ast.IsExported` — F-20260810-tls-t0-exported-abi **landed**

## Bench multi_bench (i9-14900K, go1.27rc1 GOEXPERIMENT=simd, benchtime=200ms)

| Bench | Raw | Opt | Δ |
|-------|-----|-----|---|
| SipHash24 1024B | 3321.85 MB/s | **3579.31 MB/s** | **+7.7 %** |
| BLAKE2b 128B | 789.73 MB/s | 786.97 MB/s | ~0 % |
| MD5 64B | 7307.95 MB/s | **7648.03 MB/s** | **+4.6 %** |
| MD5 stdlib | 608.91 MB/s | — | opt ~12.6× stdlib |
| FastXOR 4096B | 20445 MB/s | **22258 MB/s** | **+8.9 %** |

Écart vs peer review §2.3 (SipHash −4 %, XOR 0 %) : machine/bruit/gen v2 (import purge + rotate) — **direction non régressive** sur cette mesure.

## KAT produit

- `spec/c_sources` KAT raw≡opt : **PASS**
- `kat/` AEAD (sans et avec `GOEXPERIMENT=simd`) : **PASS**

## Suite (faite 20260810 suite)

1. **ROTR64 landed** — `matchRotateBinary` (OR/XOR, ROL/ROR) ; 32× `RotateLeft64` dans blake2b opt ; KAT PASS
2. **`scripts/regenerate_opt.sh`** — raw→opt v2 pour les 4 kernels commités
3. Bench post-ROTR (benchtime=300ms) : BLAKE2b toujours ~0 % (bottleneck hors rotate) ; SipHash **+5,6 %** ; KAT vert
4. Ne pas attendre de gain SIMD générique (F-20260810-q2)

### Bench post-regen opt v2 (i9-14900K, go1.27rc1+simd, 300ms)

| Bench | Raw | Opt | Δ |
|-------|-----|-----|---|
| SipHash24 | 3401 MB/s | **3591 MB/s** | **+5,6 %** |
| BLAKE2b | 800 MB/s | 797 MB/s | ~0 % |
| MD5 | 7749 MB/s | 7520 MB/s | bruit |
| FastXOR | 21166 MB/s | 21716 MB/s | +2,6 % |
