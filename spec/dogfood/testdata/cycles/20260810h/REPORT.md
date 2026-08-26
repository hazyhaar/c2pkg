# Dogfood 20260810h — re-gen post `__ccgo_up` + uintptr 2e tour

## Objectif

Rejouer les kernels volume (cycle g) avec gen courant : rewrite `__ccgo_up`, polish index tableau, second tour `uintptr(-N)`.

## Résultats

| Kernel | raw `__ccgo_up` | opt `__ccgo_up` | `uintptr(-` | build opt |
|--------|----------------:|----------------:|------------:|-----------|
| tiny_regex | 56 | **0** | 0 | OK |
| cjson | 198 | **0** | 0 | OK |
| lz4 | 115 | **0** | 0 | OK |
| mpc | 356 | **0** | 0 | OK |

## Bugs gen découverts / fermés dans cette passe

1. **Second tour `uintptr(-N)`** — le BinaryExpr `x + uintptr(-…)` capturé comme arg de `__ccgo_up` puis recollé dans `unsafe.Pointer(E)` n'était pas réécrit au 1er tour Apply.
2. **Polish index** `(*(*[N]T)(p))[i]` → `(*[N]T)(p)[i]` (2e tour IndexExpr).

## A1 FastXOR

Hot path raw≡opt sémantiquement (loads u64 + queue u8). Régression bench ~−12 % stable sous go1.27rc1+simd : **waiver mem-bound** (pas de rollback). Finding `F-20260810-ccgo-up-goulot` notes.
