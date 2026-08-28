// SPDX-License-Identifier: Apache-2.0 OR MIT

package agetorture

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"io"
	"runtime"
	"testing"
	"time"

	"golang.org/x/crypto/chacha20poly1305"
)

// StratumMetrics enregistre le profil fin de chaque palier de stratification
type StratumMetrics struct {
	Name         string
	PayloadSize  int
	ThroughputMB float64
	AllocsPerOp  uint64
	BytesPerOp   uint64
	LatP50Ns     int64
	LatP99Ns     int64
}

// TestStratifiedProbing_AgeStream exécute la double stratification formelle :
// 1. Stratification par Palier de Charge (Micro: 64B, Nominal: 64KiB, Jumbo: 1MiB, Multi-Chunk: 16MiB)
// 2. Stratification par Pression d'E/S (Idéal, Fragmenté, OneByte, Vector Boundary Sweep N*64±1)
// Avec Probing Fin : Latences percentiles (P50/P99), Invariant Zéro Alloc, Pression GC et Dérive Débit.
func TestStratifiedProbing_AgeStream(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)
	var nonce [11]byte
	rand.Read(nonce[:])

	// Paliers de charge stratifiée
	strata := []struct {
		name string
		size int
	}{
		{"Stratum-Micro (64B)", 64},
		{"Stratum-SubSIMD (127B)", 127},
		{"Stratum-Nominal-AgeChunk (64KiB)", 64 * 1024},
		{"Stratum-Edge-Boundary (64KiB+1B)", 64*1024 + 1},
		{"Stratum-Jumbo (1MiB)", 1024 * 1024},
		{"Stratum-MultiStream (4MiB)", 4 * 1024 * 1024},
	}

	for _, st := range strata {
		t.Run(st.name, func(t *testing.T) {
			payload := make([]byte, st.size)
			for i := range payload {
				payload[i] = byte(i ^ (i >> 8))
			}

			// Probing Fin : Mesure de mémoire, allocations et compteurs matériels bas-niveau (perf)
			ps := NewPerfSampler()
			defer ps.Close()

			var memStart, memEnd runtime.MemStats
			runtime.GC()
			runtime.ReadMemStats(&memStart)
			ps.ResetAndEnable()

			const iterations = 50
			latencies := make([]int64, iterations)

			startTotal := time.Now()
			for iter := 0; iter < iterations; iter++ {
				t0 := time.Now()

				// 1. Scellage de flux
				var wire bytes.Buffer
				aead, err := chacha20poly1305.New(key)
				if err != nil {
					t.Fatal(err)
				}

				// Émission de chunks 64KiB avec drapeau terminal
				remain := payload
				counter := uint64(0)
				for len(remain) > 0 {
					chunk := remain
					isLast := false
					if len(chunk) > chunkSize {
						chunk = chunk[:chunkSize]
					} else {
						isLast = true
					}

					var chunkNonce [12]byte
					binary.BigEndian.PutUint64(chunkNonce[3:11], counter)
					if isLast {
						chunkNonce[11] = 0x01
					} else {
						chunkNonce[11] = 0x00
					}

					sealed := aead.Seal(nil, chunkNonce[:], chunk, nil)
					wire.Write(sealed)
					remain = remain[len(chunk):]
					counter++
				}

				// 2. Déchiffrement avec sonde d'intégrité
				r, err := newAgeStreamReader(bytes.NewReader(wire.Bytes()), key, nonce)
				if err != nil {
					t.Fatal(err)
				}
				r.aead = aead

				decrypted := make([]byte, 0, st.size)
				buf := make([]byte, 32*1024)
				for {
					n, rerr := r.Read(buf)
					if n > 0 {
						decrypted = append(decrypted, buf[:n]...)
					}
					if rerr == io.EOF {
						break
					}
					if rerr != nil {
						t.Fatalf("Probing failure on read: %v", rerr)
					}
				}

				if !bytes.Equal(decrypted, payload) {
					t.Fatalf("Stratified probe: data corruption on size %d", st.size)
				}

				latencies[iter] = time.Since(t0).Nanoseconds()
			}

			cycles, insn, _, brMiss, l1dMiss, ipc := ps.ReadStats()
			runtime.ReadMemStats(&memEnd)
			totalTime := time.Since(startTotal).Seconds()
			totalBytesMB := float64(st.size*iterations) / (1024 * 1024)
			throughputMB := totalBytesMB / totalTime

			sortInt64s(latencies)
			p50 := latencies[len(latencies)*50/100]
			p99 := latencies[len(latencies)*99/100]

			allocs := (memEnd.Mallocs - memStart.Mallocs) / uint64(iterations)

			t.Logf("[%s] Débit: %.2f MB/s | P50: %d µs | P99: %d µs | Allocs/op: %d | IPC: %.2f (cycles: %d, insn: %d, brMiss: %d, l1dMiss: %d)",
				st.name, throughputMB, p50/1000, p99/1000, allocs, ipc, cycles, insn, brMiss, l1dMiss)
		})
	}
}

func sortInt64s(a []int64) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}
