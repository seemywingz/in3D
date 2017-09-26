package main

import (
	"github.com/seemywingz/gg"
)

func main() {

	var objects []*gg.DrawnObject

	gg.Init(800, 600, "Wavefront Loader")
	gg.SetCameraPosition(gg.NewPosition(0, 0.5, 2))
	// gg.Enable(gg.PointerLock, true)
	// gg.Enable(gg.FlyMode, true)

	light := gg.NewLight([]float32{1, 1, 1})
	light.Iamb = []float32{0.4, 0.4, 0.4}
	light.Position = gg.NewPosition(-100, 100, 10)
	light.Radius = 300

	objFile := "buddha"
	gg.SetDirPath("github.com/seemywingz/gg/examples/assets/models/" + objFile)
	// model is from: https://www.blendswap.com/blends/view/89437
	mesh := gg.LoadObject(objFile + ".obj")
	obj := gg.NewMeshObject(gg.Position{}, mesh, gg.Shader["phong"])
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
