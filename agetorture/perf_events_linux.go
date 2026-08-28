// SPDX-License-Identifier: Apache-2.0 OR MIT

//go:build linux

package agetorture

import (
	"encoding/binary"
	"syscall"
	"unsafe"
)

// Constantes Linux perf_event_open
const (
	PERF_TYPE_HARDWARE   = 0
	PERF_TYPE_HW_CACHE   = 3

	PERF_COUNT_HW_CPU_CYCLES          = 0
	PERF_COUNT_HW_INSTRUCTIONS        = 1
	PERF_COUNT_HW_CACHE_MISSES        = 3
	PERF_COUNT_HW_BRANCH_INSTRUCTIONS = 4
	PERF_COUNT_HW_BRANCH_MISSES       = 5

	PERF_COUNT_HW_CACHE_L1D  = 0
	PERF_COUNT_HW_CACHE_OP_READ = 0
	PERF_COUNT_HW_CACHE_RESULT_MISS = 1

	PERF_FORMAT_TOTAL_TIME_ENABLED = 1 << 0
	PERF_FORMAT_TOTAL_TIME_RUNNING = 1 << 1
)

type perfEventAttr [136]byte

func setAttr(attr *perfEventAttr, typ uint32, config uint64) {
	binary.LittleEndian.PutUint32(attr[0:4], typ)
	binary.LittleEndian.PutUint32(attr[4:8], 136) // size = 136 bytes
	binary.LittleEndian.PutUint64(attr[8:16], config)
	// bits à l'offset 40 : 0x61 = disabled (1) | exclude_kernel (0x20) | exclude_hv (0x40)
	attr[40] = 0x61
}

// PerfSampler capture les compteurs CPU natifs via l'appel système perf_event_open (sans fork/exec)
type PerfSampler struct {
	fdCycles   int
	fdInsn     int
	fdBranches int
	fdBrMiss   int
	fdL1dMiss  int
	supported  bool
}

func perfEventOpen(attr *perfEventAttr, pid int, cpu int, groupFd int, flags uintptr) (int, error) {
	fd, _, errno := syscall.Syscall6(
		syscall.SYS_PERF_EVENT_OPEN,
		uintptr(unsafe.Pointer(attr)),
		uintptr(pid),
		^uintptr(0), // cpu: -1
		^uintptr(0), // group_fd: -1
		flags,
		0,
	)
	if errno != 0 {
		return -1, errno
	}
	return int(fd), nil
}

func openHwCounter(typ uint32, config uint64) (int, error) {
	var attr perfEventAttr
	setAttr(&attr, typ, config)

	return perfEventOpen(&attr, 0, -1, -1, 0)
}

// NewPerfSampler initialise l'écouteur de compteurs matériels bas-niveau sur le thread courant
func NewPerfSampler() *PerfSampler {
	ps := &PerfSampler{
		fdCycles:   -1,
		fdInsn:     -1,
		fdBranches: -1,
		fdBrMiss:   -1,
		fdL1dMiss:  -1,
	}

	var err error
	ps.fdCycles, err = openHwCounter(PERF_TYPE_HARDWARE, PERF_COUNT_HW_CPU_CYCLES)
	if err != nil {
		return ps
	}
	ps.fdInsn, _ = openHwCounter(PERF_TYPE_HARDWARE, PERF_COUNT_HW_INSTRUCTIONS)
	ps.fdBranches, _ = openHwCounter(PERF_TYPE_HARDWARE, PERF_COUNT_HW_BRANCH_INSTRUCTIONS)
	ps.fdBrMiss, _ = openHwCounter(PERF_TYPE_HARDWARE, PERF_COUNT_HW_BRANCH_MISSES)
	
	l1dMissConfig := uint64(PERF_COUNT_HW_CACHE_L1D) | (uint64(PERF_COUNT_HW_CACHE_OP_READ) << 8) | (uint64(PERF_COUNT_HW_CACHE_RESULT_MISS) << 16)
	ps.fdL1dMiss, _ = openHwCounter(PERF_TYPE_HW_CACHE, l1dMissConfig)

	ps.supported = true
	return ps
}

func readCounter(fd int) uint64 {
	if fd < 0 {
		return 0
	}
	var buf [8]byte
	n, err := syscall.Read(fd, buf[:])
	if err != nil || n < 8 {
		return 0
	}
	return binary.LittleEndian.Uint64(buf[:])
}

// ResetAndEnable remet à zéro et active les compteurs matériels
func (ps *PerfSampler) ResetAndEnable() {
	if !ps.supported {
		return
	}
	fds := []int{ps.fdCycles, ps.fdInsn, ps.fdBranches, ps.fdBrMiss, ps.fdL1dMiss}
	for _, fd := range fds {
		if fd >= 0 {
			_, _, _ = syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), 0x2402 /* PERF_EVENT_IOC_RESET */, 0)
			_, _, _ = syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), 0x2400 /* PERF_EVENT_IOC_ENABLE */, 0)
		}
	}
}

// ReadStats capture l'état instantané des compteurs matériels
func (ps *PerfSampler) ReadStats() (cycles, insn, branches, brMiss, l1dMiss uint64, ipc float64) {
	if !ps.supported {
		return 0, 0, 0, 0, 0, 0
	}
	cycles = readCounter(ps.fdCycles)
	insn = readCounter(ps.fdInsn)
	branches = readCounter(ps.fdBranches)
	brMiss = readCounter(ps.fdBrMiss)
	l1dMiss = readCounter(ps.fdL1dMiss)

	if cycles > 0 {
		ipc = float64(insn) / float64(cycles)
	}
	return
}

// Close libère les descripteurs de fichiers système
func (ps *PerfSampler) Close() {
	fds := []int{ps.fdCycles, ps.fdInsn, ps.fdBranches, ps.fdBrMiss, ps.fdL1dMiss}
	for _, fd := range fds {
		if fd >= 0 {
			_ = syscall.Close(fd)
		}
	}
}
