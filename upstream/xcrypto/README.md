# Harnais amont x/crypto/chacha20poly1305 (G9)

But : produire les seuls chiffres publiables sur golang/go#80881 — ceux que
l'équipe Go peut rejouer en une commande — et prouver que la greffe pur Go
AVX2 passe la suite de tests amont **sans qu'une ligne des tests soit
modifiée**.

## Ce que contient ce dossier

| Fichier | Rôle |
|---|---|
| `chacha20poly1305_archsimd.go` | Greffe : `seal`/`open` du chemin IETF délégués à `c2simd.AEADSubkeyLockDst` / `AEADSubkeyUnlockDst`. Tag `archsimd && goexperiment.simd && amd64 && !purego`. Reprend les contrôles d'aliasing du générique. |
| `harness.sh` | Copie `x/crypto@v0.54.0` depuis le cache de modules dans `_build/crypto`, ajoute la greffe, complète le tag de l'assembleur (`&& !archsimd`), pose un `go.work` local, vérifie par `cmp` que les tests amont sont intacts, joue tests et bancs dans trois configurations, écrit `results/`. |
| `results/` | `asm.txt`, `purego.txt`, `archsimd.txt` (sorties brutes `go test -bench`) et `benchstat.md`. Versionnés : ce sont les chiffres cités. |
| `_build/` | Matérialisation, ignorée par git. |

## Configurations comparées

- **asm** : `chacha20poly1305_amd64.s`, l'implémentation que la proposition
  vise à remplacer.
- **purego** : `-tags purego`, chemin générique Go (repli amont).
- **archsimd** : `-tags archsimd` + `GOEXPERIMENT=simd`, la greffe.

Bancs amont : `BenchmarkChacha20Poly1305`, longueurs 64, 1350 et 8192
octets, Seal/Open, nonce 12 et 24 octets (`-X`). Aucun point à 32 octets ni
à 1 Mo n'existe dans le banc amont ; un chiffre à ces tailles ne vient pas
d'ici.

## Usage

```
./harness.sh test            # matérialise + tests des trois configurations
PIN="taskset -c 0-15" ./harness.sh bench   # bancs épinglés, benchstat
./harness.sh all
```

Variables : `XCRYPTO_VERSION`, `BENCH_COUNT` (10), `BENCH_TIME` (300ms), `PIN`.
