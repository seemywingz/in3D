package main

import (
	"github.com/seemywingz/gg"
)

func main() {

	var objects []*gg.DrawnObject

	gg.Init(800, 600, "Wavefront Loader")
	gg.SetCameraPosition(gg.NewPosition(0, 0.5, 2))

	l := gg.NewLight([]float32{1, 1, 1})
	l.Position = gg.NewPosition(-10, 10, 10)
	l.Radius = 300

	gg.SetDirPath("github.com/seemywingz/gg/examples/buddha/assets")
	mesh := gg.LoadObject("buddha.obj")
	obj := gg.NewMeshObject(gg.Position{}, mesh, gg.NewTexture("buddha.jpg"), gg.Shader["phong"])
	obj.SceneLogic = func(s *gg.SceneData) {
		s.YRotation++
	}
	objects = append(objects, obj)

	for !gg.ShouldClose() {
		gg.Update()
		for _, o := range objects {
			o.Draw()
		}
		gg.SwapBuffers()
	}
}
