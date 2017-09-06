package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Camera : struct to store camera matrices
type Camera struct {
	Projection mgl32.Mat4
	Model      mgl32.Mat4
	View       mgl32.Mat4
	MVP        mgl32.Mat4
	MVPID      int32
	Position
}

// Update : update camera
func (c *Camera) Update() {
	// c.Z -= 0.1
	// c.X -= 0.1

	translateMatrix := mgl32.Translate3D(c.X, c.Y, c.Z)
	model := translateMatrix.Mul4(c.Model)
	// TODO: update view from camera rotation
	c.MVP = c.Projection.Mul4(c.View.Mul4(model))
	gl.UniformMatrix4fv(0, 1, false, &c.MVP[0])
}

// New : return new Camera
func (c Camera) New(position Position) Camera {

	mvPointer, free := gl.Strs("MVP")
	defer free()
	mvpid := gl.GetUniformLocation(shaders[0], *mvPointer)

	//Projection matrix : 45° Field of View, width:height ratio, display range : 0.1 unit <-> 100 units
	projection := mgl32.Perspective(mgl32.DegToRad(45.0), width/height, 0.1, 100)
	//model matrix : and identity matrix (model will be at te origin)
	model := mgl32.Ident4()
	view := mgl32.Ident4()
	// view := mgl32.LookAt(
	// 	c.X, c.Y, c.Z, //Camera is at (x, y, z), in world space
	// 	0, 0, 0, //and looks at the origin
	// 	0, 1, 0, //head is up (set to 0, -1, 0 to look upside-down)
	// )
	mvp := projection.Mul4(view.Mul4(model))

	return Camera{projection, model, view, mvp, mvpid, position}
}
