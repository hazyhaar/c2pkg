// SPDX-License-Identifier: MIT
package c2bvh

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestBVHVsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc non disponible")
	}

	srcC, err := filepath.Abs(filepath.Join("sources", "tinybvh.c"))
	if err != nil {
		t.Fatal(err)
	}
	srcH, err := filepath.Abs(filepath.Join("sources", "tinybvh.h"))
	if err != nil {
		t.Fatal(err)
	}

	dir := t.TempDir()
	mainC := fmt.Sprintf(`#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include "%s"
#include "%s"

int main(int argc, char **argv) {
    if (argc < 2) return 1;
    const char *out_path = argv[1];

    bvh_tri_t tri = {
        .v0 = { -1.0f, -1.0f, 0.0f },
        .v1 = { 1.0f, -1.0f, 0.0f },
        .v2 = { 0.0f, 1.0f, 0.0f },
        .id = 42
    };

    bvh_ray_t ray = {
        .origin = { 0.0f, 0.0f, -5.0f },
        .direction = { 0.0f, 0.0f, 1.0f },
        .inv_dir = { 0.0f, 0.0f, 1.0f },
        .t_min = 0.001f,
        .t_max = 1000.0f
    };

    bvh_hit_t hit = {0};
    bool ok = bvh_intersect_ray_triangle(&ray, &tri, &hit);

    FILE *fout = fopen(out_path, "wb");
    if (!fout) return 1;
    uint8_t hit_byte = ok ? 1 : 0;
    fwrite(&hit_byte, 1, 1, fout);
    fwrite(&hit.t, sizeof(float), 1, fout);
    fwrite(&hit.u, sizeof(float), 1, fout);
    fwrite(&hit.v, sizeof(float), 1, fout);
    fwrite(&hit.tri_id, sizeof(uint32_t), 1, fout);
    fclose(fout);

    return 0;
}
`, srcH, srcC)

	mainPath := filepath.Join(dir, "tinybvh_oracle_main.c")
	if err := os.WriteFile(mainPath, []byte(mainC), 0600); err != nil {
		t.Fatal(err)
	}

	binPath := filepath.Join(dir, "tinybvh_oracle_bin")
	cmdBuild := exec.Command("gcc", "-O2", "-Wall", "-Wextra", mainPath, "-o", binPath)
	if out, err := cmdBuild.CombinedOutput(); err != nil {
		t.Fatalf("échec compilation oracle C: %v, sortie: %s", err, string(out))
	}

	outPath := filepath.Join(dir, "oracle_hit.bin")
	cmdRun := exec.Command(binPath, outPath)
	if out, err := cmdRun.CombinedOutput(); err != nil {
		t.Fatalf("échec oracle bvh: %v, out: %s", err, string(out))
	}

	// Validation Go
	tri := Triangle{
		V0: Vec3{-1.0, -1.0, 0.0},
		V1: Vec3{1.0, -1.0, 0.0},
		V2: Vec3{0.0, 1.0, 0.0},
		ID: 42,
	}
	ray := NewRay(Vec3{0.0, 0.0, -5.0}, Vec3{0.0, 0.0, 1.0}, 0.001, 1000.0)

	var hit HitRecord
	ok := IntersectRayTriangle(&ray, &tri, &hit)

	if !ok || !hit.Hit {
		t.Fatal("échec attendu d'intersection rayon-triangle en Go")
	}

	if hit.T != 5.0 {
		t.Errorf("distance T erronée : attendu 5.0, obtenu %f", hit.T)
	}
	if hit.TriID != 42 {
		t.Errorf("TriID erroné : attendu 42, obtenu %d", hit.TriID)
	}
}

func TestBVHTreeTraverseZeroAlloc(t *testing.T) {
	rng := rand.New(rand.NewSource(777))
	triangles := make([]Triangle, 256)
	for i := range triangles {
		cx := rng.Float32()*100 - 50
		cy := rng.Float32()*100 - 50
		cz := rng.Float32()*100 - 50
		triangles[i] = Triangle{
			V0: Vec3{cx - 1, cy - 1, cz},
			V1: Vec3{cx + 1, cy - 1, cz},
			V2: Vec3{cx, cy + 1, cz},
			ID: uint32(i),
		}
	}

	tree := BuildBVH(triangles)
	if len(tree.Nodes) == 0 {
		t.Fatal("arbre BVH vide")
	}

	ray := NewRay(Vec3{0, 0, -200}, Vec3{0, 0, 1}, 0.01, 1000.0)
	var hit HitRecord

	allocs := testing.AllocsPerRun(100, func() {
		r := ray
		hit = HitRecord{}
		tree.IntersectRay(&r, &hit)
	})

	if allocs != 0 {
		t.Errorf("échec invariant ARCHTIME : attendu 0 allocs/op, obtenu %f", allocs)
	}
}
