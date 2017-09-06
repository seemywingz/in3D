package main

import (
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
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

// Position : struct to store 3D coords
type Position struct {
	X float32
	Y float32
	Z float32
}

// MouseControls : control the camera via the mouse
func (c *Camera) MouseControls() {

	if c.PointerLock && window.GetMouseButton(glfw.MouseButton1) == glfw.Release {
		// fmt.Println("PointerLock Disabled")
		window.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
		c.PointerLock = false
	}

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
			// fmt.Println("PointerLock Enabled")
			x, y := window.GetCursorPos()
			c.LastX = x
			c.LastY = y
			window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
			c.PointerLock = true
		}
	}
}

// KeyControls : control the camera via the keyboard
func (c *Camera) KeyControls() {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		os.Exit(1)
	}
	// Press w
	if window.GetKey(glfw.KeyW) == glfw.Press {
		camera.Z += 0.1
	}
	// Press A
	if window.GetKey(glfw.KeyA) == glfw.Press {
	}
	// Press s
	if window.GetKey(glfw.KeyS) == glfw.Press {
		camera.Z -= 0.1
	}
	// Press d
	if window.GetKey(glfw.KeyD) == glfw.Press {
	}
	// Press q
	if window.GetKey(glfw.KeyQ) == glfw.Press {
	}
	// Press e
	if window.GetKey(glfw.KeyE) == glfw.Press {
	}
}

// Update : update camera
func (c *Camera) Update() {
	c.MouseControls()
	c.KeyControls()

	translateMatrix := mgl32.Translate3D(c.X, c.Y, c.Z)
	model := translateMatrix.Mul4(c.Model)

	xrotMatrix := mgl32.HomogRotate3D(mgl32.DegToRad(c.XRotation), mgl32.Vec3{1, 0, 0})
	yrotMatrix := mgl32.HomogRotate3D(mgl32.DegToRad(c.YRotation), mgl32.Vec3{0, 1, 0})
	c.View = xrotMatrix.Mul4(yrotMatrix.Mul4(c.Model))

	c.MVP = c.Projection.Mul4(c.View.Mul4(model))
	gl.UniformMatrix4fv(c.MVPID, 1, false, &c.MVP[0])
}

// New : return new Camera
func (Camera) New(position Position) Camera {

	mvPointer, free := gl.Strs("MVP")
	defer free()
	mvpid := gl.GetUniformLocation(shaders[0], *mvPointer)

	//Projection matrix : 45° Field of View, width:height ratio, display range : 0.1 unit <-> 1000 units
	projection := mgl32.Perspective(mgl32.DegToRad(45.0), width/height, 0.1, 1000)
	model := mgl32.Ident4()
	view := mgl32.LookAt(
		position.X, position.Y, position.Z, //Camera is at (x, y, z), in world space
		0, 0, 0, //and looks at the origin
		0, 1, 0, //head is up (set to 0, -1, 0 to look upside-down)
	)
	mvp := projection.Mul4(view.Mul4(model))

	cam := Camera{projection, model, view, mvp, mvpid, position, CameraData{}}
	x, y := window.GetCursorPos()
	cam.LastX = x
	cam.LastY = y

	return cam
}
