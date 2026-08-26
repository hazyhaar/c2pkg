#!/usr/bin/env bash
# prove35.sh — preuve consommateur vierge hors workspace via proxy file://.
# Simule la publication v0.1.0 ; ne la remplace pas.
set -euo pipefail

export GOTOOLCHAIN=go1.27.0
export GOEXPERIMENT=simd

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
C2SIMD_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"
DEVHOROS="$(cd "$C2SIMD_ROOT/.." && pwd)"
SCRATCH="$C2SIMD_ROOT/bin/scratch/proxy35"
SRC="$SCRATCH/src"
PROXY="$SCRATCH/proxy"
CONSUMER="$SCRATCH/consumer"
GOCACHE_ISO="$SCRATCH/gocache"
VERSION=v0.1.0

log() { printf '\n======== %s ========\n' "$*"; }

fail() {
  printf '\nROUGE: %s\n' "$*" >&2
  exit 1
}

log "prépare scratch $SCRATCH"
rm -rf "$SRC" "$PROXY" "$CONSUMER" "$GOCACHE_ISO"
mkdir -p "$SRC" "$PROXY" "$CONSUMER" "$GOCACHE_ISO"

log "copie monocypher55 (sans exclusion volumineuse)"
rsync -a --delete \
  --exclude '.git/' \
  "$DEVHOROS/pkg/monocypher55/" "$SRC/monocypher55/"

log "copie secretstream55 (sans exclusion volumineuse)"
rsync -a --delete \
  --exclude '.git/' \
  "$DEVHOROS/pkg/secretstream55/" "$SRC/secretstream55/"

log "copie c2pkg (sans exclusion volumineuse)"
rsync -a --delete \
  --exclude '.git/' \
  "$C2SIMD_ROOT/c2pkg/" "$SRC/c2pkg/"

log "copie c2simd (exclut bin/, .cache/, .config/, go/, spec/dogfood/testdata/cycles, upstream/xcrypto/_build)"
rsync -a --delete \
  --exclude '.git/' \
  --exclude 'bin/' \
  --exclude '.cache/' \
  --exclude '.config/' \
  --exclude 'go/' \
  --exclude 'spec/dogfood/testdata/cycles/' \
  --exclude 'upstream/xcrypto/_build/' \
  "$C2SIMD_ROOT/" "$SRC/c2simd/"

log "copie secretstream55 : suppression du replace"
if ! grep -q '^replace ' "$SRC/secretstream55/go.mod"; then
  fail "replace attendu absent de la copie secretstream55"
fi
grep -v '^replace ' "$SRC/secretstream55/go.mod" > "$SRC/secretstream55/go.mod.tmp"
mv "$SRC/secretstream55/go.mod.tmp" "$SRC/secretstream55/go.mod"
if grep -q '^replace ' "$SRC/secretstream55/go.mod"; then
  fail "replace encore présent après suppression"
fi
# v0.0.0 n'existe pas hors replace : la copie exige la version publiée sur ce proxy.
sed -i 's|code.hazyhaar.fr/devhoros/pkg/monocypher55 v0.0.0|code.hazyhaar.fr/devhoros/pkg/monocypher55 v0.1.0|' \
  "$SRC/secretstream55/go.mod"
echo '--- secretstream55/go.mod (copie) ---'
cat "$SRC/secretstream55/go.mod"

log "copie c2simd : require github.com/hazyhaar/c2pkg v0.1.0"
if grep -q 'github.com/hazyhaar/c2pkg' "$SRC/c2simd/go.mod"; then
  echo "require c2pkg déjà présent"
else
  awk '
    /^require \($/ && !done {
      print
      print "\tgithub.com/hazyhaar/c2pkg v0.1.0"
      done=1
      next
    }
    { print }
  ' "$SRC/c2simd/go.mod" > "$SRC/c2simd/go.mod.tmp"
  mv "$SRC/c2simd/go.mod.tmp" "$SRC/c2simd/go.mod"
fi
echo '--- c2simd/go.mod (copie) ---'
cat "$SRC/c2simd/go.mod"

log "construit mkproxy (GOWORK=off GOPROXY=off)"
(
  cd "$SCRIPT_DIR"
  GOPROXY=off GOWORK=off go build -o "$SCRATCH/mkproxy" .
)

log "produit les zips v0.1.0 et l'arbre proxy"
"$SCRATCH/mkproxy" -proxy "$PROXY" -version "$VERSION" \
  "code.hazyhaar.fr/devhoros/pkg/monocypher55=$SRC/monocypher55" \
  "code.hazyhaar.fr/devhoros/pkg/secretstream55=$SRC/secretstream55" \
  "code.hazyhaar.fr/devhoros/c2simd=$SRC/c2simd" \
  "github.com/hazyhaar/c2pkg=$SRC/c2pkg"

log "ls -R proxy | head"
ls -R "$PROXY" | head

log "purge ciblée des quatre modules@v0.1.0 dans le cache (piège d'extraction antérieure)"
GOMODCACHE="$(GOWORK=off go env GOMODCACHE)"
purge_mod() {
  local p
  for p in "$@"; do
    [ -e "$p" ] || continue
    chmod -R u+w "$p" >/dev/null 2>&1 || true
    rm -rf "$p" >/dev/null 2>&1 || fail "purge impossible: $p"
  done
}
purge_mod \
  "$GOMODCACHE/code.hazyhaar.fr/devhoros/pkg/monocypher55@${VERSION}" \
  "$GOMODCACHE/code.hazyhaar.fr/devhoros/pkg/secretstream55@${VERSION}" \
  "$GOMODCACHE/code.hazyhaar.fr/devhoros/c2simd@${VERSION}" \
  "$GOMODCACHE/github.com/hazyhaar/c2pkg@${VERSION}"
# les globs du cache download peuvent ne rien matcher
shopt -s nullglob
purge_mod \
  "$GOMODCACHE/cache/download/code.hazyhaar.fr/devhoros/pkg/monocypher55/@v/${VERSION}."* \
  "$GOMODCACHE/cache/download/code.hazyhaar.fr/devhoros/pkg/secretstream55/@v/${VERSION}."* \
  "$GOMODCACHE/cache/download/code.hazyhaar.fr/devhoros/c2simd/@v/${VERSION}."* \
  "$GOMODCACHE/cache/download/github.com/hazyhaar/c2pkg/@v/${VERSION}."*
shopt -u nullglob
echo "purge ok"

log "crée consommateur example.org/consumer"
cat > "$CONSUMER/go.mod" <<'EOF'
module example.org/consumer

go 1.27
EOF

cat > "$CONSUMER/main.go" <<'EOF'
package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"code.hazyhaar.fr/devhoros/c2simd/aeadengine"
	"code.hazyhaar.fr/devhoros/pkg/secretstream55"
)

func roundtrip(msg []byte) ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	enc, err := secretstream55.NewEncryptorWithEngine(&buf, key, aeadengine.Engine{})
	if err != nil {
		return nil, err
	}
	if _, err := enc.Write(msg); err != nil {
		return nil, err
	}
	if err := enc.Close(); err != nil {
		return nil, err
	}
	dec, err := secretstream55.NewDecryptorWithEngine(bytes.NewReader(buf.Bytes()), key, aeadengine.Engine{})
	if err != nil {
		return nil, err
	}
	return io.ReadAll(dec)
}

func main() {
	msg := []byte("proxy35 consumer")
	out, err := roundtrip(msg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if !bytes.Equal(out, msg) {
		fmt.Fprintln(os.Stderr, "roundtrip mismatch")
		os.Exit(1)
	}
	fmt.Println("ok")
}
EOF

cat > "$CONSUMER/main_test.go" <<'EOF'
package main

import (
	"bytes"
	"testing"
)

func TestRoundtrip(t *testing.T) {
	msg := []byte("test payload for proxy35")
	out, err := roundtrip(msg)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(out, msg) {
		t.Fatalf("got %q", out)
	}
}
EOF

PROXY_URL="file://${PROXY}"
export GOPROXY="${PROXY_URL},off"
export GONOSUMDB='*'
export GONOSUMCHECK=1
export GOFLAGS=-mod=mod
export GOWORK=off
export GOCACHE="$GOCACHE_ISO"

echo "GOPROXY=$GOPROXY"
echo "GOCACHE=$GOCACHE"
echo "GOWORK=$GOWORK"

cd "$CONSUMER"

log "go mod tidy"
go mod tidy || fail "go mod tidy"
echo '--- consumer go.mod après tidy ---'
cat go.mod

log "go build ./..."
go build ./... || fail "go build ./..."
echo "go build ./... VERT"

log "go test ./..."
go test ./... || fail "go test ./..."

log "cross GOOS=linux GOARCH=arm64 go build ./..."
GOOS=linux GOARCH=arm64 go build ./... || fail "cross linux/arm64"

log "cross GOOS=js GOARCH=wasm go build ./..."
GOOS=js GOARCH=wasm go build ./... || fail "cross js/wasm"

log "cross GOOS=windows GOARCH=amd64 go build ./..."
GOOS=windows GOARCH=amd64 go build ./... || fail "cross windows/amd64"

log "preuve proxy35 VERTE"
echo "consumer go.mod:"
cat go.mod
