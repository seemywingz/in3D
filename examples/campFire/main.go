package main

import (
	in3d "in3D"
)

func main() {

	var objects []*in3d.DrawnObject

	in3d.Init(800, 600, "Wavefront Loader")
	in3d.SetCameraPosition(in3d.NewPosition(0, 40, 50))
	in3d.Enable(in3d.PointerLock, true)
	in3d.Enable(in3d.FlyMode, true)

	moonColor := []float32{0.8, 0.8, 1}
	moon := in3d.NewLight()
	moon.Diffuse = moonColor
	moon.Position = in3d.NewPosition(100, 800, 0)
	moon.Radius = 1000

	fireLight := in3d.NewLight()
	fireLight.Diffuse = []float32{10, 5, 1}
	fireLight.Specular = []float32{1, 0, 0}
	fireLight.Position = in3d.NewPosition(0, 45, 10)
	fireLight.Radius = 25
	// fireLight.Draw = true

	flickerLight := in3d.NewLight()
	flickerLight.Position = in3d.NewPosition(0, 45, 10)
	flickerLight.Radius = 25
	// flickerLight.Draw = true

	// all models are from: https://www.blendswap.com/
	objFile := "campFire"
	in3d.SetRelPath("../assets/models/" + objFile)
	mesh := in3d.LoadObject(objFile+".obj", in3d.Shader["phong"])
	obj := in3d.NewMeshObject(in3d.Position{}, mesh, in3d.Shader["phong"])
	obj.YRotation = 110
	objects = append(objects, obj)

	in3d.MojaveWorkaround()

	for !in3d.ShouldClose() {
		in3d.Update()
		if in3d.Random(0, 100)%2 == 0 {
			flickerLight.Radius = in3d.RandomF() * 20
		}
		for _, o := range objects {
			o.Draw()
		}
		in3d.SwapBuffers()
	}
}
