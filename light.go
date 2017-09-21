package gg

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Light : struct to hold light data
type Light struct {
	LPosID  int32
	IambID  int32
	IdifID  int32
	IspecID int32
	LightData
}

// LightData :
type LightData struct {
	LPos  *float32
	Iamb  *float32
	Idif  *float32
	Ispec *float32
}

// NewLight :
func NewLight(pos, iamb, idif, ispec []float32) *Light {

	LPos := &[]float32{1, 1, 1}[0]
	Iamb := &[]float32{0, 0, 0}[0]
	Idif := &[]float32{0.4, 0.4, 0.4}[0]
	Ispec := &[]float32{10, 10, 10}[0]

	LPosID := gl.GetUniformLocation(Shader["singleLight"], gl.Str("lightPos\x00"))
	IambID := gl.GetUniformLocation(Shader["singleLight"], gl.Str("Iamb\x00"))
	IdifID := gl.GetUniformLocation(Shader["singleLight"], gl.Str("Idif\x00"))
	IspecID := gl.GetUniformLocation(Shader["singleLight"], gl.Str("Ispec\x00"))

	data := LightData{
		LPos,
		Iamb,
		Idif,
		Ispec,
	}

	return &Light{
		LPosID,
		IambID,
		IdifID,
		IspecID,
		data,
	}
}

// Draw :
func (l *Light) Draw() {
	gl.UseProgram(Shader["singleLight"])
	gl.Uniform3fv(l.LPosID, 1, l.LPos)
	gl.Uniform3fv(l.IambID, 1, l.Iamb)
	gl.Uniform3fv(l.IdifID, 1, l.Idif)
	gl.Uniform3fv(l.IspecID, 1, l.Ispec)

}
