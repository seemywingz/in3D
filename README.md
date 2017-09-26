# GG
A relatively simple Go powered OpenGL Graphics Engine

Create a new Window, Get OpenGL Context, Setup Camera Projection,  
create 3D gopher from an .obj file, Draw!  
Go Ahead, you can do it yourself...
`go get github.com/seemywingz/gg`
```go
package main

import (
	"github.com/seemywingz/gg"
)

func main() {

	var objects []*gg.DrawnObject

	gg.Init(800, 600, "Wavefront Loader")
	gg.SetCameraPosition(gg.NewPosition(0, 0.55, 2))
	// gg.Enable(gg.PointerLock, true)
	// gg.Enable(gg.FlyMode, true)

	light := gg.NewLight()
	light.Position = gg.NewPosition(-10, 10, 10)
	light.Radius = 20

	light = gg.NewLight()
	light.Position = gg.NewPosition(5, 1, 1)
	light.Radius = 100

	model := "buddha"
	gg.SetDirPath("github.com/seemywingz/gg/examples/assets/models/" + model)
	// all models are from: https://www.blendswap.com/
	mesh := gg.LoadObject(model + ".obj")
	obj := gg.NewMeshObject(gg.Position{}, mesh, gg.Shader["phong"])
	obj.SceneLogic = func(s *gg.SceneData) {
		s.YRotation++
	}
	objects = append(objects, obj)

	for !gg.ShouldClose() {
		gg.Update()
		for _, o := range objects {
			o.Draw()
		}
		gg.SwapBuffers()
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

