package main

import (
	"fmt"

	"github.com/seemywingz/gg"
)

func main() {

	gg.Init(800, 600, "Wavefront Loader")
	gg.SetCameraPosition(gg.NewPosition(0, 0, 5))

	light := gg.NewLight([]float32{1, 1, 1})
	light.Iamb = []float32{0.8, 0.8, 0.8}
	light.Position = gg.NewPosition(-10, 10, 10)
	light.Radius = 300

	light = gg.NewLight([]float32{1, 1, 1})
	light.Iamb = []float32{0.001, 0.001, 0.001}
	light.Position = gg.NewPosition(10, 0, 10)
	light.Radius = 19

	gg.SetDirPath("github.com/seemywingz/gg/examples/assets/models/gopher")
	gopherMesh := gg.LoadObject("gopher.obj")
	gopher := gg.NewMeshObject(gg.Position{}, gopherMesh, gg.Shader["phong"])
	gopher.ZRotation = -90 // this .obj was exported sideways lol
	gopher.SceneLogic = func(s *gg.SceneData) {
		s.YRotation++
	}
	fmt.Println(len(gopher.Mesh.MaterialGroups))

	for !gg.ShouldClose() {
		gg.Update()
		gopher.Draw()
		gg.SwapBuffers()
	}
}
