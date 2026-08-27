// SPDX-License-Identifier: MIT
package c2gltf

import (
	"encoding/binary"
	"errors"
	"math"
)

const (
	MagicGLB  = 0x46546C67 // 'glTF'
	ChunkJSON = 0x4E4F534A // 'JSON'
	ChunkBIN  = 0x004E4942 // 'BIN\0'
)

type ComponentType int

const (
	ComponentTypeR8   ComponentType = 5120
	ComponentTypeR8U  ComponentType = 5121
	ComponentTypeR16  ComponentType = 5122
	ComponentTypeR16U ComponentType = 5123
	ComponentTypeR32U ComponentType = 5125
	ComponentTypeR32F ComponentType = 5126
)

type Type int

const (
	TypeScalar Type = 1
	TypeVec2   Type = 2
	TypeVec3   Type = 3
	TypeVec4   Type = 4
	TypeMat2   Type = 5
	TypeMat3   Type = 6
	TypeMat4   Type = 7
)

// GLBHeader représente l'en-tête et les pointeurs vers les tronçons JSON et BIN d'un fichier .glb.
type GLBHeader struct {
	Magic        uint32
	Version      uint32
	Length       uint32
	JSONChunkLen uint32
	JSONChunkTyp uint32
	JSONData     []byte
	BINChunkLen  uint32
	BINChunkTyp  uint32
	BINData      []byte
}

// ParseGLBHeader parse les tronçons GLB binaires avec 0 allocation dynamique.
func ParseGLBHeader(data []byte) (GLBHeader, error) {
	if len(data) < 12 {
		return GLBHeader{}, errors.New("taille glb invalide")
	}

	magic := binary.LittleEndian.Uint32(data[0:4])
	version := binary.LittleEndian.Uint32(data[4:8])
	length := binary.LittleEndian.Uint32(data[8:12])

	if magic != MagicGLB || version != 2 || int(length) > len(data) {
		return GLBHeader{}, errors.New("magic ou version glb invalide")
	}

	hdr := GLBHeader{
		Magic:   magic,
		Version: version,
		Length:  length,
	}

	offset := 12
	if offset+8 > len(data) {
		return GLBHeader{}, errors.New("tronçon json tronqué")
	}
	c0Len := binary.LittleEndian.Uint32(data[offset : offset+4])
	c0Typ := binary.LittleEndian.Uint32(data[offset+4 : offset+8])
	offset += 8

	if c0Typ != ChunkJSON || offset+int(c0Len) > len(data) {
		return GLBHeader{}, errors.New("tronçon json invalide")
	}

	hdr.JSONChunkLen = c0Len
	hdr.JSONChunkTyp = c0Typ
	hdr.JSONData = data[offset : offset+int(c0Len)]
	offset += int(c0Len)

	// Alignement 4-octets
	offset = (offset + 3) & ^3

	for offset+8 <= len(data) {
		cLen := binary.LittleEndian.Uint32(data[offset : offset+4])
		cTyp := binary.LittleEndian.Uint32(data[offset+4 : offset+8])
		offset += 8

		if cTyp == ChunkBIN && offset+int(cLen) <= len(data) {
			hdr.BINChunkLen = cLen
			hdr.BINChunkTyp = cTyp
			hdr.BINData = data[offset : offset+int(cLen)]
			break
		}
		offset += int(cLen)
		offset = (offset + 3) & ^3
	}

	return hdr, nil
}

func numComponents(t Type) int {
	switch t {
	case TypeScalar:
		return 1
	case TypeVec2:
		return 2
	case TypeVec3:
		return 3
	case TypeVec4, TypeMat2:
		return 4
	case TypeMat3:
		return 9
	case TypeMat4:
		return 16
	default:
		return 0
	}
}

// AccessorReadFloat lit des composantes flottantes depuis un buffer d'accesseur sans allocation dans le slice fourni.
func AccessorReadFloat(compType ComponentType, t Type, src []byte, count int, stride int, out []float32) int {
	numComps := numComponents(t)
	if numComps == 0 || count == 0 {
		return 0
	}

	elemSize := 0
	switch compType {
	case ComponentTypeR8, ComponentTypeR8U:
		elemSize = 1
	case ComponentTypeR16, ComponentTypeR16U:
		elemSize = 2
	case ComponentTypeR32U, ComponentTypeR32F:
		elemSize = 4
	default:
		return 0
	}

	if stride < numComps*elemSize {
		stride = numComps * elemSize
	}

	outIdx := 0
	for i := 0; i < count; i++ {
		elemOff := i * stride
		for c := 0; c < numComps; c++ {
			cOff := elemOff + c*elemSize
			if cOff+elemSize > len(src) || outIdx >= len(out) {
				return outIdx
			}
			var val float32
			switch compType {
			case ComponentTypeR8:
				val = float32(int8(src[cOff]))
			case ComponentTypeR8U:
				val = float32(src[cOff])
			case ComponentTypeR16:
				v16 := int16(binary.LittleEndian.Uint16(src[cOff : cOff+2]))
				val = float32(v16)
			case ComponentTypeR16U:
				v16u := binary.LittleEndian.Uint16(src[cOff : cOff+2])
				val = float32(v16u)
			case ComponentTypeR32U:
				v32u := binary.LittleEndian.Uint32(src[cOff : cOff+4])
				val = float32(v32u)
			case ComponentTypeR32F:
				u := binary.LittleEndian.Uint32(src[cOff : cOff+4])
				val = math.Float32frombits(u)
			}
			out[outIdx] = val
			outIdx++
		}
	}

	return outIdx
}
