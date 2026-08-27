// SPDX-License-Identifier: Apache-2.0 OR MIT
package c2raster

import (
	"math"
	"runtime"
)

const (
	SubpixelBits  = 4
	SubpixelScale = 1 << SubpixelBits
)

// Vertex3D représente un sommet géométrique avec coordonnées homogènes écran, texture, normale et couleur.
type Vertex3D struct {
	X, Y, Z, InvW float32
	U, V          float32
	Nx, Ny, Nz    float32
	Color         uint32
}

// Triangle3D représente un triangle 3D défini par 3 sommets ordonnés.
type Triangle3D struct {
	V0 Vertex3D
	V1 Vertex3D
	V2 Vertex3D
}

// Tile64Bounds définit les bornes 2D d'une tuile de calcul 64x64.
type Tile64Bounds struct {
	X0, Y0 int
	X1, Y1 int
}

type workerTask struct {
	tStart, tEnd int
	tris         []Triangle3D
}

// RasterContext3D encapsule les tampons de rendu, la grille de tuiles et les structures de travail réutilisables.
// Invariant ARCHTIME : 0 allocation dynamique par trame une fois initialisé.
type RasterContext3D struct {
	Width       int
	Height      int
	Stride      int
	ColorBuffer []uint32
	DepthBuffer []float32

	NumTilesX int
	NumTilesY int
	Tiles     []Tile64Bounds
	TileMinZ  []float32
	TileMaxZ  []float32

	// Pool de workers persistants pour 0 alloc/trame en parallèle
	numWorkers int
	taskChans  []chan workerTask
	workerDone []chan struct{}
	stopped    bool
}

// NewRasterContext3D alloue et initialise un contexte de rastérisation tuilée de résolution W x H.
func NewRasterContext3D(width, height, stride int) *RasterContext3D {
	if stride < width {
		stride = width
	}
	totalPixels := stride * height

	numTilesX := (width + 63) / 64
	numTilesY := (height + 63) / 64
	numTiles := numTilesX * numTilesY

	tiles := make([]Tile64Bounds, numTiles)
	tileMinZ := make([]float32, numTiles)
	tileMaxZ := make([]float32, numTiles)

	for ty := 0; ty < numTilesY; ty++ {
		for tx := 0; tx < numTilesX; tx++ {
			idx := ty*numTilesX + tx
			x0 := tx * 64
			y0 := ty * 64
			x1 := x0 + 64
			y1 := y0 + 64
			if x1 > width {
				x1 = width
			}
			if y1 > height {
				y1 = height
			}
			tiles[idx] = Tile64Bounds{X0: x0, Y0: y0, X1: x1, Y1: y1}
			tileMinZ[idx] = 1.0
			tileMaxZ[idx] = 1.0
		}
	}

	workers := runtime.GOMAXPROCS(0)
	if workers < 1 {
		workers = 1
	}

	ctx := &RasterContext3D{
		Width:       width,
		Height:      height,
		Stride:      stride,
		ColorBuffer: make([]uint32, totalPixels),
		DepthBuffer: make([]float32, totalPixels),
		NumTilesX:   numTilesX,
		NumTilesY:   numTilesY,
		Tiles:       tiles,
		TileMinZ:    tileMinZ,
		TileMaxZ:    tileMaxZ,
		numWorkers:  workers,
		taskChans:   make([]chan workerTask, workers),
		workerDone:  make([]chan struct{}, workers),
	}

	for w := 0; w < workers; w++ {
		ctx.taskChans[w] = make(chan workerTask, 1)
		ctx.workerDone[w] = make(chan struct{}, 1)
		go ctx.workerLoop(w)
	}

	return ctx
}

func (ctx *RasterContext3D) workerLoop(workerID int) {
	for task := range ctx.taskChans[workerID] {
		for tIdx := task.tStart; tIdx < task.tEnd; tIdx++ {
			tile := ctx.Tiles[tIdx]
			for i := range task.tris {
				tri := &task.tris[i]
				minX := min3f(tri.V0.X, tri.V1.X, tri.V2.X)
				maxX := max3f(tri.V0.X, tri.V1.X, tri.V2.X)
				minY := min3f(tri.V0.Y, tri.V1.Y, tri.V2.Y)
				maxY := max3f(tri.V0.Y, tri.V1.Y, tri.V2.Y)

				if maxX < float32(tile.X0) || minX >= float32(tile.X1) ||
					maxY < float32(tile.Y0) || minY >= float32(tile.Y1) {
					continue
				}

				minZ := min3f(tri.V0.Z, tri.V1.Z, tri.V2.Z)
				if minZ >= ctx.TileMaxZ[tIdx] {
					continue
				}

				RasterizeTriangleOnTile(tri, ctx.ColorBuffer, ctx.DepthBuffer, ctx.Width, ctx.Height, ctx.Stride, tile.X0, tile.Y0, tile.X1, tile.Y1)
			}
		}
		ctx.workerDone[workerID] <- struct{}{}
	}
}

// Close libère les goroutines de travail persistantes.
func (ctx *RasterContext3D) Close() {
	if ctx.stopped {
		return
	}
	ctx.stopped = true
	for w := 0; w < ctx.numWorkers; w++ {
		close(ctx.taskChans[w])
	}
}

// Clear réinitialise les tampons couleur et profondeur avec les valeurs spécifiées.
func (ctx *RasterContext3D) Clear(clearColor uint32, clearDepth float32) {
	for i := range ctx.ColorBuffer {
		ctx.ColorBuffer[i] = clearColor
	}
	for i := range ctx.DepthBuffer {
		ctx.DepthBuffer[i] = clearDepth
	}
	for i := range ctx.TileMinZ {
		ctx.TileMinZ[i] = clearDepth
		ctx.TileMaxZ[i] = clearDepth
	}
}

func min3f(a, b, c float32) float32 {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}

func max3f(a, b, c float32) float32 {
	m := a
	if b > m {
		m = b
	}
	if c > m {
		m = c
	}
	return m
}

func clampF(v, minV, maxV float32) float32 {
	if v < minV {
		return minV
	}
	if v > maxV {
		return maxV
	}
	return v
}

func clampInt(v, minV, maxV int) int {
	if v < minV {
		return minV
	}
	if v > maxV {
		return maxV
	}
	return v
}

func edgeFuncFixed(ax, ay, bx, by, px, py int64) int64 {
	return (bx-ax)*(py-ay) - (by-ay)*(px-ax)
}

func isTopLeftFixed(ax, ay, bx, by int64) bool {
	dx := bx - ax
	dy := by - ay
	return (dy < 0) || (dy == 0 && dx < 0)
}

// RasterizeTriangleOnTile rastérise un triangle 3D borné strictement à une tuile donnée.
func RasterizeTriangleOnTile(
	tri *Triangle3D,
	colorBuf []uint32,
	depthBuf []float32,
	width, height, stride int,
	tileX0, tileY0, tileX1, tileY1 int,
) {
	if tri.V0.InvW <= 0 || tri.V1.InvW <= 0 || tri.V2.InvW <= 0 {
		return
	}

	x0 := int64(math.Round(float64(tri.V0.X * float32(SubpixelScale))))
	y0 := int64(math.Round(float64(tri.V0.Y * float32(SubpixelScale))))
	x1 := int64(math.Round(float64(tri.V1.X * float32(SubpixelScale))))
	y1 := int64(math.Round(float64(tri.V1.Y * float32(SubpixelScale))))
	x2 := int64(math.Round(float64(tri.V2.X * float32(SubpixelScale))))
	y2 := int64(math.Round(float64(tri.V2.Y * float32(SubpixelScale))))

	areaFixed := edgeFuncFixed(x0, y0, x1, y1, x2, y2)
	if areaFixed <= 0 {
		return
	}

	minX := min3f(tri.V0.X, tri.V1.X, tri.V2.X)
	maxX := max3f(tri.V0.X, tri.V1.X, tri.V2.X)
	minY := min3f(tri.V0.Y, tri.V1.Y, tri.V2.Y)
	maxY := max3f(tri.V0.Y, tri.V1.Y, tri.V2.Y)

	minIX := clampInt(int(math.Floor(float64(minX))), tileX0, tileX1-1)
	maxIX := clampInt(int(math.Ceil(float64(maxX))), tileX0, tileX1-1)
	minIY := clampInt(int(math.Floor(float64(minY))), tileY0, tileY1-1)
	maxIY := clampInt(int(math.Ceil(float64(maxY))), tileY0, tileY1-1)

	if minIX > maxIX || minIY > maxIY {
		return
	}

	invArea := 1.0 / float64(areaFixed)

	var bias0 int64 = -1
	if isTopLeftFixed(x1, y1, x2, y2) {
		bias0 = 0
	}
	var bias1 int64 = -1
	if isTopLeftFixed(x2, y2, x0, y0) {
		bias1 = 0
	}
	var bias2 int64 = -1
	if isTopLeftFixed(x0, y0, x1, y1) {
		bias2 = 0
	}

	c0 := tri.V0.Color
	c1 := tri.V1.Color
	c2 := tri.V2.Color

	r0 := float32(c0&0xff) * tri.V0.InvW
	g0 := float32((c0>>8)&0xff) * tri.V0.InvW
	b0 := float32((c0>>16)&0xff) * tri.V0.InvW
	a0 := float32((c0>>24)&0xff) * tri.V0.InvW

	r1 := float32(c1&0xff) * tri.V1.InvW
	g1 := float32((c1>>8)&0xff) * tri.V1.InvW
	b1 := float32((c1>>16)&0xff) * tri.V1.InvW
	a1 := float32((c1>>24)&0xff) * tri.V1.InvW

	r2 := float32(c2&0xff) * tri.V2.InvW
	g2 := float32((c2>>8)&0xff) * tri.V2.InvW
	b2 := float32((c2>>16)&0xff) * tri.V2.InvW
	a2 := float32((c2>>24)&0xff) * tri.V2.InvW

	z0 := tri.V0.Z
	z1 := tri.V1.Z
	z2 := tri.V2.Z

	for y := minIY; y <= maxIY; y++ {
		py := (int64(y) << SubpixelBits) + (SubpixelScale / 2)
		row := y * stride

		for x := minIX; x <= maxIX; x++ {
			px := (int64(x) << SubpixelBits) + (SubpixelScale / 2)

			w0 := edgeFuncFixed(x1, y1, x2, y2, px, py)
			w1 := edgeFuncFixed(x2, y2, x0, y0, px, py)
			w2 := edgeFuncFixed(x0, y0, x1, y1, px, py)

			if (w0+bias0 >= 0) && (w1+bias1 >= 0) && (w2+bias2 >= 0) {
				l0 := float32(float64(w0) * invArea)
				l1 := float32(float64(w1) * invArea)
				l2 := float32(float64(w2) * invArea)

				zInterp := l0*z0 + l1*z1 + l2*z2

				idx := row + x
				if depthBuf != nil {
					if zInterp >= depthBuf[idx] {
						continue
					}
					depthBuf[idx] = zInterp
				}

				invW := l0*tri.V0.InvW + l1*tri.V1.InvW + l2*tri.V2.InvW
				if invW <= 0 {
					continue
				}
				wInterp := float32(1.0) / invW

				if colorBuf != nil {
					r := (l0*r0 + l1*r1 + l2*r2) * wInterp
					g := (l0*g0 + l1*g1 + l2*g2) * wInterp
					b := (l0*b0 + l1*b1 + l2*b2) * wInterp
					a := (l0*a0 + l1*a1 + l2*a2) * wInterp

					ur := uint8(clampF(r, 0, 255))
					ug := uint8(clampF(g, 0, 255))
					ub := uint8(clampF(b, 0, 255))
					ua := uint8(clampF(a, 0, 255))

					colorBuf[idx] = uint32(ur) | (uint32(ug) << 8) | (uint32(ub) << 16) | (uint32(ua) << 24)
				}
			}
		}
	}
}

// RasterizeTriangle rastérise un triangle 3D sur l'ensemble de l'écran avec tuilage adaptatif.
func (ctx *RasterContext3D) RasterizeTriangle(tri *Triangle3D) {
	minX := min3f(tri.V0.X, tri.V1.X, tri.V2.X)
	maxX := max3f(tri.V0.X, tri.V1.X, tri.V2.X)
	minY := min3f(tri.V0.Y, tri.V1.Y, tri.V2.Y)
	maxY := max3f(tri.V0.Y, tri.V1.Y, tri.V2.Y)

	for i := range ctx.Tiles {
		t := ctx.Tiles[i]
		if maxX < float32(t.X0) || minX >= float32(t.X1) ||
			maxY < float32(t.Y0) || minY >= float32(t.Y1) {
			continue
		}
		RasterizeTriangleOnTile(tri, ctx.ColorBuffer, ctx.DepthBuffer, ctx.Width, ctx.Height, ctx.Stride, t.X0, t.Y0, t.X1, t.Y1)
	}
}

// RasterizeTrianglesParallel rastérise un tableau de triangles 3D en parallèle sur les tuiles L1 avec 0 allocation.
func (ctx *RasterContext3D) RasterizeTrianglesParallel(tris []Triangle3D) {
	if len(tris) == 0 {
		return
	}

	numTiles := len(ctx.Tiles)
	tilesPerWorker := (numTiles + ctx.numWorkers - 1) / ctx.numWorkers

	activeWorkers := 0
	for w := 0; w < ctx.numWorkers; w++ {
		tStart := w * tilesPerWorker
		tEnd := tStart + tilesPerWorker
		if tEnd > numTiles {
			tEnd = numTiles
		}
		if tStart >= tEnd {
			break
		}

		activeWorkers++
		ctx.taskChans[w] <- workerTask{tStart: tStart, tEnd: tEnd, tris: tris}
	}

	for w := 0; w < activeWorkers; w++ {
		<-ctx.workerDone[w]
	}
}
