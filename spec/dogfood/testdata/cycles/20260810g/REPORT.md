# Dogfood 20260810g — C extrême légitime (pas hack)

## Cartographie « quoi d’autre ? »

| Domaine | Candidat | Extrême en | Résultat cycle |
|---------|----------|------------|----------------|
| **JSON / scraping data** | cJSON | parse récursif, alloc, gotos | build OK ; 266× `__ccgo_up`, 74 goto |
| **Compression** | lz4 | memmove builtin, bit packing | **FAIL→OK** fixette iqlibc + (pas de -1 ici sur build final) |
| **Regex** | tiny-regex-c | backtracking, `ptr-1` | **FAIL→OK** `uintptr(-1)` |
| **AST / parsers** | mpc (combinators) | gros graphe d’appels | build OK ; 427× `__ccgo_up`, 276 tls |
| HTML scraping | gumbo tokenizer | multi-file lourd | non enchaîné (deps) |
| ZIP/deflate | miniz header | amalgamation | non enchaîné |

## Fixette gen (landed)

1. **`x + uintptr(-N)` → `x - uintptr(N)`** — tiny-regex `str-1`
2. **`iqlibc.__builtin_memmove` → `libc.Xmemmove`** — lz4

## Métriques (ordre de grandeur)

| Kernel | L Go | tls raw→opt | __ccgo_up | bits.Rotate |
|--------|------|-------------|-----------|-------------|
| tiny_regex | 944 | 18→8 | 75 | 0 |
| cjson | 3440 | 112→81 | 266 | 0 |
| lz4 | 3084 | 87→85 | 151 | 0 |
| mpc | 6010 | 276→257 | 427 | 0 |

## Leçons

- L’extrême utile pour le gen n’est **pas** le red-team : c’est **parseurs / compress / regex / AST**.
- Peu de rotates ; beaucoup d’**ABI ccgo** (uintptr, builtins, tls, data).
- mpc/cJSON = stress T0 + `__ccgo_up` (article « patterns extrêmes »).
