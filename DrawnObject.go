package in3D

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// DrawnObject : a struct to hold openGL object data
type DrawnObject struct {
	Mesh           *Mesh
	MVPID          int32
	ModelMatrixID  int32
	NormalMatrixID int32
	IambID         int32
	IdifID         int32
	IspecID        int32
	ShininessID    int32
	Scale          float32
	SceneData
}

// NewPointsObject :
func NewPointsObject(position Position, points []float32, texture uint32, color []float32, program uint32) *DrawnObject {
	vao := MakeVAO(points, program)
	mat := &Material{
		"default",
		color,
		color,
		color,
		1,
		texture,
	}
	mg := make(map[string]*MaterialGroup)
	mg["dfault"] = &MaterialGroup{
		mat,
		[]*Face{},
		vao,
		int32(len(points)),
	}
	mesh := &Mesh{mg}
	return NewMeshObject(position, mesh, program)
}

// NewMeshObject : Create new DrawnObject
func NewMeshObject(position Position, mesh *Mesh, program uint32) *DrawnObject {

	ModelMatrixID := gl.GetUniformLocation(program, gl.Str("MODEL\x00"))
	NormalMatrixID := gl.GetUniformLocation(program, gl.Str("NormalMatrix\x00"))
	MVPID := gl.GetUniformLocation(program, gl.Str("MVP\x00"))

	uniform := fmt.Sprintf("Material")
	IambID := gl.GetUniformLocation(program, gl.Str(uniform+".Iamb\x00"))
	IdifID := gl.GetUniformLocation(program, gl.Str(uniform+".Idif\x00"))
	IspecID := gl.GetUniformLocation(program, gl.Str(uniform+".Ispec\x00"))
	ShininessID := gl.GetUniformLocation(program, gl.Str(uniform+".shininess\x00"))

	d := &DrawnObject{
		mesh,
		MVPID,
		ModelMatrixID,
		NormalMatrixID,
		IambID,
		IdifID,
		IspecID,
		ShininessID,
		1,
		SceneData{},
	}
	d.Position = position
	d.Program = program
	return d
}

func (d *DrawnObject) translateRotate() *mgl32.Mat4 {
	model := mgl32.Translate3D(d.X, d.Y, d.Z).
		Mul4(mgl32.Scale3D(d.Scale, d.Scale, d.Scale))
	xrotMatrix := mgl32.HomogRotate3DX(mgl32.DegToRad(d.XRotation))
	yrotMatrix := mgl32.HomogRotate3DY(mgl32.DegToRad(d.YRotation))
	zrotMatrix := mgl32.HomogRotate3DZ(mgl32.DegToRad(d.ZRotation))
	final := model.Mul4(xrotMatrix.Mul4(yrotMatrix.Mul4(zrotMatrix)))
	return &final
}

// Draw : draw the object
func (d *DrawnObject) Draw() {

	if d.SceneLogic != nil {
		d.SceneLogic(&d.SceneData)
	}

	modelMatrix := d.translateRotate()
	normalMatrix := modelMatrix.Inv().Transpose()

	gl.UseProgram(d.Program)
	gl.UniformMatrix4fv(d.MVPID, 1, false, &camera.MVP[0])
	gl.UniformMatrix4fv(d.ModelMatrixID, 1, false, &modelMatrix[0])
	gl.UniformMatrix4fv(d.NormalMatrixID, 1, false, &normalMatrix[0])

	for _, m := range d.Mesh.MaterialGroups {

		// Material
		gl.Uniform3fv(d.IambID, 1, &m.Material.Ambient[0])
		gl.Uniform3fv(d.IspecID, 1, &m.Material.Specular[0])
		gl.Uniform3fv(d.IdifID, 1, &m.Material.Diffuse[0])
		gl.Uniform1f(d.ShininessID, m.Material.Shininess)

		gl.BindVertexArray(m.VAO)
		gl.Enable(gl.TEXTURE_2D)
		gl.BindTexture(gl.TEXTURE_2D, m.Material.Texture)

		gl.DrawArrays(gl.TRIANGLES, 0, m.VertCount)
		gl.BindTexture(gl.TEXTURE_2D, 0)
	}

}
