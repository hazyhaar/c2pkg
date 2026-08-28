// SPDX-License-Identifier: Apache-2.0 OR MIT

//go:build !linux

package agetorture

type PerfSampler struct{}

func NewPerfSampler() *PerfSampler { return &PerfSampler{} }
func (ps *PerfSampler) ResetAndEnable() {}
func (ps *PerfSampler) ReadStats() (cycles, insn, branches, brMiss, l1dMiss uint64, ipc float64) {
	return 0, 0, 0, 0, 0, 0
}
func (ps *PerfSampler) Close() {}
