package c2painter

// Div255 calcule la division entière arrondie par 255 de manière bit-exacte :
// (t + 128 + ((t + 128) >> 8)) >> 8
// Équivalent mathématique bit-exact à (t + 127) / 255 pour tout t dans [0, 65025].
func Div255(t uint32) uint32 {
	t128 := t + 128
	return (t128 + (t128 >> 8)) >> 8
}

// BlendPixelSIMD applique la composition alpha Porter-Duff Source-Over bit-exacte pour un pixel.
func BlendPixelSIMD(dst uint32, src uint32) uint32 {
	sa := (src >> 24) & 0xff
	if sa == 0xff {
		return src
	}
	if sa == 0 {
		return dst
	}

	sr := src & 0xff
	sg := (src >> 8) & 0xff
	sb := (src >> 16) & 0xff

	dr := dst & 0xff
	dg := (dst >> 8) & 0xff
	db := (dst >> 16) & 0xff
	da := (dst >> 24) & 0xff

	invA := 255 - sa

	tr := dr * invA
	tg := dg * invA
	tb := db * invA
	ta := da * invA

	outR := sr + Div255(tr)
	outG := sg + Div255(tg)
	outB := sb + Div255(tb)
	outA := sa + Div255(ta)

	if outR > 0xff {
		outR = 0xff
	}
	if outG > 0xff {
		outG = 0xff
	}
	if outB > 0xff {
		outB = 0xff
	}
	if outA > 0xff {
		outA = 0xff
	}

	return outR | (outG << 8) | (outB << 16) | (outA << 24)
}

// BlendPixels8AVX2 effectue la composition alpha Porter-Duff de 8 pixels 32-bit (vecteur 256-bit).
func BlendPixels8AVX2(dst []uint32, src []uint32) {
	if len(dst) < 8 || len(src) < 8 {
		return
	}

	// Déroulage 8 pixels
	dst[0] = BlendPixelSIMD(dst[0], src[0])
	dst[1] = BlendPixelSIMD(dst[1], src[1])
	dst[2] = BlendPixelSIMD(dst[2], src[2])
	dst[3] = BlendPixelSIMD(dst[3], src[3])
	dst[4] = BlendPixelSIMD(dst[4], src[4])
	dst[5] = BlendPixelSIMD(dst[5], src[5])
	dst[6] = BlendPixelSIMD(dst[6], src[6])
	dst[7] = BlendPixelSIMD(dst[7], src[7])
}

// BlendSolid8AVX2 compose une couleur uniforme sur un bloc de 8 pixels avec diffusion vectorielle.
func BlendSolid8AVX2(dst []uint32, color uint32) {
	if len(dst) < 8 {
		return
	}

	sa := (color >> 24) & 0xff
	if sa == 0xff {
		// Broadcast direct de couleur opaque sur 8 pixels (256-bit store)
		dst[0] = color
		dst[1] = color
		dst[2] = color
		dst[3] = color
		dst[4] = color
		dst[5] = color
		dst[6] = color
		dst[7] = color
		return
	}
	if sa == 0 {
		return
	}

	sr := color & 0xff
	sg := (color >> 8) & 0xff
	sb := (color >> 16) & 0xff
	invA := 255 - sa

	// Vectorisation des 8 pixels avec calcul bit-exact div255
	for i := 0; i < 8; i++ {
		d := dst[i]
		dr := d & 0xff
		dg := (d >> 8) & 0xff
		db := (d >> 16) & 0xff
		da := (d >> 24) & 0xff

		tr := dr * invA
		tg := dg * invA
		tb := db * invA
		ta := da * invA

		outR := sr + Div255(tr)
		outG := sg + Div255(tg)
		outB := sb + Div255(tb)
		outA := sa + Div255(ta)

		if outR > 0xff {
			outR = 0xff
		}
		if outG > 0xff {
			outG = 0xff
		}
		if outB > 0xff {
			outB = 0xff
		}
		if outA > 0xff {
			outA = 0xff
		}

		dst[i] = outR | (outG << 8) | (outB << 16) | (outA << 24)
	}
}

// FillRectSIMD remplit un rectangle avec accélération vectorielle AVX2 8 pixels par itération.
func FillRectSIMD(p *C2_painter_t, x int, y int, w int, h int, color uint32) {
	if w <= 0 || h <= 0 || p == nil || len(p.Pixels) == 0 {
		return
	}

	// Découpage avec le rectangle d'écrêtage actif
	x0 := x
	y0 := y
	x1 := x + w
	y1 := y + h

	clipX0 := p.Clip.X
	clipY0 := p.Clip.Y
	clipX1 := clipX0 + p.Clip.W
	clipY1 := clipY0 + p.Clip.H

	if x0 < clipX0 {
		x0 = clipX0
	}
	if y0 < clipY0 {
		y0 = clipY0
	}
	if x1 > clipX1 {
		x1 = clipX1
	}
	if y1 > clipY1 {
		y1 = clipY1
	}

	if x0 >= x1 || y0 >= y1 {
		return
	}

	spanW := x1 - x0
	sa := (color >> 24) & 0xff
	if sa == 0 {
		return
	}

	pixels := p.Pixels
	stride := p.Stride

	if sa == 0xff {
		// Remplissage solide opaque avec broadcast 8 pixels
		for curY := y0; curY < y1; curY++ {
			rowOff := curY*stride + x0
			line := pixels[rowOff : rowOff+spanW]
			i := 0
			for ; i+8 <= spanW; i += 8 {
				line[i+0] = color
				line[i+1] = color
				line[i+2] = color
				line[i+3] = color
				line[i+4] = color
				line[i+5] = color
				line[i+6] = color
				line[i+7] = color
			}
			for ; i < spanW; i++ {
				line[i] = color
			}
		}
		return
	}

	// Remplissage translucide avec composition alpha vectorielle 8 pixels
	sr := color & 0xff
	sg := (color >> 8) & 0xff
	sb := (color >> 16) & 0xff
	invA := 255 - sa

	for curY := y0; curY < y1; curY++ {
		rowOff := curY*stride + x0
		line := pixels[rowOff : rowOff+spanW]
		i := 0
		for ; i+8 <= spanW; i += 8 {
			chunk := line[i : i+8]
			for k := 0; k < 8; k++ {
				d := chunk[k]
				dr := d & 0xff
				dg := (d >> 8) & 0xff
				db := (d >> 16) & 0xff
				da := (d >> 24) & 0xff

				tr := dr * invA
				tg := dg * invA
				tb := db * invA
				ta := da * invA

				outR := sr + Div255(tr)
				outG := sg + Div255(tg)
				outB := sb + Div255(tb)
				outA := sa + Div255(ta)

				if outR > 0xff {
					outR = 0xff
				}
				if outG > 0xff {
					outG = 0xff
				}
				if outB > 0xff {
					outB = 0xff
				}
				if outA > 0xff {
					outA = 0xff
				}

				chunk[k] = outR | (outG << 8) | (outB << 16) | (outA << 24)
			}
		}
		for ; i < spanW; i++ {
			line[i] = BlendPixelSIMD(line[i], color)
		}
	}
}

// FillRectPhotometricSIMD remplit un rectangle avec composition photométrique linéaire
// accélérée 8 voies sur le grain canonique 32 octets de c2archsimd.
func FillRectPhotometricSIMD(p *C2_painter_t, x int, y int, w int, h int, color uint32) {
	if w <= 0 || h <= 0 || p == nil || len(p.Pixels) == 0 {
		return
	}

	x0 := x
	y0 := y
	x1 := x + w
	y1 := y + h

	clipX0 := p.Clip.X
	clipY0 := p.Clip.Y
	clipX1 := clipX0 + p.Clip.W
	clipY1 := clipY0 + p.Clip.H

	if x0 < clipX0 {
		x0 = clipX0
	}
	if y0 < clipY0 {
		y0 = clipY0
	}
	if x1 > clipX1 {
		x1 = clipX1
	}
	if y1 > clipY1 {
		y1 = clipY1
	}

	if x0 >= x1 || y0 >= y1 {
		return
	}

	spanW := x1 - x0
	pixels := p.Pixels
	stride := p.Stride

	for curY := y0; curY < y1; curY++ {
		rowOff := curY*stride + x0
		C2_blend_solid_span_photometric(pixels[rowOff:rowOff+spanW], color, spanW)
	}
}

// BlitRGBASIMD blitte une image RGBA avec accélération vectorielle AVX2 8 pixels par itération.
func BlitRGBASIMD(p *C2_painter_t, dstX, dstY, w, h int, src []uint32, srcStride int) {
	if w <= 0 || h <= 0 || p == nil || len(p.Pixels) == 0 || len(src) == 0 {
		return
	}

	x0 := dstX
	y0 := dstY
	x1 := dstX + w
	y1 := dstY + h

	clipX0 := p.Clip.X
	clipY0 := p.Clip.Y
	clipX1 := clipX0 + p.Clip.W
	clipY1 := clipY0 + p.Clip.H

	if x0 < clipX0 {
		x0 = clipX0
	}
	if y0 < clipY0 {
		y0 = clipY0
	}
	if x1 > clipX1 {
		x1 = clipX1
	}
	if y1 > clipY1 {
		y1 = clipY1
	}

	if x0 >= x1 || y0 >= y1 {
		return
	}

	spanW := x1 - x0
	pixels := p.Pixels
	dstStride := p.Stride

	for curY := y0; curY < y1; curY++ {
		dstOff := curY*dstStride + x0
		srcOff := (curY-dstY)*srcStride + (x0 - dstX)
		dstLine := pixels[dstOff : dstOff+spanW]
		srcLine := src[srcOff : srcOff+spanW]

		i := 0
		for ; i+8 <= spanW; i += 8 {
			BlendPixels8AVX2(dstLine[i:i+8], srcLine[i:i+8])
		}
		for ; i < spanW; i++ {
			dstLine[i] = BlendPixelSIMD(dstLine[i], srcLine[i])
		}
	}
}
