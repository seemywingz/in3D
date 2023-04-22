package main

import (
	"os"

	in3d "github.com/seemywingz/in3D"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {

	var objects []*in3d.DrawnObject

	in3d.Init(800, 600, "Wavefront Loader")
	window := in3d.GetWindow()
	// in3d.SetClearColor(float32(1.0), float32(1.0), float32(1.0), float32(1.0))
	in3d.SetCameraPosition(in3d.NewPosition(0, 2.5, 5))
	in3d.SetCameraSpeed(0.1)
	in3d.Enable(in3d.PointerLock, true)
	in3d.Enable(in3d.FlyMode, true)

	sun := in3d.NewLight()
	sun.Position = in3d.NewPosition(100, 100, 3)
	sun.Ambient = []float32{1, 1, 1}
	sun.Diffuse = []float32{1, 1, 1}
	sun.Specular = []float32{1, 1, 1}
	sun.Draw = true
	sun.Radius = 10000

	light := in3d.NewLight()
	light.Position = in3d.NewPosition(0, 0, 1.5)
	light.Ambient = []float32{0.7, 0.7, 0.7}
	light.Diffuse = []float32{10, 10, 10}
	light.Specular = []float32{10, 10, 10}
	light.Draw = true
	light.Radius = 10

	// all models are from: https://www.blendswap.com/
	model := "crate"
	in3d.SetRelPath("../assets/models/" + model)
	meshShader := in3d.Shader["normalMap"]
	mesh := in3d.LoadObject(model+".obj", meshShader)
	meshObject := in3d.NewMeshObject(in3d.NewPosition(0, 0, 0), mesh, meshShader)
	meshObject.Scale = 0.5
	meshObject.SceneLogic = func(s *in3d.SceneData) {
		s.YRotation += 0.5
	}
	objects = append(objects, meshObject)

	for !in3d.ShouldClose() {
		in3d.Update()
		for _, o := range objects {
			o.Draw()
		}
		if window.GetKey(glfw.KeyEscape) == in3d.Press {
			os.Exit(0)
		}
		in3d.SwapBuffers()
	}
}
