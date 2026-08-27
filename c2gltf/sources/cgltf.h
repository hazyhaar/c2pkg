/* SPDX-License-Identifier: MIT */
#ifndef CGLTF_H
#define CGLTF_H

#include <stdint.h>
#include <stddef.h>
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef enum {
    CGLTF_COMPONENT_TYPE_INVALID = 0,
    CGLTF_COMPONENT_TYPE_R_8 = 5120,
    CGLTF_COMPONENT_TYPE_R_8U = 5121,
    CGLTF_COMPONENT_TYPE_R_16 = 5122,
    CGLTF_COMPONENT_TYPE_R_16U = 5123,
    CGLTF_COMPONENT_TYPE_R_32U = 5125,
    CGLTF_COMPONENT_TYPE_R_32F = 5126
} cgltf_component_type;

typedef enum {
    CGLTF_TYPE_INVALID = 0,
    CGLTF_TYPE_SCALAR = 1,
    CGLTF_TYPE_VEC2 = 2,
    CGLTF_TYPE_VEC3 = 3,
    CGLTF_TYPE_VEC4 = 4,
    CGLTF_TYPE_MAT2 = 5,
    CGLTF_TYPE_MAT3 = 6,
    CGLTF_TYPE_MAT4 = 7
} cgltf_type;

typedef struct {
    uint32_t magic;
    uint32_t version;
    uint32_t length;
    uint32_t json_chunk_len;
    uint32_t json_chunk_type;
    const char *json_data;
    uint32_t bin_chunk_len;
    uint32_t bin_chunk_type;
    const uint8_t *bin_data;
} cgltf_glb_header_t;

/* Parse l'en-tête GLB binaire sans aucune allocation dynamique */
bool cgltf_parse_glb_header(const uint8_t *data, size_t size, cgltf_glb_header_t *out_header);

/* Lit des composantes flottantes depuis un buffer d'accesseur binaire */
size_t cgltf_accessor_read_float(
    cgltf_component_type comp_type,
    cgltf_type type,
    const uint8_t *src,
    size_t count,
    size_t stride,
    float *out_floats
);

#ifdef __cplusplus
}
#endif

#endif /* CGLTF_H */
