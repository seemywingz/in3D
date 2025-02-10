package in3d

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/seemywingz/go-toolbox"
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
	VertID []int
	UVID   []int
	NormID []int
}

func buildVAOforMatGroup(group *MaterialGroup, vertexs, uvs, normals [][]float32, program uint32) {
	var (
		vao []float32
	)

	for _, f := range group.Faces { // use face data to construct GL VAO: XYZ UV [3]normal [3]tangent
		// This is UGLY!!
		vec0 := mgl32.NewVecNFromData(vertexs[f.VertID[0]-1])
		vec1 := mgl32.NewVecNFromData(vertexs[f.VertID[1]-1])
		vec2 := mgl32.NewVecNFromData(vertexs[f.VertID[2]-1])

		normal0 := normals[f.NormID[0]-1]
		normal1 := normals[f.NormID[1]-1]
		normal2 := normals[f.NormID[2]-1]

		uv0 := mgl32.NewVecNFromData([]float32{0, 0})
		uv1 := mgl32.NewVecNFromData([]float32{0, 0})
		uv2 := mgl32.NewVecNFromData([]float32{0, 0})

		tangent := mgl32.NewVecNFromData([]float32{0, 0, 0})

		if f.UVID[0] >= 0 {
			// if we have UV mappings, calculate tangentent and bitangent for normal map
			uv0 = mgl32.NewVecNFromData(uvs[f.UVID[0]-1])
			uv1 = mgl32.NewVecNFromData(uvs[f.UVID[1]-1])
			uv2 = mgl32.NewVecNFromData(uvs[f.UVID[2]-1])

			e1 := vec1.Sub(nil, vec0)
			e2 := vec2.Sub(nil, vec0)

			dUV1 := uv1.Sub(nil, uv0)
			dUV2 := uv2.Sub(nil, uv0)
			x, y, z := 0, 1, 2
			f := 1.0 / (dUV1.Get(x)*dUV2.Get(y) - dUV2.Get(x)*dUV1.Get(y))
			// print(f)

			tangent.Set(x, f*(dUV2.Get(y)*e1.Get(x)-dUV1.Get(y)*e2.Get(x)))
			tangent.Set(y, f*(dUV2.Get(y)*e1.Get(y)-dUV1.Get(y)*e2.Get(y)))
			tangent.Set(z, f*(dUV2.Get(y)*e1.Get(z)-dUV1.Get(y)*e2.Get(z)))
			tangent = tangent.Normalize(nil)
			// println(tangent)
		}

		// This is UGLY!!
		vao = append(vao, vec0.Raw()...)
		vao = append(vao, uv0.Raw()...)
		vao = append(vao, normal0...)
		vao = append(vao, tangent.Raw()...)

		vao = append(vao, vec1.Raw()...)
		vao = append(vao, uv1.Raw()...)
		vao = append(vao, normal1...)
		vao = append(vao, tangent.Raw()...)

		vao = append(vao, vec2.Raw()...)
		vao = append(vao, uv2.Raw()...)
		vao = append(vao, normal2...)
		vao = append(vao, tangent.Raw()...)
	}

	group.VAO = MakeVAO(vao, program)
	group.VertCount = int32(len(vao))
}

// LoadObject : opens a wavefront file and parses it into Material Groups
// TODO: Fix  UV coords, they are upside down...
func LoadObject(filename string, program uint32) *Mesh {
	file, ferr := os.Open(filename)
	toolbox.EoE(ferr, "Error Opening File")
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
				toolbox.EoE(errors.New(filename), "Error Parsing Vertex too few feilds ")
			}
			var v []float32
			for i := 1; i < 4; i++ {
				f, err := strconv.ParseFloat(fields[i], 32)
				toolbox.EoE(err, "Failed to parse float")
				v = append(v, float32(f))
			}
			vertexs = append(vertexs, v)
		case "vt":
			if len(fields) != 3 {
				toolbox.EoE(errors.New(filename), "Error Parsing UV coords")
			}
			var uv []float32
			for i := 1; i < 3; i++ {
				f, err := strconv.ParseFloat(fields[i], 32)
				toolbox.EoE(err, "Failed to parse float")
				uv = append(uv, float32(f))
			}
			uvs = append(uvs, uv)
		case "vn":
			if len(fields) != 4 {
				toolbox.EoE(errors.New(filename+" "+line), "unsupported vertex normal line")
			}
			var n []float32
			for i := 1; i < 4; i++ {
				f, err := strconv.ParseFloat(fields[i], 32)
				toolbox.EoE(err, "cannot parse float")
				n = append(n, float32(f))
			}
			normals = append(normals, n)
		case "f":
			if len(fields) != 4 {
				toolbox.EoE(errors.New(filename), "unsupported face:"+string(rune(len(fields)))+" "+line)
			}
			var (
				vi, ui, ni []int
			)
			for i := 1; i < 4; i++ {
				faceStr := strings.Split(fields[i], "/")
				svi, err := strconv.Atoi(faceStr[0])
				vi = append(vi, svi)
				toolbox.EoE(err, "unsupported face vertex index")
				sni, err := strconv.Atoi(faceStr[2])
				ni = append(ni, sni)
				toolbox.EoE(err, "unsupported face normal index")
				if faceStr[1] == "" {
					// set negative value as placeholder for .obj with no UV mapping
					faceStr[1] = "-1"
				}
				sui, err := strconv.Atoi(faceStr[1])
				ui = append(ui, sui)
				toolbox.EoE(err, "unsupported face uv index")
			}
			materialGroups[currentGroup].Faces = append(materialGroups[currentGroup].Faces, &Face{vi, ui, ni})
		}
	}

	for _, g := range materialGroups {
		buildVAOforMatGroup(g, vertexs, uvs, normals, program)
	}

	return &Mesh{materialGroups}
}

// LoadMaterials : create material groups from wavefront data
func LoadMaterials(filename string) map[string]*MaterialGroup {

	file, ferr := os.Open(filename)
	toolbox.EoE(ferr, "Error Opening Material File")
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
				toolbox.EoE(errors.New(filename), "unsupported material definition")
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
				toolbox.EoE(errors.New(filename), "unsupported ambient color line")
			}
			for i := 0; i < 3; i++ {
				f, err := strconv.ParseFloat(fields[i+1], 32)
				toolbox.EoE(err, "Error parsing float")
				materialGroups[currentMat].Material.Ambient[i] = float32(f)
			}
		case "Kd":
			if len(fields) != 4 {
				toolbox.EoE(errors.New(filename), "Error Diffuse Parse")
			}
			for i := 0; i < 3; i++ {
				f, err := strconv.ParseFloat(fields[i+1], 32)
				toolbox.EoE(err, "Error parsing float")
				materialGroups[currentMat].Material.Diffuse[i] = float32(f)
			}
		case "Ks":
			if len(fields) != 4 {
				toolbox.EoE(errors.New(filename), "Error KS Parse")
			}
			for i := 0; i < 3; i++ {
				f, err := strconv.ParseFloat(fields[i+1], 32)
				toolbox.EoE(err, "Error parsing float")
				materialGroups[currentMat].Material.Specular[i] = float32(f)
			}
		case "Ns":
			if len(fields) != 2 {
				toolbox.EoE(errors.New(filename), "Error NS Parse")
			}
			f, err := strconv.ParseFloat(fields[1], 32)
			toolbox.EoE(err, "Error parsing float")
			materialGroups[currentMat].Material.Shininess = float32(f / 1000 * 128)
		// case "d":
		// 	if len(fields) != 2 {
		// 	    toolbox.EoE("Error d Parse", errors.New(filename))
		// 	}
		// 	f, err := strconv.ParseFloat(fields[1], 32)
		//     toolbox.EoE("Error parsing float", err)
		// 	materialGroups[currentMat].Material.Shininess = float32(f)
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

	toolbox.EoE(scanner.Err(), "Scann Error")
	return materialGroups
}
