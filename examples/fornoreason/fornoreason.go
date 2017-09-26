package main

import (
	"math/rand"
	"time"

	"github.com/seemywingz/gg"
)

var (
	texture      map[string]uint32
	sceneObjects []*gg.DrawnObject
)

func randObjects(numberOfObjects, min, max int, points []float32, textr, shadr uint32) {
	for i := 0; i < numberOfObjects; i++ {

		rand.Seed(time.Now().UnixNano())
		x := float32(rand.Intn(max-min) + min)
		y := float32(rand.Intn(max-min) + min)
		z := float32(rand.Intn(max-min) + min)
		color := []float32{
			rand.Float32(),
			rand.Float32(),
			rand.Float32(),
		}
		d := gg.NewPointsObject(gg.NewPosition(x, y, z), points, textr, color, shadr)
		// dy := rand.Float32() * 1
		// dx := rand.Float32() * 1
		// d.SceneLogic = func(d *gg.SceneData) {
		// 	d.XRotation += dx
		// 	d.YRotation += dy
		// }
		sceneObjects = append(sceneObjects, d)
	}
}

func loadTextures() {
	texture = make(map[string]uint32)
	texture["none"] = gg.NoTexture
	gg.SetDirPath("github.com/seemywingz/gg/examples/assets/textures")
	texture["box"] = gg.NewTexture("box.jpg")
}

func main() {

	gg.Init(800, 600, "Good Game")
	gg.SetCameraPosition(gg.NewPosition(0, 10, 100))
	gg.Enable(gg.PointerLock, true)
	gg.Enable(gg.FlyMode, true)

	loadTextures()
	min, max := -20, 20
	randObjects(200, min, max, gg.Cube, texture["none"], gg.Shader["fixedLight"])
	randObjects(700, min, max, gg.Cube, texture["box"], gg.Shader["fixedLight"])
	lightBox := gg.NewPointsObject(gg.Position{}, gg.Cube, gg.NoTexture, []float32{1, 1, 1}, gg.Shader["basic"])
	sceneObjects = append(sceneObjects, lightBox)

	for !gg.ShouldClose() {
		gg.Update()

		for _, obj := range sceneObjects {
			obj.Draw()
		}

		gg.SwapBuffers()
	}
}
