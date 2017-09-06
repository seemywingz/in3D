package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	z    float32
	view = mgl32.LookAt(
		0, 0, 1, //Camera is at (0, 0, -1), in world space
		0, 0, 0, //and looks at the origin
		0, 1, 0, //head is up (set to 0, -1, 0 to look upside-down)
	)
)

// Camera : struct to store camera matrices
type Camera struct {
	Projection mgl32.Mat4
	Model      mgl32.Mat4
	MVP        mgl32.Mat4
	MVPID      int32
}

// Draw : makes camera a DrawnObject interface and updates each tick
func (c Camera) Draw() {
	// gl.UseProgram(shaders[0])
	z -= 0.01
	translateMatrix := mgl32.Translate3D(0.5, -0.2, z)
	model := translateMatrix.Mul4(c.Model)
	// c.MVP = c.Projection.Mul4(view.Mul4(model))
	c.MVP = c.Projection.Mul4(model)
	gl.UniformMatrix4fv(0, 1, false, &c.MVP[0])
}

// New : return new Camera
func (c Camera) New() Camera {

	mvPointer, free := gl.Strs("MVP")
	defer free()
	mvpid := gl.GetUniformLocation(shaders[0], *mvPointer)

	//Projection matrix : 45° Field of View, width:heigh ratio, display range : 0.1 unit <-> 100 units
	projection := mgl32.Perspective(mgl32.DegToRad(45.0), width/height, 0.1, 100)
	//model matrix : and identity matrix (model will be at te origin)
	model := mgl32.Ident4()
	//our ModelViewProjection : multiplication of our 3 matrices
	mvp := projection.Mul4(view.Mul4(model))

	return Camera{projection, model, mvp, mvpid}
}
