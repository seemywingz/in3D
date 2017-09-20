package main

import (
	"math/rand"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/seemywingz/gt"
)

func main() {
	runtime.LockOSThread()

	var windowWidth = 00
	var windowHeight = 600
	window = initGlfw(windowWidth, windowHeight, "go-gl Boiler")
	defer glfw.Terminate()

	initGL()
	loadShaders()
	loadTextures()

	camera = Camera{}.New(Position{0, 0, 0}, false)

	randObject(100, -200, 200, cube, texture["none"], shader["basic"])
	randObject(200, -200, 200, cube, texture["box"], shader["color"])
	// randObject(200, -200, 200, cube, texture["box"], shader["texture"])
	randObject(700, -200, 200, cube, texture["box"], shader["phong"])
	drawnObjects = append(drawnObjects, Card{}.New(Position{0, 0, -5}, "tk"))
	drawnObjects = append(drawnObjects, Card{}.New(Position{-5, 0, -5}, "imn"))

	// d := DrawnObjectData{}.New(Position{0, 0, -5}, cube, texture["box"], shader["phong"])
	// d.DrawLogic = func(d *DrawnObjectData) {
	// 	// d.XRotation++
	// 	d.YRotation++
	// }
	// drawnObjects = append(drawnObjects, d)

	for !window.ShouldClose() {
		camera.Update()
		update()
	}
}

func randObject(numberOfObjects, min, max int, points []float32, textr, shadr uint32) {
	for i := 0; i < numberOfObjects; i++ {
		rand.Seed(time.Now().UnixNano())
		x := float32(rand.Intn(max-min) + min)
		y := float32(rand.Intn(max-min) + min)
		z := float32(rand.Intn(max-min) + min)
		d := DrawnObjectData{}.New(Position{x, y, z}, points, textr, shadr)
		dy := rand.Float32() * 10
		// dx := rand.Float32() * 10
		d.DrawLogic = func(d *DrawnObjectData) {
			// d.XRotation += dx
			d.YRotation += dy
		}
		d.Color = Color{gt.Randomf(), gt.Randomf(), gt.Randomf(), 0}
		drawnObjects = append(drawnObjects, d)
	}
}

func loadShaders() {
	shader = make(map[string]uint32)
	gt.SetDirPath("github.com/seemywingz/gg/shaders")
	shader["basic"] = createGLprogram("basicVect.glsl", "basicFrag.glsl")
	shader["color"] = createGLprogram("basicVect.glsl", "colorFrag.glsl")
	shader["texture"] = createGLprogram("textureVect.glsl", "textureFrag.glsl")
	shader["phong"] = createGLprogram("blinnPhongVect.glsl", "blinnPhongFrag.glsl")
}

func loadTextures() {
	texture = make(map[string]uint32)
	texture["none"] = 99999
	gt.SetDirPath("github.com/seemywingz/gg/textures")
	texture["lifion"] = newTexture("lifion.png")
	texture["box"] = newTexture("square.jpg")
	texture["tk"] = newTexture("tk.jpg")
	texture["imn"] = newTexture("imn.jpg")
	texture["back"] = newTexture("back.jpg")
}

func update() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, obj := range drawnObjects {
		obj.Draw()
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
