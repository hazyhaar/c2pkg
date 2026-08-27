// SPDX-License-Identifier: MIT
package c2bcdec

func expand5(v uint8) uint8 {
	return (v << 3) | (v >> 2)
}

func expand6(v uint8) uint8 {
	return (v << 2) | (v >> 4)
}

func decodeRGB565(c uint16, rgb *[4]uint8) {
	rgb[0] = expand5(uint8((c >> 11) & 0x1F))
	rgb[1] = expand6(uint8((c >> 5) & 0x3F))
	rgb[2] = expand5(uint8(c & 0x1F))
	rgb[3] = 0xFF
}

// DecodeBC1Block décode un bloc compressé de 8 octets en une tuile 4x4 RGBA (4 octets/pixel).
func DecodeBC1Block(src []byte, dst []byte, destinationPitch int) {
	if len(src) < 8 || len(dst) < 4*destinationPitch {
		return
	}

	c0 := uint16(src[0]) | (uint16(src[1]) << 8)
	c1 := uint16(src[2]) | (uint16(src[3]) << 8)
	indices := uint32(src[4]) | (uint32(src[5]) << 8) | (uint32(src[6]) << 16) | (uint32(src[7]) << 24)

	var pal [4][4]uint8
	decodeRGB565(c0, &pal[0])
	decodeRGB565(c1, &pal[1])

	if c0 > c1 {
		pal[2][0] = uint8((2*int(pal[0][0]) + int(pal[1][0]) + 1) / 3)
		pal[2][1] = uint8((2*int(pal[0][1]) + int(pal[1][1]) + 1) / 3)
		pal[2][2] = uint8((2*int(pal[0][2]) + int(pal[1][2]) + 1) / 3)
		pal[2][3] = 0xFF

		pal[3][0] = uint8((int(pal[0][0]) + 2*int(pal[1][0]) + 1) / 3)
		pal[3][1] = uint8((int(pal[0][1]) + 2*int(pal[1][1]) + 1) / 3)
		pal[3][2] = uint8((int(pal[0][2]) + 2*int(pal[1][2]) + 1) / 3)
		pal[3][3] = 0xFF
	} else {
		pal[2][0] = uint8((int(pal[0][0]) + int(pal[1][0])) / 2)
		pal[2][1] = uint8((int(pal[0][1]) + int(pal[1][1])) / 2)
		pal[2][2] = uint8((int(pal[0][2]) + int(pal[1][2])) / 2)
		pal[2][3] = 0xFF

		pal[3][0] = 0
		pal[3][1] = 0
		pal[3][2] = 0
		pal[3][3] = 0
	}

	for y := 0; y < 4; y++ {
		rowOff := y * destinationPitch
		for x := 0; x < 4; x++ {
			idx := (indices >> ((y*4 + x) * 2)) & 0x03
			pOff := rowOff + x*4
			dst[pOff+0] = pal[idx][0]
			dst[pOff+1] = pal[idx][1]
			dst[pOff+2] = pal[idx][2]
			dst[pOff+3] = pal[idx][3]
		}
	}
}

// DecodeBC3Block décode un bloc compressé de 16 octets en une tuile 4x4 RGBA.
func DecodeBC3Block(src []byte, dst []byte, destinationPitch int) {
	if len(src) < 16 || len(dst) < 4*destinationPitch {
		return
	}

	// Décodage couleur BC1 (octets 8 à 15)
	DecodeBC1Block(src[8:], dst, destinationPitch)

	// Décodage alpha (octets 0 à 7)
	a0 := src[0]
	a1 := src[1]
	var aIndices uint64
	for i := 0; i < 6; i++ {
		aIndices |= uint64(src[2+i]) << (i * 8)
	}

	var aPal [8]uint8
	aPal[0] = a0
	aPal[1] = a1

	if a0 > a1 {
		for i := 1; i <= 6; i++ {
			aPal[1+i] = uint8(((7-i)*int(a0) + i*int(a1) + 3) / 7)
		}
	} else {
		for i := 1; i <= 4; i++ {
			aPal[1+i] = uint8(((5-i)*int(a0) + i*int(a1) + 2) / 5)
		}
		aPal[6] = 0x00
		aPal[7] = 0xFF
	}

	for y := 0; y < 4; y++ {
		rowOff := y * destinationPitch
		for x := 0; x < 4; x++ {
			aIdx := (aIndices >> ((y*4 + x) * 3)) & 0x07
			dst[rowOff+x*4+3] = aPal[aIdx]
		}
	}
}

// DecodeBC4Block décode un bloc compressé de 8 octets en une tuile 4x4 1-canal (R).
func DecodeBC4Block(src []byte, dst []byte, destinationPitch int) {
	if len(src) < 8 || len(dst) < 4*destinationPitch {
		return
	}

	r0 := src[0]
	r1 := src[1]
	var rIndices uint64
	for i := 0; i < 6; i++ {
		rIndices |= uint64(src[2+i]) << (i * 8)
	}

	var rPal [8]uint8
	rPal[0] = r0
	rPal[1] = r1

	if r0 > r1 {
		for i := 1; i <= 6; i++ {
			rPal[1+i] = uint8(((7-i)*int(r0) + i*int(r1) + 3) / 7)
		}
	} else {
		for i := 1; i <= 4; i++ {
			rPal[1+i] = uint8(((5-i)*int(r0) + i*int(r1) + 2) / 5)
		}
		rPal[6] = 0x00
		rPal[7] = 0xFF
	}

	for y := 0; y < 4; y++ {
		rowOff := y * destinationPitch
		for x := 0; x < 4; x++ {
			rIdx := (rIndices >> ((y*4 + x) * 3)) & 0x07
			dst[rowOff+x] = rPal[rIdx]
		}
	}
}

// DecodeBC5Block décode un bloc compressé de 16 octets en une tuile 4x4 2-canaux (RG).
func DecodeBC5Block(src []byte, dst []byte, destinationPitch int) {
	if len(src) < 16 || len(dst) < 4*destinationPitch {
		return
	}

	var tempR [16]byte
	var tempG [16]byte

	DecodeBC4Block(src[0:8], tempR[:], 4)
	DecodeBC4Block(src[8:16], tempG[:], 4)

	for y := 0; y < 4; y++ {
		rowOff := y * destinationPitch
		for x := 0; x < 4; x++ {
			dst[rowOff+x*2+0] = tempR[y*4+x]
			dst[rowOff+x*2+1] = tempG[y*4+x]
		}
	}
}
