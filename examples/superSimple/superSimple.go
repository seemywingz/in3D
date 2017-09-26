package main

import (
	"github.com/seemywingz/in3D"
)

func main() {

	in3D.Init(800, 600, "Simple Triangle")

	tri := in3D.NewPointsObject(
		in3D.NewPosition(0, 0, -5),
		in3D.Triangle,
		0,
		[]float32{1, 0, 1},
		in3D.Shader["color"],
	)

	for !in3D.ShouldClose() {
		in3D.Update()
		tri.XRotation++
		tri.YRotation++
		tri.Draw()
		in3D.SwapBuffers()
	}
}
