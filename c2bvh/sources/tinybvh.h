/* SPDX-License-Identifier: MIT */
#ifndef TINYBVH_H
#define TINYBVH_H

#include <stdint.h>
#include <stddef.h>
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    float x, y, z;
} bvh_vec3_t;

typedef struct {
    bvh_vec3_t min;
    bvh_vec3_t max;
} bvh_aabb_t;

typedef struct {
    bvh_vec3_t v0, v1, v2;
    uint32_t id;
} bvh_tri_t;

typedef struct {
    bvh_vec3_t origin;
    bvh_vec3_t direction;
    bvh_vec3_t inv_dir;
    float t_min;
    float t_max;
} bvh_ray_t;

typedef struct {
    float t;
    float u, v;
    uint32_t tri_id;
    bool hit;
} bvh_hit_t;

typedef struct {
    bvh_aabb_t aabb;
    uint32_t left_first; // index nœud gauche ou premier triangle
    uint32_t tri_count;  // 0 si nœud interne, >0 si feuille
} bvh_node_t;

/* Test d'intersection rayon-triangle Möller-Trumbore */
bool bvh_intersect_ray_triangle(const bvh_ray_t *ray, const bvh_tri_t *tri, bvh_hit_t *hit);

/* Test d'intersection rayon-AABB (Slab test) */
bool bvh_intersect_ray_aabb(const bvh_ray_t *ray, const bvh_aabb_t *aabb, float *t_near);

#ifdef __cplusplus
}
#endif

#endif /* TINYBVH_H */
