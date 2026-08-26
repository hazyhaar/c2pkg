package findings

"F-20260824-photometric-avx2-gather-asm": #Finding & {
	id:      "F-20260824-photometric-avx2-gather-asm"
	kernel:  "c2painter_photometric"
	stage:   "handwrite"
	symptom: "L'API standard Go simd/archsimd ne disposant pas d'instruction matérielle gather, la transposition Go pure impose un déversement en pile à chaque lecture de LUT (1.29 Go/s). Seul un noyau assembleur direct AVX2 exploitant VPGATHERDD permet d'atteindre 3.76 Go/s."
	evidence: {
		file_line:    "pkg/c2painter/c2archsimd_avx2_amd64.s:137"
		bench_before: "1320 MB/s (sgoiter portable scalar)"
		bench_after:  "3763 MB/s (AVX2 ASM scanline span)"
		kat:          "pass"
	}
	lever:  "handwrite"
	action: "Déploiement d'une architecture à deux étages : noyau assembleur direct amd64.s pour la saturation AVX2 sur x86-64, et repli universel Go pur transpilé par sgoiter sans CGO pour les cibles Wasm, ARM64, RISC-V et x86 sans AVX2."
	status: "landed"
	notes:  "Parité bit-exacte validée contre l'oracle GCC -O2 et le span sgoiter sur l'ensemble des largeurs de trame (1 à 3840 pixels). Offsets de structure validés formellement par unsafe.Offsetof."
}
