package main

import (
	in3d "github.com/seemywingz/in3D"
)

func main() {

	var objects []*in3d.DrawnObject

	in3d.Init(800, 600, "Wavefront Loader")
	in3d.SetClearColor(0.1, 0.1, 0.1, 1)
	in3d.SetCameraPosition(in3d.NewPosition(0, 0.55, 2))
	in3d.SetCameraSpeed(0.1)
	in3d.Enable(in3d.PointerLock, true)
	in3d.Enable(in3d.FlyMode, true)

	light := in3d.NewLight()
	// light.Specular = []float32{50, 50, 50}
	light.Position = in3d.NewPosition(0, 1, 1)
	light.Draw = true
	light.DrawnObject.Scale = 0.05
	light.Radius = 10

	dx := float32(0.01)
	n := float32(0)
	light.SceneLogic = func(s *in3d.SceneData) {
		n += dx
		max := float32(5)
		if n > max || n < -max {
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

	sky.SceneLogic = func(s *in3d.SceneData) {
		s.YRotation += 0.01
	}

	// all models are from: https://www.blendswap.com/
	model = "buddha"
	meshShader := in3d.Shader["in3d"]
	in3d.SetRelPath("../assets/models/" + model)
	mesh := in3d.LoadObject("buddha.obj", meshShader)
	meshObject := in3d.NewMeshObject(in3d.NewPosition(-0.5, 0, 0), mesh, meshShader)
	meshObject.YRotation = 90
	objects = append(objects, meshObject)

	meshShader = in3d.Shader["phong"]
	mesh = in3d.LoadObject("buddha.obj", meshShader)
	buddha := in3d.NewMeshObject(in3d.NewPosition(0.5, 0, 0), mesh, meshShader)
	buddha.YRotation = -90
	objects = append(objects, buddha)

	plane := in3d.NewPointsObject(
		in3d.NewPosition(0, 0, 0),
		in3d.Plane, in3d.NoTexture,
		[]float32{1, 1, 1},
		in3d.Shader["phong"])
	plane.XRotation = -90
	plane.Scale = 500
	objects = append(objects, plane)

	for !in3d.ShouldClose() {
		in3d.Update()
		for _, object := range objects {
			object.Draw()
		}
		in3d.SwapBuffers()
	}
}
