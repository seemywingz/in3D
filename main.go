package main

import (
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width  = 800
	height = 800
	title  = "go-gl Boiler"
)

var (
	window       *glfw.Window
	drawnObjects []DrawnObject
	camera       Camera
	shaders      []uint32

	triangle = []float32{
		0, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}
	square = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,

		-0.5, 0.5, 0,
		0.5, 0.5, 0,
		0.5, -0.5, 0,
	}
)

func main() {
	runtime.LockOSThread()

	window = initGlfw(width, height, title)
	defer glfw.Terminate()

	initGL()
	loadShaders()

	camera = Camera{}.New()

	drawnObjects = append(drawnObjects, DrawnObjectData{}.New(Position{0, 0, 1}, Color{0, 0, 1}, triangle, shaders[0]))

	for !window.ShouldClose() {
		draw()
	}
}

func loadShaders() {
	shaders = append(
		shaders,
		createGLprogram(
			readShaderFile("./shaders/vertex.glsl"),
			readShaderFile("./shaders/fragment.glsl"),
		))
}

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	camera.Draw()

	for _, obj := range drawnObjects {
		obj.Draw()
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
