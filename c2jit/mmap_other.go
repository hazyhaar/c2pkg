// SPDX-License-Identifier: Apache-2.0 OR MIT
//go:build !(linux || darwin || freebsd || openbsd || netbsd)

package c2jit

import (
	"errors"
	"unsafe"
)

var ErrUnsupportedPlatform = errors.New("c2jit: exécution JIT non supportée sur cette plateforme")

type ExecutableBuffer struct {
	mem []byte
}

func AllocateBuffer(size int) (*ExecutableBuffer, error) {
	return nil, ErrUnsupportedPlatform
}

func (b *ExecutableBuffer) Write(code []byte) error {
	return ErrUnsupportedPlatform
}

func (b *ExecutableBuffer) SealExecutable() error {
	return ErrUnsupportedPlatform
}

func (b *ExecutableBuffer) UnsealWritable() error {
	return ErrUnsupportedPlatform
}

func (b *ExecutableBuffer) Ptr() uintptr {
	return 0
}

func (b *ExecutableBuffer) Free() error {
	return nil
}

type jitFuncVal struct {
	fn uintptr
}

func AsFunc2(codePtr uintptr) func(a, b uintptr) uintptr {
	return nil
}

func AsFunc3(codePtr uintptr) func(a, b, c uintptr) uintptr {
	return nil
}
