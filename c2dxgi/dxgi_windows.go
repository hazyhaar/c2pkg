//go:build windows

package c2dxgi

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	modD3D11 = syscall.NewLazyDLL("d3d11.dll")
	modDXGI  = syscall.NewLazyDLL("dxgi.dll")

	procD3D11CreateDevice = modD3D11.NewProc("D3D11CreateDevice")
	procCreateDXGIFactory = modDXGI.NewProc("CreateDXGIFactory1")
)

type Device struct {
	ptr uintptr
}

type DeviceContext struct {
	ptr uintptr
}

type SwapChain struct {
	ptr uintptr
}

// CallVTable invoque directement le slot de vtable d'une interface COM sans overhead runtime.
func CallVTable(comPtr uintptr, slot uint32, args ...uintptr) (uintptr, uintptr, error) {
	if comPtr == 0 {
		return 0, 0, fmt.Errorf("c2dxgi: comPtr est nul")
	}
	// Déréférencement de la vtable : *(**uintptr)(comPtr)
	vtablePtr := *(*uintptr)(unsafe.Pointer(comPtr))
	methodPtr := *(*uintptr)(unsafe.Pointer(vtablePtr + uintptr(slot)*unsafe.Sizeof(uintptr(0))))

	callArgs := make([]uintptr, 0, len(args)+1)
	callArgs = append(callArgs, comPtr)
	callArgs = append(callArgs, args...)

	r1, r2, err := syscall.SyscallN(methodPtr, callArgs...)
	return r1, r2, err
}

// Present effectue la présentation matérielle synchronisée au VBlank via DXGI.
func (sc *SwapChain) Present(syncInterval, flags uint32) error {
	slot := C2_dxgi_get_slot(2, 0) // Slot 8 = Present
	r1, _, _ := CallVTable(sc.ptr, slot, uintptr(syncInterval), uintptr(flags))
	if int32(r1) < 0 {
		return fmt.Errorf("c2dxgi: SwapChain.Present failed (HRESULT 0x%08X)", r1)
	}
	return nil
}

// Release libère une référence COM (Slot 2).
func ReleaseCOM(comPtr uintptr) uint32 {
	if comPtr == 0 {
		return 0
	}
	slot := C2_dxgi_get_slot(1, 2) // Slot 2 = Release
	r1, _, _ := CallVTable(comPtr, slot)
	return uint32(r1)
}
