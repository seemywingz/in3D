package in3d

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
