package main

import (
	"github.com/seemywingz/in3D"
)

func main() {

	var objects []*in3D.DrawnObject

	in3D.Init(800, 600, "Wavefront Loader")
	in3D.SetClearColor(0.1, 0.1, 0.1, 1)
	in3D.SetCameraPosition(in3D.NewPosition(0, 0.55, 2))
	in3D.Enable(in3D.PointerLock, true)
	in3D.Enable(in3D.FlyMode, true)

	light := in3D.NewLight()
	light.Ambient = []float32{0.6, 0.6, 0.6}
	light.Position = in3D.NewPosition(0, 100, -100)
	light.Radius = 100000

	// all models are from: https://www.blendswap.com/
	model := "sky"
	in3D.SetRelPath("../assets/models/" + model)

	skymesh := in3D.LoadObject(model + ".obj")
	sky := in3D.NewMeshObject(in3D.Position{}, skymesh, in3D.Shader["texture"])
	sky.Scale = 10000
	objects = append(objects, sky)

	model = "buddha"
	in3D.SetRelPath("../assets/models/" + model)
	bmesh := in3D.LoadObject(model + ".obj")
	buddha := in3D.NewMeshObject(in3D.Position{}, bmesh, in3D.Shader["phong"])
	objects = append(objects, buddha)

	for !in3D.ShouldClose() {
		in3D.Update()
		for _, o := range objects {
			o.Draw()
		}
		in3D.SwapBuffers()
	}
}
