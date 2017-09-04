package main

import (
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width  = 800
	height = 600
	title  = "go-gl Boiler"
)

var (
	window    *glfw.Window
	triangels []Triangle
)

func main() {
	runtime.LockOSThread()

	window = initGlfw(width, height, title)
	defer glfw.Terminate()

	initGL()

	triangels = append(triangels, Triangle{}.New(1, 1, 1))

	for !window.ShouldClose() {
		draw()
	}
}

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, obj := range triangels {
		obj.Draw()
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
