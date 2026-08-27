package c2raster

// RasterizeBezierQuad calcule la couverture anti-crénelée (0-255) d'une courbe de Bézier quadratique
// définie par les points de contrôle (x0, y0), (x1, y1), (x2, y2) dans un masque d'octets 8-bit de dimensions w x h.
// Exécution analytique zéro-allocation.
func RasterizeBezierQuad(x0, y0, x1, y1, x2, y2 float32, coverage []byte, w, h int) {
	if w <= 0 || h <= 0 || coverage == nil || len(coverage) < w*h {
		return
	}
	Rasterize_bezier_quad(x0, y0, x1, y1, x2, y2, coverage, w, h)
}

// RasterizeRoundedRect trace un rectangle à coins arrondis analytiquement anti-crénelé
// dans le tampon de pixels 32-bit framebuffer de largeur/stride donnés.
// Exécution analytique zéro-allocation.
func RasterizeRoundedRect(framebuffer []uint32, stride, x, y, w, h, radius int, color uint32) {
	if stride <= 0 || w <= 0 || h <= 0 || framebuffer == nil || (y+h-1)*stride+(x+w) > len(framebuffer) {
		return
	}
	Rasterize_rounded_rect(framebuffer, stride, x, y, w, h, radius, color)
}
