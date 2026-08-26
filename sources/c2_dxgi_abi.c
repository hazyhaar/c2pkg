#include "c2_dxgi_abi.h"

uint32_t c2_dxgi_get_swapchain_desc_size(void) {
    return (uint32_t)sizeof(c2_dxgi_swap_chain_desc_t);
}

uint32_t c2_dxgi_get_swapchain_desc1_size(void) {
    return (uint32_t)sizeof(c2_dxgi_swap_chain_desc1_t);
}

uint32_t c2_d3d11_get_texture2d_desc_size(void) {
    return (uint32_t)sizeof(c2_d3d11_texture2d_desc_t);
}

uint32_t c2_d3d11_get_subresource_data_size(void) {
    return (uint32_t)sizeof(c2_d3d11_subresource_data_t);
}

/* Interface IDs: 1=IUnknown, 2=IDXGISwapChain, 3=IDXGISwapChain1, 4=ID3D11Device, 5=ID3D11DeviceContext */
uint32_t c2_dxgi_get_slot(uint32_t interface_id, uint32_t method_id) {
    if (interface_id == 1) {
        /* IUnknown */
        if (method_id == 0) return C2_SLOT_IUNKNOWN_QUERY_INTERFACE;
        if (method_id == 1) return C2_SLOT_IUNKNOWN_ADD_REF;
        if (method_id == 2) return C2_SLOT_IUNKNOWN_RELEASE;
    } else if (interface_id == 2) {
        /* IDXGISwapChain */
        if (method_id == 0) return C2_SLOT_SWAPCHAIN_PRESENT;
        if (method_id == 1) return C2_SLOT_SWAPCHAIN_GET_BUFFER;
        if (method_id == 2) return C2_SLOT_SWAPCHAIN_RESIZE_BUFFERS;
    } else if (interface_id == 3) {
        /* IDXGISwapChain1 */
        if (method_id == 0) return C2_SLOT_SWAPCHAIN1_GET_DESC1;
        if (method_id == 1) return C2_SLOT_SWAPCHAIN1_PRESENT1;
    } else if (interface_id == 4) {
        /* ID3D11Device */
        if (method_id == 0) return C2_SLOT_D3D11_DEVICE_CREATE_TEXTURE2D;
        if (method_id == 1) return C2_SLOT_D3D11_DEVICE_CREATE_RENDER_TARGET_VIEW;
        if (method_id == 2) return C2_SLOT_D3D11_DEVICE_GET_IMMEDIATE_CONTEXT;
    } else if (interface_id == 5) {
        /* ID3D11DeviceContext */
        if (method_id == 0) return C2_SLOT_D3D11_CONTEXT_UPDATE_SUBRESOURCE;
        if (method_id == 1) return C2_SLOT_D3D11_CONTEXT_MAP;
        if (method_id == 2) return C2_SLOT_D3D11_CONTEXT_UNMAP;
        if (method_id == 3) return C2_SLOT_D3D11_CONTEXT_FLUSH;
    }
    return 0xFFFFFFFF;
}

/* GUIDs ARCHTIME connus (Tables de bits exactes) */
void c2_dxgi_get_iid(uint32_t iid_type, c2_guid_t *out_guid) {
    if (!out_guid) return;

    if (iid_type == 1) {
        /* IID_IUnknown: {00000000-0000-0000-C000-000000000046} */
        out_guid->data1 = 0x00000000;
        out_guid->data2 = 0x0000;
        out_guid->data3 = 0x0000;
        out_guid->data4[0] = 0xC0; out_guid->data4[1] = 0x00;
        out_guid->data4[2] = 0x00; out_guid->data4[3] = 0x00;
        out_guid->data4[4] = 0x00; out_guid->data4[5] = 0x00;
        out_guid->data4[6] = 0x00; out_guid->data4[7] = 0x46;
    } else if (iid_type == 2) {
        /* IID_IDXGISwapChain: {310d36a0-d2e7-4c0a-aa04-6a9d23b8886a} */
        out_guid->data1 = 0x310d36a0;
        out_guid->data2 = 0xd2e7;
        out_guid->data3 = 0x4c0a;
        out_guid->data4[0] = 0xaa; out_guid->data4[1] = 0x04;
        out_guid->data4[2] = 0x6a; out_guid->data4[3] = 0x9d;
        out_guid->data4[4] = 0x23; out_guid->data4[5] = 0xb8;
        out_guid->data4[6] = 0x88; out_guid->data4[7] = 0x6a;
    } else if (iid_type == 3) {
        /* IID_IDXGISwapChain1: {790a45f8-93a3-4f86-aec3-7b643bf7afa7} */
        out_guid->data1 = 0x790a45f8;
        out_guid->data2 = 0x93a3;
        out_guid->data3 = 0x4f86;
        out_guid->data4[0] = 0xae; out_guid->data4[1] = 0xc3;
        out_guid->data4[2] = 0x7b; out_guid->data4[3] = 0x64;
        out_guid->data4[4] = 0x3b; out_guid->data4[5] = 0xf7;
        out_guid->data4[6] = 0xaf; out_guid->data4[7] = 0xa7;
    } else if (iid_type == 4) {
        /* IID_ID3D11Texture2D: {6f15aaf2-d208-4e89-9ab4-489535d34f9c} */
        out_guid->data1 = 0x6f15aaf2;
        out_guid->data2 = 0xd208;
        out_guid->data3 = 0x4e89;
        out_guid->data4[0] = 0x9a; out_guid->data4[1] = 0xb4;
        out_guid->data4[2] = 0x48; out_guid->data4[3] = 0x95;
        out_guid->data4[4] = 0x35; out_guid->data4[5] = 0xd3;
        out_guid->data4[6] = 0x4f; out_guid->data4[7] = 0x9c;
    } else if (iid_type == 5) {
        /* IID_ID3D11Device: {db6f6ddb-ac77-4e88-8253-819df9bbf140} */
        out_guid->data1 = 0xdb6f6ddb;
        out_guid->data2 = 0xac77;
        out_guid->data3 = 0x4e88;
        out_guid->data4[0] = 0x82; out_guid->data4[1] = 0x53;
        out_guid->data4[2] = 0x81; out_guid->data4[3] = 0x9d;
        out_guid->data4[4] = 0xf9; out_guid->data4[5] = 0xbb;
        out_guid->data4[6] = 0xf1; out_guid->data4[7] = 0x40;
    }
}
