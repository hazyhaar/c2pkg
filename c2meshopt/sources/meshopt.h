/* SPDX-License-Identifier: MIT */
#ifndef MESHOPT_H
#define MESHOPT_H

#include <stdint.h>
#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

/*
 * meshopt_optimize_vertex_cache réordonne les indices de triangles pour maximiser
 * le taux de succès (hit rate) dans le cache matériel de sommets (L1 / Post-Transform Cache).
 */
void meshopt_optimize_vertex_cache(
    uint32_t *destination,
    const uint32_t *indices,
    size_t index_count,
    size_t vertex_count
);

/*
 * meshopt_calc_vertex_cache_stats calcule le ratio ACMR (Average Cache Miss Ratio)
 * sur un cache de taille cache_size (ex: 16 ou 32 entrées).
 */
float meshopt_calc_vertex_cache_stats(
    const uint32_t *indices,
    size_t index_count,
    size_t vertex_count,
    unsigned int cache_size
);

#ifdef __cplusplus
}
#endif

#endif /* MESHOPT_H */
