package main

import (
	in3d "in3D"
)

func main() {

	var objects []*in3d.DrawnObject

	in3d.Init(800, 600, "Wavefront Loader")
	in3d.SetClearColor(0.1, 0.1, 0.1, 1)
	in3d.SetCameraPosition(in3d.NewPosition(0, 0.55, 2))
	in3d.SetCameraSpeed(0.1)
	in3d.Enable(in3d.PointerLock, true)
	in3d.Enable(in3d.FlyMode, true)
	// Close Window When Escape is Pressed
	in3d.KeyAction[in3d.KeyEscape] = func() {
		in3d.Window.SetShouldClose(true)
	}

	light := in3d.NewLight()
	// light.Specular = []float32{50, 50, 50}
	light.Position = in3d.NewPosition(0, 1, 1)
	light.Draw = true
	light.DrawnObject.Scale = 0.05
	light.Radius = 10

	// Close Window When Escape is Pressed
	in3d.KeyAction[in3d.KeyEscape] = func() {
		in3d.Exit()
	}

	dx := float32(0.01)
	n := float32(0)
	light.SceneLogic = func(s *in3d.SceneData) {
		n += dx
		if n > 2 || n < -2 {
			dx = -dx
		}
		s.Position.Z = n
	}

	model := "sky"
	skyShader := in3d.Shader["texture"]
	in3d.SetRelPath("../assets/models/" + model)
	skymesh := in3d.LoadObject(model+".obj", skyShader)
	sky := in3d.NewMeshObject(in3d.Position{}, skymesh, skyShader)
	sky.Scale = 10000
	objects = append(objects, sky)

	// rotateY := func(s *in3d.SceneData) {
	// 	s.YRotation += 0.1
	// }

	// all models are from: https://www.blendswap.com/
	model = "buddha"
	meshShader := in3d.Shader["normalMap"]
	in3d.SetRelPath("../assets/models/" + model)
	mesh := in3d.LoadObject(model+".obj", meshShader)
	meshObject := in3d.NewMeshObject(in3d.NewPosition(-0.5, 0, 0), mesh, meshShader)
	meshObject.YRotation = 90
	objects = append(objects, meshObject)

	model = "buddha"
	meshShader = in3d.Shader["phong"]
	mesh = in3d.LoadObject(model+".obj", meshShader)
	buddha := in3d.NewMeshObject(in3d.NewPosition(0.5, 0, 0), mesh, meshShader)
	buddha.YRotation = -90
	objects = append(objects, buddha)

	for !in3d.ShouldClose() {
		in3d.Update()
		for _, o := range objects {
			o.Draw()
		}
		in3d.SwapBuffers()
	}
}
