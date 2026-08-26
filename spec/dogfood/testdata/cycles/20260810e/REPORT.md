# Dogfood 20260810e — DL C légitimes (refus red-team)

## Politique

**Refus** : C pentest/red-team/exploit/shellcode/C2.  
**OK** : crypto/hash open-source permissif pour exercer le gen.

## Téléchargés / portés

| Artefact | Origine | Résultat |
|----------|---------|----------|
| `tweetnacl_dogfood.c` + `.h` | tweetnacl.cr.yp.to 20140427 + stub `randombytes` | ccgo 2787 lignes ; **L32** → 4× RotateLeft32 ; build OK |
| `murmur3_x86_32.c` | algo public domain, port C lab | 3× RotateLeft32 ; tls×3 élidé ; build OK |

ATTRIBUTION : `spec/c_sources/testdata/c_sources/ATTRIBUTION_DOWNLOADED.md`

## Findings de la passe

1. **L32 tweetnacl** — helper + `u32`=ulong LP64 → cast `uint32`/`u32` (landed)
2. Murmur confirme ROTL via fonction `rotl32` locale + DeadCode

## Relecture

- tweetnacl `core` : `u32(bits.RotateLeft32(uint32(...), int(7|9|13|18)))` — sémantique 32-bit OK
- murmur opt : rotates 15/13/15 OK ; `int(int8(15))` bruit cast ccgo (mineur)

## Non fait (volontaire)

- xxHash header 250 Ko — trop lourd pour cycle rapide
- crc32c Google = C++ 
- tout corpus offensif
