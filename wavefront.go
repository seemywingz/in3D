package gg

// import (
// 	"bufio"
// 	"errors"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// 	"strings"
// )
//
// //  Mesh : contains vertexes, normals and the material
// type Mesh struct {
// 	Vertexes []float32
// 	Normals  []float32
// 	Material *Material
// }
//
// // Material represents a material
// type Material struct {
// 	Name      string
// 	Ambient   []float32
// 	Diffuse   []float32
// 	Specular  []float32
// 	Shininess float32
// }
//
// // Read opens a wavefront file and parses it into a map of objects
// func Read(filename string) *Mesh {
// 	file, err := os.Open(filename)
// 	EoE("Failed to Open Fiile", err)
// 	defer file.Close()
//
// 	var (
// 		materials map[string]*Material
// 		// objects   = make(map[string]*Mesh)
// 		// object    *Mesh
// 		mesh   *Mesh
// 		vertex []float32
// 		normal []float32
// 	)
//
// 	lno := 0
// 	line := ""
// 	scanner := bufio.NewScanner(file)
//
// 	fail := func(msg string) {
// 		EoE(fmt.Sprintf(msg+" at %s:%d: %s", filename, lno, line), errors.New(""))
// 	}
//
// 	mesh = &Mesh{}
// 	for scanner.Scan() {
// 		lno++
// 		line = scanner.Text()
//
// 		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, " ") {
// 			continue
// 		}
//
// 		fields := strings.Fields(line)
// 		if len(fields) == 0 {
// 			continue
// 		}
//
// 		switch fields[0] {
// 		case "mtllib":
// 			if len(fields) != 2 {
// 				fail("unsupported materials library line")
// 			}
// 			materials = readMaterials(filepath.Join(filepath.Dir(filename), fields[1]))
// 			EoE("Error Loading MTL File", err)
// 			continue
// 			// case "o":
// 			// 	if len(fields) != 2 {
// 			// 		EoE("unsupported object line", errors.New(""))
// 			// 	}
// 			// 	object = &Mesh{Name: fields[1]}
// 			// 	objects[object.Name] = object
// 			// 	mesh = nil
// 			// 	continue
// 		}
//
// 		switch fields[0] {
// 		case "usemtl":
// 			if len(fields) != 2 {
// 				fail("unsupported material usage line")
// 			}
// 			mesh.Material = materials[fields[1]]
// 			if mesh.Material == nil {
// 				EoE(fmt.Sprintf("material %q not defined", fields[1]), errors.New(""))
// 			}
// 		case "v":
// 			if len(fields) != 4 {
// 				fail("unsupported vertex line")
// 			}
// 			for i := 0; i < 3; i++ {
// 				f, err := strconv.ParseFloat(fields[i+1], 32)
// 				if err != nil {
// 					fail("cannot parse float")
// 				}
// 				vertex = append(vertex, float32(f))
// 			}
// 		case "vn":
// 			if len(fields) != 4 {
// 				fail("unsupported vertex normal line")
// 			}
// 			for i := 0; i < 3; i++ {
// 				f, err := strconv.ParseFloat(fields[i+1], 32)
// 				if err != nil {
// 					fail("cannot parse float")
// 				}
// 				normal = append(normal, float32(f))
// 			}
// 		case "f":
// 			if len(fields) != 4 {
// 				fail("unsupported face line")
// 			}
// 			for i := 0; i < 3; i++ {
// 				face := strings.Split(fields[i+1], "/")
// 				if len(face) != 3 {
// 					fail("unsupported face shape (not a triangle)")
// 				}
// 				vi, err := strconv.Atoi(face[0])
// 				if err != nil {
// 					fail("unsupported face vertex index")
// 				}
// 				ni, err := strconv.Atoi(face[2])
// 				if err != nil {
// 					fail("unsupported face normal index")
// 				}
// 				vi = (vi - 1) * 3
// 				ni = (ni - 1) * 3
// 				mesh.Vertexes = append(mesh.Vertexes, vertex[vi], vertex[vi+1], vertex[vi+2])
// 				mesh.Normals = append(mesh.Normals, normal[ni], normal[ni+1], normal[ni+2])
// 			}
// 		}
// 	}
// 	EoE("Scanner Error", scanner.Err())
//
// 	return mesh
// }
//
// func readMaterials(filename string) map[string]*Material {
// 	file, err := os.Open(filename)
// 	EoE("cannot read referenced material library:", err)
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
// 				fail("unsupported material definition")
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
// 			fail("found data before material")
// 		}
//
// 		switch fields[0] {
// 		case "Ka":
// 			if len(fields) != 4 {
// 				fail("unsupported ambient color line")
// 			}
// 			for i := 0; i < 3; i++ {
// 				f, err := strconv.ParseFloat(fields[i+1], 32)
// 				if err != nil {
// 					fail("cannot parse float")
// 				}
// 				material.Ambient[i] = float32(f)
// 			}
// 		case "Kd":
// 			if len(fields) != 4 {
// 				fail("unsupported diffuse color line")
// 			}
// 			for i := 0; i < 3; i++ {
// 				f, err := strconv.ParseFloat(fields[i+1], 32)
// 				if err != nil {
// 					fail("cannot parse float")
// 				}
// 				material.Diffuse[i] = float32(f)
// 			}
// 		case "Ks":
// 			if len(fields) != 4 {
// 				fail("unsupported specular color line")
// 			}
// 			for i := 0; i < 3; i++ {
// 				f, err := strconv.ParseFloat(fields[i+1], 32)
// 				if err != nil {
// 					fail("cannot parse float")
// 				}
// 				material.Specular[i] = float32(f)
// 			}
// 		case "Ns":
// 			if len(fields) != 2 {
// 				fail("unsupported shininess line")
// 			}
// 			f, err := strconv.ParseFloat(fields[1], 32)
// 			if err != nil {
// 				fail("cannot parse float")
// 			}
// 			material.Shininess = float32(f / 1000 * 128)
// 		case "d":
// 			if len(fields) != 2 {
// 				fail("unsupported transparency line")
// 			}
// 			f, err := strconv.ParseFloat(fields[1], 32)
// 			if err != nil {
// 				fail("cannot parse float")
// 			}
// 			material.Ambient[3] = float32(f)
// 			material.Diffuse[3] = float32(f)
// 			material.Specular[3] = float32(f)
// 		}
// 	}
//
// 	EoE("Scaner Error", scanner.Err())
//
// 	// Exporting from blender seems to show everything too dark in
// 	// practice, so hack colors to look closer to what we see there.
// 	// TODO This needs more real world checking.
// 	for _, material := range materials {
// 		if material.Ambient[0] == 0 &&
// 			material.Ambient[1] == 0 &&
// 			material.Ambient[2] == 0 &&
// 			material.Ambient[3] == 1 {
// 			material.Ambient[0] = material.Diffuse[0] * 0.7
// 			material.Ambient[1] = material.Diffuse[1] * 0.7
// 			material.Ambient[2] = material.Diffuse[2] * 0.7
// 		}
//
// 		for i := 0; i < 3; i++ {
// 			material.Diffuse[i] *= 1.3
// 			if material.Diffuse[i] > 1 {
// 				material.Diffuse[i] = 1
// 			}
// 		}
// 	}
//
// 	return materials
// }
