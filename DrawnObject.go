package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
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
	Translation int32
	Rotation    int32
}

// New : Create new DrawnObjectData
func (DrawnObjectData) New(position Position, points []float32, program uint32) *DrawnObjectData {

	ptrT, freeT := gl.Strs("translation")
	defer freeT()
	transloc := gl.GetUniformLocation(program, *ptrT)

	ptrR, freeR := gl.Strs("rotation")
	defer freeR()
	rotloc := gl.GetUniformLocation(program, *ptrR)

	return &DrawnObjectData{
		makeVao(points),
		program,
		points,
		position,
		transloc,
		rotloc,
	}
}

var n float32

// Draw : draw the vertecies
func (d *DrawnObjectData) Draw() {
	gl.UseProgram(d.Program)
	gl.BindVertexArray(d.Vao)
	n += 0.001
	// xrotMatrix := mgl32.HomogRotate3D(mgl32.DegToRad(n), mgl32.Vec3{1, 0, 0})
	// yrotMatrix := mgl32.HomogRotate3D(mgl32.DegToRad(n), mgl32.Vec3{0, 1, 0})
	// model := xrotMatrix.Mul4(yrotMatrix.Mul4(mgl32.Ident4()))
	gl.Uniform4f(d.Translation, d.X, d.Y, d.Z, 1.0)
	// gl.UniformMatrix4fv(d.Rotation, 1, false, &model[0])
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(d.Points)/3))
}
