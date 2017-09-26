# GG
A relatively simple Go powered OpenGL Graphics Engine

Create a new Window, Get OpenGL Context, Setup Camera Projection, create 3D gopher, Draw!  
Go Ahead, you can do it yourself...
`go get github.com/seemywingz/gg`
```go
package main

import (
	"fmt"

	"github.com/seemywingz/gg"
)

func main() {

	gg.Init(800, 600, "Wavefront Loader")
	gg.SetCameraPosition(gg.NewPosition(0, 0, 5))

	light := gg.NewLight([]float32{1, 1, 1})
	light.Iamb = []float32{0.8, 0.8, 0.8}
	light.Position = gg.NewPosition(-10, 10, 10)
	light.Radius = 30

	gg.SetDirPath("github.com/seemywingz/gg/examples/assets/models/gopher")
	gopherMesh := gg.LoadObject("gopher.obj")
	gopher := gg.NewMeshObject(gg.Position{}, gopherMesh, gg.Shader["phong"])
	gopher.ZRotation = -90 // this .obj was exported sideways lol
	gopher.SceneLogic = func(s *gg.SceneData) {
		s.YRotation++
	}
	fmt.Println(len(gopher.Mesh.MaterialGroups))

	for !gg.ShouldClose() {
		gg.Update()
		gopher.Draw()
		gg.SwapBuffers()
	}
}

```
### METODO:
Update the Mesh loader to build textures from .mtl  
  
### YOUTODO:
Checkout the other examples to see some more basic functionality

#### Note:
##### Some Names and method may change until version 1.0 is tagged
##### Also, texture UVs are, for some reason, imported upside down. ( flip your texture vertiacally to render correctly  )

