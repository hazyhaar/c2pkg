package blake3archtsim

import (
	"encoding/binary"
	"math/bits"
)

const (
	BlockSize = 64
	ChunkSize = 1024
	OutLen    = 32
	KeyLen    = 32

	ChunkStart        = 1 << 0
	ChunkEnd          = 1 << 1
	Parent            = 1 << 2
	Root              = 1 << 3
	KeyedHash         = 1 << 4
	DeriveKeyContext  = 1 << 5
	DeriveKeyMaterial = 1 << 6
)

var iv = [8]uint32{
	0x6A09E667, 0xBB67AE85, 0x3C6EF372, 0xA54FF53A,
	0x510E527F, 0x9B05688C, 0x1F83D9AB, 0x5BE0CD19,
}

var msgPermutation = [16]uint8{
	2, 6, 3, 10, 7, 0, 4, 13, 1, 11, 12, 5, 9, 14, 15, 8,
}

// CompressNodeFallback executes 7 rounds of BLAKE3 compression in constant time.
func CompressNodeFallback(cv *[8]uint32, block *[64]byte, blockLen uint8, counter uint64, flags uint8, out *[16]uint32) {
	var state = [16]uint32{
		cv[0], cv[1], cv[2], cv[3], cv[4], cv[5], cv[6], cv[7],
		iv[0], iv[1], iv[2], iv[3],
		uint32(counter), uint32(counter >> 32), uint32(blockLen), uint32(flags),
	}

	var msg [16]uint32
	for i := 0; i < 16; i++ {
		msg[i] = binary.LittleEndian.Uint32(block[i*4 : (i+1)*4])
	}

	for r := 0; r < 7; r++ {
		// Column round
		gScalar(&state, 0, 4, 8, 12, msg[0], msg[1])
		gScalar(&state, 1, 5, 9, 13, msg[2], msg[3])
		gScalar(&state, 2, 6, 10, 14, msg[4], msg[5])
		gScalar(&state, 3, 7, 11, 15, msg[6], msg[7])

		// Diagonal round
		gScalar(&state, 0, 5, 10, 15, msg[8], msg[9])
		gScalar(&state, 1, 6, 11, 12, msg[10], msg[11])
		gScalar(&state, 2, 7, 8, 13, msg[12], msg[13])
		gScalar(&state, 3, 4, 9, 14, msg[14], msg[15])

		var newMsg [16]uint32
		for i := 0; i < 16; i++ {
			newMsg[i] = msg[msgPermutation[i]]
		}
		msg = newMsg
	}

	for i := 0; i < 8; i++ {
		out[i] = state[i] ^ state[i+8]
		out[i+8] = state[i+8] ^ cv[i]
	}
}

func gScalar(state *[16]uint32, a, b, c, d int, mx, my uint32) {
	state[a] = state[a] + state[b] + mx
	state[d] = bits.RotateLeft32(state[d]^state[a], -16)
	state[c] = state[c] + state[d]
	state[b] = bits.RotateLeft32(state[b]^state[c], -12)
	state[a] = state[a] + state[b] + my
	state[d] = bits.RotateLeft32(state[d]^state[a], -8)
	state[c] = state[c] + state[d]
	state[b] = bits.RotateLeft32(state[b]^state[c], -7)
}

// Digest represents an active BLAKE3 hashing instance.
type Digest struct {
	keyWords    [8]uint32
	hasherFlags uint8
	stack       [54][8]uint32
	stackLen    uint8
	chunk       chunkState
}

type chunkState struct {
	cv               [8]uint32
	chunkCounter     uint64
	buf              [BlockSize]byte
	bufLen           uint8
	blocksCompressed uint8
	flags            uint8
}

// New returns an initialized BLAKE3 hasher.
func New() *Digest {
	d := new(Digest)
	d.Reset()
	return d
}

func wordsFromKey(key [32]byte) (w [8]uint32) {
	for i := 0; i < 8; i++ {
		w[i] = binary.LittleEndian.Uint32(key[i*4 : (i+1)*4])
	}
	return w
}

func (d *Digest) init(keyWords [8]uint32, flags uint8) {
	d.keyWords = keyWords
	d.hasherFlags = flags
	d.stackLen = 0
	d.chunk.cv = keyWords
	d.chunk.chunkCounter = 0
	d.chunk.bufLen = 0
	d.chunk.blocksCompressed = 0
	d.chunk.flags = flags
	d.chunk.buf = [BlockSize]byte{}
}

func (d *Digest) Reset() {
	d.init(iv, 0)
}

func DeriveKey(context string, keyMaterial []byte) [32]byte {
	var ctx Digest
	ctx.init(iv, DeriveKeyContext)
	_, _ = ctx.Write([]byte(context))
	ck := ctx.SumFinal()
	var mat Digest
	mat.init(wordsFromKey(ck), DeriveKeyMaterial)
	_, _ = mat.Write(keyMaterial)
	return mat.SumFinal()
}

// Write absorbs data.
func (d *Digest) Write(p []byte) (int, error) {
	total := len(p)
	for len(p) > 0 {
		if d.chunk.bufLen == BlockSize {
			if d.chunk.blocksCompressed == 15 {
				chunkOut := d.chunk.output()
				d.pushStack(chunkOut)
				d.chunk.chunkCounter++
				d.chunk.cv = d.keyWords
				d.chunk.bufLen = 0
				d.chunk.blocksCompressed = 0
				d.chunk.flags = d.hasherFlags
				d.chunk.buf = [BlockSize]byte{}
			} else {
				var out [16]uint32
				flags := d.chunk.flags
				if d.chunk.blocksCompressed == 0 {
					flags |= ChunkStart
				}
				CompressNode(&d.chunk.cv, &d.chunk.buf, BlockSize, d.chunk.chunkCounter, flags, &out)
				copy(d.chunk.cv[:], out[:8])
				d.chunk.blocksCompressed++
				d.chunk.bufLen = 0
				d.chunk.buf = [BlockSize]byte{}
			}
		}

		want := BlockSize - int(d.chunk.bufLen)
		if want > len(p) {
			want = len(p)
		}
		copy(d.chunk.buf[d.chunk.bufLen:], p[:want])
		d.chunk.bufLen += uint8(want)
		p = p[want:]
	}
	return total, nil
}

func (c *chunkState) output() [8]uint32 {
	flags := c.flags | ChunkEnd
	if c.blocksCompressed == 0 {
		flags |= ChunkStart
	}
	var out [16]uint32
	CompressNode(&c.cv, &c.buf, c.bufLen, c.chunkCounter, flags, &out)
	var cv [8]uint32
	copy(cv[:], out[:8])
	return cv
}

func (d *Digest) pushStack(cv [8]uint32) {
	totalChunks := d.chunk.chunkCounter + 1
	for (totalChunks & 1) == 0 {
		parentBlock := makeParentBlock(d.stack[d.stackLen-1], cv)
		d.stackLen--
		var parentOut [16]uint32
		CompressNode(&d.keyWords, &parentBlock, BlockSize, 0, Parent|d.hasherFlags, &parentOut)
		copy(cv[:], parentOut[:8])
		totalChunks >>= 1
	}
	d.stack[d.stackLen] = cv
	d.stackLen++
}

func makeParentBlock(l, r [8]uint32) [64]byte {
	var b [64]byte
	for i := 0; i < 8; i++ {
		binary.LittleEndian.PutUint32(b[i*4:(i+1)*4], l[i])
		binary.LittleEndian.PutUint32(b[32+i*4:32+(i+1)*4], r[i])
	}
	return b
}

// Sum256 computes BLAKE3 hash in one pass without heap allocation.
func Sum256(data []byte) [32]byte {
	var d Digest
	d.Reset()
	_, _ = d.Write(data)
	return d.SumFinal()
}

// SumFinal finishes hash calculation.
func (d *Digest) SumFinal() [32]byte {
	var out [32]byte
	if d.stackLen == 0 {
		// Single chunk root
		flags := d.chunk.flags | ChunkEnd | Root
		if d.chunk.blocksCompressed == 0 {
			flags |= ChunkStart
		}
		var out16 [16]uint32
		CompressNode(&d.chunk.cv, &d.chunk.buf, d.chunk.bufLen, d.chunk.chunkCounter, flags, &out16)
		for i := 0; i < 8; i++ {
			binary.LittleEndian.PutUint32(out[i*4:(i+1)*4], out16[i])
		}
		return out
	}

	currentCV := d.chunk.output()
	for d.stackLen > 0 {
		parentBlock := makeParentBlock(d.stack[d.stackLen-1], currentCV)
		d.stackLen--
		flags := uint8(Parent) | d.hasherFlags
		if d.stackLen == 0 {
			flags |= Root
		}
		var parentOut [16]uint32
		CompressNode(&d.keyWords, &parentBlock, BlockSize, 0, flags, &parentOut)
		copy(currentCV[:], parentOut[:8])
	}

	for i := 0; i < 8; i++ {
		binary.LittleEndian.PutUint32(out[i*4:(i+1)*4], currentCV[i])
	}
	return out
}
