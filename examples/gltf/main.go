package main

import (
	in3d "in3D"
)

func main() {

	var objects []*in3d.DrawnObject

	in3d.Init(800, 600, "GLTF Loader")
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
	light.Radius = 100

	dx := float32(0.01)
	n := float32(0)
	light.SceneLogic = func(s *in3d.SceneData) {
		n += dx
		if n > 2 || n < -2 {
			dx = -dx
		}
		s.Position.Z = n
	}

	// model := "baofeng"
	// meshShader := in3d.Shader["in3D"]
	// in3d.SetDir("../assets/models/")
	// mesh := in3d.LoadGLTF(model, meshShader)
	// meshObject := in3d.NewMeshObject(in3d.NewPosition(-0.5, 0, 0), mesh, meshShader)
	// meshObject.Scale = 1.0
	// meshObject.YRotation = 90
	// objects = append(objects, meshObject)

	meshShader := in3d.Shader["in3D"]
	in3d.SetDir("../assets/models/buddha")
	mesh := in3d.LoadOBJ("buddha.obj", meshShader)
	meshObject := in3d.NewMeshObject(in3d.NewPosition(-0.5, 0, 0), mesh, meshShader)
	meshObject.YRotation = 90
	objects = append(objects, meshObject)

	for !in3d.ShouldClose() {
		in3d.Update()
		in3d.DrawObjects(objects)
		in3d.SwapBuffers()
	}
}
