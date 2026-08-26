// SPDX-License-Identifier: Apache-2.0 OR MIT

// Package c2fused est un noyau ChaCha20 8 blocs et Poly1305 scalaire 64 bits
// entrelacés, émis par sgoiter depuis sources/c2fused/c2fused_simd.c.
// Le code de production n'existe que sous goexperiment.simd && amd64 (AVX2) :
// sur toute autre cible le paquet est vide et ne doit pas être importé — les
// consommateurs portent la même contrainte de build.
package c2fused
