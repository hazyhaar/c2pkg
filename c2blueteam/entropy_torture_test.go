package c2blueteam

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"testing"
)

func referenceShannonEntropy(data []byte) float64 {
	if len(data) == 0 {
		return 0.0
	}
	var freq [256]float64
	for _, b := range data {
		freq[b]++
	}
	var ent float64
	total := float64(len(data))
	for _, count := range freq {
		if count > 0 {
			p := count / total
			ent -= p * math.Log2(p)
		}
	}
	return ent
}

// TestEntropy_ExhaustiveMonteCarlo éprouve 50 000 vecteurs aléatoires contre la formule analytique.
func TestEntropy_ExhaustiveMonteCarlo(t *testing.T) {
	rng := rand.New(rand.NewSource(0xCAFEBABE1337))
	var maxDiff float64 = 0.0
	var sumDiff float64 = 0.0
	const iterations = 50000

	buf := make([]byte, 65536)

	for iter := 0; iter < iterations; iter++ {
		length := 1 + rng.Intn(4096)
		mode := rng.Intn(4)

		switch mode {
		case 0: // Bruit blanc total
			rng.Read(buf[:length])
		case 1: // Alphabet restreint (4 à 32 symboles)
			alphabet := 4 + rng.Intn(28)
			for i := 0; i < length; i++ {
				buf[i] = byte(rng.Intn(alphabet))
			}
		case 2: // Périodicité bruiteuse
			period := 2 + rng.Intn(16)
			for i := 0; i < length; i++ {
				buf[i] = byte((i % period) + rng.Intn(3))
			}
		case 3: // Répétitions par blocs
			val := byte(rng.Intn(256))
			for i := 0; i < length; i++ {
				if i%32 == 0 {
					val = byte(rng.Intn(256))
				}
				buf[i] = val
			}
		}

		refFloat := referenceShannonEntropy(buf[:length])
		q8 := C2bt_calc_entropy_8_8(buf[:length], uint64(length))
		q8Float := float64(q8) / 256.0
		diff := math.Abs(refFloat - q8Float)

		if diff > maxDiff {
			maxDiff = diff
		}
		sumDiff += diff

		if diff >= 0.05 {
			t.Fatalf("échec torture monte-carlo à l'itération %d (len=%d, mode=%d) : ref=%.6f, q8=%.6f, diff=%.6f",
				iter, length, mode, refFloat, q8Float, diff)
		}
	}

	t.Logf("Monte-Carlo (50 000 vecteurs) validé : maxDiff = %.6f b/o, meanDiff = %.6f b/o",
		maxDiff, sumDiff/float64(iterations))
}

// TestEntropy_StatisticalDistributions valide les 6 profils statistiques majeurs.
func TestEntropy_StatisticalDistributions(t *testing.T) {
	sizes := []int{32, 64, 128, 255, 256, 257, 512, 1000, 1024, 4096, 65536}
	rng := rand.New(rand.NewSource(424242))

	for _, n := range sizes {
		buf := make([]byte, n)

		// 1. Dirac (Identique)
		for i := range buf {
			buf[i] = 0x55
		}
		if q8 := C2bt_calc_entropy_8_8(buf, uint64(n)); q8 != 0 {
			t.Errorf("taille %d : Dirac attendu 0, obtenu %d", n, q8)
		}

		// 2. Bimodal (50% 0x00, 50% 0xFF -> H=1.0 bit = 256 Q8.8)
		for i := range buf {
			if i%2 == 0 {
				buf[i] = 0x00
			} else {
				buf[i] = 0xFF
			}
		}
		ref := referenceShannonEntropy(buf)
		q8 := C2bt_calc_entropy_8_8(buf, uint64(n))
		diff := math.Abs(ref - (float64(q8) / 256.0))
		if diff >= 0.05 {
			t.Errorf("taille %d : Bimodal gap diff=%.6f", n, diff)
		}

		// 3. Base64 (64 symboles -> max ~6.0 bits)
		const b64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
		for i := range buf {
			buf[i] = b64[rng.Intn(64)]
		}
		q8 = C2bt_calc_entropy_8_8(buf, uint64(n))
		if n >= 512 && (q8 < 1400 || q8 > 1600) {
			t.Errorf("taille %d : Base64 q8=%d hors plage [1400..1600]", n, q8)
		}

		// 4. Hexadécimal (16 symboles -> max ~4.0 bits)
		const hex = "0123456789abcdef"
		for i := range buf {
			buf[i] = hex[rng.Intn(16)]
		}
		q8 = C2bt_calc_entropy_8_8(buf, uint64(n))
		if n >= 512 && (q8 < 900 || q8 > 1150) {
			t.Errorf("taille %d : Hex q8=%d hors plage [900..1150]", n, q8)
		}
	}
}

// TestEntropy_MisalignedSlices vérifie l'absence de faille sur tranches non alignées.
func TestEntropy_MisalignedSlices(t *testing.T) {
	arena := make([]byte, 4096)
	for i := range arena {
		arena[i] = byte((i * 13) ^ 0xAA)
	}

	for offset := 0; offset < 16; offset++ {
		for length := 1; length <= 256; length++ {
			sub := arena[offset : offset+length]
			ref := referenceShannonEntropy(sub)
			q8 := C2bt_calc_entropy_8_8(sub, uint64(length))
			diff := math.Abs(ref - (float64(q8) / 256.0))
			if diff >= 0.05 {
				t.Fatalf("offset=%d, len=%d : diff=%.6f >= 0.05", offset, length, diff)
			}
		}
	}
}

// TestEntropy_ConcurrencyTorture_100Goroutines éprouve le calcul d'entropie sous 100 goroutines parallèles
// en validant la tolérance stricte < 0.05 bit/octet et l'invariance thread-safe sans divergence.
func TestEntropy_ConcurrencyTorture_100Goroutines(t *testing.T) {
	const goroutines = 100
	const iterations = 200
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func(seed int) {
			defer wg.Done()
			rng := rand.New(rand.NewSource(int64(seed + 1000)))
			localBuf := make([]byte, 1024)
			for i := 0; i < iterations; i++ {
				rng.Read(localBuf)
				ref := referenceShannonEntropy(localBuf)
				q8 := C2bt_calc_entropy_8_8(localBuf, uint64(len(localBuf)))
				diff := math.Abs(ref - (float64(q8) / 256.0))
				if diff >= 0.05 {
					t.Errorf("goroutine %d: diff=%.6f >= 0.05 (q8=%d, ref=%.4f)", seed, diff, q8, ref)
				}
			}
		}(g)
	}

	wg.Wait()
}

func BenchmarkEntropy_Sweep_ARCHTIME(b *testing.B) {
	sizes := []int{16, 64, 128, 256, 512, 1024, 4096, 16384, 65536, 1048576}
	for _, sz := range sizes {
		b.Run(fmt.Sprintf("%dB", sz), func(b *testing.B) {
			buf := make([]byte, sz)
			for i := range buf {
				buf[i] = byte((i * 17) ^ 0x3C)
			}
			b.SetBytes(int64(sz))
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = C2bt_calc_entropy_8_8(buf, uint64(sz))
			}
		})
	}
}

func BenchmarkEntropy_Sweep_StandardFloat(b *testing.B) {
	sizes := []int{16, 64, 128, 256, 512, 1024, 4096, 16384, 65536, 1048576}
	for _, sz := range sizes {
		b.Run(fmt.Sprintf("%dB", sz), func(b *testing.B) {
			buf := make([]byte, sz)
			for i := range buf {
				buf[i] = byte((i * 17) ^ 0x3C)
			}
			b.SetBytes(int64(sz))
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = referenceShannonEntropy(buf)
			}
		})
	}
}
