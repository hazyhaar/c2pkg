#!/usr/bin/env bash
# Porte de Phase 0 : la suite complète du pôle dans un conteneur vierge où
# /devhoros n'existe pas. Les arbres sont COPIÉS (jamais montés en écriture)
# sous /work, le cache de modules Go de l'hôte est monté en lecture seule
# (aucun réseau requis pour les dépendances), gcc et libsodium sont installés
# dans le conteneur (seul accès réseau : apt).
#
# Usage : scripts/container_gate.sh [copie] — la copie est faite dans
# $GATE_DIR (défaut : un répertoire temporaire du scratch de session).
set -euo pipefail
HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"          # /devhoros/c2simd
ROOT="$(cd "$HERE/.." && pwd)"                                     # /devhoros
GATE_DIR="${GATE_DIR:-/tmp/claude-1000/-devhoros/f658c5bc-3885-40f1-bfc7-22b3982b7763/scratchpad/gate}"
IMAGE="${IMAGE:-golang:1.27}"
MODCACHE="$(go env GOMODCACHE)"

echo "== copie des arbres du workspace c2simd vers $GATE_DIR (sans .git, sans bin/)"
# Le conteneur tourne en root : les artefacts qu'il écrit dans la copie (bin/sgoiter
# construit par le garde) appartiennent à root ; la purge tente sans, puis avec sudo.
rm -rf "$GATE_DIR" 2>/dev/null || sudo -n rm -rf "$GATE_DIR"; mkdir -p "$GATE_DIR/pkg"
rsync -a --exclude '.git' --exclude 'bin/' --exclude '_build/' "$ROOT/c2simd/" "$GATE_DIR/c2simd/"
for m in monocypher55 secretstream55 c2quic c2archsimd lzw55 ccitt55 pdfast55 tt55; do
  [ -d "$ROOT/pkg/$m" ] && rsync -a --exclude '.git' "$ROOT/pkg/$m/" "$GATE_DIR/pkg/$m/"
done
echo "== chemins hôte résiduels dans la copie (hors vecteurs de données) :"
grep -rn --include='*.go' --include='*.sh' --include='*.c' --include='Makefile' -E '"/devhoros|"/home/cl-ment|"/tmp/claude' "$GATE_DIR" | grep -v 'evasion_vectors_test.go' | head -20 || true

echo "== exécution dans $IMAGE"
docker run --rm \
  -v "$GATE_DIR:/work" \
  -v "$MODCACHE:/go/pkg/mod:ro" \
  -e GOTOOLCHAIN=local -e GOEXPERIMENT=simd -e GOPROXY=off \
  -e SECRETSTREAM_REQUIRE_SIMD=1 -e HOME=/tmp \
  -w /work/c2simd "$IMAGE" bash -c '
set -o pipefail
RC=0; rm -f /tmp/gate_rc; note(){ echo "$1-exit=$2"; [ "$2" = 0 ] || echo 1 >> /tmp/gate_rc; }  # fichier : les sous-shells ne remontent pas RC
apt-get update -qq >/dev/null && apt-get install -y -qq gcc libsodium-dev pkg-config >/dev/null 2>&1 || echo "apt: échec (gcc/libsodium absents)"
go version
echo "== c2simd (GOEXPERIMENT=simd, -race)"; go test -race -count=1 ./... 2>&1 | grep -vE "no test files" | tail -40; note c2simd ${PIPESTATUS[0]}
echo "== c2pkg"; (cd c2pkg && go test -race -count=1 ./... 2>&1 | grep -vE "no test files" | tail -30; note c2pkg ${PIPESTATUS[0]})
echo "== monocypher55 (avec SIMD)"; (cd ../pkg/monocypher55 && go test -race -count=1 ./... 2>&1 | tail -5; note mc55 ${PIPESTATUS[0]})
echo "== secretstream55 (avec SIMD, recette)"; (cd ../pkg/secretstream55 && go test -race -count=1 ./... 2>&1 | grep -vE "no test files" | tail -8; note ss55 ${PIPESTATUS[0]})
echo "== secretstream55 / monocypher55 / aeadengine SANS SIMD (repli)"; (cd ../pkg/secretstream55 && GOEXPERIMENT= SECRETSTREAM_REQUIRE_SIMD= go test -count=1 ./... 2>&1 | grep -vE "no test files" | tail -5; note ss55-nosimd ${PIPESTATUS[0]}); (cd ../pkg/monocypher55 && GOEXPERIMENT= go test -count=1 ./... 2>&1 | tail -3; note mc55-nosimd ${PIPESTATUS[0]}); (GOEXPERIMENT= go test -count=1 ./aeadengine/ 2>&1 | tail -2; note aeadengine-nosimd ${PIPESTATUS[0]})
[ -s /tmp/gate_rc ] && RC=1; echo "== verdict conteneur : RC=$RC"; exit $RC
'
