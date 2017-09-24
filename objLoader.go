package gg

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// Mesh contains vertexes, normals and faces
type Mesh struct {
	Vertexes []float32
	Normals  []float32
	Faces    []*Face
	VAO      []float32
}

// Face :
type Face struct {
	Vertex []float32
	Normal []float32
}

// Material represents a material
type Material struct {
	Name      string
	Ambient   []float32
	Diffuse   []float32
	Specular  []float32
	Shininess float32
}

// LoadObject : opens a wavefront file and parses it into a map of objects
func LoadObject(filename string) *Mesh {
	file, err := os.Open(filename)
	EoE("Error Opening File", err)
	defer file.Close()

	var (
		vertexs []float32
		normals []float32
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
			for i := 1; i < 4; i++ {
				f, err := strconv.ParseFloat(fields[i], 32)
				EoE("Failed to parse float", err)
				vertexs = append(vertexs, float32(f))
			}
		case "vn":
			if len(fields) != 4 {
				EoE("unsupported vertex normal line", errors.New(filename+" "+line))
			}
			for i := 1; i < 4; i++ {
				f, err := strconv.ParseFloat(fields[i], 32)
				EoE("cannot parse float", err)
				normals = append(normals, float32(f))
			}
		case "f":
			if len(fields) != 4 {
				EoE("unsupported face: "+line, errors.New(filename))
			}
			var vertex []float32
			var normal []float32
			for i := 1; i < 4; i++ {
				faceStr := strings.Split(fields[i], "/")
				if len(faceStr) == 3 {
					vi, err := strconv.Atoi(faceStr[0])
					EoE("unsupported face vertex index", err)
					vertex = append(vertex, vertexs[vi])
					ni, err := strconv.Atoi(faceStr[2])
					EoE("unsupported face normal index", err)
					normal = append(normal, normals[ni])
				} else {
					EoE("Error Parsing Face (expected triangle)", errors.New(filename))
				}
			}
			face := &Face{vertex, normal}
			faces = append(faces, face)
		}
	}

	for _, f := range faces {
		for _, v := range f.Vertex {
			vao = append(vao, v)
		}
		for i := 0; i < 2; i++ {
			vao = append(vao, 0)
		}
		for _, n := range f.Normal {
			vao = append(vao, n)
		}
	}

	return &Mesh{vertexs, normals, faces, vao}
}

// func readMaterials(filename string) (map[string]*Material, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return nil, fmt.Errorf("cannot read referenced material library: %v", err)
// 	}
// 	defer file.Close()
//
// 	var (
// 		materials = make(map[string]*Material)
// 		material  *Material
// 	)
//
// 	lno := 0
// 	line := ""
// 	scanner := bufio.NewScanner(file)
//
// 	fail := func(msg string) error {
// 		return fmt.Errorf(msg+" at %s:%d: %s", filename, lno, line)
// 	}
//
// 	for scanner.Scan() {
// 		lno++
// 		line = scanner.Text()
// 		if strings.HasPrefix(line, "#") {
// 			continue
// 		}
//
// 		fields := strings.Fields(line)
// 		if len(fields) == 0 {
// 			continue
// 		}
//
// 		if fields[0] == "newmtl" {
// 			if len(fields) != 2 {
// 				return nil, fail("unsupported material definition")
// 			}
//
// 			material = &Material{Name: fields[1]}
// 			material.Ambient = []float32{0.2, 0.2, 0.2, 1.0}
// 			material.Diffuse = []float32{0.8, 0.8, 0.8, 1.0}
// 			material.Specular = []float32{0.0, 0.0, 0.0, 1.0}
// 			materials[material.Name] = material
//
// 			continue
// 		}
//
// 		if material == nil {
// 			return nil, fail("found data before material")
// 		}
//
// 		switch fields[0] {
// 		case "Ka":
// 			if len(fields) != 4 {
// 				return nil, fail("unsupported ambient color line")
// 			}
// 			for i := 0; i < 3; i++ {
// 				f, err := strconv.ParseFloat(fields[i+1], 32)
// 				if err != nil {
// 					return nil, fail("cannot parse float")
// 				}
// 				material.Ambient[i] = float32(f)
// 			}
// 		case "Kd":
// 			if len(fields) != 4 {
// 				return nil, fail("unsupported diffuse color line")
// 			}
// 			for i := 0; i < 3; i++ {
// 				f, err := strconv.ParseFloat(fields[i+1], 32)
// 				if err != nil {
// 					return nil, fail("cannot parse float")
// 				}
// 				material.Diffuse[i] = float32(f)
// 			}
// 		case "Ks":
// 			if len(fields) != 4 {
// 				return nil, fail("unsupported specular color line")
// 			}
// 			for i := 0; i < 3; i++ {
// 				f, err := strconv.ParseFloat(fields[i+1], 32)
// 				if err != nil {
// 					return nil, fail("cannot parse float")
// 				}
// 				material.Specular[i] = float32(f)
// 			}
// 		case "Ns":
// 			if len(fields) != 2 {
// 				return nil, fail("unsupported shininess line")
// 			}
// 			f, err := strconv.ParseFloat(fields[1], 32)
// 			if err != nil {
// 				return nil, fail("cannot parse float")
// 			}
// 			material.Shininess = float32(f / 1000 * 128)
// 		case "d":
// 			if len(fields) != 2 {
// 				return nil, fail("unsupported transparency line")
// 			}
// 			f, err := strconv.ParseFloat(fields[1], 32)
// 			if err != nil {
// 				return nil, fail("cannot parse float")
// 			}
// 			material.Ambient[3] = float32(f)
// 			material.Diffuse[3] = float32(f)
// 			material.Specular[3] = float32(f)
// 		}
// 	}
//
// 	if err := scanner.Err(); err != nil {
// 		return nil, err
// 	}
//
// 		for i := 0; i < 3; i++ {
// 			material.Diffuse[i] *= 1.3
// 			if material.Diffuse[i] > 1 {
// 				material.Diffuse[i] = 1
// 			}
// 		}
// 	}
//
// 	return materials, nil
// }
