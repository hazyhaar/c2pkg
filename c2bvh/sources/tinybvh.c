/* SPDX-License-Identifier: MIT */
#include "tinybvh.h"
#include <math.h>

static inline bvh_vec3_t bvh_sub(bvh_vec3_t a, bvh_vec3_t b) {
    bvh_vec3_t r = { a.x - b.x, a.y - b.y, a.z - b.z };
    return r;
}

static inline bvh_vec3_t bvh_cross(bvh_vec3_t a, bvh_vec3_t b) {
    bvh_vec3_t r = {
        a.y * b.z - a.z * b.y,
        a.z * b.x - a.x * b.z,
        a.x * b.y - a.y * b.x
    };
    return r;
}

static inline float bvh_dot(bvh_vec3_t a, bvh_vec3_t b) {
    return a.x * b.x + a.y * b.y + a.z * b.z;
}

bool bvh_intersect_ray_triangle(const bvh_ray_t *ray, const bvh_tri_t *tri, bvh_hit_t *hit) {
    if (!ray || !tri || !hit) return false;

    const float EPSILON = 1e-7f;
    bvh_vec3_t edge1 = bvh_sub(tri->v1, tri->v0);
    bvh_vec3_t edge2 = bvh_sub(tri->v2, tri->v0);
    bvh_vec3_t h = bvh_cross(ray->direction, edge2);
    float a = bvh_dot(edge1, h);

    if (a > -EPSILON && a < EPSILON) {
        return false; // Rayon parallèle au triangle
    }

    float f = 1.0f / a;
    bvh_vec3_t s = bvh_sub(ray->origin, tri->v0);
    float u = f * bvh_dot(s, h);

    if (u < 0.0f || u > 1.0f) {
        return false;
    }

    bvh_vec3_t q = bvh_cross(s, edge1);
    float v = f * bvh_dot(ray->direction, q);

    if (v < 0.0f || u + v > 1.0f) {
        return false;
    }

    float t = f * bvh_dot(edge2, q);
    if (t > ray->t_min && t < ray->t_max) {
        hit->t = t;
        hit->u = u;
        hit->v = v;
        hit->tri_id = tri->id;
        hit->hit = true;
        return true;
    }

    return false;
}

bool bvh_intersect_ray_aabb(const bvh_ray_t *ray, const bvh_aabb_t *aabb, float *t_near) {
    if (!ray || !aabb) return false;

    float tx1 = (aabb->min.x - ray->origin.x) * ray->inv_dir.x;
    float tx2 = (aabb->max.x - ray->origin.x) * ray->inv_dir.x;

    float tmin = tx1 < tx2 ? tx1 : tx2;
    float tmax = tx1 > tx2 ? tx1 : tx2;

    float ty1 = (aabb->min.y - ray->origin.y) * ray->inv_dir.y;
    float ty2 = (aabb->max.y - ray->origin.y) * ray->inv_dir.y;

    float tymin = ty1 < ty2 ? ty1 : ty2;
    float tymax = ty1 > ty2 ? ty1 : ty2;

    if ((tmin > tymax) || (tymin > tmax)) return false;

    if (tymin > tmin) tmin = tymin;
    if (tymax < tmax) tmax = tymax;

    float tz1 = (aabb->min.z - ray->origin.z) * ray->inv_dir.z;
    float tz2 = (aabb->max.z - ray->origin.z) * ray->inv_dir.z;

    float tzmin = tz1 < tz2 ? tz1 : tz2;
    float tzmax = tz1 > tz2 ? tz1 : tz2;

    if ((tmin > tzmax) || (tzmin > tmax)) return false;

    if (tzmin > tmin) tmin = tzmin;
    if (tzmax < tmax) tmax = tzmax;

    if (tmax < ray->t_min || tmin > ray->t_max) return false;

    if (t_near) *t_near = tmin > ray->t_min ? tmin : ray->t_min;
    return true;
}
