package c2dxgi

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"testing"
	"unsafe"
)

func hashBytes(data []byte) uint64 {
	var h uint64 = 0xcbf29ce484222325
	for _, b := range data {
		h ^= uint64(b)
		h *= 0x100000001b3
	}
	return h
}

func TestDXGI_ArchtimeConstantsAndSizes(t *testing.T) {
	// 1. Validation des tailles de descripteurs GPU
	szDesc1 := C2_dxgi_get_swapchain_desc1_size()
	if szDesc1 != 48 {
		t.Fatalf("Taille DXGI_SWAP_CHAIN_DESC1 = %d, attendu 48", szDesc1)
	}

	szTex2D := C2_d3d11_get_texture2d_desc_size()
	if szTex2D != 44 {
		t.Fatalf("Taille D3D11_TEXTURE2D_DESC = %d, attendu 44", szTex2D)
	}

	szSubres := C2_d3d11_get_subresource_data_size()
	if szSubres != 16 {
		t.Fatalf("Taille D3D11_SUBRESOURCE_DATA = %d, attendu 16", szSubres)
	}

	// 2. Validation des vtable slots
	slotPresent := C2_dxgi_get_slot(2, 0)
	if slotPresent != 8 {
		t.Fatalf("Slot Present = %d, attendu 8", slotPresent)
	}

	slotGetBuf := C2_dxgi_get_slot(2, 1)
	if slotGetBuf != 9 {
		t.Fatalf("Slot GetBuffer = %d, attendu 9", slotGetBuf)
	}

	slotUpdateSub := C2_dxgi_get_slot(5, 0)
	if slotUpdateSub != 48 {
		t.Fatalf("Slot UpdateSubresource = %d, attendu 48", slotUpdateSub)
	}

	slotMap := C2_dxgi_get_slot(5, 1)
	if slotMap != 14 {
		t.Fatalf("Slot Map = %d, attendu 14", slotMap)
	}
}

func TestDXGI_GUIDsAndHashes(t *testing.T) {
	var guid C2_guid_t
	C2_dxgi_get_iid(3, &guid)

	guidBytes := (*[16]byte)(unsafe.Pointer(&guid))[:]
	h := hashBytes(guidBytes)

	const expectedHash uint64 = 0x1D6CC985A0D85512
	if h != expectedHash {
		t.Fatalf("IID_IDXGISwapChain1 Hash = 0x%016X, attendu 0x%016X (oracle C gcc -O2)", h, expectedHash)
	}
}

func TestDXGIVsCOracle(t *testing.T) {
	tmpDir := t.TempDir()
	oracleBin := filepath.Join(tmpDir, "test_dxgi_oracle")

	srcs := []string{
		"/devhoros/c2simd/sources/test_dxgi_oracle.c",
		"/devhoros/c2simd/sources/c2_dxgi_abi.c",
	}

	cmd := exec.Command("gcc", append([]string{
		"-O2", "-Wall", "-Wextra", "-std=gnu99",
		"-I/devhoros/c2simd/sources",
		"-o", oracleBin,
	}, srcs...)...)

	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Échec compilation Oracle C: %v\n%s", err, string(out))
	}

	cmdRun := exec.Command(oracleBin)
	outRun, errRun := cmdRun.CombinedOutput()
	if errRun != nil {
		t.Fatalf("Échec exécution Oracle C: %v\n%s", errRun, string(outRun))
	}

	if !bytes.Contains(outRun, []byte("100% des vérifications ARCHTIME DXGI/D3D11 validées")) {
		t.Fatalf("Sortie Oracle C inattendue:\n%s", string(outRun))
	}

	t.Logf("Sortie Oracle C (gcc -O2):\n%s", string(outRun))
}
