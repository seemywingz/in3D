package main

import (
	"strconv"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// Square : a struct to hold openGL triangle data
type DrawnObject struct {
	Vao     *uint32
	Program uint32
	Points  []float32
}

var (
	vao    uint32
	points = []float32{
		0, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}
)

// New : Create a new Square
func (d DrawnObject) New(r, g, b float32) DrawnObject {
	rs := strconv.FormatFloat(float64(r), 'f', 6, 32)
	rg := strconv.FormatFloat(float64(g), 'f', 6, 32)
	rb := strconv.FormatFloat(float64(b), 'f', 6, 32)

	if vao == 0 {
		vao = makeVao(points)
	}

	vertexShaderSource := `
  		#version 410
  		in vec3 vp;
  		void main() {
  			gl_Position = vec4(vp, 1.0);
  		}
  	` + "\x00"

	fragmentShaderSource := `
  		#version 410
  		out vec4 frag_colour;
  		void main() {
  			frag_colour = vec4(` + rs + `, ` + rg + `, ` + rb + `, 1.0);
  		}
  	` + "\x00"

	program := createGLprogram(vertexShaderSource, fragmentShaderSource)
	return DrawnObject{&vao, program, points}
}

// Draw : draw the triangle
func (d DrawnObject) Draw() {
	gl.UseProgram(d.Program)
	gl.BindVertexArray(*d.Vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(d.Points)/3))
}
