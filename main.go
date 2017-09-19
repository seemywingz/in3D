package main

import (
	"math/rand"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	gt "github.com/seemywingz/gtils"
)

func main() {
	runtime.LockOSThread()
	gt.SetDirPath("github.com/seemywingz/go-gl_boiler")

	var windowWidth = 800
	var windowHeight = 600
	window = initGlfw(windowWidth, windowHeight, "go-gl Boiler")
	defer glfw.Terminate()

	initGL()
	loadShaders()
	loadTextures()
	loadLights()

	camera = Camera{}.New(Position{0, 0, 0}, false)

	// randObject(1000, -200, 200, cube, texture["box"])
	card1 := DrawnObjectData{}.New(Position{0, 0, -4}, card, texture["tk"], shaders["phong"])
	card1.Logic = func(d *DrawnObjectData) {
		d.XRotation++
	}
	drawnObjects = append(drawnObjects, card1)

	// ground := DrawnObjectData{}.New(Position{0, 0, 0}, card, texture["tk"], shaders["phong"], nil)
	// ground.XRotation = -90
	// drawnObjects = append(drawnObjects, ground)

	for !window.ShouldClose() {
		camera.Update()
		update()
	}
}

func randObject(numberOfObjects, min, max int, points []float32, textr uint32) {
	for i := 0; i < numberOfObjects; i++ {
		rand.Seed(time.Now().UnixNano())
		x := float32(rand.Intn(max-min) + min)
		y := float32(rand.Intn(max-min) + min)
		z := float32(rand.Intn(max-min) + min)
		if i == numberOfObjects/2 {
			println("Adding Lifion Box")
			drawnObjects = append(drawnObjects, DrawnObjectData{}.New(Position{x, y, z}, points, texture["lifion"], shaders["phong"]))
		}
		drawnObjects = append(drawnObjects, DrawnObjectData{}.New(Position{x, y, z}, points, textr, shaders["phong"]))
	}
}

func loadLights() {
}

func loadShaders() {
	shaders = make(map[string]uint32)
	shaders["phong"] = createGLprogram(basicVertexSRC, basicFragmentSRC)
}

func loadTextures() {
	texture = make(map[string]uint32)
	texture["lifion"] = newTexture("textures/lifion.png")
	texture["box"] = newTexture("textures/square.jpg")
	texture["tk"] = newTexture("textures/tk.jpg")
}

func update() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, obj := range drawnObjects {
		obj.Draw()
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
