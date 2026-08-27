/* SPDX-License-Identifier: MIT */
#include "cgltf.h"
#include <string.h>

#define CGLTF_MAGIC_GLB 0x46546C67 // 'glTF'
#define CGLTF_CHUNK_JSON 0x4E4F534A // 'JSON'
#define CGLTF_CHUNK_BIN  0x004E4942 // 'BIN\0'

bool cgltf_parse_glb_header(const uint8_t *data, size_t size, cgltf_glb_header_t *out_header) {
    if (!data || size < 12 || !out_header) return false;
    memset(out_header, 0, sizeof(cgltf_glb_header_t));

    uint32_t magic = (uint32_t)(data[0] | (data[1] << 8) | (data[2] << 16) | (data[3] << 24));
    uint32_t version = (uint32_t)(data[4] | (data[5] << 8) | (data[6] << 16) | (data[7] << 24));
    uint32_t length = (uint32_t)(data[8] | (data[9] << 8) | (data[10] << 16) | (data[11] << 24));

    if (magic != CGLTF_MAGIC_GLB || version != 2 || length > size) {
        return false;
    }

    out_header->magic = magic;
    out_header->version = version;
    out_header->length = length;

    size_t offset = 12;
    // Chunk 0 : JSON
    if (offset + 8 > size) return false;
    uint32_t chunk0_len = (uint32_t)(data[offset] | (data[offset+1] << 8) | (data[offset+2] << 16) | (data[offset+3] << 24));
    uint32_t chunk0_type = (uint32_t)(data[offset+4] | (data[offset+5] << 8) | (data[offset+6] << 16) | (data[offset+7] << 24));
    offset += 8;

    if (chunk0_type != CGLTF_CHUNK_JSON || offset + chunk0_len > size) return false;
    out_header->json_chunk_len = chunk0_len;
    out_header->json_chunk_type = chunk0_type;
    out_header->json_data = (const char *)(data + offset);
    offset += chunk0_len;

    // Aligner à 4 octets
    offset = (offset + 3) & ~3;

    // Chunk 1 : BIN (Optionnel)
    if (offset + 8 <= size) {
        uint32_t chunk1_len = (uint32_t)(data[offset] | (data[offset+1] << 8) | (data[offset+2] << 16) | (data[offset+3] << 24));
        uint32_t chunk1_type = (uint32_t)(data[offset+4] | (data[offset+5] << 8) | (data[offset+6] << 16) | (data[offset+7] << 24));
        offset += 8;

        if (chunk1_type == CGLTF_CHUNK_BIN && offset + chunk1_len <= size) {
            out_header->bin_chunk_len = chunk1_len;
            out_header->bin_chunk_type = chunk1_type;
            out_header->bin_data = data + offset;
        }
    }

    return true;
}

static inline size_t cgltf_num_components(cgltf_type type) {
    switch (type) {
        case CGLTF_TYPE_SCALAR: return 1;
        case CGLTF_TYPE_VEC2: return 2;
        case CGLTF_TYPE_VEC3: return 3;
        case CGLTF_TYPE_VEC4: return 4;
        case CGLTF_TYPE_MAT2: return 4;
        case CGLTF_TYPE_MAT3: return 9;
        case CGLTF_TYPE_MAT4: return 16;
        default: return 0;
    }
}

size_t cgltf_accessor_read_float(
    cgltf_component_type comp_type,
    cgltf_type type,
    const uint8_t *src,
    size_t count,
    size_t stride,
    float *out_floats
) {
    if (!src || !out_floats || count == 0) return 0;
    size_t num_comps = cgltf_num_components(type);
    if (num_comps == 0) return 0;

    size_t elem_size = 0;
    switch (comp_type) {
        case CGLTF_COMPONENT_TYPE_R_8:
        case CGLTF_COMPONENT_TYPE_R_8U: elem_size = 1; break;
        case CGLTF_COMPONENT_TYPE_R_16:
        case CGLTF_COMPONENT_TYPE_R_16U: elem_size = 2; break;
        case CGLTF_COMPONENT_TYPE_R_32U:
        case CGLTF_COMPONENT_TYPE_R_32F: elem_size = 4; break;
        default: return 0;
    }

    if (stride < num_comps * elem_size) {
        stride = num_comps * elem_size;
    }

    size_t out_idx = 0;
    for (size_t i = 0; i < count; i++) {
        const uint8_t *elem_ptr = src + i * stride;
        for (size_t c = 0; c < num_comps; c++) {
            const uint8_t *c_ptr = elem_ptr + c * elem_size;
            float val = 0.0f;
            switch (comp_type) {
                case CGLTF_COMPONENT_TYPE_R_8:
                    val = (float)(*(const int8_t *)c_ptr);
                    break;
                case CGLTF_COMPONENT_TYPE_R_8U:
                    val = (float)(*(const uint8_t *)c_ptr);
                    break;
                case CGLTF_COMPONENT_TYPE_R_16: {
                    int16_t v16 = (int16_t)(c_ptr[0] | (c_ptr[1] << 8));
                    val = (float)v16;
                    break;
                }
                case CGLTF_COMPONENT_TYPE_R_16U: {
                    uint16_t v16u = (uint16_t)(c_ptr[0] | (c_ptr[1] << 8));
                    val = (float)v16u;
                    break;
                }
                case CGLTF_COMPONENT_TYPE_R_32U: {
                    uint32_t v32u = (uint32_t)(c_ptr[0] | (c_ptr[1] << 8) | (c_ptr[2] << 16) | (c_ptr[3] << 24));
                    val = (float)v32u;
                    break;
                }
                case CGLTF_COMPONENT_TYPE_R_32F: {
                    uint32_t v32f = (uint32_t)(c_ptr[0] | (c_ptr[1] << 8) | (c_ptr[2] << 16) | (c_ptr[3] << 24));
                    memcpy(&val, &v32f, 4);
                    break;
                }
                default: break;
            }
            out_floats[out_idx++] = val;
        }
    }

    return out_idx;
}
