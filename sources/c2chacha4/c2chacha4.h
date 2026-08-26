/* c2chacha4 — quatre blocs ChaCha20-IETF (RFC 8439) sur huit registres 256 bits.
 *
 * Noyau étroit destiné au cycle de dogfooding sgoiter « 4 blocs sans
 * déversement » (ROADMAP 4.1) : appariement « une ligne de deux blocs par
 * registre » — chaque moitié 128 bits porte une ligne 4×32 bits d'un bloc,
 * soit 4 __m256i pour 2 blocs et 8 __m256i pour 4 blocs. Le quart de tour
 * est l'opération de c2chacha1 portée sur 256 bits ; la diagonalisation
 * utilise _mm256_shuffle_epi32 ; la sortie reconstitue 32 octets d'un bloc
 * par _mm256_permute2x128_si256. Deux passes couvrent 512 octets.
 */
#ifndef C2CHACHA4_H
#define C2CHACHA4_H

#include <stddef.h>
#include <stdint.h>

/* XOR de `n` octets (n ≤ 512) de `in` avec le keystream ChaCha20-IETF
 * des blocs `counter` … `counter+7` (clé 32 o, nonce 12 o, compteur 32 bits)
 * vers `out`. `in` et `out` peuvent être identiques. Renvoie n (borné à 512). */
size_t c2chacha4_xor_blocks(uint8_t *out, const uint8_t *in, size_t n,
                            const uint8_t key[32], const uint8_t nonce[12],
                            uint32_t counter);

#endif
