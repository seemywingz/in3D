package main

import (
	"math/rand"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	gt "github.com/seemywingz/gtils"
)

var (
	camera       Camera
	window       *glfw.Window
	shaders      []uint32
	drawnObjects []DrawnObject
)

func main() {
	runtime.LockOSThread()

	var windowWidth = 800
	var windowHeight = 600
	window = initGlfw(windowWidth, windowHeight, "go-gl Boiler")
	defer glfw.Terminate()

	initGL()
	loadShaders()
	loadLights()

	gt.SetDirPath("github.com/seemywingz/go-gl_boiler")
	boxTexture := newTexture("textures/square.jpg")
	defer gl.DeleteTextures(1, &boxTexture)

	camera = Camera{}.New(Position{0, 0, 0}, false)

	randObject(10000, -200, 200, cube, boxTexture)
	drawnObjects = append(drawnObjects, DrawnObjectData{}.New(Position{0, 0, -4}, cube, boxTexture, shaders[0]))

	for !window.ShouldClose() {
		camera.Update()
		update()
	}
}

func loadLights() {
	// ambient := []float32{0.5, 0.5, 0.5, 1}
	// diffuse := []float32{1, 1, 1, 1}
	// lightPosition := []float32{-5, 5, 10, 0}
	// gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambient[0])
	// gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &diffuse[0])
	// gl.Lightfv(gl.LIGHT0, gl.POSITION, &lightPosition[0])
	// gl.Enable(gl.LIGHT0)
}

func randObject(numberOfObjects, min, max int, points []float32, texture uint32) {
	for i := 0; i < numberOfObjects; i++ {
		rand.Seed(time.Now().UnixNano())
		x := float32(rand.Intn(max-min) + min)
		y := float32(rand.Intn(max-min) + min)
		z := float32(rand.Intn(max-min) + min)
		if i == numberOfObjects/2 {
			println("Adding Lifion Box")
			lifionTexture := newTexture("textures/lifion.png")
			// defer gl.DeleteTextures(1, &lifionTexture)
			drawnObjects = append(drawnObjects, DrawnObjectData{}.New(Position{x, y, z}, points, lifionTexture, shaders[0]))
		} else {
			drawnObjects = append(drawnObjects, DrawnObjectData{}.New(Position{x, y, z}, points, texture, shaders[0]))
		}
	}
}

func loadShaders() {
	shaders = append(shaders, createGLprogram(basicVertexSRC, basicFragmentSRC))
}

func update() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, obj := range drawnObjects {
		obj.Draw()
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
