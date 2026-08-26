# Dogfood 20260810d — nouveaux blocs + relecture + recherche ccgo

## Réponses aux questions session

### La boucle relisait-elle le transpilé à chaque passe ?
**Non, pas systématiquement avant 20260810d.** Métriques + build + KAT oui ; Read humain du hot path opt **non**.  
Corrigé : `dogfood_cycle.sh` émet `REVIEW.md` (extrait hot path + checklist) ; cette passe a **lu** chacha QR, crc32, poly1305_block5 opt.

### Recherche en ligne pièges ccgo ?
**Pas avant cette passe.** Fait 20260810 : pkg.go.dev README ABI, issues GitLab cznic/ccgo (#43 va_list, #45 vet unsafe, #46 union bloat, #11 volatile), libc CLAUDE (TLS/goroutine).  
Finding : `F-20260810-ccgo-pitfalls-research` (codified).

## Nouveaux blocs C (écrits lab, pas DL massif)

| Kernel | Intent | bits.Rotate gained | tls élidé | build |
|--------|--------|--------------------:|----------:|:-----:|
| crc32_ieee | shifts/xor, ctrl négatif rotate | 0 | 1 | OK |
| chacha20_qr | ROTL32 dense crypto | **4** | **2** (fixed-point) | OK |
| fnv1a_64 | boucle mul, ctrl négatif | 0 | 1 | OK |
| poly1305_block5 | mul 5×26 limbs | 0 | 1 | OK |

Pas de clone SimSIMD/upstream (gitignore) — blocs **bornés** sous `testdata/c_sources/`.

## Relecture hot path (agent)

### chacha20_qr opt (20260810d)
- 4× `bits.RotateLeft32` (16,12,8,7) — **correct** ROL
- `double_round` sans tls après point fixe T0
- **Goulot** : chaque word via `**(**uint32_t)(__ccgo_up(...))` — finding proposed `ccgo-up-goulot`
- `__ccgo_up` = idiom ccgo (`return unsafe.Pointer(&n)`), sémantique OK si addr est heap C

### crc32_ieee opt
- Pas de rotate à gagner (shift+mask) — normal
- Boucles `goto _1/_2` ccgo — bruit SSA
- `libc.Uint32FromInt32` encore présent (fold partiel)

### poly1305_block5 opt
- Lignes d0–d4 monstrueuses (une expr par limb) — pas de règle gen applicable sans hand-write
- Confirme : poly dual-chain reste **handwrite_pointer**, pas gen

## Fix gen dans cette passe
- **T0 point fixe** (≤8 rounds) — `F-20260810-tls-t0-fixed-point` landed
- opt/ commités régénérés ; KAT sources + kat AEAD **PASS**

## Checklist pièges ccgo (opposable)

Voir finding `F-20260810-ccgo-pitfalls-research` : Xmalloc only, TLS/goroutine, vet noise, union bloat, va_list #43, volatile, symboles minuscules v4.
