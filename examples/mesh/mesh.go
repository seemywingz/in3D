package main

import (
	"github.com/seemywingz/gg"
)

func main() {

	var objects []*gg.DrawnObject

	gg.Init(800, 600, "Wavefront Loader")
	gg.SetCameraPosition(gg.NewPosition(0, 40, 50))
	gg.Enable(gg.PointerLock, true)
	gg.Enable(gg.FlyMode, true)

	light := gg.NewLight([]float32{1, 1, 1})
	light.Iamb = []float32{0.4, 0.4, 0.4}
	// light.Idif = []float32{0.9, 0.9, 0.9}
	light.Ispec = []float32{0.9, 0.9, 0.9}
	light.Position = gg.NewPosition(0, 40, 10)
	light.Radius = 25
	light.Draw = true

	objFile := "campFire"
	gg.SetDirPath("github.com/seemywingz/gg/examples/assets/models/" + objFile)
	// all models are from: https://www.blendswap.com/  -- except the gopher
	mesh := gg.LoadObject(objFile + ".obj")
	obj := gg.NewMeshObject(gg.Position{}, mesh, gg.Shader["phong"])
	obj.YRotation = 45
	// obj.SceneLogic = func(s *gg.SceneData) {
	// 	s.YRotation++
	// }
	objects = append(objects, obj)

	for !gg.ShouldClose() {
		gg.Update()
		for _, o := range objects {
			o.Draw()
		}
		gg.SwapBuffers()
	}
}
