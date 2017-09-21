package gg

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// DrawnObject : interface for opengl drawable object
type DrawnObject interface {
	Draw()
}

// DrawLogic : extra logic to perform durring DrawnObject Draw phase
type DrawLogic func(d *DrawnObjectData)

// DrawnObjectData : a struct to hold openGL object data
type DrawnObjectData struct {
	Vao     uint32
	Program uint32
	Points  *[]float32
	Position
	MVPID          int32
	ModelMatrixID  int32
	NormalMatrixID int32
	ColorID        int32
	Color          Color
	Texture        uint32
	DrawnObjectDefaults
}

// DrawnObjectDefaults :
type DrawnObjectDefaults struct {
	XRotation float32
	YRotation float32
	DrawLogic DrawLogic
}

// NewDrawnObject : Create new DrawnObjectData
func NewDrawnObject(position Position, points []float32, texture uint32, program uint32) *DrawnObjectData {

	ModelMatrixID := gl.GetUniformLocation(program, gl.Str("MODEL\x00"))
	NormalMatrixID := gl.GetUniformLocation(program, gl.Str("NormalMatrix\x00"))
	MVPID := gl.GetUniformLocation(program, gl.Str("MVP\x00"))
	ColorID := gl.GetUniformLocation(program, gl.Str("COLOR\x00"))

	return &DrawnObjectData{
		makeVao(points, program),
		program,
		&points,
		position,
		MVPID,
		ModelMatrixID,
		NormalMatrixID,
		ColorID,
		NewColor(1, 1, 1, 1),
		texture,
		DrawnObjectDefaults{},
	}
}

func (d *DrawnObjectData) translateRotate() *mgl32.Mat4 {
	model := mgl32.Translate3D(d.X, d.Y, d.Z)
	yrotMatrix := mgl32.HomogRotate3DY(mgl32.DegToRad(d.YRotation))
	xrotMatrix := mgl32.HomogRotate3DX(mgl32.DegToRad(d.XRotation))
	rotation := model.Mul4(xrotMatrix.Mul4(yrotMatrix))
	return &rotation
}

// Draw : draw the object
func (d *DrawnObjectData) Draw() {

	if d.DrawLogic != nil {
		d.DrawLogic(d)
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
		// gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, d.Texture)
	}
	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
	gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.Disable(gl.TEXTURE_2D)

}
