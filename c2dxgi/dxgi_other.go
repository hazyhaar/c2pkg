// SPDX-License-Identifier: Apache-2.0 OR MIT

//go:build !windows

package c2dxgi

import "fmt"

type Device struct {
	_ uintptr
}

type DeviceContext struct {
	_ uintptr
}

type SwapChain struct {
	_ uintptr
}

func CallVTable(comPtr uintptr, slot uint32, args ...uintptr) (uintptr, uintptr, error) {
	return 0, 0, fmt.Errorf("c2dxgi: COM/DXGI non supporté hors environnement Windows")
}

func (sc *SwapChain) Present(syncInterval, flags uint32) error {
	return fmt.Errorf("c2dxgi: SwapChain.Present non supporté hors environnement Windows")
}

func ReleaseCOM(comPtr uintptr) uint32 {
	return 0
}
