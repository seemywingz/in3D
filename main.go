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
)

func main() {
	runtime.LockOSThread()

	window = initGlfw(width, height, title)
	defer glfw.Terminate()

	initGL()

	square := []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,

		-0.5, 0.5, 0,
		0.5, 0.5, 0,
		0.5, -0.5, 0,
	}

	// triangle := []float32{
	// 	0, 0.5, 0,
	// 	-0.5, -0.5, 0,
	// 	0.5, -0.5, 0,
	// }

	drawnObjects = append(drawnObjects, DrawnObjectData{}.New(1, 0, 1, square))

	for !window.ShouldClose() {
		draw()
	}
}

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, obj := range drawnObjects {
		obj.Draw()
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
