# in3D
A relatively simple Go powered OpenGL Graphics Engine

Create a new Window, Get OpenGL Context, Setup Camera Projection,  
create 3D Mesh from an .obj file, Draw!  
Go Ahead, you can do it yourself...
`go get github.com/seemywingz/in3D`
```go
package main

import "github.com/seemywingz/in3D"

func main() {

	var objects []*in3D.DrawnObject

	in3D.Init(800, 600, "Wavefront Loader")
	in3D.SetCameraPosition(in3D.NewPosition(0, 0.55, 2))
	// in3D.Enable(in3D.PointerLock, true)
	// in3D.Enable(in3D.FlyMode, true)

	light := in3D.NewLight()
	light.Position = in3D.NewPosition(-10, 10, 10)
	light.Radius = 20

	light = in3D.NewLight()
	light.Position = in3D.NewPosition(5, 1, 1)
	light.Radius = 100

	model := "buddha"
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

```
### ME-TODO:
  •  Optimize all the things!  
  •  Add Shadows, Ambient Occulsion and other light related things  
  • Have more fun making weird examples! 

### YOU-TODO:
Checkout the other examples to see some more basic functionality

##### Note:
###### Some Names and method may change until version 1.0 is tagged
###### Also, texture UVs are, for some reason, imported upside down. ( flip your texture vertiacally to render correctly  )

