/* c2chacha8 — huit blocs ChaCha20-IETF (RFC 8439) sur registres 256 bits.
 *
 * Noyau étroit destiné au cycle de dogfooding sgoiter « un mot par registre »
 * (ROADMAP 4.1) : seize __m256i, le registre k portant le mot k des huit
 * blocs (voie j = bloc j). Pas de diagonalisation par permutation : les
 * quarts de tour s'écrivent sur les indices colonnes puis diagonales.
 * La transposition de sortie reste en registres (unpack + permute2x128).
 */
#ifndef C2CHACHA8_H
#define C2CHACHA8_H

#include <stddef.h>
#include <stdint.h>

/* XOR de `n` octets (n ≤ 512) de `in` avec le keystream ChaCha20-IETF
 * des blocs `counter` … `counter+7` (clé 32 o, nonce 12 o, compteur 32 bits)
 * vers `out`. `in` et `out` peuvent être identiques. Renvoie n (borné à 512). */
size_t c2chacha8_xor_blocks(uint8_t *out, const uint8_t *in, size_t n,
                            const uint8_t key[32], const uint8_t nonce[12],
                            uint32_t counter);

#endif
