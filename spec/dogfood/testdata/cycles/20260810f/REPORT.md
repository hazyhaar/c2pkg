# Dogfood 20260810f — motifs adversariaux + OSS défensif

## Périmètre (tranché)

| Demandé | Décision |
|---------|----------|
| Corpus C exploit / red-team / shellcode / C2 | **Refusé** (politique offensive cyber) |
| Lab C **adversarial** (punning, stack gros, dispatch, tls depth) | **Fait** — mêmes classes de motifs pour leçons transpile |
| OSS **défensif** bas niveau (libinjection SQLi detector, BSD) | **Fait** — patterns ressources / tables / __ccgo_up |

Un article « sécurisation du transpile » s’appuie sur AVOID + métriques ci-dessous, pas sur la reproduction d’armes.

## Lab adversarial

| Kernel | Build | Leçon |
|--------|-------|-------|
| adv_pointer_alias | OK | `*(uint32*)p` → `__ccgo_up` ; **pas de fix alignement** |
| adv_tls_depth | OK | 6/6 tls élidé (point fixe T0) |
| adv_computed_goto | OK | ROL **variable** non réécrit (attendu) |
| adv_stack_buf | OK | `[4096]uint8` stack Go conservé — **AVOID** services |

## libinjection_sqli (défensif BSD)

- ccgo : **40 444** lignes Go (data header dominant)
- tls : 60 → 41
- `__ccgo_up` : **305**
- `bits.RotateLeft` : 0
- `go build` opt : **OK**

Confirme ccgo #46 / data-heavy : dogfood utile, prod embarquée déconseillée sans compactage.

## Finding

`F-20260810-avoid-patterns-adversarial` (codified) — checklist AVOID opposable.
