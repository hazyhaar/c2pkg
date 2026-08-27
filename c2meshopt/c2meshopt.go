// SPDX-License-Identifier: MIT
package c2meshopt

import (
	"math"
)

type vertexScoreInfo struct {
	cacheTag              int
	score                 float32
	activeTrianglesCount  uint32
	activeTrianglesOffset uint32
}

func calcVertexScore(cachePos int, activeTriangles uint32) float32 {
	if activeTriangles == 0 {
		return -1.0
	}

	var score float32
	if cachePos < 0 {
		// Pas dans le cache
	} else if cachePos < 3 {
		score += 0.75
	} else {
		const scaler = float32(1.0) / float32(32-3)
		diff := float32(1.0) - float32(cachePos-3)*scaler
		score += float32(math.Pow(float64(diff), 1.5))
	}

	valenceBoost := float32(math.Pow(float64(activeTriangles), -0.5))
	score += 2.0 * valenceBoost
	return score
}

// OptimizeVertexCache réordonne les indices du maillage pour minimiser les défauts de cache de sommets (L1 / Post-Transform Cache).
// Exécution déterministe et conforme bit-exacte avec l'algorithme de Forsyth.
func OptimizeVertexCache(destination []uint32, indices []uint32, vertexCount int) {
	indexCount := len(indices)
	if len(destination) < indexCount || indexCount == 0 || vertexCount == 0 || indexCount%3 != 0 {
		return
	}

	triangleCount := indexCount / 3

	vertexTriangles := make([]uint32, indexCount)
	vertexOffsets := make([]uint32, vertexCount+1)
	vertexCounts := make([]uint32, vertexCount)
	triangleEmitted := make([]byte, triangleCount)
	triangleScores := make([]float32, triangleCount)

	// 1. Valences
	for _, v := range indices {
		if int(v) < vertexCount {
			vertexCounts[v]++
		}
	}

	// 2. Offsets
	var offset uint32
	for i := 0; i < vertexCount; i++ {
		vertexOffsets[i] = offset
		offset += vertexCounts[i]
		vertexCounts[i] = 0
	}
	vertexOffsets[vertexCount] = offset

	// 3. Remplissage adjacence
	for i := 0; i < triangleCount; i++ {
		v0 := indices[i*3+0]
		v1 := indices[i*3+1]
		v2 := indices[i*3+2]

		if int(v0) < vertexCount {
			vertexTriangles[vertexOffsets[v0]+vertexCounts[v0]] = uint32(i)
			vertexCounts[v0]++
		}
		if int(v1) < vertexCount {
			vertexTriangles[vertexOffsets[v1]+vertexCounts[v1]] = uint32(i)
			vertexCounts[v1]++
		}
		if int(v2) < vertexCount {
			vertexTriangles[vertexOffsets[v2]+vertexCounts[v2]] = uint32(i)
			vertexCounts[v2]++
		}
	}

	// 4. Initialisation des scores sommets
	vertexInfo := make([]vertexScoreInfo, vertexCount)
	for i := 0; i < vertexCount; i++ {
		vertexInfo[i].cacheTag = -1
		vertexInfo[i].activeTrianglesCount = vertexCounts[i]
		vertexInfo[i].activeTrianglesOffset = vertexOffsets[i]
		vertexInfo[i].score = calcVertexScore(-1, vertexCounts[i])
	}

	// 5. Calcul initial scores triangles
	for i := 0; i < triangleCount; i++ {
		v0 := indices[i*3+0]
		v1 := indices[i*3+1]
		v2 := indices[i*3+2]
		var score float32
		if int(v0) < vertexCount {
			score += vertexInfo[v0].score
		}
		if int(v1) < vertexCount {
			score += vertexInfo[v1].score
		}
		if int(v2) < vertexCount {
			score += vertexInfo[v2].score
		}
		triangleScores[i] = score
	}

	var cache [32]uint32
	cacheSize := 0
	outIndex := 0

	for outIndex < indexCount {
		bestScore := float32(-1.0)
		bestTriangle := -1

		// Recherche dans les voisins du cache
		for c := 0; c < cacheSize; c++ {
			v := cache[c]
			if int(v) >= vertexCount {
				continue
			}
			start := vertexInfo[v].activeTrianglesOffset
			count := vertexInfo[v].activeTrianglesCount
			for ti := uint32(0); ti < count; ti++ {
				tri := vertexTriangles[start+ti]
				if triangleEmitted[tri] == 0 && triangleScores[tri] > bestScore {
					bestScore = triangleScores[tri]
					bestTriangle = int(tri)
				}
			}
		}

		// Recherche linéaire si aucun trouvé dans le cache
		if bestTriangle == -1 {
			for i := 0; i < triangleCount; i++ {
				if triangleEmitted[i] == 0 {
					bestTriangle = i
					break
				}
			}
		}

		if bestTriangle == -1 {
			break
		}

		triangleEmitted[bestTriangle] = 1
		tv0 := indices[bestTriangle*3+0]
		tv1 := indices[bestTriangle*3+1]
		tv2 := indices[bestTriangle*3+2]

		destination[outIndex] = tv0
		destination[outIndex+1] = tv1
		destination[outIndex+2] = tv2
		outIndex += 3

		triVerts := [3]uint32{tv0, tv1, tv2}
		for _, v := range triVerts {
			if int(v) < vertexCount && vertexInfo[v].activeTrianglesCount > 0 {
				start := vertexInfo[v].activeTrianglesOffset
				count := vertexInfo[v].activeTrianglesCount
				for ti := uint32(0); ti < count; ti++ {
					if vertexTriangles[start+ti] == uint32(bestTriangle) {
						vertexTriangles[start+ti] = vertexTriangles[start+count-1]
						vertexInfo[v].activeTrianglesCount--
						break
					}
				}
			}
		}

		// Mise à jour file LRU
		for _, v := range triVerts {
			foundPos := -1
			for c := 0; c < cacheSize; c++ {
				if cache[c] == v {
					foundPos = c
					break
				}
			}
			if foundPos >= 0 {
				for c := foundPos; c > 0; c-- {
					cache[c] = cache[c-1]
				}
				cache[0] = v
			} else {
				if cacheSize < 32 {
					cacheSize++
				}
				for c := cacheSize - 1; c > 0; c-- {
					cache[c] = cache[c-1]
				}
				cache[0] = v
			}
		}

		// Recalcul tags et scores sommets
		for c := 0; c < cacheSize; c++ {
			v := cache[c]
			if int(v) < vertexCount {
				vertexInfo[v].cacheTag = c
				vertexInfo[v].score = calcVertexScore(c, vertexInfo[v].activeTrianglesCount)
			}
		}

		// Recalcul scores triangles impactés
		for c := 0; c < cacheSize; c++ {
			v := cache[c]
			if int(v) >= vertexCount {
				continue
			}
			start := vertexInfo[v].activeTrianglesOffset
			count := vertexInfo[v].activeTrianglesCount
			for ti := uint32(0); ti < count; ti++ {
				tri := vertexTriangles[start+ti]
				if triangleEmitted[tri] == 0 {
					v0 := indices[tri*3+0]
					v1 := indices[tri*3+1]
					v2 := indices[tri*3+2]
					var sc float32
					if int(v0) < vertexCount {
						sc += vertexInfo[v0].score
					}
					if int(v1) < vertexCount {
						sc += vertexInfo[v1].score
					}
					if int(v2) < vertexCount {
						sc += vertexInfo[v2].score
					}
					triangleScores[tri] = sc
				}
			}
		}
	}
}

// CalcVertexCacheStats calcule le ratio de défauts moyen par triangle (ACMR) pour une taille de cache donnée.
func CalcVertexCacheStats(indices []uint32, vertexCount int, cacheSize int) float32 {
	if len(indices) == 0 || vertexCount == 0 || cacheSize == 0 {
		return 0.0
	}

	cache := make([]uint32, cacheSize)
	for i := range cache {
		cache[i] = math.MaxUint32
	}

	misses := 0
	for _, v := range indices {
		hit := false
		for c := 0; c < cacheSize; c++ {
			if cache[c] == v {
				hit = true
				for j := c; j > 0; j-- {
					cache[j] = cache[j-1]
				}
				cache[0] = v
				break
			}
		}
		if !hit {
			misses++
			for j := cacheSize - 1; j > 0; j-- {
				cache[j] = cache[j-1]
			}
			cache[0] = v
		}
	}

	return float32(misses) / float32(len(indices)/3)
}
