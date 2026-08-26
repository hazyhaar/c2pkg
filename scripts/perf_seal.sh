#!/usr/bin/env bash
# Compteurs matériels (cœurs P) sur un sous-banc amont, pour un binaire de
# test donné : cycles, instructions, décomposition topdown, assistances
# SSE/AVX (transitions coûteuses), branch-misses, L1D misses.
# Usage : sudo scripts/perf_seal.sh <binaire.test> <regex-bench> [benchtime]
set -euo pipefail
BIN="$1"; BENCH="$2"; BT="${3:-2s}"
EV="cpu_core/cycles/,cpu_core/instructions/,cpu_core/topdown-retiring/,cpu_core/topdown-fe-bound/,cpu_core/topdown-be-bound/,cpu_core/topdown-bad-spec/,cpu_core/assists.sse_avx_mix/,cpu_core/branch-misses/,cpu_core/L1-dcache-load-misses/"
perf stat -x, -e "$EV" -- taskset -c 0-15 "$BIN" -test.run '^$' -test.bench "$BENCH" -test.benchtime "$BT" -test.benchmem 2>&1
