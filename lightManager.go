package in3d

import (
	"errors"
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// LightManager :
type LightManager struct {
	Lights []*Light
}

// Light : struct to hold light data
type Light struct {
	Radius   float32
	Ambient  []float32
	Difffuse []float32
	Specular []float32
	SceneData
	Draw        bool
	DrawnObject *DrawnObject
}

// NewLightManager :
func NewLightManager() *LightManager {
	lightManager = &LightManager{
		[]*Light{},
	}
	return lightManager
}

// NewLight :
func NewLight() *Light {
	return BuildLight(
		NewPosition(0, 0, 0),     // position
		50,                       // radius
		[]float32{0.2, 0.2, 0.2}, // ambiant intensity
		[]float32{1, 1, 1},       // diffuse intensity
		[]float32{1, 1, 1},       // specular intensity
		false,                    // draw
	)
}

// NewColorLight :
func NewColorLight(amb, dif, spec []float32) *Light {
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
func BuildLight(position Position, radius float32, amb, dif, spec []float32, draw bool) *Light {
	n := len(lightManager.Lights)
	if n > MaxLights {
		EoE("Error adding New Light:", errors.New("Max lights reached "+string(rune(MaxLights))))
	}
	fmt.Println("Adding Light:", n)

	drawnObject := NewPointsObject(position, Cube, NoTexture, dif, Shader["color"])
	drawnObject.Scale = 0.05

	light := &Light{
		radius,
		amb,
		dif,
		spec,
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

	for n, light := range manager.Lights {
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
		for _, program := range Shader {

			uniform := fmt.Sprintf("Light[%v]", n)
			LRadID := gl.GetUniformLocation(program, gl.Str(uniform+".lightRad\x00"))
			LPosID := gl.GetUniformLocation(program, gl.Str(uniform+".lightPos\x00"))
			AmbID := gl.GetUniformLocation(program, gl.Str(uniform+".Iamb\x00"))
			DifID := gl.GetUniformLocation(program, gl.Str(uniform+".Idif\x00"))
			SpecID := gl.GetUniformLocation(program, gl.Str(uniform+".Ispec\x00"))

			gl.UseProgram(program)
			gl.Uniform1f(LRadID, light.Radius)
			gl.Uniform3fv(LPosID, 1, &[]float32{light.X, light.Y, light.Z}[0])
			gl.Uniform3fv(AmbID, 1, &light.Ambient[0])
			gl.Uniform3fv(DifID, 1, &light.Difffuse[0])
			gl.Uniform3fv(SpecID, 1, &light.Specular[0])
		}
	}
}
