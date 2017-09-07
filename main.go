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

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			drawnObjects = append(drawnObjects, DrawnObjectData{}.New(Position{float32(i * 2), float32(j * 2), 0}, triangle, shaders[0]))
		}
	}

	camera.PointerLock = true
	for !window.ShouldClose() {
		camera.Update()
		draw()
	}
}

func randObject(points []float32) {
	min, max := -200, 200
	for i := 0; i < 2000; i++ {
		rand.Seed(time.Now().UnixNano())
		x := float32(rand.Intn(max-min) + min)
		y := float32(rand.Intn(max-min) + min)
		z := float32(rand.Intn(max-min) + min)
		drawnObjects = append(drawnObjects, DrawnObjectData{}.New(Position{x, y, z}, points, shaders[0]))
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

	for _, obj := range drawnObjects {
		obj.Draw()
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
