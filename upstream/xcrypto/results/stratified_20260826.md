# Banc stratifié — chemin sous-clé c2simd vs assembleur x/crypto — état du 2026-08-26 (HEAD c95084b5)

i9-14900K (cœurs P 0-15 épinglés), go1.27.0, GOEXPERIMENT=simd, `BenchmarkStrat`
(`c2simd/stratified_subkey_bench_test.go`), 5 répétitions × 300 ms, médianes,
synthèse `scripts/stratified_summary.py`. Référence antérieure : `stratified.md`
(2026-08-25, avant 4.0 (a)(b)(c), c2chacha8 et 2.3).

## 1. Décomposition par étage (ns, médiane ; part du total c2simd)

| Longueur | asm ns | total ns | total/asm | s1 clé Poly | s2 ctor MAC | s3 keystream | s4 MAC (propre) | résidu (attente) | coût GC | découpage 4096 |
|---|---|---|---|---|---|---|---|---|---|---|
| 64 B | 99 | 184 | ×1.86 | 67 (37 %) | 8 (4 %) | 67 (36 %) | 46 (25 %) | -3 (-2 %) | -2 | -41 |
| 1350 B | 494 | 974 | ×1.97 | 68 (7 %) | 94 (10 %) | 605 (62 %) | 305 (31 %) | -97 (-10 %) | +18 | +2 |
| 8192 B | 2440 | 3632 | ×1.49 | 67 (2 %) | 95 (3 %) | 3312 (91 %) | 1129 (31 %) | -971 (-27 %) | +2 | +107 |

| Série | coût fixe a (ns/appel) | coût par octet b (ns/o) | débit asymptotique 1/b |
|---|---|---|---|
| ref_asm | 92 | 0.2869 | 3.49 Go/s |
| total | 276 | 0.4124 | 2.43 Go/s |
| total_gcoff | 269 | 0.4128 | 2.42 Go/s |
| s3_keystream | 53 | 0.3981 | 2.51 Go/s |
| s4_mac | 122 | 0.1363 | 7.34 Go/s |

Allocations (B/op, allocs/op) : {}

Le résidu négatif à 1 350 et 8 192 B signifie que les étages mesurés isolément
(s3 + s4) dépassent le total réel : dans le chemin réel, le MAC d'une tranche
recouvre le keystream de la suivante (exécution dans le désordre) ; les parts
d'étage sont donc des bornes hautes à grande taille.

## 2. Delta avec le 2026-08-25

| Série | 25 août | 26 août | delta |
|---|---|---|---|
| total 64 B | 485 ns (×4,92 asm) | 184 ns (×1,86) | −62 % |
| total 1 350 B | 1 272 ns (×2,49) | 974 ns (×1,97) | −23 % |
| total 8 192 B | 4 806 ns (×1,91) | 3 632 ns (×1,49) | −24 % |
| coût fixe total | 502 ns | 276 ns | −45 % (asm : 92) |
| coût par octet total | 0,527 ns/o | 0,412 ns/o | −22 % (asm : 0,287) |
| s3 keystream par octet | 0,392 ns/o | 0,398 ns/o | inchangé |
| s4 MAC par octet | 0,132 ns/o | 0,136 ns/o | inchangé |
| s2 constructeur MAC (chemin long) | 154 ns | 94 ns | −39 % |
| s1 clé Poly (bloc 0) | 109 ns | 67 ns | −39 % |

## 3. Goulets restants, par ordre de poids

1. **Keystream ChaCha20 par octet : 0,398 ns/o contre 0,287 ns/o pour l'AEAD assembleur complet** — 91 % du temps à 8 192 B, 62 % à 1 350 B. Sonde codegen du noyau émis `c2pkg/c2chacha8.C2chacha8_xor_blocks` (`scripts/codegen_probe.sh`) : 2 840 instructions, **306 accès pile de registres YMM** (déversements), 30 appels non inlinés, 73 contrôles de bornes, 320 rotations émulées (VPSLLD/VPSRLD/VPOR, rotations 12 et 7) contre 160 `VPSHUFB` (16 et 8), 0 VZEROUPPER. L'état de huit blocs occupe seize YMM ; les temporaires déversent. C'est l'item 4.1 : la forme émise de `c2chacha4` (quatre blocs, 0 accès pile dans les tours) est le modèle, mais son banc n'a pas battu `c2chacha8` — le gain doit venir d'un noyau 8 blocs sans déversement (deux passes de 4 ou réordonnancement des temporaires), pas d'un retour à 4 blocs.
2. **64 octets : deux blocs ChaCha séquentiels** — s1 clé Poly (bloc 0) 67 ns + s3 keystream (bloc 1) 67 ns = 134 des 184 ns ; l'assembleur fait l'AEAD entier en 92 ns. Un appel unique produisant les blocs 0 et 1 (`c2chacha4` avec compteur 0..3, ou un noyau 2 blocs) supprime un coût fixe d'appel et une moitié de la latence ; cible ≈ −50 ns sur les messages ≤ 64 B.
3. **MAC, coût fixe 122 ns et constructeur 94 ns sur le chemin long** (`NewPoly1305QuadChain`, 658 instructions, 15 appels ; précalcul r²…r⁴ par appel — choix de sûreté, cf. mutations.md) — 13 % du total à 1 350 B. Le coût par octet du MAC (0,136 ns/o, 7,3 Go/s) n'est pas un goulet.
4. **Découpage 4 096** : +107 ns à 8 192 B (`s3b_keystream_onecall` 3 205 ns contre s3 3 312) — 3 %, secondaire.

Coût fixe : 276 ns contre 92 (×3,0, était ×5,3) ; coût par octet : ×1,44 (était ×1,78). À grande taille, seul le noyau ChaCha20 sépare encore la greffe de l'assembleur.

## 4. ERRATUM (2026-08-26, 16 h) — l'étage s3 mesurait un noyau mort

Les étages `s3_keystream` et `s3b_keystream_onecall` de `BenchmarkStrat` appelaient
`EncryptChaCha20_SIMD256` (`simd_stream.go`, noyau manuscrit, 923 instructions), qui
n'est plus appelé par `AEADSubkeyLockDst`/`UnlockDst` depuis le branchement de
`c2chacha8` (chemin réel : `xorKeystream8` → `c2pkg/c2chacha8`). Le « keystream
0,398 ns/o » des sections 1–3 est donc celui d'un noyau hors chemin ; la
sonde codegen de la section 3 (306 accès pile, etc.) porte, elle, bien sur
`c2chacha8`. Banc corrigé (`stratified_subkey_bench_test.go`, étage s3 →
`xorKeystream8`, s3b → seize appels directs) et rejoué, mêmes conditions :

| Longueur | asm ns | total ns | total/asm | s1 clé Poly | s2 ctor MAC | s3 keystream | s4 MAC (propre) | résidu (attente) | coût GC | découpage 4096 |
|---|---|---|---|---|---|---|---|---|---|---|
| 64 B | 100 | 177 | ×1.76 | 71 (40 %) | 4 (2 %) | 68 (38 %) | 45 (26 %) | -12 (-7 %) | +1 | -82 |
| 1350 B | 531 | 1004 | ×1.89 | 68 (7 %) | 95 (9 %) | 460 (46 %) | 318 (32 %) | 63 (6 %) | +12 | +8 |
| 8192 B | 2537 | 3724 | ×1.47 | 66 (2 %) | 95 (3 %) | 2479 (67 %) | 1185 (32 %) | -101 (-3 %) | +97 | +91 |

| Série | coût fixe a (ns/appel) | coût par octet b (ns/o) | débit asymptotique 1/b |
|---|---|---|---|
| ref_asm | 103 | 0.2976 | 3.36 Go/s |
| total | 279 | 0.4235 | 2.36 Go/s |
| total_gcoff | 280 | 0.4116 | 2.43 Go/s |
| s3_keystream | 54 | 0.2962 | 3.38 Go/s |
| s4_mac | 123 | 0.1432 | 6.99 Go/s |

Allocations (B/op, allocs/op) : {}

Lecture corrigée :
- le keystream réel est à **0,296 ns/o**, égal à l'AEAD assembleur complet
  (0,298) ; le noyau 8 blocs émis est au plafond par octet — `perf` topdown
  (rapport `rapport_c2simd_perf_topdown_20260826.md`) : IPC 2,84,
  `memory_bound` 15 %, ports vectoriels chargés par les rotations émulées ;
- à 8 192 B, l'écart total à l'assembleur (1 187 ns) est le **MAC non
  recouvert** (s4 = 1 185 ns, 0,145 ns/o) ; à 1 350 B, c'est le MAC (318 ns)
  plus le coût fixe (s1 + s2 ≈ 163 ns contre 103 ns d'AEAD asm complet) ;
- les verdicts des chantiers 3 et 4 (« aucun gain de keystream ») tiennent :
  ils reposaient aussi sur le banc du noyau `c2chacha8` seul (152,8 → 147,8 ns,
  149,9 → 151,7 ns à 512 B), qui n'a pas bougé ;
- le levier restant à grande taille est la **fusion ChaCha–Poly** (item 4.x,
  témoin gcc puis cycle sgoiter) : plafond −32 % à 8 Ko ; à petite taille,
  le coût fixe (`c2chacha2` pour bloc 0 + bloc 1, constructeur MAC).
