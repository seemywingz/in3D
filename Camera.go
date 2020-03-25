package in3D

import (
	"math"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

// Camera : struct to store camera matrices
type Camera struct {
	Projection mgl32.Mat4
	CameraData
}

// CameraData : struct to hold CameraData
type CameraData struct {
	XRotation float32
	YRotation float32
	LastX     float64
	LastY     float64
	MVP       mgl32.Mat4
	Speed     float32
	Position
}

// NewCamera : return new Camera
func NewCamera() *Camera {

	// Projection matrix :
	//    45° Field of View,
	//    width:height ratio,
	//    display range : 0.1 unit <-> 1000 units
	w, h := Window.GetSize()
	ratio := float32(w) / float32(h)
	projection := mgl32.Perspective(mgl32.DegToRad(45.0), ratio, 0.1, 11000)

	// Create new Camera instance
	camera = &Camera{projection, CameraData{}}
	// camera.Speed = 0.1
	camera.Speed = 1
	return camera
}

// MouseControls : control the camera via the mouse
func (c *Camera) MouseControls() {
	glfw.PollEvents()

	if Feature[Look] {
		x, y := Window.GetCursorPos()

		sensitivity := float32(0.1)
		c.YRotation += -float32(c.LastX-x) * sensitivity
		c.XRotation += -float32(c.LastY-y) * sensitivity

		xmax := float32(90)
		if c.XRotation < -xmax {
			c.XRotation = -xmax
		}
		if c.XRotation > xmax {
			c.XRotation = xmax
		}

		c.LastX = x
		c.LastY = y
	}
}

// KeyControls : control the camera via the keyboard
func (c *Camera) KeyControls() {
	if !Feature[KeyControls] {
		return
	}
	// Press w
	if Window.GetKey(glfw.KeyW) == glfw.Press {
		// move forward
		c.X += float32(math.Sin(float64(mgl32.DegToRad(c.YRotation)))) * c.Speed
		c.Z -= float32(math.Cos(float64(mgl32.DegToRad(c.YRotation)))) * c.Speed
		c.Y -= float32(math.Sin(float64(mgl32.DegToRad(c.XRotation)))) * c.Speed
	}
	// Press A
	if Window.GetKey(glfw.KeyA) == glfw.Press {
		// Move left
		c.X -= float32(math.Cos(float64(mgl32.DegToRad(c.YRotation)))) * c.Speed
		c.Z -= float32(math.Sin(float64(mgl32.DegToRad(c.YRotation)))) * c.Speed
	}
	// Press s
	if Window.GetKey(glfw.KeyS) == glfw.Press {
		// Move Backward
		c.X -= float32(math.Sin(float64(mgl32.DegToRad(c.YRotation)))) * c.Speed
		c.Z += float32(math.Cos(float64(mgl32.DegToRad(c.YRotation)))) * c.Speed
		c.Y += float32(math.Sin(float64(mgl32.DegToRad(c.XRotation)))) * c.Speed
	}
	// Press d
	if Window.GetKey(glfw.KeyD) == glfw.Press {
		// Move Right
		c.X += float32(math.Cos(float64(mgl32.DegToRad(c.YRotation)))) * c.Speed
		c.Z += float32(math.Sin(float64(mgl32.DegToRad(c.YRotation)))) * c.Speed
	}
	// Press space
	if Window.GetKey(glfw.KeySpace) == glfw.Press {
		if Feature[FlyMode] {
			if Window.GetKey(glfw.KeyLeftShift) == glfw.Press {
				c.Y -= c.Speed
			} else {
				c.Y += c.Speed
			}
		} else {
			//jump
		}
	}
}

// Update : update camera
func (c *Camera) Update() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	c.MouseControls()
	c.KeyControls()

	modelMatrix := mgl32.Translate3D(-c.X, -c.Y, -c.Z)
	// modelMatrix := translateMatrix.Mul4(mgl32.Ident4())

	xrotMatrix := mgl32.HomogRotate3DX(mgl32.DegToRad(c.XRotation))
	yrotMatrix := mgl32.HomogRotate3DY(mgl32.DegToRad(c.YRotation))
	view := xrotMatrix.Mul4(yrotMatrix.Mul4(mgl32.Ident4()))
	c.MVP = c.Projection.Mul4(view.Mul4(modelMatrix))
}
