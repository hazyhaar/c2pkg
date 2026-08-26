#!/usr/bin/env bash
# Harnais amont G9 : exécute les tests et les bancs de
# golang.org/x/crypto/chacha20poly1305 (version épinglée) tels quels, dans
# trois configurations, et compare par benchstat.
#
#   asm      : chacha20poly1305_amd64.s (référence à remplacer)
#   purego   : chemin générique Go (-tags purego)
#   archsimd : greffe pur Go AVX2 c2simd (-tags archsimd, GOEXPERIMENT=simd)
#
# Règle : aucune ligne des fichiers *_test.go amont n'est modifiée. Seules
# modifications de la copie : (1) ajout de chacha20poly1305_archsimd.go,
# (2) le tag de l'assembleur reçoit `&& !archsimd`, (3) un go.work local
# rend le module c2simd visible.
#
# Usage : ./harness.sh [test|bench|all]   (défaut : all)
#   XCRYPTO_VERSION (défaut v0.54.0), BENCH_COUNT (défaut 10),
#   BENCH_TIME (défaut 300ms), PIN (ex. "taskset -c 0-15" pour épingler).
set -euo pipefail

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
if [ -d "$HERE/../../cmd/c2simd-probe" ]; then
  C2SIMD="$(cd "$HERE/../.." && pwd)"
elif [ -d "$HERE/../c2chacha8" ]; then
  C2SIMD="$(cd "$HERE/.." && pwd)"
else
  C2SIMD="$(cd "$HERE/../.." && pwd)"
fi
VER="${XCRYPTO_VERSION:-v0.54.0}"
BUILD="$HERE/_build"
CRYPTO="$BUILD/crypto"
PKG="$CRYPTO/chacha20poly1305"
RESULTS="$HERE/results"
COUNT="${BENCH_COUNT:-10}"
BTIME="${BENCH_TIME:-300ms}"
PIN="${PIN:-}"
export GOTOOLCHAIN="${GOTOOLCHAIN:-go1.27.0}"

SRC="$(go env GOMODCACHE)/golang.org/x/crypto@${VER}"
[ -d "$SRC" ] || { echo "module cache sans x/crypto@$VER : go mod download golang.org/x/crypto@$VER" >&2; exit 2; }

materialize() {
  rm -rf "$BUILD"; mkdir -p "$BUILD"
  cp -r "$SRC" "$CRYPTO"; chmod -R u+w "$CRYPTO"
  cp "$HERE/chacha20poly1305_archsimd.go" "$PKG/"
  # (2) l'assembleur cède la place sous le tag archsimd
  sed -i 's#^//go:build gc && !purego$#//go:build gc \&\& !purego \&\& !archsimd#' "$PKG/chacha20poly1305_amd64.go"
  grep -q '!archsimd' "$PKG/chacha20poly1305_amd64.go" || { echo "tag assembleur non patché" >&2; exit 3; }
  # (3) workspace local : la copie de x/crypto + c2simd / c2pkg
  cat > "$BUILD/go.work" <<EOF
go 1.27.0

toolchain go1.27.0

use (
	./crypto
	$C2SIMD
EOF
  if [ -d "$C2SIMD/c2pkg" ]; then
    echo "	$C2SIMD/c2pkg" >> "$BUILD/go.work"
  fi
  echo ")" >> "$BUILD/go.work"
  # preuve : fichiers de test identiques à l'amont
  for f in chacha20poly1305_test.go chacha20poly1305_vectors_test.go; do
    cmp -s "$SRC/chacha20poly1305/$f" "$PKG/$f" || { echo "test amont modifié : $f" >&2; exit 4; }
  done
  echo "matérialisé : $PKG (x/crypto $VER, tests amont intacts)"
}

run_tests() {
  cd "$PKG"
  echo "== tests asm";      go test -count=1 .
  echo "== tests purego";   go test -count=1 -tags purego .
  echo "== tests archsimd"; GOEXPERIMENT=simd go test -count=1 -tags archsimd .
  echo "== garde : archsimd sélectionné (le fichier greffé est bien compilé)"
  GOEXPERIMENT=simd go list -tags archsimd -f '{{join .GoFiles " "}}' . | tr ' ' '\n' | grep -q '^chacha20poly1305_archsimd.go$'
  GOEXPERIMENT=simd go list -tags archsimd -f '{{join .GoFiles " "}}' . | tr ' ' '\n' | grep -q '^chacha20poly1305_amd64.go$' && { echo "assembleur encore compilé sous archsimd" >&2; exit 5; }
  echo "ok : archsimd compile la greffe et pas l'assembleur"
}

run_bench() {
  cd "$PKG"; mkdir -p "$RESULTS"
  local common=(-run '^$' -bench 'BenchmarkChacha20Poly1305' -benchmem -count="$COUNT" -benchtime="$BTIME")
  echo "== bench asm";      $PIN go test "${common[@]}" . | tee "$RESULTS/asm.txt"
  echo "== bench purego";   $PIN go test -tags purego "${common[@]}" . | tee "$RESULTS/purego.txt"
  echo "== bench archsimd"; GOEXPERIMENT=simd $PIN go test -tags archsimd "${common[@]}" . | tee "$RESULTS/archsimd.txt"
  {
    echo "# x/crypto $VER — $(date -u +%Y-%m-%dT%H:%MZ) — $(go version) — count=$COUNT benchtime=$BTIME pin='${PIN}'"
    echo; echo "## asm (référence) vs archsimd"; benchstat "$RESULTS/asm.txt" "$RESULTS/archsimd.txt"
    echo; echo "## purego (générique) vs archsimd"; benchstat "$RESULTS/purego.txt" "$RESULTS/archsimd.txt"
  } | tee "$RESULTS/benchstat.md"
}

run_strata() {
  echo "== banc stratifié du pipeline (étages s1..s5 du scellement)"
  (cd "$C2SIMD" && ./bin/c2simd-probe strata)
}

run_mutations() {
  echo "== profilage micro-unitaire des 10 mutations de données (M1..M10)"
  (cd "$C2SIMD" && ./bin/c2simd-probe mutations)
}

case "${1:-all}" in
  test)      materialize; run_tests ;;
  bench)     [ -d "$PKG" ] || materialize; run_bench ;;
  strata)    run_strata ;;
  mutations) run_mutations ;;
  all)       materialize; run_tests; run_bench; run_strata; run_mutations ;;
  *) echo "usage: $0 [test|bench|strata|mutations|all]" >&2; exit 1 ;;
esac
