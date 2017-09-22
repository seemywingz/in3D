package gg

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// LightLogic :
type LightLogic func(l *Light)

// LightManager :
type LightManager struct {
	Lights  []*Light
	LRadID  int32
	LPosID  int32
	IambID  int32
	IdifID  int32
	IspecID int32
}

// Light : struct to hold light data
type Light struct {
	Radius float32
	Iamb   *float32
	Idif   *float32
	Ispec  *float32
	StdData
}

// NewLightManager :
func NewLightManager() *LightManager {

	LRadID := gl.GetUniformLocation(Shader["multiLight"], gl.Str("lightRad\x00"))
	LPosID := gl.GetUniformLocation(Shader["multiLight"], gl.Str("lightPos\x00"))
	IambID := gl.GetUniformLocation(Shader["multiLight"], gl.Str("Iamb\x00"))
	IdifID := gl.GetUniformLocation(Shader["multiLight"], gl.Str("Idif\x00"))
	IspecID := gl.GetUniformLocation(Shader["multiLight"], gl.Str("Ispec\x00"))
	lightManager = &LightManager{
		[]*Light{},
		LRadID,
		LPosID,
		IambID,
		IdifID,
		IspecID,
	}
	return lightManager
}

// NewLight :
func NewLight(pos Position, radius float32, iamb, idif, ispec []float32) *Light {
	l := &Light{
		radius,
		&iamb[0],
		&idif[0],
		&ispec[0],
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
		gl.UseProgram(Shader["multiLight"])
		gl.Uniform1f(l.LRadID, light.Radius)
		gl.Uniform3fv(l.LPosID, 1, &[]float32{light.X, light.Y, light.Z}[0])
		gl.Uniform3fv(l.IambID, 1, light.Iamb)
		gl.Uniform3fv(l.IdifID, 1, light.Idif)
		gl.Uniform3fv(l.IspecID, 1, light.Ispec)
	}
}
