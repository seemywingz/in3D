package main

import (
	"github.com/seemywingz/in3D"
)

func main() {

	var objects []*in3D.DrawnObject

	in3D.Init(800, 600, "Wavefront Loader")
	// in3D.SetClearColor(0.1, 0.1, 0.1, 1)
	in3D.SetCameraPosition(in3D.NewPosition(0, 0.55, 2))
	in3D.SetCameraSpeed(0.01)
	in3D.Enable(in3D.PointerLock, true)
	in3D.Enable(in3D.FlyMode, true)

	light := in3D.NewLight()
	light.Position = in3D.NewPosition(0, 1, 1)
	light.Ambient = []float32{0.5, 0.5, 0.5}
	light.Specular = []float32{10, 10, 10}
	light.Draw = true
	light.Radius = 10

	model := "sky"
	skyShader := in3D.Shader["texture"]
	in3D.SetRelPath("../assets/models/" + model)
	skymesh := in3D.LoadObject(model+".obj", skyShader)
	sky := in3D.NewMeshObject(in3D.Position{}, skymesh, skyShader)
	sky.Scale = 10000
	objects = append(objects, sky)

	// all models are from: https://www.blendswap.com/
	model = "buddha"
	in3D.SetRelPath("../assets/models/" + model)

	meshShader := in3D.Shader["phong"]
	mesh := in3D.LoadObject(model+".obj", meshShader)
	buddha := in3D.NewMeshObject(in3D.NewPosition(-0.5, 0, 0), mesh, meshShader)

	meshShader = in3D.Shader["normalMap"]
	mesh = in3D.LoadObject(model+".obj", meshShader)
	buddha1 := in3D.NewMeshObject(in3D.NewPosition(0.5, 0, 0), mesh, meshShader)

	rotate := func(s *in3D.SceneData) {
		s.YRotation += 0.5
	}

	buddha.SceneLogic = rotate
	buddha1.SceneLogic = rotate
	objects = append(objects, buddha)
	objects = append(objects, buddha1)

	for !in3D.ShouldClose() {
		in3D.Update()
		for _, o := range objects {
			o.Draw()
		}
		in3D.SwapBuffers()
	}
}
