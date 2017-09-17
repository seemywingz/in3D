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
	LocalMVP int32
}

// New : Create new DrawnObjectData
func (DrawnObjectData) New(position Position, points []float32, program uint32) *DrawnObjectData {

	ptr, free1 := gl.Strs("localRotation")
	defer free1()
	lmvploc := gl.GetUniformLocation(program, *ptr)

	return &DrawnObjectData{
		makeVao(points),
		program,
		points,
		position,
		lmvploc,
	}
}

var n float32

// Draw : draw the vertecies
func (d *DrawnObjectData) Draw() {

	n += 0.001
	if n > 360 {
		n = 0
	}

	// translate to obj position
	m := mgl32.Translate3D(d.X, d.Y, d.Z)

	// rotataton
	yrotMatrix := mgl32.HomogRotate3DY(mgl32.DegToRad(n))
	rotation := m.Mul4(yrotMatrix)

	gl.UseProgram(d.Program)
	gl.BindVertexArray(d.Vao)

	// println(d.LocalMVP)
	gl.UniformMatrix4fv(d.LocalMVP, 1, false, &rotation[0])

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(d.Points)/3))
}
