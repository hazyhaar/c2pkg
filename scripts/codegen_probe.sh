#!/usr/bin/env bash
# Sonde statique de génération de code Go 1.27 (GOEXPERIMENT=simd) pour un
# paquet : par fonction, nombre d'instructions, accès pile de registres YMM
# (déversements et passage d'arguments vectoriels), appels vers le paquet
# (frontières de fonction non inlinées), sites de contrôle de bornes,
# VZEROUPPER, rotations émulées (VPSLLD/VPSRLD) et rotations par permutation
# (VPSHUFB). Oracle décidable des fiches de thésaurisation « forme émise ».
#
# Usage : scripts/codegen_probe.sh <répertoire-paquet> [seuil-instructions]
set -euo pipefail
PKG="$1"; MIN="${2:-60}"
export GOTOOLCHAIN="${GOTOOLCHAIN:-go1.27.0}" GOEXPERIMENT=simd
TMP="$(mktemp -d)"; trap 'rm -rf "$TMP"' EXIT
( cd "$PKG" && go build -gcflags='-S' . ) > "$TMP/asm.txt" 2>&1 || true
( cd "$PKG" && go build -gcflags='-m=2' . ) 2>&1 | grep 'cannot inline.*exceeds budget' > "$TMP/noinline.txt" || true

printf "%-48s %6s %7s %8s %7s %4s %9s %8s\n" fonction instr SP-YMM CALLpkg bounds VZU shift-rot vpshufb
grep -oE '^[^ ]+ STEXT' "$TMP/asm.txt" | sed 's/ STEXT//' | grep -v '\.init$\|\.func[0-9]' | while read -r fn; do
  awk -v f="$fn STEXT" 'index($0,f)==1{p=1; next} p && /STEXT/{p=0} p' "$TMP/asm.txt" > "$TMP/f.txt"
  # Seules les lignes d'instruction (préfixe adresse) sont comptées : le
  # dump -S répète chaque CALL sur une ligne de relocation qu'il faut ignorer.
  grep -E '^\s+0x[0-9a-f]+ [0-9]+ \(' "$TMP/f.txt" > "$TMP/i.txt" || true
  tot=$(wc -l < "$TMP/i.txt")
  [ "${tot:-0}" -lt "$MIN" ] && continue
  sp=$(grep -E 'VMOV' "$TMP/i.txt" | grep -c '(SP)' || true)
  cp=$(grep 'CALL' "$TMP/i.txt" | grep -vc 'runtime\.' || true)
  bd=$(grep -c 'panicBounds' "$TMP/i.txt" || true)
  vz=$(grep -c VZEROUPPER "$TMP/i.txt" || true)
  sh=$(grep -cE 'VPS[LR]L[DQ]' "$TMP/i.txt" || true)
  pb=$(grep -c VPSHUFB "$TMP/i.txt" || true)
  printf "%-48s %6s %7s %8s %7s %4s %9s %8s\n" "${fn##*/}" "$tot" "$sp" "$cp" "$bd" "$vz" "$sh" "$pb"
done | sort -k2 -nr
echo
echo "inlining refusé (coût > budget) : $(wc -l < "$TMP/noinline.txt")"
sed 's/^/  /' "$TMP/noinline.txt" | head -20
