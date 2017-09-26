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

	fireColor := []float32{1, 0.5, 0}
	fireLight := gg.NewLight()
	fireLight.Dif = fireColor
	fireLight.Position = gg.NewPosition(0, 40, 10)
	fireLight.Radius = 15
	fireLight.Draw = true

	moonColor := []float32{1, 1, 1}
	moon := gg.NewLight()
	moon.Dif = moonColor
	moon.Position = gg.NewPosition(100, 800, 0)
	moon.Radius = 2000

	objFile := "campFire"
	gg.SetDirPath("github.com/seemywingz/gg/examples/assets/models/" + objFile)
	// all models are from: https://www.blendswap.com/  -- except the gopher
	mesh := gg.LoadObject(objFile + ".obj")
	obj := gg.NewMeshObject(gg.Position{}, mesh, gg.Shader["phong"])
	obj.YRotation = 110
	objects = append(objects, obj)

	for !gg.ShouldClose() {
		gg.Update()
		for _, o := range objects {
			o.Draw()
		}
		gg.SwapBuffers()
	}
}
