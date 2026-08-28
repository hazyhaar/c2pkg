// SPDX-License-Identifier: Apache-2.0 OR MIT
//go:build linux || darwin || freebsd || openbsd || netbsd

package c2jit

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

var (
	ErrBufferSealed   = errors.New("c2jit: écriture impossible sur un tampon scellé en exécution (RX)")
	ErrBufferUnsealed = errors.New("c2jit: exécution impossible sur un tampon non scellé (RW)")
)

// ExecutableBuffer gère une plage de mémoire allouée dynamiquement sous le cycle de vie W^X.
type ExecutableBuffer struct {
	mem    []byte
	size   int
	isExec bool
}

// AllocateBuffer alloue un tampon de mémoire anonyme aligné sur les frontières de pages physiques (W^X).
func AllocateBuffer(size int) (*ExecutableBuffer, error) {
	if size <= 0 {
		size = 4096
	}
	pageSize := syscall.Getpagesize()
	allocSize := ((size + pageSize - 1) / pageSize) * pageSize

	mem, err := syscall.Mmap(
		-1,
		0,
		allocSize,
		syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_ANON|syscall.MAP_PRIVATE,
	)
	if err != nil {
		return nil, fmt.Errorf("c2jit: mmap RW a échoué: %w", err)
	}

	return &ExecutableBuffer{
		mem:    mem,
		size:   allocSize,
		isExec: false,
	}, nil
}

// Write écrit des octets machines dans le tampon RW.
func (b *ExecutableBuffer) Write(code []byte) error {
	if b.isExec {
		return ErrBufferSealed
	}
	if len(code) > len(b.mem) {
		return ErrBufferTooSmall
	}
	copy(b.mem, code)
	return nil
}

// SealExecutable bascule irréversiblement les permissions de la page en PROT_READ | PROT_EXEC (RX).
func (b *ExecutableBuffer) SealExecutable() error {
	if b.isExec {
		return nil
	}
	err := syscall.Mprotect(b.mem, syscall.PROT_READ|syscall.PROT_EXEC)
	if err != nil {
		return fmt.Errorf("c2jit: mprotect RX a échoué: %w", err)
	}
	b.isExec = true
	return nil
}

// UnsealWritable repasse les permissions de la page en PROT_READ | PROT_WRITE (RW) pour recyclage.
func (b *ExecutableBuffer) UnsealWritable() error {
	if !b.isExec {
		return nil
	}
	err := syscall.Mprotect(b.mem, syscall.PROT_READ|syscall.PROT_WRITE)
	if err != nil {
		return fmt.Errorf("c2jit: mprotect RW a échoué: %w", err)
	}
	b.isExec = false
	return nil
}

// Ptr retourne l'adresse mémoire du début du code exécutable.
func (b *ExecutableBuffer) Ptr() uintptr {
	if len(b.mem) == 0 {
		return 0
	}
	return uintptr(unsafe.Pointer(&b.mem[0]))
}

// Free libère définitivement la mémoire virtuelle via munmap.
func (b *ExecutableBuffer) Free() error {
	if len(b.mem) == 0 {
		return nil
	}
	err := syscall.Munmap(b.mem)
	b.mem = nil
	b.size = 0
	b.isExec = false
	return err
}

type jitFuncVal struct {
	fn uintptr
}

// AsFunc2 convertit un pointeur de code machine 64-bit en fonction Go à 2 arguments entiers sans allocation.
func AsFunc2(codePtr uintptr) func(a, b uintptr) uintptr {
	fv := &jitFuncVal{fn: codePtr}
	return *(*func(a, b uintptr) uintptr)(unsafe.Pointer(&fv))
}

// AsFunc3 convertit un pointeur de code machine 64-bit en fonction Go à 3 arguments entiers sans allocation.
func AsFunc3(codePtr uintptr) func(a, b, c uintptr) uintptr {
	fv := &jitFuncVal{fn: codePtr}
	return *(*func(a, b, c uintptr) uintptr)(unsafe.Pointer(&fv))
}
