package main

import (
	"fmt"

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

	ptrT, freeT := gl.Strs("trans")
	defer freeT()
	transloc := gl.GetUniformLocation(program, *ptrT)

	ptrR, freeR := gl.Strs("rot")
	defer freeR()
	rotloc := gl.GetUniformLocation(program, *ptrR)
	fmt.Println(rotloc)

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
	n += 0.1
	if n > 180 {
		n = 0
	}

	gl.Uniform4f(d.Translation, d.X, d.Y, d.Z, 1.0)
	gl.Uniform4f(d.Rotation, n, d.Y, d.Z, 1.0)

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(d.Points)/3))
}
