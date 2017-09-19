package main

import (
	"math"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

// Camera : struct to store camera matrices
type Camera struct {
	Projection mgl32.Mat4
	MVPID      int32
	POSID      int32
	Position
	CameraData
}

// CameraData : struct to hold CameraData
type CameraData struct {
	XRotation   float32
	YRotation   float32
	LastX       float64
	LastY       float64
	PointerLock bool
}

// MouseControls : control the camera via the mouse
func (c *Camera) MouseControls() {

	if c.PointerLock {
		x, y := window.GetCursorPos()

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
	} else { // no PointerLock
		if window.GetMouseButton(glfw.MouseButton1) == glfw.Press {
			c.EnablePointerLock()
		}
	}
}

// EnablePointerLock :
func (c *Camera) EnablePointerLock() {
	// fmt.Println("PointerLock Enabled")
	x, y := window.GetCursorPos()
	c.LastX = x
	c.LastY = y
	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	c.PointerLock = true
}

// DisablePointerLock :
func (c *Camera) DisablePointerLock() {
	window.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
	c.PointerLock = false
}

// KeyControls : control the camera via the keyboard
func (c *Camera) KeyControls() {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		// os.Exit(1)
		c.DisablePointerLock()
	}
	// Press w
	if window.GetKey(glfw.KeyW) == glfw.Press {
		// move forward
		c.X -= float32(math.Sin(float64(mgl32.DegToRad(c.YRotation))))
		c.Z += float32(math.Cos(float64(mgl32.DegToRad(c.YRotation))))
		c.Y += float32(math.Sin(float64(mgl32.DegToRad(c.XRotation))))
	}
	// Press A
	if window.GetKey(glfw.KeyA) == glfw.Press {
		// Move left
		c.X += float32(math.Cos(float64(mgl32.DegToRad(c.YRotation))))
		c.Z += float32(math.Sin(float64(mgl32.DegToRad(c.YRotation))))
	}
	// Press s
	if window.GetKey(glfw.KeyS) == glfw.Press {
		// Move Backward
		c.X += float32(math.Sin(float64(mgl32.DegToRad(c.YRotation))))
		c.Z -= float32(math.Cos(float64(mgl32.DegToRad(c.YRotation))))
		c.Y -= float32(math.Sin(float64(mgl32.DegToRad(c.XRotation))))
	}
	// Press d
	if window.GetKey(glfw.KeyD) == glfw.Press {
		// Move Right
		c.X -= float32(math.Cos(float64(mgl32.DegToRad(c.YRotation))))
		c.Z -= float32(math.Sin(float64(mgl32.DegToRad(c.YRotation))))
	}
	// Press space
	if window.GetKey(glfw.KeySpace) == glfw.Press {
		if window.GetKey(glfw.KeyLeftShift) == glfw.Press {
			c.Y++
		} else {
			c.Y--
		}
	}
	// Press q
	if window.GetKey(glfw.KeyQ) == glfw.Press {
	}
	// Press e
	if window.GetKey(glfw.KeyE) == glfw.Press {
	}
}

// New : return new Camera
func (Camera) New(position Position, pointerLock bool) Camera {

	mvpid := gl.GetUniformLocation(shader["phong"], gl.Str("MVP\x00"))
	posid := gl.GetUniformLocation(shader["phong"], gl.Str("CPOS\x00"))

	// Projection matrix : 45° Field of View, width:height ratio, display range : 0.1 unit <-> 1000 units
	w, h := window.GetSize()
	ratio := float32(w) / float32(h)
	projection := mgl32.Perspective(mgl32.DegToRad(45.0), ratio, 0.1, 1000)

	// Create new Camera instance
	cam := Camera{projection, mvpid, posid, position, CameraData{}}
	if pointerLock {
		cam.EnablePointerLock()
	} else {
		cam.DisablePointerLock()
	}

	return cam
}

// Update : update camera
func (c *Camera) Update() {
	c.MouseControls()
	c.KeyControls()

	translateMatrix := mgl32.Translate3D(c.X, c.Y, c.Z)
	model := translateMatrix.Mul4(mgl32.Ident4())

	xrotMatrix := mgl32.HomogRotate3DX(mgl32.DegToRad(c.XRotation))
	yrotMatrix := mgl32.HomogRotate3DY(mgl32.DegToRad(c.YRotation))
	view := xrotMatrix.Mul4(yrotMatrix.Mul4(mgl32.Ident4()))
	MVP := c.Projection.Mul4(view.Mul4(model))
	gl.UniformMatrix4fv(c.MVPID, 1, false, &MVP[0])
}
