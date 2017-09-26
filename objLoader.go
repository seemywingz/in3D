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
	Name      string
	Ambient   []float32
	Diffuse   []float32
	Specular  []float32
	Shininess float32
}

// Face :
type Face struct {
	VertIdx int
	UVIdx   int
	NormIdx int
}

var defaultMaterial = Material{
	"default",
	[]float32{0.1, 0.1, 0.1},
	[]float32{1, 1, 1},
	[]float32{0.8, 0.8, 0.8},
	1,
}

func buildVAOforMatGroup(group *MaterialGroup, vertexs, uvs, normals [][]float32) {
	vao := []float32{}
	for _, f := range group.Faces { // use face data to construct GL VAO XYZUVNXNYNZ
		vao = append(vao, vertexs[f.VertIdx-1]...)
		if f.UVIdx >= 0 {
			vao = append(vao, uvs[f.UVIdx-1]...)
		} else {
			vao = append(vao, []float32{0, 0}...)
		}
		vao = append(vao, normals[f.NormIdx-1]...)
	}
	group.VAO = MakeVAO(vao, Shader["phong"])
	group.VertCount = int32(len(vao))
}

// LoadObject : opens a wavefront file and parses it into ObjData
func LoadObject(filename string) *Mesh {
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
			// fmt.Println(materialGroups)
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
		buildVAOforMatGroup(g, vertexs, uvs, normals)
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
		}
	}

	EoE("Scann Error", scanner.Err())
	return materialGroups
}
