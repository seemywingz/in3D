package main

import (
	"github.com/seemywingz/in3D"
)

func main() {

	var objects []*in3D.DrawnObject

	in3D.Init(800, 600, "Wavefront Loader")
	in3D.SetCameraPosition(in3D.NewPosition(0, 40, 50))
	in3D.Enable(in3D.PointerLock, true)
	in3D.Enable(in3D.FlyMode, true)

	moonColor := []float32{0.3, 0.3, 0.8}
	moon := in3D.NewLight()
	moon.Difffuse = moonColor
	moon.Position = in3D.NewPosition(100, 800, 0)
	moon.Radius = 1000

	fireLight := in3D.NewLight()
	fireLight.Difffuse = []float32{10, 5, 1}
	fireLight.Specular = []float32{1, 0, 0}
	fireLight.Position = in3D.NewPosition(0, 45, 10)
	fireLight.Radius = 25
	fireLight.Draw = true

	objFile := "campFire"
	in3D.SetDirPath("github.com/seemywingz/in3D/examples/assets/models/" + objFile)
	// all models are from: https://www.blendswap.com/
	mesh := in3D.LoadObject(objFile + ".obj")
	obj := in3D.NewMeshObject(in3D.Position{}, mesh, in3D.Shader["phong"])
	obj.YRotation = 110
	objects = append(objects, obj)

	for !in3D.ShouldClose() {
		in3D.Update()
		for _, o := range objects {
			o.Draw()
		}
		in3D.SwapBuffers()
	}
}
