package main

import (
	"math/rand"
	"time"

	"github.com/seemywingz/gg"
)

var (
	texture      map[string]uint32
	lights       []gg.Light
	drawnObjects []*gg.DrawnObject
)

func randObjects(numberOfObjects, min, max int, points []float32, textr, shadr uint32) {

	for i := 0; i < numberOfObjects; i++ {

		rand.Seed(time.Now().UnixNano())
		x, y, z := gg.Random(min, max), gg.Random(min, max), gg.Random(min, max)
		rx, ry, rz := gg.Randomf(), gg.Randomf(), gg.Randomf()
		if i%101 == 0 {
			roamingLight := gg.NewLight()
			roamingLight.Position = gg.NewPosition(float32(x), float32(y), float32(z))
			roamingLight.Draw = true
			roamingLight.Idif = []float32{1, 0, 1}
			// roamingLight.Idif = []float32{rx, ry, rz}
			roamingLight.SceneLogic = func(s *gg.SceneData) {
				s.X += rx * 3
				s.Y += ry * 3
				s.Z += rz * 3
				if s.X > float32(max) || s.X < float32(min) {
					rx = -rx
				}
				if s.Y > float32(max) || s.Y < float32(min) {
					ry = -ry
				}
				if s.Z > float32(max) || s.Z < float32(min) {
					rz = -rz
				}
			}
		} else {
			d := gg.NewPointsObject(
				gg.NewPosition(float32(x), float32(y), float32(z)),
				points,
				textr,
				shadr)
			// d.SceneLogic = func(s *gg.SceneData) {
			// s.XRotation += rx
			// s.YRotation += ry
			// s.ZRotation += rz
			// }
			if textr == gg.NoTexture {
				// d.Color = gg.NewColor(rx, ry, rz, 1)
			}
			drawnObjects = append(drawnObjects, d)
		}
	}
}

func loadTextures() {
	texture = make(map[string]uint32)
	texture["none"] = gg.NoTexture
	gg.SetDirPath("github.com/seemywingz/gg/examples/assets/textures")
	texture["box"] = gg.NewTexture("box.jpg")
	texture["box1"] = gg.NewTexture("box1.jpg")
}

func main() {

	gg.Init(000, 600, "Roaming Light")
	gg.SetCameraPosition(gg.NewPosition(0, 50, 300))
	gg.Enable(gg.PointerLock, true)
	gg.Enable(gg.FlyMode, true)

	loadTextures()

	min, max := -100, 100
	randObjects(800, min, max, gg.Cube, texture["none"], gg.Shader["phong"])
	randObjects(100, min, max, gg.Cube, texture["box"], gg.Shader["phong"])
	randObjects(100, min, max, gg.Cube, texture["box1"], gg.Shader["phong"])

	centerLight := gg.NewLight()
	centerLight.Draw = true
	// centerLight.Idif = []float32{1, 1, 0} // R G B

	for !gg.ShouldClose() {
		gg.Update()

		for _, obj := range drawnObjects {
			obj.Draw()
		}

		gg.SwapBuffers()
	}
}
