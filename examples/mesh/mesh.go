package main

import "github.com/seemywingz/in3D"

func main() {

	var objects []*in3D.DrawnObject

	in3D.Init(800, 600, "Wavefront Loader")
	in3D.SetClearColor(0.1, 0.1, 0.1, 1)
	in3D.SetCameraPosition(in3D.NewPosition(0, 0.55, 2))
	in3D.Enable(in3D.PointerLock, true)
	in3D.Enable(in3D.FlyMode, true)

	light := in3D.NewLight()
	light.Ambient = []float32{0, 0, 0}
	light.Position = in3D.NewPosition(-10, 10, 10)
	light.Radius = 200

	light = in3D.NewLight()
	light.Position = in3D.NewPosition(5, 10, 1)
	light.Radius = 1000

	model := "trex"
	in3D.SetDirPath("github.com/seemywingz/in3D/examples/assets/models/" + model)
	// all models are from: https://www.blendswap.com/
	mesh := in3D.LoadObject(model + ".obj")
	obj := in3D.NewMeshObject(in3D.Position{}, mesh, in3D.Shader["phong"])
	obj.SceneLogic = func(s *in3D.SceneData) {
		s.YRotation++
	}
	objects = append(objects, obj)

	for !in3D.ShouldClose() {
		in3D.Update()
		for _, o := range objects {
			o.Draw()
		}
		in3D.SwapBuffers()
	}
}
