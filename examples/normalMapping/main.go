package main

import (
	in3d "in3D"
)

func main() {

	var objects []*in3d.DrawnObject

	in3d.Init(800, 600, "Wavefront Loader")
	// in3d.SetClearColor(0.1, 0.1, 0.1, 1)
	in3d.SetCameraPosition(in3d.NewPosition(0, 0.55, 2))
	in3d.SetCameraSpeed(0.01)
	in3d.Enable(in3d.PointerLock, true)
	in3d.Enable(in3d.FlyMode, true)

	light := in3d.NewLight()
	light.Position = in3d.NewPosition(0, 1, 1)
	light.Ambient = []float32{0.5, 0.5, 0.5}
	light.Specular = []float32{10, 10, 10}
	light.Draw = true
	light.Radius = 10

	model := "sky"
	skyShader := in3d.Shader["texture"]
	in3d.SetRelPath("../assets/models/" + model)
	skymesh := in3d.LoadObject(model+".obj", skyShader)
	sky := in3d.NewMeshObject(in3d.Position{}, skymesh, skyShader)
	sky.Scale = 10000
	objects = append(objects, sky)

	// all models are from: https://www.blendswap.com/
	model = "buddha"
	in3d.SetRelPath("../assets/models/" + model)

	meshShader := in3d.Shader["phong"]
	mesh := in3d.LoadObject(model+".obj", meshShader)
	buddha := in3d.NewMeshObject(in3d.NewPosition(-0.5, 0, 0), mesh, meshShader)

	meshShader = in3d.Shader["normalMap"]
	mesh = in3d.LoadObject(model+".obj", meshShader)
	buddha1 := in3d.NewMeshObject(in3d.NewPosition(0.5, 0, 0), mesh, meshShader)

	rotate := func(s *in3d.SceneData) {
		s.YRotation += 0.5
	}

	buddha.SceneLogic = rotate
	buddha1.SceneLogic = rotate
	objects = append(objects, buddha)
	objects = append(objects, buddha1)

	for !in3d.ShouldClose() {
		in3d.Update()
		for _, o := range objects {
			o.Draw()
		}
		in3d.SwapBuffers()
	}
}
