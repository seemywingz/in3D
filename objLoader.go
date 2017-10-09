package in3D

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// Mesh :
type Mesh struct {
	MaterialGroups map[string]*MaterialGroup
}

// MaterialGroup :
type MaterialGroup struct {
	Material  *Material
	Faces     []*Face
	VAO       uint32
	VertCount int32
}

// Material represents a material
type Material struct {
	Name       string
	Ambient    []float32
	Diffuse    []float32
	Specular   []float32
	Shininess  float32
	DiffuseTex uint32
	NormalTex  uint32
}

// Face :
type Face struct {
	VertIdx int
	UVIdx   int
	NormIdx int
}

func buildVAOforMatGroup(group *MaterialGroup, vertexs, uvs, normals [][]float32, program uint32) {
	var (
		vao       []float32
		vec       []float32
		uv        []float32
		normal    []float32
		tangent   []float32
		bitangent []float32
	)

	for _, f := range group.Faces { // use face data to construct GL VAO XYZ UV [3]normal [3]tangent

		vec = vertexs[f.VertIdx-1]
		normal = normals[f.NormIdx-1]

		if f.UVIdx >= 0 {
			uv = uvs[f.UVIdx-1]
			tangent = []float32{0, 0, 0}
			bitangent = []float32{0, 0, 0}
		} else {
			uv = []float32{0, 0}
			tangent = []float32{0, 0, 0}
			bitangent = []float32{0, 0, 0}
		}

		vao = append(vao, vec...)
		vao = append(vao, uv...)
		vao = append(vao, normal...)
		vao = append(vao, tangent...)
		vao = append(vao, bitangent...)
	}

	group.VAO = MakeVAO(vao, program)
	group.VertCount = int32(len(vao))

	// p0 := mgl32.NewVecNFromData(vertexs[f.VertIdx-1])
	// i := f.VertIdx
	// if i == len(vertexs) {
	// 	i = 0
	// }
	// p1 := mgl32.NewVecNFromData(vertexs[i])

	// edge := p1.Sub(nil, p0)
	// fmt.Println(edge)
	// var uv1 []float32

	// i := f.VertIdx
	// if i == len(uvs) {
	// 	i = 0
	// }
	// uv1 = uvs[i]
}

// LoadObject : opens a wavefront file and parses it into Material Groups
// TODO: Fix  UV coords, they are upside down...
func LoadObject(filename string, program uint32) *Mesh {
	file, ferr := os.Open(filename)
	EoE("Error Opening File", ferr)
	defer file.Close()

	vertexs := [][]float32{}
	normals := [][]float32{}
	uvs := [][]float32{}

	var materialGroups map[string]*MaterialGroup

	currentGroup := "string"
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
		case "mtllib":
			materialGroups = LoadMaterials(fields[1])
		case "usemtl":
			currentGroup = fields[1]
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
				if faceStr[1] == "" {
					faceStr[1] = "-1"
				}
				ui, err = strconv.Atoi(faceStr[1])
				EoE("unsupported face uv index", err)
				materialGroups[currentGroup].Faces = append(materialGroups[currentGroup].Faces, &Face{vi, ui, ni})
			}
		}
	}

	for _, g := range materialGroups {
		buildVAOforMatGroup(g, vertexs, uvs, normals, program)
	}

	return &Mesh{materialGroups}
}

// LoadMaterials :
func LoadMaterials(filename string) map[string]*MaterialGroup {

	file, ferr := os.Open(filename)
	EoE("Error Opening Material File", ferr)
	defer file.Close()

	line := ""
	scanner := bufio.NewScanner(file)
	currentMat := ""
	materialGroups := make(map[string]*MaterialGroup)

	for scanner.Scan() {
		line = scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		if fields[0] == "newmtl" {
			if len(fields) != 2 {
				EoE("unsupported material definition", errors.New(filename))
			}
			currentMat = fields[1]
			material := &Material{
				currentMat,
				[]float32{0.1, 0.1, 0.1},
				[]float32{1, 1, 1},
				[]float32{0.8, 0.8, 0.8},
				1,
				NoTexture,
				NoTexture,
			}
			materialGroups[currentMat] = &MaterialGroup{}
			materialGroups[currentMat].Material = material

			continue
		}

		switch fields[0] {
		case "Ka":
			if len(fields) != 4 {
				EoE("unsupported ambient color line", errors.New(filename))
			}
			for i := 0; i < 3; i++ {
				f, err := strconv.ParseFloat(fields[i+1], 32)
				EoE("Error parsing float", err)
				materialGroups[currentMat].Material.Ambient[i] = float32(f)
			}
		case "Kd":
			if len(fields) != 4 {
				EoE("Error Diffuse Parse", errors.New(filename))
			}
			for i := 0; i < 3; i++ {
				f, err := strconv.ParseFloat(fields[i+1], 32)
				EoE("Error parsing float", err)
				materialGroups[currentMat].Material.Diffuse[i] = float32(f)
			}
		case "Ks":
			if len(fields) != 4 {
				EoE("Error KS Parse", errors.New(filename))
			}
			for i := 0; i < 3; i++ {
				f, err := strconv.ParseFloat(fields[i+1], 32)
				EoE("Error parsing float", err)
				materialGroups[currentMat].Material.Specular[i] = float32(f)
			}
		case "Ns":
			if len(fields) != 2 {
				EoE("Error NS Parse", errors.New(filename))
			}
			f, err := strconv.ParseFloat(fields[1], 32)
			EoE("Error parsing float", err)
			materialGroups[currentMat].Material.Shininess = float32(f / 1000 * 128)
		case "d":
			if len(fields) != 2 {
				EoE("Error d Parse", errors.New(filename))
			}
			f, err := strconv.ParseFloat(fields[1], 32)
			EoE("Error parsing float", err)
			materialGroups[currentMat].Material.Shininess = float32(f)
		case "map_Kd":
			DiffuseTexFile := fields[1]
			DiffuseTex := NewTexture(DiffuseTexFile)
			materialGroups[currentMat].Material.DiffuseTex = DiffuseTex
		case "map_Bump":
			NormalTexFile := fields[1]
			NormalTex := NewTexture(NormalTexFile)
			materialGroups[currentMat].Material.NormalTex = NormalTex
		}
	}

	EoE("Scann Error", scanner.Err())
	return materialGroups
}
