package main

//
// import (
// 	"strconv"
//
// 	"github.com/go-gl/gl/v4.1-core/gl"
// )
//
// // Square : a struct to hold openGL triangle data
// type Square struct {
// 	Vao     *uint32
// 	Program uint32
// 	Points  []float32
// }
//
// var (
// 	squareVAO    uint32
// 	squarePoints = []float32{
// 		0, 0.5, 0,
// 		-0.5, -0.5, 0,
// 		0.5, -0.5, 0,
// 	}
// )
//
// // New : Create a new Square
// func (s Square) New(r, g, b float32) Square {
// 	rs := strconv.FormatFloat(float64(r), 'f', 6, 32)
// 	rg := strconv.FormatFloat(float64(g), 'f', 6, 32)
// 	rb := strconv.FormatFloat(float64(b), 'f', 6, 32)
//
// 	if squareVAO == 0 {
// 		squareVAO = makeVao(squarePoints)
// 	}
//
// 	vertexShaderSource := `
//   		#version 410
//   		in vec3 vp;
//   		void main() {
//   			gl_Position = vec4(vp, 1.0);
//   		}
//   	` + "\x00"
//
// 	fragmentShaderSource := `
//   		#version 410
//   		out vec4 frag_colour;
//   		void main() {
//   			frag_colour = vec4(` + rs + `, ` + rg + `, ` + rb + `, 1.0);
//   		}
//   	` + "\x00"
//
// 	program := createGLprogram(vertexShaderSource, fragmentShaderSource)
// 	return Square{&squareVAO, program, squarePoints}
// }
//
// // Draw : draw the triangle
// func (s Square) Draw() {
// 	gl.UseProgram(s.Program)
// 	gl.BindVertexArray(*s.Vao)
// 	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(s.Points)/3))
// }
