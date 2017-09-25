package gg

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// Mesh :
type Mesh struct {
	MaterialGroups []*MaterialGroup
}

// MaterialGroup :
type MaterialGroup struct {
	Material  *Material
	VAO       uint32
	VertCount int32
}

// Material represents a material
type Material struct {
	Name      string
	Ambient   []float32
	Diffuse   []float32
	Specular  []float32
	Shininess float32
}

// ObjData contains vertexes, normals, faces and GL approved VAO
type ObjData struct {
	Vertexes       [][]float32
	UVs            [][]float32
	Normals        [][]float32
	Faces          []*Face
	MaterialGroups []*MaterialGroup
}

// Face :
type Face struct {
	VertIdx int
	UVIdx   int
	NormIdx int
}

func appendToVAO(vao []float32, vec []float32) []float32 {
	for _, v := range vec {
		vao = append(vao, v)
	}
	return vao
}

var defaultMaterial Material

// LoadObject : opens a wavefront file and parses it into ObjData
func LoadObject(filename string) *Mesh {
	file, ferr := os.Open(filename)
	EoE("Error Opening File", ferr)
	defer file.Close()

	var (
		vertexs        [][]float32
		uvs            [][]float32
		normals        [][]float32
		faces          []*Face
		vao            []float32
		materialGroups []*MaterialGroup
	)

	defaultMaterial = Material{
		"default",
		[]float32{0.1, 0.1, 0.1},
		[]float32{1, 1, 1},
		[]float32{0.8, 0.8, 0.8},
		1,
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, " ") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		switch fields[0] {
		case "v":
			if len(fields) != 4 {
				EoE("Error Parsing Vertex too few feilds ", errors.New(filename))
			}
			var v []float32
			for i := 1; i < 4; i++ {
				f, err := strconv.ParseFloat(fields[i], 32)
				EoE("Failed to parse float", err)
				v = append(v, float32(f))
			}
			vertexs = append(vertexs, v)
		case "vt":
			if len(fields) != 3 {
				EoE("Error Parsing UV coords", errors.New(filename))
			}
			var uv []float32
			for i := 1; i < 3; i++ {
				f, err := strconv.ParseFloat(fields[i], 32)
				EoE("Failed to parse float", err)
				uv = append(uv, float32(f))
			}
			uvs = append(uvs, uv)
		case "vn":
			if len(fields) != 4 {
				EoE("unsupported vertex normal line", errors.New(filename+" "+line))
			}
			var n []float32
			for i := 1; i < 4; i++ {
				f, err := strconv.ParseFloat(fields[i], 32)
				EoE("cannot parse float", err)
				n = append(n, float32(f))
			}
			normals = append(normals, n)
		case "f":
			if len(fields) != 4 {
				EoE("unsupported face:"+string(len(fields))+" "+line, errors.New(filename))
			}
			var (
				vi, ui, ni int
				err        error
			)
			for i := 1; i < 4; i++ {
				faceStr := strings.Split(fields[i], "/")
				vi, err = strconv.Atoi(faceStr[0])
				EoE("unsupported face vertex index", err)
				ni, err = strconv.Atoi(faceStr[2])
				EoE("unsupported face normal index", err)
				if len(faceStr) == 4 {
					ui, err = strconv.Atoi(faceStr[1])
					EoE("unsupported face uv index", err)
				} else {
					ui = 1
				}
				faces = append(faces, &Face{vi, ui, ni})
			}
		}
	}

	for _, f := range faces { // use face data to construct GL VAO XYZUVNXNYNZ
		vao = appendToVAO(vao, vertexs[f.VertIdx-1])
		if len(uvs) != 0 {
			vao = appendToVAO(vao, uvs[f.UVIdx-1])
		} else {
			vao = appendToVAO(vao, []float32{0, 0})
		}
		vao = appendToVAO(vao, normals[f.NormIdx-1])
		// TODO: parse material from mtllib *.mtl
	}
	return &Mesh{materialGroups}
}
