// SPDX-License-Identifier: Apache-2.0 OR MIT

// Package c2chacha4 est un noyau ChaCha20 émis par sgoiter depuis
// sources/c2chacha4/c2chacha4_simd.c. Le code de production n'existe que sous
// goexperiment.simd && amd64 (AVX2) : sur toute autre cible le paquet est
// vide et ne doit pas être importé — les consommateurs portent la même
// contrainte de build.
package c2chacha4
