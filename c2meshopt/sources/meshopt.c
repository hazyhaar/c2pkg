/* SPDX-License-Identifier: MIT */
#include "meshopt.h"
#include <stdlib.h>
#include <string.h>
#include <math.h>

#define MESHOPT_MAX_VALENCE 32

typedef struct {
    int cache_tag;
    float score;
    uint32_t active_triangles_count;
    uint32_t active_triangles_offset;
} VertexScoreInfo;

static inline float meshopt_calc_vertex_score(int cache_pos, uint32_t active_triangles) {
    if (active_triangles == 0) return -1.0f;

    float score = 0.0f;
    if (cache_pos < 0) {
        // Pas dans le cache
    } else if (cache_pos < 3) {
        // Dans les 3 premiers emplacements (triangle le plus récent)
        score += 0.75f;
    } else {
        // Calcul du score basé sur la position LRU (taille 32)
        const float scaler = 1.0f / (32 - 3);
        score += powf(1.0f - (float)(cache_pos - 3) * scaler, 1.5f);
    }

    // Bonus de valence faible
    float valence_boost = powf((float)active_triangles, -0.5f);
    score += 2.0f * valence_boost;
    return score;
}

void meshopt_optimize_vertex_cache(
    uint32_t *destination,
    const uint32_t *indices,
    size_t index_count,
    size_t vertex_count
) {
    if (!destination || !indices || index_count == 0 || vertex_count == 0) return;
    if (index_count % 3 != 0) return;

    size_t triangle_count = index_count / 3;

    // Table d'adjacence sommets -> triangles
    uint32_t *vertex_triangles = (uint32_t *)malloc(index_count * sizeof(uint32_t));
    uint32_t *vertex_offsets = (uint32_t *)calloc(vertex_count + 1, sizeof(uint32_t));
    uint32_t *vertex_counts = (uint32_t *)calloc(vertex_count, sizeof(uint32_t));
    uint8_t *triangle_emitted = (uint8_t *)calloc(triangle_count, sizeof(uint8_t));
    float *triangle_scores = (float *)malloc(triangle_count * sizeof(float));

    if (!vertex_triangles || !vertex_offsets || !vertex_counts || !triangle_emitted || !triangle_scores) {
        free(vertex_triangles);
        free(vertex_offsets);
        free(vertex_counts);
        free(triangle_emitted);
        free(triangle_scores);
        return;
    }

    // 1. Comptage des valences
    for (size_t i = 0; i < index_count; i++) {
        uint32_t v = indices[i];
        if (v < vertex_count) vertex_counts[v]++;
    }

    // 2. Préfixes d'offsets
    uint32_t offset = 0;
    for (size_t i = 0; i < vertex_count; i++) {
        vertex_offsets[i] = offset;
        offset += vertex_counts[i];
        vertex_counts[i] = 0; // Réinitialisé pour remplir vertex_triangles
    }
    vertex_offsets[vertex_count] = offset;

    // 3. Remplissage de l'adjacence
    for (size_t i = 0; i < triangle_count; i++) {
        uint32_t v0 = indices[i * 3 + 0];
        uint32_t v1 = indices[i * 3 + 1];
        uint32_t v2 = indices[i * 3 + 2];

        if (v0 < vertex_count) vertex_triangles[vertex_offsets[v0] + vertex_counts[v0]++] = (uint32_t)i;
        if (v1 < vertex_count) vertex_triangles[vertex_offsets[v1] + vertex_counts[v1]++] = (uint32_t)i;
        if (v2 < vertex_count) vertex_triangles[vertex_offsets[v2] + vertex_counts[v2]++] = (uint32_t)i;
    }

    // 4. Initialisation des scores de sommets
    VertexScoreInfo *vertex_info = (VertexScoreInfo *)malloc(vertex_count * sizeof(VertexScoreInfo));
    for (size_t i = 0; i < vertex_count; i++) {
        vertex_info[i].cache_tag = -1;
        vertex_info[i].active_triangles_count = vertex_counts[i];
        vertex_info[i].active_triangles_offset = vertex_offsets[i];
        vertex_info[i].score = meshopt_calc_vertex_score(-1, vertex_counts[i]);
    }

    // 5. Calcul initial des scores de triangles
    for (size_t i = 0; i < triangle_count; i++) {
        uint32_t v0 = indices[i * 3 + 0];
        uint32_t v1 = indices[i * 3 + 1];
        uint32_t v2 = indices[i * 3 + 2];
        float score = 0.0f;
        if (v0 < vertex_count) score += vertex_info[v0].score;
        if (v1 < vertex_count) score += vertex_info[v1].score;
        if (v2 < vertex_count) score += vertex_info[v2].score;
        triangle_scores[i] = score;
    }

    // File LRU du cache (taille 32)
    uint32_t cache[32];
    int cache_size = 0;

    size_t out_index = 0;
    size_t best_triangle = 0;

    while (out_index < index_count) {
        // Recherche du meilleur triangle dans le cache local ou recherche globale
        float best_score = -1.0f;
        best_triangle = (size_t)-1;

        // Recherche prioritaire parmi les voisins des sommets en cache
        for (int c = 0; c < cache_size; c++) {
            uint32_t v = cache[c];
            if (v >= vertex_count) continue;
            uint32_t start = vertex_info[v].active_triangles_offset;
            uint32_t end = start + vertex_info[v].active_triangles_count;
            for (uint32_t ti = start; ti < end; ti++) {
                uint32_t tri = vertex_triangles[ti];
                if (!triangle_emitted[tri] && triangle_scores[tri] > best_score) {
                    best_score = triangle_scores[tri];
                    best_triangle = tri;
                }
            }
        }

        // Si aucun triangle trouvé dans le cache, recherche séquentielle du prochain non émis
        if (best_triangle == (size_t)-1) {
            for (size_t i = 0; i < triangle_count; i++) {
                if (!triangle_emitted[i]) {
                    best_triangle = i;
                    break;
                }
            }
        }

        if (best_triangle == (size_t)-1) break;

        // Émission du triangle
        triangle_emitted[best_triangle] = 1;
        uint32_t tv[3];
        tv[0] = indices[best_triangle * 3 + 0];
        tv[1] = indices[best_triangle * 3 + 1];
        tv[2] = indices[best_triangle * 3 + 2];

        destination[out_index++] = tv[0];
        destination[out_index++] = tv[1];
        destination[out_index++] = tv[2];

        // Mise à jour de la valence des sommets du triangle émis
        for (int j = 0; j < 3; j++) {
            uint32_t v = tv[j];
            if (v < vertex_count && vertex_info[v].active_triangles_count > 0) {
                // Retrait du triangle de la liste active par permutation avec le dernier
                uint32_t start = vertex_info[v].active_triangles_offset;
                uint32_t count = vertex_info[v].active_triangles_count;
                for (uint32_t ti = 0; ti < count; ti++) {
                    if (vertex_triangles[start + ti] == (uint32_t)best_triangle) {
                        vertex_triangles[start + ti] = vertex_triangles[start + count - 1];
                        vertex_info[v].active_triangles_count--;
                        break;
                    }
                }
            }
        }

        // Mise à jour de la file LRU
        for (int j = 0; j < 3; j++) {
            uint32_t v = tv[j];
            int found_pos = -1;
            for (int c = 0; c < cache_size; c++) {
                if (cache[c] == v) {
                    found_pos = c;
                    break;
                }
            }
            if (found_pos >= 0) {
                for (int c = found_pos; c > 0; c--) {
                    cache[c] = cache[c - 1];
                }
                cache[0] = v;
            } else {
                if (cache_size < 32) cache_size++;
                for (int c = cache_size - 1; c > 0; c--) {
                    cache[c] = cache[c - 1];
                }
                cache[0] = v;
            }
        }

        // Recalcul des tags de cache et scores des sommets
        for (int c = 0; c < cache_size; c++) {
            uint32_t v = cache[c];
            if (v < vertex_count) {
                vertex_info[v].cache_tag = c;
                vertex_info[v].score = meshopt_calc_vertex_score(c, vertex_info[v].active_triangles_count);
            }
        }

        // Recalcul des scores des triangles impactés
        for (int c = 0; c < cache_size; c++) {
            uint32_t v = cache[c];
            if (v >= vertex_count) continue;
            uint32_t start = vertex_info[v].active_triangles_offset;
            uint32_t count = vertex_info[v].active_triangles_count;
            for (uint32_t ti = 0; ti < count; ti++) {
                uint32_t tri = vertex_triangles[start + ti];
                if (!triangle_emitted[tri]) {
                    uint32_t v0 = indices[tri * 3 + 0];
                    uint32_t v1 = indices[tri * 3 + 1];
                    uint32_t v2 = indices[tri * 3 + 2];
                    float sc = 0.0f;
                    if (v0 < vertex_count) sc += vertex_info[v0].score;
                    if (v1 < vertex_count) sc += vertex_info[v1].score;
                    if (v2 < vertex_count) sc += vertex_info[v2].score;
                    triangle_scores[tri] = sc;
                }
            }
        }
    }

    free(vertex_triangles);
    free(vertex_offsets);
    free(vertex_counts);
    free(triangle_emitted);
    free(triangle_scores);
    free(vertex_info);
}

float meshopt_calc_vertex_cache_stats(
    const uint32_t *indices,
    size_t index_count,
    size_t vertex_count,
    unsigned int cache_size
) {
    if (!indices || index_count == 0 || vertex_count == 0 || cache_size == 0) return 0.0f;

    uint32_t *cache = (uint32_t *)malloc(cache_size * sizeof(uint32_t));
    for (unsigned int i = 0; i < cache_size; i++) cache[i] = (uint32_t)-1;

    size_t misses = 0;
    for (size_t i = 0; i < index_count; i++) {
        uint32_t v = indices[i];
        int hit = 0;
        for (unsigned int c = 0; c < cache_size; c++) {
            if (cache[c] == v) {
                hit = 1;
                // Déplacement au sommet de la pile LRU
                for (unsigned int j = c; j > 0; j--) {
                    cache[j] = cache[j - 1];
                }
                cache[0] = v;
                break;
            }
        }
        if (!hit) {
            misses++;
            for (unsigned int j = cache_size - 1; j > 0; j--) {
                cache[j] = cache[j - 1];
            }
            cache[0] = v;
        }
    }

    free(cache);
    return (float)misses / (float)(index_count / 3);
}
