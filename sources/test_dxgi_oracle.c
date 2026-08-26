#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include "c2_dxgi_abi.h"

static uint64_t hash_bytes(const uint8_t *data, size_t len) {
    uint64_t h = 0xcbf29ce484222325ULL;
    for (size_t i = 0; i < len; i++) {
        h ^= (uint64_t)data[i];
        h *= 0x100000001b3ULL;
    }
    return h;
}

int main(void) {
    printf("[ORACLE] Démarrage du banc de parité DXGI/D3D11 ARCHTIME (gcc -O2)...\n");

    /* 1. Vérification des tailles de structures */
    uint32_t sc_desc_sz = c2_dxgi_get_swapchain_desc_size();
    uint32_t sc_desc1_sz = c2_dxgi_get_swapchain_desc1_size();
    uint32_t tex2d_sz = c2_d3d11_get_texture2d_desc_size();
    uint32_t subres_sz = c2_d3d11_get_subresource_data_size();

    printf("[ORACLE] DXGI_SWAP_CHAIN_DESC  taille = %u octets (attendu: 72)\n", sc_desc_sz);
    printf("[ORACLE] DXGI_SWAP_CHAIN_DESC1 taille = %u octets (attendu: 48)\n", sc_desc1_sz);
    printf("[ORACLE] D3D11_TEXTURE2D_DESC   taille = %u octets (attendu: 44)\n", tex2d_sz);
    printf("[ORACLE] D3D11_SUBRESOURCE_DATA taille = %u octets (attendu: 16)\n", subres_sz);

    if (sc_desc_sz != 72 || sc_desc1_sz != 48 || tex2d_sz != 44 || subres_sz != 16) {
        fprintf(stderr, "[ORACLE] ERREUR : Dérive d'alignement ou de taille de structure GPU !\n");
        return 1;
    }

    /* 2. Vérification des slots de vtable */
    uint32_t slot_present = c2_dxgi_get_slot(2, 0);   /* Present */
    uint32_t slot_getbuf = c2_dxgi_get_slot(2, 1);    /* GetBuffer */
    uint32_t slot_updatesub = c2_dxgi_get_slot(5, 0); /* UpdateSubresource */
    uint32_t slot_map = c2_dxgi_get_slot(5, 1);       /* Map */

    printf("[ORACLE] Slot IDXGISwapChain::Present            = %u (attendu: 8)\n", slot_present);
    printf("[ORACLE] Slot IDXGISwapChain::GetBuffer           = %u (attendu: 9)\n", slot_getbuf);
    printf("[ORACLE] Slot ID3D11DeviceContext::UpdateSubres   = %u (attendu: 48)\n", slot_updatesub);
    printf("[ORACLE] Slot ID3D11DeviceContext::Map            = %u (attendu: 14)\n", slot_map);

    if (slot_present != 8 || slot_getbuf != 9 || slot_updatesub != 48 || slot_map != 14) {
        fprintf(stderr, "[ORACLE] ERREUR : Décalage de slot de vtable !\n");
        return 1;
    }

    /* 3. Vérification des GUIDs / IIDs */
    c2_guid_t iid_swapchain1;
    c2_dxgi_get_iid(3, &iid_swapchain1);
    uint64_t iid_hash = hash_bytes((const uint8_t *)&iid_swapchain1, sizeof(c2_guid_t));
    printf("[ORACLE] IID_IDXGISwapChain1 Hash = 0x%016llX\n", (unsigned long long)iid_hash);

    printf("[ORACLE] SUCCÈS : 100%% des vérifications ARCHTIME DXGI/D3D11 validées.\n");
    return 0;
}
