# Proposal: `golang.org/x/image/raster3d`: Pure-Go, Zero-Allocation Software 3D Triangle Rasterizer

**Author:** Hazyhaar  
**Target:** Go Issue Tracker (`golang/go` / `golang.org/x/image`)  
**Status:** DRAFT / PROPOSAL (No external publication)  
**Date:** 2026-08-28  

---

## 1. Context & Motivation

In the Go standard sub-repositories, **`golang.org/x/image/vector`** provides robust, pure-Go 2D vector and Bézier rasterization (developed by Nigel Tao and Rob Pike). 

However, there is currently **no official, pure-Go package for 3D geometric rasterization**. 

Applications requiring offline rendering, headless CI validation of 3D assets (glTF, OBJ), CLI terminal graphics (TUI / Braille rendering), or embedded rendering currently must either:
1. Pull in heavy, platform-specific CGO bindings to GPU runtimes (OpenGL, Vulkan, Metal).
2. Suffer from un-optimized scalar loops with heavy per-frame memory allocation.

We propose adding **`golang.org/x/image/raster3d`** (or a subpackage under `x/image`): a lightweight, standalone 3D rendering pipeline (4,700 LOC total across mesh optimization, glTF parsing, BVH acceleration, PBR shading and the core 400 LOC subpixel rasterizer kernel) with **zero dynamic heap allocations per frame** and bit-exact deterministic rendering.

---

## 2. Core Architecture & Guarantees

```
                   PIPELINE D'EXÉCUTION TUILE 64x64 EN PUR GO
  ┌─────────────────────────────────────────────────────────────────────────────┐
  │ 1. Sub-Pixel Precision │ 4-bit fixed-point (1/16th pixel) edge functions    │
  │    & Top-Left Rule     │ Guaranteed watertight shared edges (0 seam gaps)   │
  ├────────────────────────┼────────────────────────────────────────────────────┤
  │ 2. Tiled Binning       │ 64x64 screen-space tiles distributed across CPU    │
  │    (Zero-Contention)   │ workers via persistent pre-allocated channels      │
  ├────────────────────────┼────────────────────────────────────────────────────┤
  │ 3. Interpolation       │ Perspective-correct barycentric interpolation      │
  │                        │ (Z-Buffer depth, RGB color, UV textures, Normals)  │
  ├────────────────────────┼────────────────────────────────────────────────────┤
  │ 4. Zero Alloc Invariant│ 0 B/op, 0 allocs/op once RasterContext3D is init   │
  │                        │ Reusable contiguous ColorBuffer and DepthBuffer    │
  └─────────────────────────────────────────────────────────────────────────────┘
```

---

## 3. Benchmark Data & Validation

### Measured Performance (Single & Multi-Threaded on Intel Core i9-14900K)
* **Real-time Performance:** Capable of rendering complex scenes (e.g., animated aquatic mesh with thousands of triangles) at **1080p 60 FPS** entirely on CPU.
* **Heap Allocations:** **`0 B/op` and `0 allocs/op`** during established render loops (`TestRaster3DZeroAlloc` PASS).
* **Parity Guarantee:** Bit-exact pixel parity against C reference oracle compiled under `gcc -O2 -fsanitize=address,undefined` (`TestRaster3DVsCOracle` PASS).
* **Edge Invariant:** Mathematical proof and unit tests verifying watertight boundaries (`TestRaster3DSharedEdgeWatertight` PASS).

---

## 4. API Design (Sketch)

```go
package raster3d

// Vertex3D represents a screen-space vertex with homogeneous depth and attributes.
type Vertex3D struct {
	X, Y, Z, InvW float32
	U, V          float32
	Nx, Ny, Nz    float32
	Color         uint32
}

// Triangle3D represents a 3D triangle defined by three ordered vertices.
type Triangle3D struct {
	V0, V1, V2 Vertex3D
}

// Context3D encapsulates pre-allocated framebuffers and tile worker pools.
type Context3D struct {
	Width, Height int
	Stride        int
	ColorBuffer   []uint32
	DepthBuffer   []float32
	// ... persistent worker pools ...
}

func NewContext3D(width, height int, numWorkers int) *Context3D
func (ctx *Context3D) Clear(color uint32, depth float32)
func (ctx *Context3D) DrawTriangles(tris []Triangle3D)
func (ctx *Context3D) Close()
```

---

## 5. Compatibility & Integration Path

- **Zero CGO:** 100% portable Go, compiles natively for Linux, macOS, Windows, FreeBSD, OpenBSD, and WebAssembly (`GOOS=js,wasm`).
- **Standard Library Interop:** Directly exports to `*image.RGBA` or `*image.NRGBA`.
- **Standalone Reference:** The implementation is currently battle-tested and available for review in [`github.com/hazyhaar/c2pkg/c2raster`](https://github.com/hazyhaar/c2pkg/tree/main/c2raster).
