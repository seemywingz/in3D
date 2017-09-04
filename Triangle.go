package main

import "github.com/go-gl/gl/v4.1-core/gl"

// Triangle : a struct to hold openGL triangle data
type Triangle struct {
	Points []float32
	Vao    *uint32
}

var (
	vao    uint32
	points = []float32{
		0, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}
)

// New : Create a new Triangle
func (t Triangle) New() Triangle {
	if vao == 0 {
		vao = makeVao(points)
	}
	return Triangle{points, &vao}
}

// Draw : draw the triangle
func (t Triangle) Draw() {
	gl.BindVertexArray(*t.Vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(t.Points)/3))
}
