package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// DrawnObject : interface for opengl drawable object
type DrawnObject interface {
	Draw()
}

// DrawnObjectData : a struct to hold openGL object data
type DrawnObjectData struct {
	Vao     uint32
	Program uint32
	Points  []float32
	Position
	ModelMatrix int32
	DrawnObjectDefaults
}

// DrawnObjectDefaults :
type DrawnObjectDefaults struct {
	XRotation float32
	YRotation float32
	DrawLogic DrawLogic
}

// New : Create new DrawnObjectData
func (DrawnObjectData) New(position Position, points []float32, texture uint32, program uint32) *DrawnObjectData {

	ModelMatrix := gl.GetUniformLocation(program, gl.Str("MODEL\x00"))

	return &DrawnObjectData{
		makeVao(points, program),
		program,
		points,
		position,
		ModelMatrix,
		DrawnObjectDefaults{},
	}
}

func (d *DrawnObjectData) rotate() *mgl32.Mat4 {
	model := mgl32.Translate3D(d.X, d.Y, d.Z)
	yrotMatrix := mgl32.HomogRotate3DY(mgl32.DegToRad(d.YRotation))
	xrotMatrix := mgl32.HomogRotate3DX(mgl32.DegToRad(d.XRotation))
	rotation := model.Mul4(xrotMatrix.Mul4(yrotMatrix))
	return &rotation
}

// Draw : draw the object
func (d *DrawnObjectData) Draw() {
	if d.DrawLogic != nil {
		d.DrawLogic(d)
	}
	modelMatrix := d.rotate()
	gl.UniformMatrix4fv(d.ModelMatrix, 1, false, &modelMatrix[0])
	// println(d.ModelMatrix)

	gl.UseProgram(d.Program)
	gl.BindVertexArray(d.Vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
}
