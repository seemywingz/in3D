package gg

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// LightLogic :
type LightLogic func(l *Light)

// LightManager :
type LightManager struct {
	Lights  []*Light
	Program uint32
}

// Light : struct to hold light data
type Light struct {
	Radius  float32
	Iamb    *float32
	Idif    *float32
	Ispec   *float32
	LRadID  int32
	LPosID  int32
	IambID  int32
	IdifID  int32
	IspecID int32
	StdData
}

// NewLightManager :
func NewLightManager() *LightManager {

	lightManager = &LightManager{
		[]*Light{},
		Shader["multiLight"],
	}
	return lightManager
}

// NewLight :
func NewLight(pos Position, radius float32, iamb, idif, ispec []float32) *Light {
	// n := len(lightManager.Lights)
	// LightID := gl.GetUniformLocation(Shader["multiLight"], gl.Str("Light["+string(n)+"]\x00"))
	LRadID := gl.GetUniformLocation(lightManager.Program, gl.Str("Light[0].lightRad\x00"))
	LPosID := gl.GetUniformLocation(lightManager.Program, gl.Str("Light[0].lightPos\x00"))
	IambID := gl.GetUniformLocation(lightManager.Program, gl.Str("Light[0].Iamb\x00"))
	IdifID := gl.GetUniformLocation(lightManager.Program, gl.Str("Light[0].Idif\x00"))
	IspecID := gl.GetUniformLocation(lightManager.Program, gl.Str("Light[0].Ispec\x00"))

	l := &Light{
		radius,
		&iamb[0],
		&idif[0],
		&ispec[0],
		LRadID,
		LPosID,
		IambID,
		IdifID,
		IspecID,
		StdData{},
	}
	l.Position = pos
	lightManager.Lights = append(lightManager.Lights, l)
	return l
}

// Update :
func (l *LightManager) Update() {

	for _, light := range l.Lights {
		if light.SceneLogic != nil {
			light.SceneLogic(&light.StdData)
		}
		light.X++
		gl.UseProgram(l.Program)
		gl.Uniform1f(light.LRadID, light.Radius)
		gl.Uniform3fv(light.LPosID, 1, &[]float32{light.X, light.Y, light.Z}[0])
		gl.Uniform3fv(light.IambID, 1, light.Iamb)
		gl.Uniform3fv(light.IdifID, 1, light.Idif)
		gl.Uniform3fv(light.IspecID, 1, light.Ispec)
	}
}
