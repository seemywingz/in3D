package main

import (
	"fmt"
	"os"
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

	camera = Camera{}.New(Position{0, 0, -10})

	drawnObjects = append(drawnObjects, DrawnObjectData{}.New(Position{0, 0, 1}, Color{0, 0, 1}, triangle, shaders[0]))

	for !window.ShouldClose() {
		handleMouse()
		handleKeys()
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

func handleMouse() {
	x, y := window.GetCursorPos()
	fmt.Println(x, y)
	if window.GetMouseButton(glfw.MouseButton1) == glfw.Press {
		fmt.Println("Click")
	}
}

func handleKeys() {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		os.Exit(1)
	}
	// Press w
	if window.GetKey(glfw.KeyW) == glfw.Press {
		camera.Z += 0.1
	}
	// Press A
	if window.GetKey(glfw.KeyA) == glfw.Press {
		yangle--
	}
	// Press s
	if window.GetKey(glfw.KeyS) == glfw.Press {
		camera.Z -= 0.1
	}
	// Press d
	if window.GetKey(glfw.KeyD) == glfw.Press {
		yangle++
	}
	// Press q
	if window.GetKey(glfw.KeyQ) == glfw.Press {
		xangle--
	}
	// Press e
	if window.GetKey(glfw.KeyE) == glfw.Press {
		xangle++
	}
}

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	camera.Update()

	for _, obj := range drawnObjects {
		obj.Draw()
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
