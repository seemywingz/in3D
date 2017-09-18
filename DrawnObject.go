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
	Texture     uint32
	DrawnObjectDefaults
}

// DrawnObjectDefaults :
type DrawnObjectDefaults struct {
	XRotation float32
	YRotation float32
}

// New : Create new DrawnObjectData
func (DrawnObjectData) New(position Position, points []float32, texture uint32, program uint32) *DrawnObjectData {

	ptr, free1 := gl.Strs("MODEL")
	defer free1()
	ModelMatrix := gl.GetUniformLocation(program, *ptr)

	return &DrawnObjectData{
		makeVao(points, program),
		program,
		points,
		position,
		ModelMatrix,
		texture,
		DrawnObjectDefaults{},
	}
}

func (d *DrawnObjectData) rotate() *mgl32.Mat4 {
	d.YRotation++
	d.XRotation++

	model := mgl32.Translate3D(d.X, d.Y, d.Z)
	yrotMatrix := mgl32.HomogRotate3DY(mgl32.DegToRad(d.YRotation))
	xrotMatrix := mgl32.HomogRotate3DX(mgl32.DegToRad(d.XRotation))
	rotation := model.Mul4(xrotMatrix.Mul4(yrotMatrix))
	return &rotation
}

// Draw : draw the vertecies
func (d *DrawnObjectData) Draw() {
	gl.UseProgram(d.Program)
	gl.UniformMatrix4fv(d.ModelMatrix, 1, false, &d.rotate()[0])
	gl.BindVertexArray(d.Vao)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, d.Texture)

	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
}
