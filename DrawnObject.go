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
}

// New : Create new DrawnObjectData
func (DrawnObjectData) New(position Position, points []float32, program uint32) *DrawnObjectData {
	ptr, free := gl.Strs("translation")
	defer free()
	loc := gl.GetUniformLocation(program, *ptr)

	return &DrawnObjectData{
		makeVao(points),
		program,
		points,
		position,
		loc,
	}
}

// Draw : draw the vertecies
func (d *DrawnObjectData) Draw() {
	gl.UseProgram(d.Program)
	gl.BindVertexArray(d.Vao)
	gl.Uniform4f(d.Translation, d.X, d.Y, d.Z, 1.0)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(d.Points)/3))
}
