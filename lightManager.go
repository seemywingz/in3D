package in3D

import (
	"errors"
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// LightManager :
type LightManager struct {
	Lights  []*Light
	Program uint32
}

// Light : struct to hold light data
type Light struct {
	Radius   float32
	Ambient  [3]float32
	Difffuse [3]float32
	Specular [3]float32
	LRadID   int32
	LPosID   int32
	AmbID    int32
	DifID    int32
	SpecID   int32
	SceneData
	Draw        bool
	DrawnObject *DrawnObject
}

// NewLightManager :
func NewLightManager() *LightManager {
	lightManager = &LightManager{
		[]*Light{},
		Shader["normalMap"],
	}
	return lightManager
}

// NewLight :
func NewLight() *Light {
	return BuildLight(
		NewPosition(0, 0, 0), // position
		50,                   // radius
		[3]float32{0.2, 0.2, 0.2}, // ambiant intensity
		[3]float32{1, 1, 1},       // diffuse intensity
		[3]float32{1, 1, 1},       // specular intensity
		false,                     // draw
	)
}

// NewColorLight :
func NewColorLight(amb, dif, spec [3]float32) *Light {
	return BuildLight(
		NewPosition(0, 0, 0), // position
		50,                   // radius
		amb,                  // ambiant intensity
		dif,                  // diffuse intensity
		spec,                 // specular intensity
		false,                // draw
	)
}

// BuildLight :
func BuildLight(position Position, radius float32, amb, dif, spec [3]float32, draw bool) *Light {
	n := len(lightManager.Lights)
	if n > MaxLights {
		EoE("Error adding New Light:", errors.New("Max lights reached "+string(MaxLights)))
	}
	fmt.Println("Adding Light:", n)
	uniform := fmt.Sprintf("Light[%v]", n)
	LRadID := gl.GetUniformLocation(lightManager.Program, gl.Str(uniform+".lightRad\x00"))
	LPosID := gl.GetUniformLocation(lightManager.Program, gl.Str(uniform+".lightPos\x00"))
	ambID := gl.GetUniformLocation(lightManager.Program, gl.Str(uniform+".Iamb\x00"))
	difID := gl.GetUniformLocation(lightManager.Program, gl.Str(uniform+".Idif\x00"))
	specID := gl.GetUniformLocation(lightManager.Program, gl.Str(uniform+".Ispec\x00"))

	drawnObject := NewPointsObject(position, Cube, NoTexture, dif, Shader["color"])

	light := &Light{
		radius,
		amb,
		dif,
		spec,
		LRadID,
		LPosID,
		ambID,
		difID,
		specID,
		SceneData{},
		draw,
		drawnObject,
	}
	light.Position = position
	lightManager.Lights = append(lightManager.Lights, light)
	return light
}

// Update :
func (manager *LightManager) Update() {

	for _, light := range manager.Lights {
		if light.SceneLogic != nil {
			light.SceneLogic(&light.SceneData)
		}

		if light.Draw {
			light.DrawnObject.Position = light.Position
			for _, mg := range light.DrawnObject.Mesh.MaterialGroups {
				mg.Material.Diffuse = light.Difffuse
			}
			light.DrawnObject.Draw()
		}
		for _, p := range Shader {
			gl.UseProgram(p)
			gl.Uniform1f(light.LRadID, light.Radius)
			gl.Uniform3fv(light.LPosID, 1, &[]float32{light.X, light.Y, light.Z}[0])
			gl.Uniform3fv(light.AmbID, 1, &light.Ambient[0])
			gl.Uniform3fv(light.DifID, 1, &light.Difffuse[0])
			gl.Uniform3fv(light.SpecID, 1, &light.Specular[0])
		}
	}
}
