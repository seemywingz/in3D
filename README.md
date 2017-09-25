# GG
A relatively simple Go powered OpenGL Graphics Library

Create a new Window, Get OpenGL Context, Setup Camera Projection, create 3D triangle, Draw!  
Go Ahead, you can do it yourself...
`go get github.com/seemywingz/gg`
```go
package main

import (
	"github.com/seemywingz/gg"
)

func main() {

  // Initialize gg
	gg.Init(800, 600, "Gopher Loader")
	gg.SetCameraPosition(gg.NewPosition(0, 0, 5))

  // Create a new light for the scene
	light := gg.NewLight()
	light.Position = gg.NewPosition(-10, 10, 10)
	light.Radius = 30

  // Create a gopher
	gg.SetDirPath("github.com/seemywingz/gge/assets")// this repo's location
	// Load mesh from assets obj file
	gopherMesh := gg.LoadObject("models/gopher.obj")
	// Create a new gg mesh object from the loaded obj file
	gopher := gg.NewMeshObject(gg.Position{}, gopherMesh, gg.NoTexture, gg.Shader["phong"])
	// Some tweaking -- not needed for every mesh
	gopher.ZRotation = -90
	// What to on draw tick
	gopher.SceneLogic = func(s *gg.SceneData) {
		s.YRotation++
	}

  // While the window is open
	for !gg.ShouldClose() {
		// Update all the things
		gg.Update()
		// Draw the gopher
		gopher.Draw()
		// Swap the buffers
		gg.SwapBuffers()
	}
}
```
### METODO:
Add material loading from .mtl files
### YOUTODO:
Checkout the other examples to see some more basic functionality

##### Note:
Some Names and method may change until version 1.0 is tagged
