package main

import (
	"github.com/seemywingz/in3D"
)

func main() {

	var objects []*in3D.DrawnObject

	in3D.Init(800, 600, "Wavefront Loader")
	// in3D.SetClearColor(float32(1.0), float32(1.0), float32(1.0), float32(1.0))
	in3D.SetCameraPosition(in3D.NewPosition(0, 2.5, 5))
	in3D.SetCameraSpeed(0.1)
	in3D.Enable(in3D.PointerLock, true)
	in3D.Enable(in3D.FlyMode, true)

	light := in3D.NewLight()
	light.Position = in3D.NewPosition(0, 1, 3)
	light.Ambient = []float32{0.7, 0.7, 0.7}
	light.Difffuse = []float32{0.5, 0.5, 0.5}
	light.Specular = []float32{1, 1, 1}
	light.Draw = true
	light.Radius = 10

	light = in3D.NewLight()
	light.Position = in3D.NewPosition(100, 100, 3)
	light.Ambient = []float32{1, 1, 1}
	light.Difffuse = []float32{1, 1, 1}
	light.Specular = []float32{1, 1, 1}
	light.Draw = true
	light.Radius = 10000

	// all models are from: https://www.blendswap.com/
	model := "elf"
	in3D.SetRelPath("../assets/models/" + model)
	meshShader := in3D.Shader["phong"]
	mesh := in3D.LoadObject(model+".obj", meshShader)
	meshObject := in3D.NewMeshObject(in3D.NewPosition(0, 0, 0), mesh, meshShader)
	// meshObject.Scale = 0.10
	meshObject.SceneLogic = func(s *in3D.SceneData) {
		s.YRotation += 0.5
	}
	objects = append(objects, meshObject)

	for !in3D.ShouldClose() {
		in3D.Update()
		for _, o := range objects {
			o.Draw()
		}
		in3D.SwapBuffers()
	}
}
