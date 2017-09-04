package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	gt "github.com/seemywingz/gtils"
)

// DrawnObject : interface for opengl drawable object
type DrawnObject interface {
	Draw()
}

// Color : struct to store RGB colors as float32
type Color struct {
	R float32
	G float32
	B float32
}

// Position : struct to store 3D coords
type Position struct {
	X float32
	Y float32
	Z float32
}

// DrawnObjectData : a struct to hold openGL object data
type DrawnObjectData struct {
	Vao     uint32
	Program uint32
	Points  []float32
}

// New : Create a new object
func (d DrawnObjectData) New(color Color, points []float32) DrawnObjectData {
	vertexShaderSource := `
  		#version 410
  		in vec3 vp;
  		void main() {
  			gl_Position = vec4(vp, 10.0);
  		}
  	` + "\x00"

	fragmentShaderSource := `
  		#version 410
  		out vec4 frag_colour;
  		void main() {
  			frag_colour = vec4(` + gt.FtoA(color.R) + `, ` + gt.FtoA(color.G) + `, ` + gt.FtoA(color.B) + `, 1.0);
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
