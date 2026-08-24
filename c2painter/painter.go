package c2painter

// Surface représente un tampon mémoire 2D RGBA 32-bit pour le rendu graphique.
type Surface struct {
	Pixels []uint32
	Width  int
	Height int
	Stride int
}

// NewSurface alloue une nouvelle surface de dimensions spécifiées.
func NewSurface(width, height int) *Surface {
	if width <= 0 || height <= 0 {
		return &Surface{
			Pixels: nil,
			Width:  0,
			Height: 0,
			Stride: 0,
		}
	}
	// Alignement strict de chaque scanline sur 64 octets (16 pixels uint32) pour AVX2 / AVX-512 / NEON
	stride := ((width + 15) / 16) * 16
	pixels := make([]uint32, stride*height)
	return &Surface{
		Pixels: pixels,
		Width:  width,
		Height: height,
		Stride: stride,
	}
}

// RGBA construit une couleur 32-bit RGBA (R: octet 0, G: octet 1, B: octet 2, A: octet 3).
func RGBA(r, g, b, a uint8) uint32 {
	return C2_rgba(r, g, b, a)
}

// PackRGBA construit une couleur 32-bit RGBA (R: octet 0, G: octet 1, B: octet 2, A: octet 3).
func PackRGBA(r, g, b, a uint8) uint32 {
	return C2_rgba(r, g, b, a)
}

// UnpackRGBA extrait les 4 composantes R, G, B, A d'une couleur uint32.
func UnpackRGBA(c uint32) (r, g, b, a uint8) {
	r = uint8(c & 0xFF)
	g = uint8((c >> 8) & 0xFF)
	b = uint8((c >> 16) & 0xFF)
	a = uint8((c >> 24) & 0xFF)
	return
}

// RGBAPremul construit une couleur 32-bit RGBA prémultipliée.
func RGBAPremul(r, g, b, a uint8) uint32 {
	return C2_rgba_premul(r, g, b, a)
}

// Painter encapsule un contexte de rendu graphique 2D C2_painter_t transpilé depuis C99 via sgoiter.
type Painter struct {
	ctx                 C2_painter_t
	PhotometricBlending bool
}

// NewPainter crée un peintre opérant sur la surface fournie.
func NewPainter(surface *Surface) *Painter {
	p := &Painter{}
	if surface != nil {
		C2_painter_init(&p.ctx, surface.Pixels, surface.Width, surface.Height, surface.Stride)
	}
	return p
}

// NewPainterFromBuffer initialise un peintre directement depuis un slice de pixels RGBA.
func NewPainterFromBuffer(pixels []uint32, width, height, stride int) *Painter {
	p := &Painter{}
	C2_painter_init(&p.ctx, pixels, width, height, stride)
	return p
}

// Pixels retourne le slice de pixels sous-jacent.
func (p *Painter) Pixels() []uint32 {
	return p.ctx.Pixels
}

// Width retourne la largeur du tampon.
func (p *Painter) Width() int {
	return p.ctx.Width
}

// Height retourne la hauteur du tampon.
func (p *Painter) Height() int {
	return p.ctx.Height
}

// Stride retourne le stride du tampon.
func (p *Painter) Stride() int {
	return p.ctx.Stride
}

// SetClip définit le rectangle de découpage actif.
func (p *Painter) SetClip(x, y, w, h int) {
	C2_painter_set_clip(&p.ctx, x, y, w, h)
}

// ResetClip réinitialise le rectangle de découpage aux limites complètes de la surface.
func (p *Painter) ResetClip() {
	C2_painter_reset_clip(&p.ctx)
}

// Clear remplit l'intégralité du tampon avec la couleur indiquée en exploitant le doublement exponentiel memmove.
func (p *Painter) Clear(color uint32) {
	pixels := p.ctx.Pixels
	if len(pixels) == 0 || p.ctx.Width <= 0 || p.ctx.Height <= 0 {
		return
	}
	if p.ctx.Stride == p.ctx.Width {
		total := p.ctx.Width * p.ctx.Height
		if total > len(pixels) {
			total = len(pixels)
		}
		pixels[0] = color
		for n := 1; n < total; n *= 2 {
			copy(pixels[n:total], pixels[:n])
		}
		return
	}
	C2_clear(&p.ctx, color)
}

// DrawRect remplit un rectangle avec couleur RGBA prémultipliée (ou photométrique si PhotometricBlending est activé).
func (p *Painter) DrawRect(x, y, w, h int, color uint32) {
	if p.PhotometricBlending && (color>>24)&0xFF != 255 {
		p.DrawRectPhotometric(x, y, w, h, color)
		return
	}
	C2_fill_rect(&p.ctx, x, y, w, h, color)
}

// DrawRectSIMD remplit un rectangle avec accélération vectorielle AVX2 ou photométrique.
func (p *Painter) DrawRectSIMD(x, y, w, h int, color uint32) {
	if p.PhotometricBlending && (color>>24)&0xFF != 255 {
		p.DrawRectPhotometric(x, y, w, h, color)
		return
	}
	FillRectSIMD(&p.ctx, x, y, w, h, color)
}

// DrawCircle remplit un cercle avec antialiasing (ou composition photométrique linéaire).
func (p *Painter) DrawCircle(cx, cy, radius int, color uint32) {
	if p.PhotometricBlending && (color>>24)&0xFF != 255 {
		p.DrawCirclePhotometric(cx, cy, radius, color)
		return
	}
	C2_fill_circle(&p.ctx, cx, cy, radius, color)
}

// DrawCirclePhotometric remplit un cercle avec composition photométrique linéaire.
func (p *Painter) DrawCirclePhotometric(cx, cy, radius int, color uint32) {
	if radius <= 0 || p.ctx.Pixels == nil {
		return
	}
	r := radius
	x0 := max(cx-r-1, p.ctx.Clip.X)
	x1 := min(cx+r+1, p.ctx.Clip.X+p.ctx.Clip.W)
	y0 := max(cy-r-1, p.ctx.Clip.Y)
	y1 := min(cy+r+1, p.ctx.Clip.Y+p.ctx.Clip.H)
	if x0 >= x1 || y0 >= y1 {
		return
	}

	rInner := int64(max(0, r-1)) * int64(max(0, r-1))
	rOuter := int64(r+1) * int64(r+1)

	for y := y0; y < y1; y++ {
		dy := int64(y - cy)
		dy2 := dy * dy
		rowOff := y * p.ctx.Stride
		for x := x0; x < x1; x++ {
			dx := int64(x - cx)
			dist2 := dx*dx + dy2
			if dist2 <= rInner {
				p.ctx.Pixels[rowOff+x] = C2_blend_photometric(p.ctx.Pixels[rowOff+x], color)
			} else if dist2 <= rOuter {
				dist := C2_isqrt64(dist2 << 16)
				cov := int((int64(r)<<8 + 128) - dist)
				if cov >= 255 {
					p.ctx.Pixels[rowOff+x] = C2_blend_photometric(p.ctx.Pixels[rowOff+x], color)
				} else if cov > 0 {
					p.ctx.Pixels[rowOff+x] = C2_blend_pixel_cov_photometric(p.ctx.Pixels[rowOff+x], color, uint32(cov))
				}
			}
		}
	}
}

// StrokeCircle trace le contour d'un cercle avec antialiasing.
func (p *Painter) StrokeCircle(cx, cy, radius, strokeWidth int, color uint32) {
	C2_stroke_circle(&p.ctx, cx, cy, radius, strokeWidth, color)
}

// DrawLine trace un segment de droite antialiasé (ou composition photométrique linéaire).
func (p *Painter) DrawLine(x0, y0, x1, y1, strokeWidth int, color uint32) {
	if p.PhotometricBlending && (color>>24)&0xFF != 255 {
		p.DrawLinePhotometric(x0, y0, x1, y1, strokeWidth, color)
		return
	}
	C2_draw_line(&p.ctx, x0, y0, x1, y1, strokeWidth, color)
}

// DrawLinePhotometric trace un segment de droite avec composition photométrique linéaire.
func (p *Painter) DrawLinePhotometric(x0, y0, x1, y1, strokeWidth int, color uint32) {
	if strokeWidth <= 0 || p.ctx.Pixels == nil {
		return
	}
	dx := int64(x1 - x0)
	dy := int64(y1 - y0)
	lenSq := dx*dx + dy*dy
	if lenSq == 0 {
		p.DrawCirclePhotometric(x0, y0, (strokeWidth+1)/2, color)
		return
	}

	rad := (strokeWidth + 3) / 2
	minX := max(min(x0, x1)-rad, p.ctx.Clip.X)
	maxX := min(max(x0, x1)+rad, p.ctx.Clip.X+p.ctx.Clip.W)
	minY := max(min(y0, y1)-rad, p.ctx.Clip.Y)
	maxY := min(max(y0, y1)+rad, p.ctx.Clip.Y+p.ctx.Clip.H)
	if minX >= maxX || minY >= maxY {
		return
	}

	halfWidthFixed := int64(strokeWidth) * 128
	length := C2_isqrt64(lenSq)
	if length == 0 {
		length = 1
	}

	for y := minY; y < maxY; y++ {
		rowOff := y * p.ctx.Stride
		for x := minX; x < maxX; x++ {
			px := int64(x - x0)
			py := int64(y - y0)
			proj := (px*dx + py*dy)
			if proj < 0 {
				d2 := px*px + py*py
				d := C2_isqrt64(d2 << 16)
				cov := int(halfWidthFixed + 128 - d)
				if cov >= 255 {
					p.ctx.Pixels[rowOff+x] = C2_blend_photometric(p.ctx.Pixels[rowOff+x], color)
				} else if cov > 0 {
					p.ctx.Pixels[rowOff+x] = C2_blend_pixel_cov_photometric(p.ctx.Pixels[rowOff+x], color, uint32(cov))
				}
			} else if proj > lenSq {
				px1 := int64(x - x1)
				py1 := int64(y - y1)
				d2 := px1*px1 + py1*py1
				d := C2_isqrt64(d2 << 16)
				cov := int(halfWidthFixed + 128 - d)
				if cov >= 255 {
					p.ctx.Pixels[rowOff+x] = C2_blend_photometric(p.ctx.Pixels[rowOff+x], color)
				} else if cov > 0 {
					p.ctx.Pixels[rowOff+x] = C2_blend_pixel_cov_photometric(p.ctx.Pixels[rowOff+x], color, uint32(cov))
				}
			} else {
				cross := px*dy - py*dx
				if cross < 0 {
					cross = -cross
				}
				dist := (cross << 8) / length
				cov := int(halfWidthFixed + 128 - dist)
				if cov >= 255 {
					p.ctx.Pixels[rowOff+x] = C2_blend_photometric(p.ctx.Pixels[rowOff+x], color)
				} else if cov > 0 {
					p.ctx.Pixels[rowOff+x] = C2_blend_pixel_cov_photometric(p.ctx.Pixels[rowOff+x], color, uint32(cov))
				}
			}
		}
	}
}

// DrawRectPhotometric remplit un rectangle avec composition photométrique linéaire sgoiter.
func (p *Painter) DrawRectPhotometric(x, y, w, h int, color uint32) {
	if p.ctx.Pixels == nil || w <= 0 || h <= 0 {
		return
	}

	cx, cy, cw, ch := p.ctx.Clip.X, p.ctx.Clip.Y, p.ctx.Clip.W, p.ctx.Clip.H
	x1 := x
	y1 := y
	x2 := x + w
	y2 := y + h

	if x1 < cx {
		x1 = cx
	}
	if y1 < cy {
		y1 = cy
	}
	if x2 > cx+cw {
		x2 = cx + cw
	}
	if y2 > cy+ch {
		y2 = cy + ch
	}

	if x1 >= x2 || y1 >= y2 {
		return
	}

	stride := p.ctx.Stride
	pixels := p.ctx.Pixels

	for row := y1; row < y2; row++ {
		off := row*stride + x1
		for col := 0; col < x2-x1; col++ {
			pixels[off+col] = C2_blend_photometric(pixels[off+col], color)
		}
	}
}

// Ctx retourne le pointeur vers le contexte interne C2_painter_t.
func (p *Painter) Ctx() C2_painter_t {
	return p.ctx
}

// StrokeRect trace le contour d'un rectangle.
func (p *Painter) StrokeRect(x, y, w, h, strokeWidth int, color uint32) {
	C2_stroke_rect(&p.ctx, x, y, w, h, strokeWidth, color)
}

// DrawRoundedRect remplit un rectangle à coins arrondis avec antialiasing.
func (p *Painter) DrawRoundedRect(x, y, w, h, radius int, color uint32) {
	C2_fill_rounded_rect(&p.ctx, x, y, w, h, radius, color)
}

// StrokeRoundedRect trace le contour d'un rectangle à coins arrondis avec antialiasing.
func (p *Painter) StrokeRoundedRect(x, y, w, h, radius, strokeWidth int, color uint32) {
	C2_stroke_rounded_rect(&p.ctx, x, y, w, h, radius, strokeWidth, color)
}

// DrawEllipse remplit une ellipse avec antialiasing.
func (p *Painter) DrawEllipse(cx, cy, rx, ry int, color uint32) {
	C2_fill_ellipse(&p.ctx, cx, cy, rx, ry, color)
}

// StrokeEllipse trace le contour d'une ellipse avec antialiasing.
func (p *Painter) StrokeEllipse(cx, cy, rx, ry, strokeWidth int, color uint32) {
	C2_stroke_ellipse(&p.ctx, cx, cy, rx, ry, strokeWidth, color)
}

// DrawTextGlyph blitte un masque alpha 8-bit (glyphe textuel) avec la couleur spécifiée.
func (p *Painter) DrawTextGlyph(dstX, dstY, w, h int, mask []byte, maskStride int, color uint32) {
	if p.PhotometricBlending {
		C2_blit_mask_photometric(&p.ctx, dstX, dstY, w, h, mask, maskStride, color)
		return
	}
	C2_blit_mask(&p.ctx, dstX, dstY, w, h, mask, maskStride, color)
}

// C2_blit_mask_photometric effectue le blitte d'un masque alpha 8-bit avec composition photométrique linéaire.
func C2_blit_mask_photometric(p *C2_painter_t, dstX, dstY, w, h int, mask []byte, maskStride int, color uint32) {
	if len(mask) == 0 || w <= 0 || h <= 0 {
		return
	}
	clipX := p.Clip.X
	clipY := p.Clip.Y
	clipX2 := clipX + p.Clip.W
	clipY2 := clipY + p.Clip.H

	x1 := dstX
	y1 := dstY
	x2 := dstX + w
	y2 := dstY + h

	if x1 < clipX {
		x1 = clipX
	}
	if y1 < clipY {
		y1 = clipY
	}
	if x2 > clipX2 {
		x2 = clipX2
	}
	if y2 > clipY2 {
		y2 = clipY2
	}

	if x1 >= x2 || y1 >= y2 {
		return
	}

	pixels := p.Pixels
	stride := p.Stride

	for y := y1; y < y2; y++ {
		srcY := y - dstY
		dstRow := y * stride
		maskRow := srcY * maskStride
		for x := x1; x < x2; x++ {
			srcX := x - dstX
			alpha := mask[maskRow+srcX]
			if alpha > 0 {
				pixels[dstRow+x] = C2_blend_pixel_cov_photometric(pixels[dstRow+x], color, uint32(alpha))
			}
		}
	}
}

// DrawTextGlyphRLE décode un masque alpha compressé en RLE à la volée et effectue la composition sans allocation.
func (p *Painter) DrawTextGlyphRLE(dstX, dstY, w, h int, rle []byte, color uint32) {
	C2_blit_rle_mask(&p.ctx, dstX, dstY, w, h, rle, color)
}

// C2_blit_rle_mask décode à la volée un flux RLE (PackBits) et compose directement sur les pixels sans allocation.
func C2_blit_rle_mask(p *C2_painter_t, dst_x int, dst_y int, w int, h int, rle []byte, color uint32) {
	if len(rle) == 0 || w <= 0 || h <= 0 {
		return
	}
	clipX := p.Clip.X
	clipY := p.Clip.Y
	clipX2 := clipX + p.Clip.W
	clipY2 := clipY + p.Clip.H

	if dst_x >= clipX2 || dst_y >= clipY2 || dst_x+w <= clipX || dst_y+h <= clipY {
		return
	}

	alphaColor := uint32(color>>24) & 0xff
	if alphaColor == 0 {
		return
	}

	pixels := p.Pixels
	stride := p.Stride

	curX := 0
	curY := 0
	rleLen := len(rle)
	i := 0

	for i < rleLen && curY < h {
		header := rle[i]
		i++

		if (header & 0x80) != 0 {
			// Séquence répétée
			if i >= rleLen {
				break
			}
			runLen := int(header&0x7F) + 1
			val := uint32(rle[i])
			i++

			if val == 0 {
				// Saut direct des pixels transparents
				curX += runLen
				for curX >= w {
					curX -= w
					curY++
				}
				continue
			}

			for k := 0; k < runLen && curY < h; k++ {
				py := dst_y + curY
				px := dst_x + curX
				if px >= clipX && px < clipX2 && py >= clipY && py < clipY2 {
					off := py*stride + px
					pixels[off] = C2_blend_pixel_cov(pixels[off], color, val)
				}
				curX++
				if curX >= w {
					curX = 0
					curY++
				}
			}
		} else {
			// Séquence littérale
			litLen := int(header) + 1
			for k := 0; k < litLen && i < rleLen && curY < h; k++ {
				val := uint32(rle[i])
				i++
				if val > 0 {
					py := dst_y + curY
					px := dst_x + curX
					if px >= clipX && px < clipX2 && py >= clipY && py < clipY2 {
						off := py*stride + px
						pixels[off] = C2_blend_pixel_cov(pixels[off], color, val)
					}
				}
				curX++
				if curX >= w {
					curX = 0
					curY++
				}
			}
		}
	}
}

// Blit copie une image RGBA avec composition alpha Porter-Duff Source-Over.
func (p *Painter) Blit(dstX, dstY, w, h int, src []uint32, srcStride int) {
	C2_blit_rgba(&p.ctx, dstX, dstY, w, h, src, srcStride)
}

// BlitSIMD copie une image RGBA avec accélération vectorielle AVX2.
func (p *Painter) BlitSIMD(dstX, dstY, w, h int, src []uint32, srcStride int) {
	BlitRGBASIMD(&p.ctx, dstX, dstY, w, h, src, srcStride)
}

// FillLinearGradient remplit un rectangle avec un dégradé linéaire bicolore.
func (p *Painter) FillLinearGradient(x, y, w, h int, color0, color1 uint32, vertical bool) {
	v := 0
	if vertical {
		v = 1
	}
	C2_fill_linear_gradient(&p.ctx, x, y, w, h, color0, color1, v)
}

// DrawLinearGradient est un alias de FillLinearGradient pour la compatibilité d'API.
func (p *Painter) DrawLinearGradient(x, y, w, h int, color0, color1 uint32, vertical bool) {
	p.FillLinearGradient(x, y, w, h, color0, color1, vertical)
}

// DrawCursor trace un curseur contrasté idempotent selon son style (0: Block, 1: Bar, 2: Underline, 3: Hollow).
func (p *Painter) DrawCursor(x, y, w, h, style int, color uint32) {
	C2_draw_cursor(&p.ctx, x, y, w, h, style, color)
}
