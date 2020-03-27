package main

import (
	"math/rand"
	"time"

	"github.com/seemywingz/in3D"
)

var (
	texture      map[string]uint32
	sceneObjects []*in3D.DrawnObject
)

func randObjects(numberOfObjects, min, max int, points []float32, textr, shadr uint32) {
	for i := 0; i < numberOfObjects; i++ {
		var color []float32
		rand.Seed(time.Now().UnixNano())
		x := float32(rand.Intn(max-min) + min)
		y := float32(rand.Intn(max-min) + min)
		z := float32(rand.Intn(max-min) + min)
		if textr != in3D.NoTexture {
			color = []float32{1, 1, 1}
		} else {
			color = []float32{
				rand.Float32(),
				rand.Float32(),
				rand.Float32(),
			}
		}
		d := in3D.NewPointsObject(in3D.NewPosition(x, y, z), points, textr, color, shadr)
		sceneObjects = append(sceneObjects, d)
	}
}

func loadTextures() {
	texture = make(map[string]uint32)
	texture["none"] = in3D.NoTexture
	in3D.SetRelPath("../assets/textures")
	texture["box"] = in3D.NewTexture("box.jpg")
}

func main() {

	in3D.Init(800, 600, "Good Game")
	in3D.SetCameraPosition(in3D.NewPosition(0, 5, 100))
	in3D.Enable(in3D.PointerLock, true)
	in3D.Enable(in3D.FlyMode, true)
	light := in3D.NewLight()
	light.Draw = true
	// Close Window When Escape is Pressed
	in3D.KeyAction[in3D.KeyEscape] = func() {
		in3D.Exit()
	}

	loadTextures()
	min, max := -20, 20
	randObjects(200, min, max, in3D.Cube, texture["none"], in3D.Shader["phong"])
	randObjects(700, min, max, in3D.Cube, texture["box"], in3D.Shader["phong"])

	for !in3D.ShouldClose() {
		in3D.Update()

		for _, obj := range sceneObjects {
			obj.Draw()
		}

		in3D.SwapBuffers()
	}
}
