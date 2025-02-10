package main

import (
	"math/rand"
	"time"

	in3d "in3D"
)

var (
	texture      map[string]uint32
	sceneObjects []*in3d.DrawnObject
)

func randObjects(numberOfObjects, min, max int, points []float32, textr, shadr uint32) {
	for i := 0; i < numberOfObjects; i++ {
		var color []float32
		rand.NewSource(time.Now().UnixNano())
		x := float32(rand.Intn(max-min) + min)
		y := float32(rand.Intn(max-min) + min)
		z := float32(rand.Intn(max-min) + min)
		if textr != in3d.NoTexture {
			color = []float32{1, 1, 1}
		} else {
			color = []float32{
				rand.Float32(),
				rand.Float32(),
				rand.Float32(),
			}
		}
		d := in3d.NewPointsObject(in3d.NewPosition(x, y, z), points, textr, color, shadr)
		sceneObjects = append(sceneObjects, d)
	}
}

func loadTextures() {
	texture = make(map[string]uint32)
	texture["none"] = in3d.NoTexture
	in3d.SetRelPath("../assets/textures")
	texture["box"] = in3d.NewTexture("box.jpg")
}

func main() {

	in3d.Init(800, 600, "Good Game")
	in3d.SetCameraPosition(in3d.NewPosition(0, 5, 100))
	in3d.Enable(in3d.PointerLock, true)
	in3d.Enable(in3d.FlyMode, true)
	light := in3d.NewLight()
	light.Draw = true
	// Close Window When Escape is Pressed
	in3d.KeyAction[in3d.KeyEscape] = func() {
		in3d.Exit()
	}

	loadTextures()
	min, max := -20, 20
	randObjects(200, min, max, in3d.Cube, texture["none"], in3d.Shader["phong"])
	randObjects(700, min, max, in3d.Cube, texture["box"], in3d.Shader["phong"])

	for !in3d.ShouldClose() {
		in3d.Update()

		for _, obj := range sceneObjects {
			obj.Draw()
		}

		in3d.SwapBuffers()
	}
}
