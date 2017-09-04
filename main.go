package main

import (
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width  = 800
	height = 600

	vertexShaderSource = `
		#version 410
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 410
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(1, 1, 1, 1.0);
		}
	` + "\x00"
)

var (
	triangels []Triangle
)

func main() {
	runtime.LockOSThread()

	window := initGlfw(width, height, "Go-GL Boiler")
	defer glfw.Terminate()
	program := initOpenGL()

	triangels = append(triangels, Triangle{}.New())

	for !window.ShouldClose() {
		draw(window, program)
	}
}

func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	for _, obj := range triangels {
		obj.Draw()
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
