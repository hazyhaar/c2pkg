#!/usr/bin/env python3
"""Résume la sortie de `go test -bench BenchmarkStrat -benchmem` en un tableau
par longueur : étages, résidu (temps d'attente hors étages), coût GC, part de
chaque étage, ratio contre l'assembleur, et régression coût fixe / coût par
octet sur les trois longueurs. Lit la sortie brute sur stdin ou en fichier."""
import re
import sys
from collections import defaultdict
from statistics import median

pat = re.compile(r"^BenchmarkStrat/(\w+)/(\d+)-\d+\s+\d+\s+([\d.]+) ns/op(?:\s+[\d.]+ MB/s)?(?:\s+(\d+) B/op\s+(\d+) allocs/op)?")
src = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin
runs = defaultdict(list)
allocs = {}
for line in src:
    m = pat.match(line.strip())
    if not m:
        continue
    stage, size, ns, bop, aop = m.groups()
    runs[(stage, int(size))].append(float(ns))
    if aop is not None:
        allocs[(stage, int(size))] = (int(bop), int(aop))

def med(stage, size):
    v = runs.get((stage, size))
    return median(v) if v else float("nan")

sizes = sorted({s for (_, s) in runs})
stages = ["s1_polykey", "s2_polyctor", "s3_keystream", "s4_mac"]
print("| Longueur | asm ns | total ns | total/asm | s1 clé Poly | s2 ctor MAC | s3 keystream | s4 MAC (propre) | résidu (attente) | coût GC | découpage 4096 |")
print("|---|---|---|---|---|---|---|---|---|---|---|")
for n in sizes:
    asm, tot, gcoff = med("ref_asm", n), med("total", n), med("total_gcoff", n)
    s1, s2, s3, s4 = (med(s, n) for s in stages)
    s3b = med("s3b_keystream_onecall", n)
    mac_proper = s4 - s2
    residual = tot - (s1 + s3 + s4)
    gc = tot - gcoff
    chunk = s3 - s3b
    def cell(v):
        return f"{v:.0f} ({100*v/tot:.0f} %)"
    print(f"| {n} B | {asm:.0f} | {tot:.0f} | ×{tot/asm:.2f} | {cell(s1)} | {cell(s2)} | {cell(s3)} | {cell(mac_proper)} | {cell(residual)} | {gc:+.0f} | {chunk:+.0f} |")

# Régression t(n) = a + b·n (moindres carrés sur les longueurs présentes).
def fit(stage):
    xs = [n for n in sizes if (stage, n) in runs]
    ys = [med(stage, n) for n in xs]
    k = len(xs)
    if k < 2:
        return float("nan"), float("nan")
    mx, my = sum(xs) / k, sum(ys) / k
    b = sum((x - mx) * (y - my) for x, y in zip(xs, ys)) / sum((x - mx) ** 2 for x in xs)
    return my - b * mx, b

print()
print("| Série | coût fixe a (ns/appel) | coût par octet b (ns/o) | débit asymptotique 1/b |")
print("|---|---|---|---|")
for stage in ["ref_asm", "total", "total_gcoff", "s3_keystream", "s4_mac"]:
    a, b = fit(stage)
    print(f"| {stage} | {a:.0f} | {b:.4f} | {1/b if b else float('nan'):.2f} Go/s |")
print()
print("Allocations (B/op, allocs/op) :", {f"{k[0]}/{k[1]}": v for k, v in sorted(allocs.items()) if v[1]})
