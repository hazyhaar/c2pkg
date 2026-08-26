# Banc stratifié et compteurs matériels — chemin sous-clé c2simd vs assembleur x/crypto

2026-08-25, i9-14900K (cœurs P 0-15 épinglés), go1.27.0, GOEXPERIMENT=simd.
Sources : `c2simd/stratified_subkey_bench_test.go` (5 répétitions × 300 ms, médianes,
`scripts/stratified_summary.py`), `scripts/perf_seal.sh` sur les binaires de test
amont (`asm.test`, `archsimd.test`) et sur `c2simd.test`. Les compteurs `perf` sont
des ratios (IPC, topdown, événements par opération) ; les valeurs par opération
incluent les tours d'échauffement du banc et sont des bornes hautes de ~0–30 %.

## 1. Décomposition par étage (ns, médiane ; part du total c2simd)

| Longueur | asm | total c2simd | ×asm | s1 clé Poly (bloc 0) | s2 constructeur MAC | s3 keystream | s4 MAC propre | résidu | coût GC |
|---|---|---|---|---|---|---|---|---|---|
| 64 B | 99 | 485 | 4,92 | 109 (22 %) | 186 (38 %) | 111 (23 %) | 32 (7 %) | 48 (10 %) | +31 |
| 1 350 B | 510 | 1 272 | 2,49 | 110 (9 %) | 154 (12 %) | 626 (49 %) | 324 (25 %) | 58 (5 %) | +4 |
| 8 192 B | 2 514 | 4 806 | 1,91 | 109 (2 %) | 155 (3 %) | 3 303 (69 %) | 1 168 (24 %) | 71 (1 %) | −53 |

Régression t(n) = a + b·n sur les trois longueurs :

| Série | coût fixe a (ns/appel) | coût par octet b (ns/o) | asymptote 1/b |
|---|---|---|---|
| asm (AEAD complet) | 94 | 0,296 | 3,38 Go/s |
| c2simd total | 502 | 0,527 | 1,90 Go/s |
| c2simd s3 keystream seul | 90 | 0,392 | 2,55 Go/s |
| c2simd s4 MAC seul | 251 | 0,132 | 7,58 Go/s |

Lecture : le coût fixe c2simd est 5,3× celui de l'assembleur (502 vs 94 ns) ; le coût
par octet est 1,78× (0,527 vs 0,296). Le keystream ChaCha20 seul (0,392 ns/o) est déjà
plus lent que l'AEAD assembleur complet (0,296 ns/o) : à grande taille, le noyau
ChaCha20 est la cause, pas le MAC.

## 2. Compteurs matériels sur l'interface amont (Seal, nonce 12)

| | asm 64 B | archsimd 64 B | asm 8 192 B | archsimd 8 192 B |
|---|---|---|---|---|
| ns/op | 102 | 574 | 2 574 | 4 923 |
| instructions/op | 1 385 | 10 777 (×7,8) | 46 578 (5,7/o) | 80 330 (9,8/o) |
| cycles/op | 603 | 4 443 | 13 686 | 28 535 |
| IPC | 2,30 | 2,43 | 3,40 | 2,82 |
| topdown retiring | 40,7 % | 40,5 % | 62,7 % | 46,8 % |
| topdown backend-bound | 58,8 % | 43,0 % | 36,8 % | 48,4 % |
| topdown frontend-bound | 0,5 % | 13,6 % | 0,5 % | 3,6 % |
| topdown bad-speculation | 0,0 % | 3,0 % | 0,0 % | 1,2 % |
| `assists.sse_avx_mix` / op | 0 | 2,85 | 0 | 6,42 |
| branch-misses / op | 0,008 | 2,77 | 0,21 | 8,17 |
| L1D load misses / op | 0,03 | 5,38 | 0,90 | 9,32 |
| allocations | 0 | 3 (736 B) | 0 | 3 (736 B) |

## 3. Noyaux isolés à 1 Mio (ratios robustes)

| Noyau | débit | instructions / octet | retiring | backend-bound |
|---|---|---|---|---|
| ChaCha20 c2simd `EncryptChaCha20_SIMD256` | 2,53 Go/s | 6,76 | 49 % | 51 % |
| Poly1305 c2simd QuadChain (scalaire) | 3,15 Go/s | 8,77 | 90 % | 3 % |
| Poly1305 x/crypto asm | 5,24 Go/s | 5,16 | 47 % | 50 % |
| ChaCha20 x/crypto `chacha20.XORKeyStream` (hors AEAD) | 1,04 Go/s | 8,98 | 70 % | 29 % |

Étages fixes à 64 B : s1 bloc 0 = 1 849 instr/appel pour 64 octets de keystream
(plus que l'AEAD assembleur entier, 1 385) ; s2 constructeur MAC = 3 603 instr et
~1 470 cycles/appel, 704 B alloués.

## 4. Verdict

1. **Petits messages (64 B) : coût fixe, pas noyau.** 84 % du temps c2simd est hors
   keystream : constructeur MAC (38 %), bloc 0 calculé par le chemin 8 blocs (22 %),
   résidu allocation/GC/interface (10 %). 7,8× plus d'instructions que l'assembleur
   pour le même travail.
2. **Grands messages (8 Ko) : le noyau ChaCha20.** 69 % du temps ; 6,76 instr/octet
   avec la moitié des slots bloqués en backend (dépendances mémoire : spills
   probables — à confirmer par `-gcflags=-S`). L'assembleur fait ChaCha + Poly
   fusionnés en 5,7 instr/octet à IPC 3,4.
3. **VZEROUPPER : réel mais secondaire.** 2,85 à 6,4 assistances SSE/AVX par appel
   (0 côté assembleur) ; à ~60 cycles l'assistance, cela représente ≤ 4 % des cycles
   à 64 B et ~1,4 % à 8 Ko. Ce n'est pas la cause de l'écart ; c'est un constat
   compilateur à rapporter séparément (6.3), sans en faire l'explication.
4. **Frontend et spéculation apparaissent avec les allocations** : 13,6 % frontend-bound
   et 2,8 branch-misses/op à 64 B, absents côté assembleur — signature de la
   construction d'objets (mise à zéro de 704 B, appels via interface).

## 5. Après (a)+(b)+(c) et branchement du noyau 8 blocs émis (2026-08-26, `strat_4`)

| Longueur | asm ns | total ns | total/asm | s1 clé Poly | s2 ctor MAC | s3 keystream | s4 MAC | résidu |
|---|---|---|---|---|---|---|---|---|
| 64 B | 102 | 193 | ×1,89 | 70 (36 %) | 8 (4 %) | 72 (37 %) | 49 (26 %) | −6 |
| 1350 B | 526 | 1019 | ×1,94 | 70 (7 %) | 102 (10 %) | 663 (65 %) | 338 (33 %) | −154 |
| 8192 B | 2588 | 4002 | ×1,55 | 70 (2 %) | 100 (3 %) | 3446 (86 %) | 1230 (31 %) | −844 |

Régression linéaire : coût fixe 272 ns (asm 97), coût par octet 0,458 ns/o (asm 0,304) ;
keystream seul 0,412 ns/o (2,43 Go/s), MAC seul 0,148 ns/o (6,77 Go/s). Allocations : aucune.
Le résidu négatif aux grandes tailles indique que les étages mesurés isolément ne s'additionnent
plus (recouvrement partiel pipeline entre keystream et MAC dans le chemin fusionné).

Harnais amont correspondant (`benchstat.md`) : 0 B/op, 0 allocs/op sur les 12 sous-bancs ;
Seal-64 218,7 ns (asm 109,7), Seal-1350 1 173 ns (asm 536), Seal-8192 4 039 ns (asm 2 641).
Verdict : le coût fixe est réduit de 502 à 272 ns ; le poste dominant est désormais le noyau
ChaCha 8 blocs (0,41 ns/o contre 0,30 pour l'assembleur fusionné), puis le MAC à 1 350 B.
