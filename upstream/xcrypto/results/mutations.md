# Mutations de données de la chaîne ChaCha20-Poly1305 — nature, contraintes, leviers

2026-08-26. Référence pour 4.1 (noyau 8 blocs), 4.x (fusion) et la Phase 5 (format
maison). Ce document ne remplace aucun item de la feuille de route
(`pkg/secretstream55/ROADMAP.md`) : il en fixe le critère de recevabilité et
l'ordre. Toutes les lignes s'appuient sur du code lu : greffe amont
(`upstream/xcrypto/chacha20poly1305_archsimd.go`), `c2simd/aead_subkey.go`,
`sources/c2chacha1/c2chacha1_simd.c`, `c2pkg/c2poly1305/poly1305_gen.go`,
`x/crypto/chacha20poly1305/xchacha20poly1305.go`, `pkg/secretstream55/secretstream.go`.

## 1. Critère de non-rupture (le tuyau lisse)

Une transformation de la chaîne peut changer de forme (source C, règle d'émission,
grille ARCHTIME) si et seulement si :

1. **Parité bit-exacte** : les oracles existants restent verts (RFC 8439, x/crypto
   102 vecteurs et `TestRandom`, `gcc -O2`, parité 0–1 088).
2. **Zéro allocation** : `harness.sh bench` rend 0 B/op, 0 allocs/op.
3. **Temps constant** : aucune lecture en mémoire à une adresse dérivée d'une valeur
   secrète (clé, flux, accumulateur, sous-clé), aucun branchement sur une valeur
   secrète. La ligne de cache touchée et le chemin d'exécution sont observables
   par un tiers ; une table dans un registre (`pshufb`) ne l'est pas. Un choix
   aléatoire entre tables ajoute du bruit à la fuite sans la supprimer.

Corollaire : les seuls appariements admissibles sur du secret sont ceux dont la
grille tient dans un registre et dont l'application est une instruction à
adresse fixe. Sur des données publiques (longueurs, dispositions, trames), toute
forme est admissible.

## 2. Nature des mutations

Légende : **dynamique** = dépend des entrées et se calcule à l'appel ;
**déterministe** = forme ou constante fixée avant toute donnée, aujourd'hui portée
par la source et lue à l'appel ; **archtime** = résolue à l'émission par `sgoiter`,
rien n'en subsiste à l'exécution.

| Mutation (fichier, fonction) | En clair | Nature |
|---|---|---|
| Nonce 24 o : `HChaCha20(key, nonce[0:16]) → subkey[32]` | Clé dérivée de la clé et des 16 premiers octets du nonce | dynamique |
| `derivedNonce = 0⁴ ‖ nonce[16:24]` | Nonce court de 12 octets | dynamique (disposition déterministe, RFC) |
| Greffe : `sliceForAppend(dst, len+16)` → fenêtres chiffré / étiquette | Tampon de sortie agrandi et découpé | dynamique |
| `alias.InexactOverlap` / `AnyOverlap` → panic | Refus du chevauchement partiel | dynamique |
| Contrôles `len(subkey)==32`, `len(nonce12)==12` | Forme des paramètres | dynamique |
| `c2chacha1(polyKeyBlock, zeros, 64, subkey, nonce12, 0)` | Bloc 0 : flux brut, matière de la clé d'étiquette | dynamique |
| `row3 = LE32(ctr) ‖ nonce12` | Quatrième ligne de la matrice | dynamique (disposition déterministe, RFC) |
| `v0 = load(sigma)` | Constante « expand 32-byte k » | déterministe (RFC) |
| `v1, v2 = load(key)`, `v3 = load(row3)` | Clé et ligne compteur-nonce en registres | dynamique |
| `rot16 = load(mask16)`, `rot8 = load(mask8)` | Masques de rotation par mélange d'octets | déterministe |
| Boucle des 10 double-tours → 10 corps à plat | Déroulage par `VectorLoopUnroll` | **archtime** |
| Étapes add / xor / rot16 (pshufb) / rot12 (shift-or) / rot8 / rot7 | Quart de tour sur les quatre colonnes | dynamique |
| Immédiats `0x39/0x4e/0x93` → `PermuteScalars(…)` | Décodage des permutations (`MM_ShuffleEpi32`) | **archtime** |
| Diagonalisation et retour (`pshufd`) | Alignement des diagonales | dynamique (permutation déterministe, RFC) |
| Addition finale de l'état rechargé | Non-inversibilité ; bloc de flux | dynamique |
| `n == 64` : XOR vectoriel direct ; `n < 64` : tampon `ks[64]` + XOR partiel | Mélange du message au flux | dynamique |
| `polyKey = bloc0[0:32]` | Première moitié du bloc 0 | dynamique (découpe déterministe, RFC) |
| `cap(dst) >= len` → réutilisation, sinon `make` | Tampon jugé à la capacité | dynamique |
| Branche courte si `len(pt)+len(ad) < 512` ; noyau étroit si `len(pt) ≤ 64` | Choix de branche sur message + AD ; choix de noyau sur message | dynamique (seuils déterministes, mesurés) |
| `init` : `r = LE(polyKey[0:16]) & clamp`, `pad = polyKey[16:32]`, `h = 0` | Multiplicateur borné, seconde moitié gardée | dynamique (masque déterministe, RFC) |
| `update(ad)` + padding ; `update(ct)` + padding | Absorption avec complément à 16 | dynamique |
| `h = (h + m + 2^128)·r mod 2^130−5` par bloc | Évaluation polynomiale | dynamique |
| `sizeBlock = LE64(len ad) ‖ LE64(len ct)` | Les longueurs entrent dans l'étiquette | dynamique (disposition déterministe, RFC) |
| `final` : réduction, `+ pad`, 16 octets bas | Étiquette | dynamique |
| `Crypto_wipe(ctx)` : écriture de zéros + `runtime.KeepAlive` (`poly1305_gen.go:229`) | Effacement émis, attestation `objdump` due (2.4) | dynamique |
| Branche longue : `NewPoly1305V2` (r, r², r³, r⁴ en 26 bits × 5 voies) | Puissances par appel — choix de sûreté, jamais un cache | dynamique |
| Tranches de 4 096 : keystream puis `Update` | Chiffrement et absorption par tranche en cache | dynamique (taille déterministe) |
| Greffe : `copy(tag, mac)` | Étiquette à la suite du chiffré | dynamique |
| `Open` : séparation étiquette, recalcul, `ConstantTimeCompare`, zéro sur échec, puis flux | Vérifier avant de déchiffrer | dynamique |
| `secretstream55` : `subkey` une fois par flux ; `AD = BE64(seq) ‖ ad` ; `frame = BE32(len) ‖ ct ‖ mac` | Numéro d'ordre dans les données associées, trame préfixée | dynamique (format déterministe, maison) |

Trois mutations seulement sont archtime ; les « déterministe » sont les candidats à
une génération — mais voir §4 : seules celles dont la spécification n'existe pas
ailleurs méritent un générateur.

## 3. Scission est / doit

**Doit rester dynamique** — la sortie dépend du secret ou de l'histoire :
`HChaCha20`, les tours, l'addition finale, le XOR ; le bloc 0 et la clé
d'étiquette ; l'accumulation Poly1305, la réduction, `+ pad`, l'étiquette ; les
puissances de `r` (recalcul par appel = choix de sûreté, jamais un cache) ; la
comparaison en temps constant, la mise à zéro sur échec, le déchiffrement après
vérification ; l'effacement ; le tirage du nonce et l'incrément de séquence (les
seules mutations dynamiques **et** non rejouables).

**Est dynamique par habitude** — aucune entrée secrète ni historique n'y entre :
contrôles de bornes sur `[]byte` (20 dans `c2chacha1`, 4 dans `AEADSubkeyLockDst`),
test `n == 64` et choix de noyau, longueurs de padding et bloc de tailles, seuil 512
et tranches de 4 096, jugement à la capacité, copies à gabarit fixe. C'est le
périmètre sur lequel `sgoiter` et `c2archsimd` ont prise sans toucher au tuyau.

## 4. `c2archsimd` + ARCHTIME : usages pertinents

`c2archsimd` est le générateur de grilles du pôle (LUT16/256 par `pshufb`, hex,
varint, sRGB). Sur cette chaîne, dans le périmètre du §1 :

| Cible | Grille à générer | Instruction | Statut |
|---|---|---|---|
| Noyau 8 blocs (4.1) : compteur par voie `[0…7]` | vecteur de compteurs | `Add` sur `Uint32x8` | à faire — cœur de 4.1 |
| Noyau 8 blocs : sortie de 8 blocs contigus | masques d'entrelacement (`unpack`, `permute2x128`) | permutations 256 bits | à faire — supprime le `StoreArray → pile` de `F-sgoiter-simd-direct-projections` |
| Fusion ChaCha–Poly (4.x) | table d'ordonnancement (quelle multiplication Poly entre quels pas ChaCha) | — | après 4.1 ; à générer, jamais écrite à la main |
| Format maison v2 (Phase 5) | descripteur CUE du format → constantes Go, spécification, vecteurs | — | à la Phase 5, après un oracle sur 4.1 |
| Masques `rot16`/`rot8`, `pshufd`, `sigma`, clamp | — | — | **exclus** : fixés par la RFC 8439 et tenus par l'oracle gcc ; un générateur n'ajouterait pas de preuve |
| XOR, Poly V2 26 bits, comparaison de l'étiquette | — | — | rien à griller |

Règle retenue : **ne générer que ce dont la spécification n'existe pas ailleurs.**
La RFC n'a pas besoin d'un second législateur.

## 5. Propositions sur les mutations « par habitude » (ordre arrêté après adversarial)

1. **4.1 d'abord** (source C 8 blocs, témoin `gcc -O2 -mavx2`, grilles du §4, même
   cycle que `c2chacha1`).
2. En parallèle : tests de contrat `cap` — faits, `c2simd/aead_dst_contract_test.go`.
3. `FixedArrayParam` (`T nom[N]` → `*[N]T`) seulement après une sonde sur `c2chacha1`
   seul (`codegen_probe.sh`, bornes avant/après, `s1` du banc stratifié) ; règle
   globale seulement si la sonde bouge. L'entrée `xor_block64` ne vaut que si (3) a payé.
4. Garde de dérive des seuils 512 / 4 096 sous `-tags gate` : instrument, pas optimisation.
5. Padding, bloc de tailles, dispositions RFC : rien.
