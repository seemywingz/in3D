package in3d

import (
	"encoding/binary"
	"fmt"
	"math"
	"path/filepath"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/qmuntal/gltf"
)

// LoadGLTF loads a GLTF model and returns a Mesh
func LoadGLTF(modelName string, program uint32) *Mesh {
	filePath := filepath.Join(modelName, "scene.gltf")
	doc, err := gltf.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening GLTF file: %s, %v\n", filePath, err)
		return nil
	}

	var vertices, normals, uvs []float32
	var indices []uint32
	materialGroups := make(map[string]*MaterialGroup)

	fmt.Println("Parsing GLTF:", modelName)

	for _, mesh := range doc.Meshes {
		fmt.Println("Processing mesh:", mesh.Name)

		for _, primitive := range mesh.Primitives {
			materialName := "default"
			if primitive.Material != nil {
				materialName = doc.Materials[*primitive.Material].Name
			}

			// Ensure MaterialGroup Exists
			if _, exists := materialGroups[materialName]; !exists {
				materialGroups[materialName] = &MaterialGroup{
					Material: &Material{
						Name:       materialName,
						Ambient:    []float32{0.1, 0.1, 0.1},
						Diffuse:    []float32{1, 1, 1},
						Specular:   []float32{1, 1, 1},
						Shininess:  1.0,
						DiffuseTex: NoTexture,
						NormalTex:  NoTexture,
					},
					Faces: []*Face{},
				}
			}

			// Extract Data
			if accessor := doc.Accessors[primitive.Attributes["POSITION"]]; accessor != nil {
				vertices = append(vertices, parseBuffer(doc, accessor, 3)...)
			}
			if accessor := doc.Accessors[primitive.Attributes["NORMAL"]]; accessor != nil {
				normals = append(normals, parseBuffer(doc, accessor, 3)...)
			}
			if accessor := doc.Accessors[primitive.Attributes["TEXCOORD_0"]]; accessor != nil {
				uvs = append(uvs, parseBuffer(doc, accessor, 2)...)
			}
			if primitive.Indices != nil {
				if accessor := doc.Accessors[*primitive.Indices]; accessor != nil {
					indices = append(indices, parseIndexBuffer(doc, accessor)...)
				}
			}
		}
	}

	// Ensure at least one material group is created
	if len(materialGroups) == 0 {
		fmt.Println("Error: No MaterialGroups created from GLTF")
		return nil
	}

	// Build VAO
	mesh := &Mesh{materialGroups}
	buildVAOforGLTF(mesh, vertices, uvs, normals, indices, program)

	fmt.Printf("Total vertices: %d\n", len(vertices)/3)
	fmt.Printf("Total normals: %d\n", len(normals)/3)
	fmt.Printf("Total UVs: %d\n", len(uvs)/2)
	fmt.Printf("Total indices: %d\n", len(indices))

	return mesh
}

func parseBuffer(doc *gltf.Document, accessor *gltf.Accessor, components int) []float32 {
	if accessor == nil || accessor.BufferView == nil {
		fmt.Println("Error: Nil accessor or buffer view")
		return nil
	}

	bufferView := doc.BufferViews[*accessor.BufferView]
	buffer := doc.Buffers[bufferView.Buffer]
	data := buffer.Data[bufferView.ByteOffset : bufferView.ByteOffset+bufferView.ByteLength]

	var result []float32
	for i := 0; i < len(data); i += 4 * components {
		if i+4*components > len(data) {
			fmt.Println("Warning: Buffer overflow detected while parsing buffer view")
			break
		}
		for j := 0; j < components; j++ {
			result = append(result, math.Float32frombits(binary.LittleEndian.Uint32(data[i+j*4:])))
		}
	}
	return result
}

func parseIndexBuffer(doc *gltf.Document, accessor *gltf.Accessor) []uint32 {
	if accessor == nil || accessor.BufferView == nil {
		fmt.Println("Error: Nil accessor or buffer view")
		return nil
	}

	bufferView := doc.BufferViews[*accessor.BufferView]
	buffer := doc.Buffers[bufferView.Buffer]
	data := buffer.Data[bufferView.ByteOffset : bufferView.ByteOffset+bufferView.ByteLength]

	var indices []uint32
	for i := 0; i < len(data); i += 2 {
		indices = append(indices, uint32(binary.LittleEndian.Uint16(data[i:])))
	}
	return indices
}

func buildVAOforGLTF(mesh *Mesh, vertices, uvs, normals []float32, indices []uint32, program uint32) {
	if len(vertices) == 0 {
		fmt.Println("Error: No vertex data found in GLTF mesh.")
		return
	}

	if len(normals) == 0 {
		fmt.Println("Warning: No normal data found in GLTF mesh, using default normals.")
		normals = make([]float32, len(vertices))
	}

	if len(uvs) == 0 {
		fmt.Println("Warning: No UV data found in GLTF mesh, using default UVs.")
		uvs = make([]float32, len(vertices)/3*2)
	}

	// Create VAO
	var vao, vbo, ebo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)

	gl.BindVertexArray(vao)

	// Store vertex data in a single buffer
	interleavedData := interleaveVertexData(vertices, uvs, normals)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(interleavedData)*4, gl.Ptr(interleavedData), gl.STATIC_DRAW)

	// Store index data
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	// Define vertex attributes
	stride := int32(8 * 4) // 3 position, 2 UV, 3 normal (floats)

	// Position
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// UV
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, stride, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	// Normal
	gl.VertexAttribPointer(2, 3, gl.FLOAT, false, stride, gl.PtrOffset(5*4))
	gl.EnableVertexAttribArray(2)

	gl.BindVertexArray(0)

	// Store in material groups
	for _, group := range mesh.MaterialGroups {
		group.VAO = vao
		group.VertCount = int32(len(indices))
	}
}

func interleaveVertexData(vertices, uvs, normals []float32) []float32 {
	var interleaved []float32
	for i := 0; i < len(vertices)/3; i++ {
		interleaved = append(interleaved, vertices[i*3:i*3+3]...)
		interleaved = append(interleaved, uvs[i*2:i*2+2]...)
		interleaved = append(interleaved, normals[i*3:i*3+3]...)
	}
	return interleaved
}
