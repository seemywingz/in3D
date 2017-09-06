package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

var (
	view = mgl32.LookAt(
		4, 3, 3, //Camera is at (4,3,3), in world space
		0, 0, 0, //and looks at the origin
		0, 1, 0, //head is up (set to 0, -1, 0 to look upside-down)
	)
)

// Camera : struct to store camera matrices
type Camera struct {
	Projection mgl32.Mat4
	Model      mgl32.Mat4
	MVP        mgl32.Mat4
}

// New : return new Camera
func (c Camera) New() Camera {
	//Projection matrix : 45° Field of View, 4:3 ratio, display range : 0.1 unit <-> 100 units
	projection := mgl32.Perspective(mgl32.DegToRad(45.0), 4.0/3.0, 0.1, 100)

	//model matrix : and identity matrix (model will be at te origin)
	model := mgl32.Ident4()

	//our ModelViewProjection : multiplication of our 3 matrices
	mvp := projection.Mul4(view.Mul4(model))

	return Camera{projection, model, mvp}
}
