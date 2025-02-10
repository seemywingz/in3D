package main

import (
	"math/rand"
	"time"

	in3d "in3D"
)

var (
	texture      map[string]uint32
	drawnObjects []*in3d.DrawnObject
)

func randObjects(numberOfObjects, min, max int, points []float32, textr, shadr uint32) {
	for i := 0; i < numberOfObjects; i++ {
		rand.NewSource(time.Now().UnixNano())
		x, y, z := in3d.Random(min, max), in3d.Random(min, max), in3d.Random(min, max)
		color := []float32{1, 1, 1}
		d := in3d.NewPointsObject(
			in3d.NewPosition(float32(x), float32(y), float32(z)),
			points,
			textr,
			color,
			shadr)
		// d.SceneLogic = func(s *in3d.SceneData) {
		// s.XRotation += rx
		// s.YRotation += ry
		// s.ZRotation += rz
		// }
		drawnObjects = append(drawnObjects, d)
	}
}

func randLights(numberOfLights, min, max int) {

	for i := 0; i < numberOfLights; i++ {

		rand.NewSource(time.Now().UnixNano())
		x, y, z := in3d.Random(min, max), in3d.Random(min, max), in3d.Random(min, max)
		rx, ry, rz := in3d.RandomF(), in3d.RandomF(), in3d.RandomF()
		color := []float32{rx, ry, rz}
		roamingLight := in3d.NewColorLight([]float32{0.1, 0.1, 0.1}, color, color)
		roamingLight.Position = in3d.NewPosition(float32(x), float32(y), float32(z))
		roamingLight.Draw = true
		roamingLight.SceneLogic = func(s *in3d.SceneData) {
			s.X += rx * 1.2
			s.Y += ry * 1.2
			s.Z += rz * 1.2
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
	}
}

func loadTextures() {
	texture = make(map[string]uint32)
	texture["none"] = in3d.NoTexture
	in3d.SetDir("../assets/textures")
	texture["box"] = in3d.NewTexture("box.jpg")
	texture["box1"] = in3d.NewTexture("box1.jpg")
}

func main() {

	in3d.Init(000, 600, "Roaming Light")
	in3d.SetCameraPosition(in3d.NewPosition(0, 15, 130))
	in3d.Enable(in3d.PointerLock, true)
	in3d.Enable(in3d.FlyMode, true)

	// Close Window When Escape is Pressed
	in3d.KeyAction[in3d.KeyEscape] = func() {
		in3d.Exit()
	}

	loadTextures()

	min, max := -100, 100
	randObjects(800, min, max, in3d.Cube, texture["none"], in3d.Shader["phong"])
	randObjects(100, min, max, in3d.Cube, texture["box"], in3d.Shader["phong"])
	randObjects(100, min, max, in3d.Cube, texture["box1"], in3d.Shader["phong"])
	randLights(9, min, max)

	// centerLight := in3d.NewLight()
	// centerLight.Draw = true
	// println(centerLight.DrawnObject.IdifID)

	for !in3d.ShouldClose() {
		in3d.Update()

		for _, obj := range drawnObjects {
			obj.Draw()
		}

		in3d.SwapBuffers()
	}
}
