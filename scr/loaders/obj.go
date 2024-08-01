package load

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

type Obj struct {
	Vertices []vec3.Type
	UVs      []vec2.Type
	Normals  []vec3.Type
	Indices  []uint32
}

// FileExists checks if the given file exists.
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// OBJ loads a .obj file and returns an Obj.
func OBJ(filename string) Obj {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file '%s': %v\n", filename, err)
	}
	defer file.Close()

	var positions []vec3.Type
	var uvs []vec2.Type
	var normals []vec3.Type
	var indices []uint32
	var vertexData []float32
	indexOffset := 0
	activeMeshName := "__no_active"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)

		if len(tokens) == 0 {
			continue
		}

		switch tokens[0] {
		case "o":
			// This indicates the start of a new mesh.
			// Currently, we handle only one mesh.
			if activeMeshName != "__no_active" {
				return Obj{
					Vertices: positions,
					UVs:      uvs,
					Normals:  normals,
					Indices:  indices,
				}
			}
			activeMeshName = tokens[1]
		case "v":
			if len(tokens) < 4 {
				log.Fatalf("Invalid vertex line: %s\n", line)
			}
			x, errX := strconv.ParseFloat(tokens[1], 32)
			y, errY := strconv.ParseFloat(tokens[2], 32)
			z, errZ := strconv.ParseFloat(tokens[3], 32)
			if errX != nil || errY != nil || errZ != nil {
				log.Fatalf("Error parsing vertex: %v, %v, %v\n", errX, errY, errZ)
			}
			positions = append(positions, vec3.New(float32(x), float32(y), float32(z)))
		case "vt":
			if len(tokens) < 3 {
				log.Fatalf("Invalid texture coordinate line: %s\n", line)
			}
			x, errX := strconv.ParseFloat(tokens[1], 32)
			y, errY := strconv.ParseFloat(tokens[2], 32)
			if errX != nil || errY != nil {
				log.Fatalf("Error parsing UV: %v, %v\n", errX, errY)
			}
			uvs = append(uvs, vec2.New(float32(x), float32(y)))
		case "vn":
			if len(tokens) < 4 {
				log.Fatalf("Invalid normal line: %s\n", line)
			}
			x, errX := strconv.ParseFloat(tokens[1], 32)
			y, errY := strconv.ParseFloat(tokens[2], 32)
			z, errZ := strconv.ParseFloat(tokens[3], 32)
			if errX != nil || errY != nil || errZ != nil {
				log.Fatalf("Error parsing normal: %v, %v, %v\n", errX, errY, errZ)
			}
			normals = append(normals, vec3.New(float32(x), float32(y), float32(z)))
		case "f":
			if len(tokens) < 4 {
				log.Fatalf("Face with fewer than 3 vertices: %s\n", line)
			}

			numVertices := len(tokens) - 1
			for _, entry := range tokens[1:] {
				indicesSplit := strings.Split(entry, "/")
				vI, errV := strconv.Atoi(indicesSplit[0])
				vtI, errVT := strconv.Atoi(indicesSplit[1])
				var vnI int
				if len(indicesSplit) > 2 {
					vnI, _ = strconv.Atoi(indicesSplit[2])
				}
				if errV != nil || errVT != nil {
					log.Fatalf("Error parsing face indices: %v, %v\n", errV, errVT)
				}

				// Handle 1-based index
				vIndex := vI - 1
				vtIndex := vtI - 1
				var nIndex int
				if len(indicesSplit) > 2 {
					nIndex = vnI - 1
				}

				if vIndex >= len(positions) || vtIndex >= len(uvs) || (len(indicesSplit) > 2 && nIndex >= len(normals)) {
					log.Fatalf("Index out of bounds: v=%d, vt=%d, vn=%d\n", vIndex, vtIndex, nIndex)
				}

				p := positions[vIndex]
				uv := uvs[vtIndex]
				var n vec3.Type
				if len(indicesSplit) > 2 {
					n = normals[nIndex]
				}

				// Append vertex data (position, UV, and normal if available)
				vertexData = append(vertexData, p.X, p.Y, p.Z, uv.X, uv.Y)
				if len(indicesSplit) > 2 {
					vertexData = append(vertexData, n.X, n.Y, n.Z)
				}
			}

			startIndex := uint32(indexOffset)
			for i := 0; i <= numVertices-3; i++ {
				indices = append(indices, startIndex, startIndex+uint32(1+i), startIndex+uint32(2+i))
			}

			indexOffset += numVertices
		}
	}

	// Add the last mesh
	if activeMeshName != "__no_active" {
		return Obj{
			Vertices: positions,
			UVs:      uvs,
			Normals:  normals,
			Indices:  indices,
		}
	}

	log.Fatalf("No mesh found in the file\n")
	return Obj{}
}
