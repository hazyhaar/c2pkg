//go:build !(goexperiment.simd && amd64)

package blake3archtsim

// CompressNode delegates to constant-time fallback
func CompressNode(cv *[8]uint32, block *[64]byte, blockLen uint8, counter uint64, flags uint8, out *[16]uint32) {
	CompressNodeFallback(cv, block, blockLen, counter, flags, out)
}
