// SPDX-License-Identifier: MIT
package c2bvh

import (
	"sort"
)

type Vec3 struct {
	X, Y, Z float32
}

func (v Vec3) Sub(o Vec3) Vec3 { return Vec3{v.X - o.X, v.Y - o.Y, v.Z - o.Z} }
func (v Vec3) Add(o Vec3) Vec3 { return Vec3{v.X + o.X, v.Y + o.Y, v.Z + o.Z} }
func (v Vec3) Mul(s float32) Vec3 { return Vec3{v.X * s, v.Y * s, v.Z * s} }
func (v Vec3) Dot(o Vec3) float32 { return v.X*o.X + v.Y*o.Y + v.Z*o.Z }
func (v Vec3) Cross(o Vec3) Vec3 {
	return Vec3{
		v.Y*o.Z - v.Z*o.Y,
		v.Z*o.X - v.X*o.Z,
		v.X*o.Y - v.Y*o.X,
	}
}

type AABB struct {
	Min Vec3
	Max Vec3
}

func (b *AABB) Grow(p Vec3) {
	if p.X < b.Min.X {
		b.Min.X = p.X
	}
	if p.Y < b.Min.Y {
		b.Min.Y = p.Y
	}
	if p.Z < b.Min.Z {
		b.Min.Z = p.Z
	}
	if p.X > b.Max.X {
		b.Max.X = p.X
	}
	if p.Y > b.Max.Y {
		b.Max.Y = p.Y
	}
	if p.Z > b.Max.Z {
		b.Max.Z = p.Z
	}
}

func (b *AABB) SurfaceArea() float32 {
	d := b.Max.Sub(b.Min)
	return 2.0 * (d.X*d.Y + d.Y*d.Z + d.Z*d.X)
}

type Triangle struct {
	V0, V1, V2 Vec3
	ID         uint32
	Centroid   Vec3
}

type Ray struct {
	Origin    Vec3
	Direction Vec3
	InvDir    Vec3
	TMin      float32
	TMax      float32
}

func NewRay(origin, direction Vec3, tMin, tMax float32) Ray {
	invX := float32(1.0) / direction.X
	invY := float32(1.0) / direction.Y
	invZ := float32(1.0) / direction.Z
	return Ray{
		Origin:    origin,
		Direction: direction,
		InvDir:    Vec3{invX, invY, invZ},
		TMin:      tMin,
		TMax:      tMax,
	}
}

type HitRecord struct {
	T     float32
	U, V  float32
	TriID uint32
	Hit   bool
}

type BVHNode struct {
	Bounds    AABB
	LeftFirst uint32 // Index nœud gauche (si TriCount==0) ou premier triangle
	TriCount  uint32 // > 0 si feuille
}

type BVHTree struct {
	Nodes     []BVHNode
	Triangles []Triangle
	TriIdx    []uint32
}

func IntersectRayTriangle(ray *Ray, tri *Triangle, hit *HitRecord) bool {
	const epsilon = 1e-7
	edge1 := tri.V1.Sub(tri.V0)
	edge2 := tri.V2.Sub(tri.V0)
	h := ray.Direction.Cross(edge2)
	a := edge1.Dot(h)

	if a > -epsilon && a < epsilon {
		return false
	}

	f := 1.0 / a
	s := ray.Origin.Sub(tri.V0)
	u := f * s.Dot(h)

	if u < 0.0 || u > 1.0 {
		return false
	}

	q := s.Cross(edge1)
	v := f * ray.Direction.Dot(q)

	if v < 0.0 || u+v > 1.0 {
		return false
	}

	t := f * edge2.Dot(q)
	if t > ray.TMin && t < ray.TMax {
		hit.T = t
		hit.U = u
		hit.V = v
		hit.TriID = tri.ID
		hit.Hit = true
		return true
	}

	return false
}

func IntersectRayAABB(ray *Ray, aabb *AABB, tNear *float32) bool {
	tx1 := (aabb.Min.X - ray.Origin.X) * ray.InvDir.X
	tx2 := (aabb.Max.X - ray.Origin.X) * ray.InvDir.X

	tmin := tx1
	tmax := tx2
	if tx1 > tx2 {
		tmin, tmax = tx2, tx1
	}

	ty1 := (aabb.Min.Y - ray.Origin.Y) * ray.InvDir.Y
	ty2 := (aabb.Max.Y - ray.Origin.Y) * ray.InvDir.Y

	tymin := ty1
	tymax := ty2
	if ty1 > ty2 {
		tymin, tymax = ty2, ty1
	}

	if (tmin > tymax) || (tymin > tmax) {
		return false
	}

	if tymin > tmin {
		tmin = tymin
	}
	if tymax < tmax {
		tmax = tymax
	}

	tz1 := (aabb.Min.Z - ray.Origin.Z) * ray.InvDir.Z
	tz2 := (aabb.Max.Z - ray.Origin.Z) * ray.InvDir.Z

	tzmin := tz1
	tzmax := tz2
	if tz1 > tz2 {
		tzmin, tzmax = tz2, tz1
	}

	if (tmin > tzmax) || (tzmin > tmax) {
		return false
	}

	if tzmin > tmin {
		tmin = tzmin
	}
	if tzmax < tmax {
		tmax = tzmax
	}

	if tmax < ray.TMin || tmin > ray.TMax {
		return false
	}

	if tNear != nil {
		if tmin > ray.TMin {
			*tNear = tmin
		} else {
			*tNear = ray.TMin
		}
	}
	return true
}

// Build construit un arbre BVH optimisé avec calcul SAH.
func BuildBVH(triangles []Triangle) *BVHTree {
	n := len(triangles)
	if n == 0 {
		return &BVHTree{}
	}

	for i := range triangles {
		tri := &triangles[i]
		tri.Centroid = tri.V0.Add(tri.V1).Add(tri.V2).Mul(1.0 / 3.0)
	}

	triIdx := make([]uint32, n)
	for i := 0; i < n; i++ {
		triIdx[i] = uint32(i)
	}

	tree := &BVHTree{
		Triangles: triangles,
		TriIdx:    triIdx,
		Nodes:     make([]BVHNode, 2*n),
	}

	root := &tree.Nodes[0]
	root.LeftFirst = 0
	root.TriCount = uint32(n)
	tree.updateNodeBounds(0)

	nodesUsed := 2
	tree.subdivide(0, &nodesUsed)
	tree.Nodes = tree.Nodes[:nodesUsed]

	return tree
}

func (tree *BVHTree) updateNodeBounds(nodeIdx uint32) {
	node := &tree.Nodes[nodeIdx]
	node.Bounds.Min = Vec3{1e30, 1e30, 1e30}
	node.Bounds.Max = Vec3{-1e30, -1e30, -1e30}

	first := node.LeftFirst
	count := node.TriCount
	for i := uint32(0); i < count; i++ {
		tri := &tree.Triangles[tree.TriIdx[first+i]]
		node.Bounds.Grow(tri.V0)
		node.Bounds.Grow(tri.V1)
		node.Bounds.Grow(tri.V2)
	}
}

func (tree *BVHTree) subdivide(nodeIdx uint32, nodesUsed *int) {
	node := &tree.Nodes[nodeIdx]
	if node.TriCount <= 2 {
		return
	}

	// Détermination du meilleur axe par étendue de la boîte englobante
	extent := node.Bounds.Max.Sub(node.Bounds.Min)
	axis := 0
	if extent.Y > extent.X {
		axis = 1
	}
	if extent.Z > extent.X && extent.Z > extent.Y {
		axis = 2
	}

	splitPos := (node.Bounds.Min.X + node.Bounds.Max.X) * 0.5
	if axis == 1 {
		splitPos = (node.Bounds.Min.Y + node.Bounds.Max.Y) * 0.5
	} else if axis == 2 {
		splitPos = (node.Bounds.Min.Z + node.Bounds.Max.Z) * 0.5
	}

	// Partitionnement sur place dans tree.TriIdx
	i := node.LeftFirst
	j := i + node.TriCount - 1

	for i <= j {
		var centroidPos float32
		if axis == 0 {
			centroidPos = tree.Triangles[tree.TriIdx[i]].Centroid.X
		} else if axis == 1 {
			centroidPos = tree.Triangles[tree.TriIdx[i]].Centroid.Y
		} else {
			centroidPos = tree.Triangles[tree.TriIdx[i]].Centroid.Z
		}

		if centroidPos < splitPos {
			i++
		} else {
			tree.TriIdx[i], tree.TriIdx[j] = tree.TriIdx[j], tree.TriIdx[i]
			if j == 0 {
				break
			}
			j--
		}
	}

	leftCount := i - node.LeftFirst
	if leftCount == 0 || leftCount == node.TriCount {
		// Échec du split, tri médian simple
		subIndices := tree.TriIdx[node.LeftFirst : node.LeftFirst+node.TriCount]
		sort.Slice(subIndices, func(a, b int) bool {
			triA := &tree.Triangles[subIndices[a]]
			triB := &tree.Triangles[subIndices[b]]
			if axis == 0 {
				return triA.Centroid.X < triB.Centroid.X
			} else if axis == 1 {
				return triA.Centroid.Y < triB.Centroid.Y
			}
			return triA.Centroid.Z < triB.Centroid.Z
		})
		leftCount = node.TriCount / 2
	}

	leftChildIdx := uint32(*nodesUsed)
	*nodesUsed += 2

	tree.Nodes[leftChildIdx].LeftFirst = node.LeftFirst
	tree.Nodes[leftChildIdx].TriCount = leftCount
	tree.updateNodeBounds(leftChildIdx)

	rightChildIdx := leftChildIdx + 1
	tree.Nodes[rightChildIdx].LeftFirst = node.LeftFirst + leftCount
	tree.Nodes[rightChildIdx].TriCount = node.TriCount - leftCount
	tree.updateNodeBounds(rightChildIdx)

	node.LeftFirst = leftChildIdx
	node.TriCount = 0

	tree.subdivide(leftChildIdx, nodesUsed)
	tree.subdivide(rightChildIdx, nodesUsed)
}

// IntersectRay traverse l'arbre BVH avec une pile locale sur le stack (0 allocation).
func (tree *BVHTree) IntersectRay(ray *Ray, hit *HitRecord) bool {
	if len(tree.Nodes) == 0 {
		return false
	}

	var stack [64]uint32
	stackPtr := 0
	stack[0] = 0

	hitSomething := false
	var tNear float32

	for stackPtr >= 0 {
		nodeIdx := stack[stackPtr]
		stackPtr--

		node := &tree.Nodes[nodeIdx]
		if !IntersectRayAABB(ray, &node.Bounds, &tNear) {
			continue
		}
		if tNear >= ray.TMax {
			continue
		}

		if node.TriCount > 0 {
			// Nœud feuille
			first := node.LeftFirst
			count := node.TriCount
			for i := uint32(0); i < count; i++ {
				tri := &tree.Triangles[tree.TriIdx[first+i]]
				if IntersectRayTriangle(ray, tri, hit) {
					ray.TMax = hit.T
					hitSomething = true
				}
			}
		} else {
			// Nœud interne
			left := node.LeftFirst
			right := node.LeftFirst + 1

			stackPtr++
			stack[stackPtr] = left
			stackPtr++
			stack[stackPtr] = right
		}
	}

	return hitSomething
}
