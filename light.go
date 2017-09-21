package gg

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// LightLogic :
type LightLogic func(l *Light)

// Light : struct to hold light data
type Light struct {
	Position
	Iamb    *float32
	Idif    *float32
	Ispec   *float32
	LPosID  int32
	IambID  int32
	IdifID  int32
	IspecID int32
	LightDefaults
}

// LightData :
type LightDefaults struct {
	LightLogic LightLogic
}

// NewLight :
func NewLight(pos Position, iamb, idif, ispec []float32, logic LightLogic) *Light {

	LPosID := gl.GetUniformLocation(Shader["singleLight"], gl.Str("lightPos\x00"))
	IambID := gl.GetUniformLocation(Shader["singleLight"], gl.Str("Iamb\x00"))
	IdifID := gl.GetUniformLocation(Shader["singleLight"], gl.Str("Idif\x00"))
	IspecID := gl.GetUniformLocation(Shader["singleLight"], gl.Str("Ispec\x00"))

	return &Light{
		pos,
		&iamb[0],
		&idif[0],
		&ispec[0],
		LPosID,
		IambID,
		IdifID,
		IspecID,
		LightDefaults{},
	}
}

// Draw :
func (l *Light) Draw() {
	if l.LightLogic != nil {
		l.LightLogic(l)
	}
	gl.UseProgram(Shader["singleLight"])
	gl.Uniform3fv(l.LPosID, 1, &[]float32{l.X, l.Y, l.Z}[0])
	gl.Uniform3fv(l.IambID, 1, l.Iamb)
	gl.Uniform3fv(l.IdifID, 1, l.Idif)
	gl.Uniform3fv(l.IspecID, 1, l.Ispec)
}
