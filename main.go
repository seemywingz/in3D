package main

import (
	"math/rand"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width  = 800
	height = 600
	title  = "go-gl Boiler"
)

var (
	camera       Camera
	shaders      []uint32
	window       *glfw.Window
	drawnObjects []DrawnObject

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

	camera = Camera{}.New(Position{0, 0, -10})

	randObject(1000, -200, 200, square)

	camera.PointerLock = true
	for !window.ShouldClose() {
		camera.Update()
		draw()
	}
}

func randObject(numberOfObjects, min, max int, points []float32) {
	for i := 0; i < numberOfObjects; i++ {
		rand.Seed(time.Now().UnixNano())
		x := float32(rand.Intn(max-min) + min)
		y := float32(rand.Intn(max-min) + min)
		z := float32(rand.Intn(max-min) + min)
		drawnObjects = append(drawnObjects, DrawnObjectData{}.New(Position{x, y, z}, points, shaders[0]))
	}
}

func loadShaders() {
	shaders = append(shaders, createGLprogram(basicVertexSRC, basicFragmentSRC))
}

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, obj := range drawnObjects {
		obj.Draw()
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
