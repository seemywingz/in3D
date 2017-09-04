package main

import (
	"strconv"

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
}

// New : Create a new object
func (d DrawnObjectData) New(r, g, b float32, points []float32) DrawnObjectData {
	rs := strconv.FormatFloat(float64(r), 'f', 6, 32)
	rg := strconv.FormatFloat(float64(g), 'f', 6, 32)
	rb := strconv.FormatFloat(float64(b), 'f', 6, 32)

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
	vao := makeVao(points)
	return DrawnObjectData{vao, program, points}
}

// Draw : draw the triangle
func (d DrawnObjectData) Draw() {
	gl.UseProgram(d.Program)
	gl.BindVertexArray(d.Vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(d.Points)/3))
}
