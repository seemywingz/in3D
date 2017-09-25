package main

import (
	"github.com/seemywingz/gg"
)

func main() {

	gg.Init(800, 600, "Wavefront Loader")
	gg.SetCameraPosition(gg.NewPosition(0, 0, 5))

	l := gg.NewLight()
	l.Position = gg.NewPosition(-10, 10, 10)
	l.Radius = 30

	gg.SetDirPath("github.com/seemywingz/gg/examples/gopher")
	gopherMesh := gg.LoadObject("gopher.obj")
	gopher := gg.NewMeshObject(gg.Position{}, gopherMesh, gg.NoTexture, gg.Shader["phong"])
	gopher.ZRotation = -90
	gopher.SceneLogic = func(s *gg.SceneData) {
		s.YRotation++
	}

	for !gg.ShouldClose() {
		gg.Update()
		gopher.Draw()
		gg.SwapBuffers()
	}
}
