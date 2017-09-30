package main

import (
	"math/rand"
	"time"

	"github.com/seemywingz/in3D"
)

var (
	texture map[string]uint32
	lights  []in3D.Light
	objects []*in3D.DrawnObject
)

func randObjects(numberOfObjects, min, max int, points []float32, textr, shadr uint32) {

	for i := 0; i < numberOfObjects; i++ {

		rand.Seed(time.Now().UnixNano())
		x, y, z := in3D.Random(min, max), in3D.Random(min, max), in3D.Random(min, max)
		rx, ry, rz := in3D.Randomf(), in3D.Randomf(), in3D.Randomf()
		if i%101 == 0 {
			color := []float32{rx, ry, rz}
			roamingLight := in3D.NewColorLight([]float32{0.1, 0.1, 0.1}, color, color)
			roamingLight.Position = in3D.NewPosition(float32(x), float32(y), float32(z))
			roamingLight.Draw = true
			roamingLight.SceneLogic = func(s *in3D.SceneData) {
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
			var color []float32
			// if textr == in3D.NoTexture {
			// 	color = []float32{rx, ry, rz}
			// } else {
			color = []float32{1, 1, 1}
			// }
			d := in3D.NewPointsObject(
				in3D.NewPosition(float32(x), float32(y), float32(z)),
				points,
				textr,
				color,
				shadr)
			// d.SceneLogic = func(s *in3D.SceneData) {
			// s.XRotation += rx
			// s.YRotation += ry
			// s.ZRotation += rz
			// }
			objects = append(objects, d)
		}
	}
}

func loadTextures() {
	texture = make(map[string]uint32)
	texture["none"] = in3D.NoTexture
	in3D.SetDirPath("github.com/seemywingz/in3D/examples/assets/textures")
	texture["box"] = in3D.NewTexture("box.jpg")
	texture["box1"] = in3D.NewTexture("box1.jpg")
}

func main() {

	in3D.Init(800, 600, "Roaming Light")
	in3D.SetCameraPosition(in3D.NewPosition(0, 15, 130))
	in3D.Enable(in3D.PointerLock, true)
	in3D.Enable(in3D.FlyMode, true)
	loadTextures()
	// min, max := -100, 100
	// randObjects(800, min, max, in3D.Cube, texture["none"], in3D.Shader["phong"])

	points := []float32{}
	for i := 0; i < 8*1000; i++ {
		rand.Seed(time.Now().UnixNano())
		f := rand.Float32()
		points = append(points, f)
	}
	points = in3D.Cube
	for i, f := range points {
		if i < len(points)-1 && points[i+1] < 4 {
			points[i] = points[i+1] * 3
			points[i+1] = f
		}
	}

	obj := in3D.NewPointsObject(in3D.Position{}, points, texture["box1"], []float32{1, 1, 1}, in3D.Shader["phong"])
	obj.Scale = 10
	objects = append(objects, obj)

	centerLight := in3D.NewLight()
	centerLight.Radius = 500
	centerLight.Draw = true

	light := in3D.NewLight()
	light.Position = in3D.NewPosition(100, 10, 10)
	light.Radius = 500
	light.Draw = true

	for !in3D.ShouldClose() {
		in3D.Update()

		for _, obj := range objects {
			obj.Draw()
		}

		in3D.SwapBuffers()
	}
}
