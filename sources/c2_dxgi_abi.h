#ifndef C2_DXGI_ABI_H
#define C2_DXGI_ABI_H

#include <stdint.h>
#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

/* ========================================================================= */
/* 1. ÉNUMÉRATIONS & CONSTANTES ARCHTIME                                    */
/* ========================================================================= */

/* DXGI Formats */
#define C2_DXGI_FORMAT_UNKNOWN                   0
#define C2_DXGI_FORMAT_R32G32B32A32_FLOAT        2
#define C2_DXGI_FORMAT_R16G16B16A16_FLOAT        10
#define C2_DXGI_FORMAT_R8G8B8A8_UNORM            28
#define C2_DXGI_FORMAT_R8G8B8A8_UNORM_SRGB       29
#define C2_DXGI_FORMAT_B8G8R8A8_UNORM            87
#define C2_DXGI_FORMAT_B8G8R8A8_UNORM_SRGB       91

/* DXGI Swap Effects */
#define C2_DXGI_SWAP_EFFECT_DISCARD              0
#define C2_DXGI_SWAP_EFFECT_SEQUENTIAL           1
#define C2_DXGI_SWAP_EFFECT_FLIP_SEQUENTIAL      3
#define C2_DXGI_SWAP_EFFECT_FLIP_DISCARD         4

/* DXGI Swap Chain Flags */
#define C2_DXGI_SWAP_CHAIN_FLAG_NONPREROTATED    1
#define C2_DXGI_SWAP_CHAIN_FLAG_ALLOW_MODE_SWITCH 2
#define C2_DXGI_SWAP_CHAIN_FLAG_GDI_COMPATIBLE   4
#define C2_DXGI_SWAP_CHAIN_FLAG_FRAME_LATENCY_WAITABLE_OBJECT 64
#define C2_DXGI_SWAP_CHAIN_FLAG_ALLOW_TEARING    512

/* DXGI Usage */
#define C2_DXGI_USAGE_RENDER_TARGET_OUTPUT       0x00000020
#define C2_DXGI_USAGE_SHADER_INPUT               0x00000010

/* DXGI Scaling & Alpha */
#define C2_DXGI_SCALING_STRETCH                  0
#define C2_DXGI_SCALING_NONE                     1
#define C2_DXGI_SCALING_ASPECT_RATIO_STRETCH     2

#define C2_DXGI_ALPHA_MODE_UNSPECIFIED           0
#define C2_DXGI_ALPHA_MODE_PREMULTIPLIED         1
#define C2_DXGI_ALPHA_MODE_STRAIGHT              2
#define C2_DXGI_ALPHA_MODE_IGNORE                3

/* D3D11 Driver & Device Flags */
#define C2_D3D_DRIVER_TYPE_UNKNOWN               0
#define C2_D3D_DRIVER_TYPE_HARDWARE              1
#define C2_D3D_DRIVER_TYPE_WARP                  5

#define C2_D3D11_CREATE_DEVICE_SINGLETHREADED    0x0001
#define C2_D3D11_CREATE_DEVICE_DEBUG             0x0002
#define C2_D3D11_CREATE_DEVICE_BGRA_SUPPORT      0x0020

#define C2_D3D_FEATURE_LEVEL_11_0                0xb000
#define C2_D3D_FEATURE_LEVEL_11_1                0xb100

/* D3D11 Usage & Bind */
#define C2_D3D11_USAGE_DEFAULT                   0
#define C2_D3D11_USAGE_IMMUTABLE                 1
#define C2_D3D11_USAGE_DYNAMIC                   2
#define C2_D3D11_USAGE_STAGING                   3

#define C2_D3D11_BIND_VERTEX_BUFFER              0x0001
#define C2_D3D11_BIND_INDEX_BUFFER               0x0002
#define C2_D3D11_BIND_CONSTANT_BUFFER            0x0004
#define C2_D3D11_BIND_SHADER_RESOURCE            0x0008
#define C2_D3D11_BIND_RENDER_TARGET              0x0020
#define C2_D3D11_BIND_DEPTH_STENCIL              0x0040

#define C2_D3D11_CPU_ACCESS_WRITE                0x00010000
#define C2_D3D11_CPU_ACCESS_READ                 0x00020000

#define C2_D3D11_MAP_READ                        1
#define C2_D3D11_MAP_WRITE                       2
#define C2_D3D11_MAP_READ_WRITE                  3
#define C2_D3D11_MAP_WRITE_DISCARD               4
#define C2_D3D11_MAP_WRITE_NO_OVERWRITE          5

/* ========================================================================= */
/* 2. SLOTS DE VTABLE ARCHTIME (STABLES & IMMUABLES)                         */
/* ========================================================================= */

/* IUnknown Slots */
#define C2_SLOT_IUNKNOWN_QUERY_INTERFACE         0
#define C2_SLOT_IUNKNOWN_ADD_REF                 1
#define C2_SLOT_IUNKNOWN_RELEASE                 2

/* IDXGISwapChain Slots */
#define C2_SLOT_SWAPCHAIN_PRESENT                8
#define C2_SLOT_SWAPCHAIN_GET_BUFFER             9
#define C2_SLOT_SWAPCHAIN_SET_FULLSCREEN_STATE   10
#define C2_SLOT_SWAPCHAIN_GET_FULLSCREEN_STATE   11
#define C2_SLOT_SWAPCHAIN_GET_DESC               12
#define C2_SLOT_SWAPCHAIN_RESIZE_BUFFERS         13
#define C2_SLOT_SWAPCHAIN_RESIZE_TARGET          14
#define C2_SLOT_SWAPCHAIN_GET_CONTAINING_OUTPUT  15
#define C2_SLOT_SWAPCHAIN_GET_FRAME_STATISTICS   16
#define C2_SLOT_SWAPCHAIN_GET_LAST_PRESENT_COUNT 17

/* IDXGISwapChain1 Slots */
#define C2_SLOT_SWAPCHAIN1_GET_DESC1             18
#define C2_SLOT_SWAPCHAIN1_GET_FULLSCREEN_DESC   19
#define C2_SLOT_SWAPCHAIN1_GET_HWND              20
#define C2_SLOT_SWAPCHAIN1_GET_CORE_WINDOW       21
#define C2_SLOT_SWAPCHAIN1_PRESENT1              22
#define C2_SLOT_SWAPCHAIN1_SET_BACKGROUND_COLOR  25
#define C2_SLOT_SWAPCHAIN1_SET_ROTATION          27

/* ID3D11Device Slots */
#define C2_SLOT_D3D11_DEVICE_CREATE_BUFFER              3
#define C2_SLOT_D3D11_DEVICE_CREATE_TEXTURE2D           5
#define C2_SLOT_D3D11_DEVICE_CREATE_SHADER_RESOURCE_VIEW 7
#define C2_SLOT_D3D11_DEVICE_CREATE_RENDER_TARGET_VIEW  9
#define C2_SLOT_D3D11_DEVICE_GET_IMMEDIATE_CONTEXT      40

/* ID3D11DeviceContext Slots */
#define C2_SLOT_D3D11_CONTEXT_MAP                       14
#define C2_SLOT_D3D11_CONTEXT_UNMAP                     15
#define C2_SLOT_D3D11_CONTEXT_OM_SET_RENDER_TARGETS     33
#define C2_SLOT_D3D11_CONTEXT_RS_SET_VIEWPORTS          44
#define C2_SLOT_D3D11_CONTEXT_UPDATE_SUBRESOURCE        48
#define C2_SLOT_D3D11_CONTEXT_CLEAR_RENDER_TARGET_VIEW  50
#define C2_SLOT_D3D11_CONTEXT_FLUSH                     111

/* ========================================================================= */
/* 3. STRUCTURES C99 DE DESCRIPTEURS GPU                                    */
/* ========================================================================= */

typedef struct {
    uint32_t count;
    uint32_t quality;
} c2_dxgi_sample_desc_t;

typedef struct {
    uint32_t numerator;
    uint32_t denominator;
} c2_dxgi_rational_t;

typedef struct {
    uint32_t width;
    uint32_t height;
    c2_dxgi_rational_t refresh_rate;
    uint32_t format;
    uint32_t scanline_ordering;
    uint32_t scaling;
} c2_dxgi_mode_desc_t;

typedef struct {
    c2_dxgi_mode_desc_t buffer_desc;
    c2_dxgi_sample_desc_t sample_desc;
    uint32_t buffer_usage;
    uint32_t buffer_count;
    uint64_t output_window; /* HWND uintptr (8 octets sur x86_64) */
    int32_t  windowed;
    uint32_t swap_effect;
    uint32_t flags;
} c2_dxgi_swap_chain_desc_t;

typedef struct {
    uint32_t width;
    uint32_t height;
    uint32_t format;
    int32_t  stereo;
    c2_dxgi_sample_desc_t sample_desc;
    uint32_t buffer_usage;
    uint32_t buffer_count;
    uint32_t scaling;
    uint32_t swap_effect;
    uint32_t alpha_mode;
    uint32_t flags;
} c2_dxgi_swap_chain_desc1_t;

typedef struct {
    uint32_t width;
    uint32_t height;
    uint32_t mip_levels;
    uint32_t array_size;
    uint32_t format;
    c2_dxgi_sample_desc_t sample_desc;
    uint32_t usage;
    uint32_t bind_flags;
    uint32_t cpu_access_flags;
    uint32_t misc_flags;
} c2_d3d11_texture2d_desc_t;

typedef struct {
    uint64_t p_sys_mem;       /* const void* uintptr */
    uint32_t sys_mem_pitch;
    uint32_t sys_mem_slice_pitch;
} c2_d3d11_subresource_data_t;

typedef struct {
    uint64_t p_data;          /* void* uintptr */
    uint32_t row_pitch;
    uint32_t depth_pitch;
} c2_d3d11_mapped_subresource_t;

/* GUID / IID (16 octets) */
typedef struct {
    uint32_t data1;
    uint16_t data2;
    uint16_t data3;
    uint8_t  data4[8];
} c2_guid_t;

/* ========================================================================= */
/* 4. FONCTIONS DE VALIDATION ARCHTIME                                       */
/* ========================================================================= */

uint32_t c2_dxgi_get_swapchain_desc_size(void);
uint32_t c2_dxgi_get_swapchain_desc1_size(void);
uint32_t c2_d3d11_get_texture2d_desc_size(void);
uint32_t c2_d3d11_get_subresource_data_size(void);

uint32_t c2_dxgi_get_slot(uint32_t interface_id, uint32_t method_id);
void     c2_dxgi_get_iid(uint32_t iid_type, c2_guid_t *out_guid);

#ifdef __cplusplus
}
#endif

#endif /* C2_DXGI_ABI_H */
