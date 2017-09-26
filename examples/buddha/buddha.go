package main

import (
	"github.com/seemywingz/in3D"
)

func main() {

	var objects []*in3D.DrawnObject

	in3D.Init(800, 600, "Wavefront Loader")
	in3D.SetCameraPosition(in3D.NewPosition(0, 0.5, 2))

	l := in3D.NewLight()
	l.Position = in3D.NewPosition(-10, 10, 10)
	l.Radius = 300

	in3D.SetDirPath("github.com/seemywingz/in3D/examples/assets/models/buddha")
	mesh := in3D.LoadObject("buddha.obj")
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
