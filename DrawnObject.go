package gg

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// DrawLogic : extra logic to perform durring SceneObject Draw phase
type DrawLogic func(d *DrawnObject)

// DrawnObject : a struct to hold openGL object data
type DrawnObject struct {
	Vao            uint32
	Points         *[]float32
	MVPID          int32
	ModelMatrixID  int32
	NormalMatrixID int32
	ColorID        int32
	Color          Color
	Texture        uint32
	Scale          float32
	StdData
}

// NewDrawnObject : Create new DrawnObject
func NewDrawnObject(position Position, points []float32, texture uint32, program uint32) *DrawnObject {

	ModelMatrixID := gl.GetUniformLocation(program, gl.Str("MODEL\x00"))
	NormalMatrixID := gl.GetUniformLocation(program, gl.Str("NormalMatrix\x00"))
	MVPID := gl.GetUniformLocation(program, gl.Str("MVP\x00"))
	ColorID := gl.GetUniformLocation(program, gl.Str("COLOR\x00"))

	d := &DrawnObject{
		makeVao(points, program),
		&points,
		MVPID,
		ModelMatrixID,
		NormalMatrixID,
		ColorID,
		NewColor(1, 1, 1, 1),
		texture,
		1,
		StdData{},
	}
	d.Position = position
	d.Program = program
	return d
}

func (d *DrawnObject) translateRotate() *mgl32.Mat4 {
	model := mgl32.Translate3D(d.X, d.Y, d.Z).
		Mul4(mgl32.Scale3D(d.Scale, d.Scale, d.Scale))
	yrotMatrix := mgl32.HomogRotate3DY(mgl32.DegToRad(d.YRotation))
	xrotMatrix := mgl32.HomogRotate3DX(mgl32.DegToRad(d.XRotation))
	rotation := model.Mul4(xrotMatrix.Mul4(yrotMatrix))
	return &rotation
}

// Draw : draw the object
func (d *DrawnObject) Draw() {

	if d.SceneLogic != nil {
		d.SceneLogic(&d.StdData)
	}

	modelMatrix := d.translateRotate()
	normalMatrix := modelMatrix.Inv()
	normalMatrix = normalMatrix.Transpose()

	gl.UseProgram(d.Program)
	gl.UniformMatrix4fv(d.MVPID, 1, false, &camera.MVP[0])
	gl.UniformMatrix4fv(d.ModelMatrixID, 1, false, &modelMatrix[0])
	gl.UniformMatrix4fv(d.NormalMatrixID, 1, false, &normalMatrix[0])
	gl.Uniform4f(d.ColorID, d.Color.R, d.Color.G, d.Color.B, d.Color.A)

	gl.BindVertexArray(d.Vao)
	if d.Texture != NoTexture {
		gl.Enable(gl.TEXTURE_2D)
		gl.BindTexture(gl.TEXTURE_2D, d.Texture)
	}

	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
	gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.Disable(gl.TEXTURE_2D)

}
