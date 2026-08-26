/* c2chacha1 — un bloc ChaCha20 (RFC 8439) sur registres 128 bits.
 *
 * Noyau étroit destiné au cycle de dogfooding sgoiter « forme émise sans
 * déversement » (ROADMAP 4.1 / F-sgoiter-chacha1-nospill) : la matrice 4×4
 * tient dans quatre __m128i, les rotations 16 et 8 passent par une permutation
 * d'octets (pshufb), 12 et 7 par décalage-ou, la diagonalisation par
 * _mm_shuffle_epi32. Aucune boucle sur plusieurs blocs : ce noyau sert le
 * bloc 0 (clé Poly1305) et les messages d'au plus 64 octets.
 */
#ifndef C2CHACHA1_H
#define C2CHACHA1_H

#include <stddef.h>
#include <stdint.h>

/* XOR de `n` octets (n ≤ 64) de `in` avec le bloc de keystream
 * ChaCha20-IETF (clé 32 o, nonce 12 o, compteur 32 bits) vers `out`.
 * `in` et `out` peuvent être identiques. Renvoie n. */
size_t c2chacha1_xor_block(uint8_t *out, const uint8_t *in, size_t n,
                           const uint8_t key[32], const uint8_t nonce[12],
                           uint32_t counter);

#endif
