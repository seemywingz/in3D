package in3d

import (
	"fmt"
	_ "image/jpeg" // include jpeg support
	_ "image/png"  // include png support
	"os"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/seemywingz/go-toolbox"
)

// Init : initializes glfw and returns a Window to use, then initGL
func Init(width, height int, title string) {
	runtime.LockOSThread()

	toolbox.EoE(glfw.Init(), "Error Initializing GLFW")
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	var err error
	if width == 0 {
		mode := glfw.GetPrimaryMonitor().GetVideoMode()
		width = mode.Width
		height = mode.Height
		Window, err = glfw.CreateWindow(width, height, title, glfw.GetPrimaryMonitor(), nil)
	} else {
		Window, err = glfw.CreateWindow(width, height, title, nil, nil)
	}
	toolbox.EoE(err, "Error Creating GLFW Window")
	Window.MakeContextCurrent()
	initGL()
	initShaders()
	NewCamera()
	NewLightManager()
}

// initGL : initialize GL setting and print version
func initGL() {
	toolbox.EoE(gl.Init(), "Error Initializing OpenGL")

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	gl.Enable(gl.CULL_FACE) // Enable Backface Culling
	gl.CullFace(gl.BACK)    // Cull Backfaces

	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.BLEND)

	version := gl.GoStr(gl.GetString(gl.VERSION))
	println("OpenGL version", version)
}

// initShaders :
func initShaders() {
	toolbox.SetRelPath("shaders")
	Shader["basic"] = NewShader("Vert.glsl", "basicFrag.glsl")
	Shader["color"] = NewShader("Vert.glsl", "colorFrag.glsl")
	Shader["texture"] = NewShader("Vert.glsl", "textureFrag.glsl")
	Shader["phong"] = NewShader("Vert.glsl", "blinnPhongFrag.glsl")
	Shader["normalMap"] = NewShader("normalMapVert.glsl", "normalMapFrag.glsl")
	Shader["in3D"] = NewShader("in3dVert.glsl", "in3DFrag.glsl")
}

// MakeVAO : initializes and returns a vertex array from the points provided.
func MakeVAO(points []float32, program uint32) uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(points)*4, gl.Ptr(points), gl.STATIC_DRAW)

	vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointerWithOffset(vertAttrib, 3, gl.FLOAT, false, 11*4, 0)

	vertTexCoordAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(vertTexCoordAttrib)
	gl.VertexAttribPointerWithOffset(vertTexCoordAttrib, 2, gl.FLOAT, false, 11*4, 3*4)

	vertNormalAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertNormal\x00")))
	gl.EnableVertexAttribArray(vertNormalAttrib)
	gl.VertexAttribPointerWithOffset(vertNormalAttrib, 3, gl.FLOAT, true, 11*4, 5*4)

	vertTangentAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertTangent\x00")))
	gl.EnableVertexAttribArray(vertTangentAttrib)
	gl.VertexAttribPointerWithOffset(vertTangentAttrib, 3, gl.FLOAT, true, 11*4, 8*4)

	return vao
}

// CompileShader :
func CompileShader(source string, shaderType uint32) uint32 {
	shader := gl.CreateShader(shaderType)

	sources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, sources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		toolbox.EoE(fmt.Errorf("failed to compile %v: %v", source, log), "Failed to Compile Source ")
	}

	return shader
}

// CompileShaderFromFile : create gl shader from source string
func CompileShaderFromFile(sourceFile string, shaderType uint32) uint32 {

	source, err := os.ReadFile(sourceFile)
	toolbox.EoE(err, "Error Reading Source File")

	return CompileShader(string(source)+"\x00", shaderType)
}

// NewShader : create GL shader program from provided GLSL source files
func NewShader(vertexShaderSourceFile, fragmentShaderSourceFile string) uint32 {
	vertexShader := CompileShaderFromFile(vertexShaderSourceFile, gl.VERTEX_SHADER)
	fragmentShader := CompileShaderFromFile(fragmentShaderSourceFile, gl.FRAGMENT_SHADER)

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)

	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		fmt.Println("Error Linking Shader Program:\n", log)
		os.Exit(1)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
	return program
}

func DrawObjects(objects []*DrawnObject) {
	for _, object := range objects {
		object.Draw()
	}
}

// ShouldClose : wrapper for glfw
func ShouldClose() bool {
	return Window.ShouldClose()
}

// Exit : Initiate GLFW Shutdown Sequence
func Exit() {
	Window.SetShouldClose(true)
}

// SwapBuffers : wrapper for glfw
func SwapBuffers() {
	Window.SwapBuffers()
}

// Update : Update OpenGL Scene by apply camera, then light object models
func Update() {
	camera.Update()
	lightManager.Update()
}

// GetCamera : return pointer to gg camera
func GetCamera() *Camera {
	return camera
}

// GetWindow : return pointer to gg Window
func GetWindow() *glfw.Window {
	return Window
}

// SetCameraPosition :
func SetCameraPosition(position Position) {
	camera.Position = position
}

// SetCameraSpeed :
func SetCameraSpeed(speed float32) {
	camera.Speed = speed
}

// SetClearColor :
func SetClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}

// TogglePointerLock :
func TogglePointerLock() {
	if Feature[PointerLock] {
		fmt.Println("PointerLock Enabled")
		x, y := Window.GetCursorPos()
		camera.LastX = x
		camera.LastY = y
		Window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	} else {
		fmt.Println("PointerLock Disabled")
		Window.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
	}
}

// MojaveWorkaround : move window 1 pixel as OpenGL workaround -- https://github.com/glfw/glfw/issues/1334
func MojaveWorkaround() {
	// macOS Mojave workaround
	x, y := Window.GetPos()
	Update()
	Window.SetPos(x+1, y)
}

// SetFlyModeControls : Set Default Key controls for FlyMode
func SetFlyModeControls() {
	KeyAction[KeyW] = camera.MoveForward
	KeyAction[KeyS] = camera.MoveBackward
	KeyAction[KeyA] = camera.StrafeLeft
	KeyAction[KeyD] = camera.StrafeRight
	KeyAction[KeySpace] = camera.Fly
}

// Enable :
func Enable(feature int, enabled bool) {

	Feature[feature] = enabled
	switch feature {
	case MouseControls:
		x, y := Window.GetCursorPos()
		camera.LastX = x
		camera.LastY = y
	case KeyControls:
		SetFlyModeControls()
	case PointerLock:
		TogglePointerLock()
	case FlyMode:
		Feature[MouseControls] = enabled
		Feature[KeyControls] = enabled
		SetFlyModeControls()
	}
}
