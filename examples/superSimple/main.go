package main

import (
	in3d "in3D"
)

func main() {

	in3d.Init(800, 600, "Simple Cube in3D")
	// in3d.NewLight().Position =
	// 	in3d.Position{X: 10, Y: 1, Z: 10}

	in3d.SetRelPath("../assets/textures")
	texture := in3d.NewTexture("seemywingz.jpg")

	obj := in3d.NewPointsObject(
		in3d.NewPosition(0, 0, -7),
		in3d.Cube,
		texture,
		in3d.White,
		in3d.Shader["texture"],
	)
	obj.SceneLogic = func(s *in3d.SceneData) {
		s.XRotation += 0.1
		s.YRotation += 0.1
	}

	for !in3d.ShouldClose() {
		in3d.Update()
		obj.Draw()
		in3d.SwapBuffers()
	}
}
