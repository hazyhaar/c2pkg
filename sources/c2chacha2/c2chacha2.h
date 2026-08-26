/* c2chacha2 — deux blocs ChaCha20-IETF (RFC 8439) sur quatre registres 256 bits.
 *
 * Noyau étroit destiné au cycle de dogfooding sgoiter « deux blocs sans
 * déversement » : appariement « une ligne de deux blocs par registre » —
 * chaque moitié 128 bits porte une ligne 4×32 bits d'un bloc, soit
 * 4 __m256i pour 2 blocs (seize mots d'état). Le quart de tour est
 * l'opération de c2chacha1 portée sur 256 bits ; la diagonalisation
 * utilise _mm256_shuffle_epi32 (intra-voie 128 bits) ; la sortie
 * reconstitue 32 octets d'un bloc par _mm256_permute2x128_si256.
 * Un appel produit les blocs `counter` et `counter+1`, au plus 128 octets.
 */
#ifndef C2CHACHA2_H
#define C2CHACHA2_H

#include <stddef.h>
#include <stdint.h>

/* XOR de `n` octets (n ≤ 128) de `in` avec le keystream ChaCha20-IETF
 * des blocs `counter` et `counter+1` (clé 32 o, nonce 12 o, compteur 32 bits)
 * vers `out`. `in` et `out` peuvent être identiques. Renvoie n (borné à 128). */
size_t c2chacha2_xor_blocks(uint8_t *out, const uint8_t *in, size_t n,
                            const uint8_t key[32], const uint8_t nonce[12],
                            uint32_t counter);

#endif
