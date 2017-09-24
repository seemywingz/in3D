package gg

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// Mesh contains vertexes, normals, faces and GL approved VAO
type Mesh struct {
	Vertexes []*Vec3
	Normals  []*Vec3
	Faces    []*Face
	VAO      []float32
}

// Vec3 :
type Vec3 struct {
	// TODO: ensure [3]float32... add append func
	Coords []float32
}

// Face :
type Face struct {
	VertIdx int
	NormIdx int
}

// Material represents a material
type Material struct {
	Name      string
	Ambient   []float32
	Diffuse   []float32
	Specular  []float32
	Shininess float32
}

func appendToVAO(vao []float32, vec *Vec3) []float32 {
	for _, v := range vec.Coords {
		vao = append(vao, v)
	}
	return vao
}

// LoadObject : opens a wavefront file and parses it into mesh
func LoadObject(filename string) *Mesh {
	file, err := os.Open(filename)
	EoE("Error Opening File", err)
	defer file.Close()

	var (
		vertexs []*Vec3
		normals []*Vec3
		faces   []*Face
		vao     []float32
	)

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
			vertexs = append(vertexs, &Vec3{v})
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
			normals = append(normals, &Vec3{n})
		case "f":
			if len(fields) != 4 {
				EoE("unsupported face: "+line, errors.New(filename))
			}
			for i := 1; i < 4; i++ {
				faceStr := strings.Split(fields[i], "/")
				if len(faceStr) == 3 {
					vi, err := strconv.Atoi(faceStr[0])
					EoE("unsupported face vertex index", err)
					ni, err := strconv.Atoi(faceStr[2])
					EoE("unsupported face normal index", err)
					faces = append(faces, &Face{vi, ni})
				} else {
					EoE("Error Parsing Face (expected triangle)", errors.New(filename))
				}
			}
		}
	}

	for _, f := range faces { // use face data to construct GL VAO XYZUVNXNYNZ
		vao = appendToVAO(vao, vertexs[f.VertIdx-1])
		// TODO: parse material from mtllib *.mtl
		vao = append(vao, 0)
		vao = append(vao, 0)
		vao = appendToVAO(vao, normals[f.NormIdx-1])
	}
	return &Mesh{vertexs, normals, faces, vao}
}
