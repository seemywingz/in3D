package gg

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg" // include jpeg support
	_ "image/png"  // include png support
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Init : initializes glfw and returns a Window to use, then InitGL
func Init(width, height int, pointerLock bool, title string) {
	runtime.LockOSThread()

	EoE("Error Initializing GLFW", glfw.Init())
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	var err error
	if width == 0 {
		mode := glfw.GetPrimaryMonitor().GetVideoMode()
		width = mode.Width
		height = mode.Height
		window, err = glfw.CreateWindow(width, height, title, glfw.GetPrimaryMonitor(), nil)
	} else {
		window, err = glfw.CreateWindow(width, height, title, nil, nil)
	}
	EoE("Error Creating GLFW Window", err)
	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	window.SetInputMode(glfw.StickyMouseButtonsMode, 1)
	window.MakeContextCurrent()
	InitGL()
	NewCamera(pointerLock)
}

// InitGL : initialize GL setting and print version
func InitGL() {
	EoE("Error Initializing OpenGL", gl.Init())

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	version := gl.GoStr(gl.GetString(gl.VERSION))
	println("OpenGL version", version)
	loadShaders()
}

func loadShaders() {
	Shader = make(map[string]uint32)
	SetDirPath("github.com/seemywingz/gg/shaders")
	Shader["basic"] = NewShader("basicVect.glsl", "basicFrag.glsl")
	Shader["color"] = NewShader("basicVect.glsl", "colorFrag.glsl")
	Shader["texture"] = NewShader("textureVect.glsl", "textureFrag.glsl")
	Shader["phong"] = NewShader("blinnPhongVect.glsl", "blinnPhongFrag.glsl")
}

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32, program uint32) uint32 {

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(points)*4, gl.Ptr(points), gl.STATIC_DRAW)

	vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))

	vertNormalAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertNormal\x00")))
	gl.EnableVertexAttribArray(vertNormalAttrib)
	gl.VertexAttribPointer(vertNormalAttrib, 3, gl.FLOAT, true, 8*4, gl.PtrOffset(5*4))

	return vao
}

// CompileShader :
func CompileShader(source string, shaderType uint32) uint32 {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		EoE("Failed to Compile Source ", fmt.Errorf("failed to compile %v: %v", source, log))
	}

	return shader
}

// CompileShaderFromFile : create gl shader from source string
func CompileShaderFromFile(sourceFile string, shaderType uint32) uint32 {

	source, err := ioutil.ReadFile(sourceFile)
	EoE("Error Reading Source File", err)

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

		EoE("Error Linking Shader Program", fmt.Errorf("failed to link program: %v", log))
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
	return program
}

// NewTexture : greate GL reference to provided texture
func NewTexture(file string) uint32 {
	imgFile, err := os.Open(file)
	EoE("Error Loading Texture", err)

	img, _, err := image.Decode(imgFile)
	EoE("Error Decoding Image", err)

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		EoE("Error Getting RGB Strride", errors.New("unsupported stride"))
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture
}

// ShouldClose : wraper for glfw
func ShouldClose() bool {
	return window.ShouldClose()
}

// SwapBuffers : wrapper for glfw
func SwapBuffers() {
	window.SwapBuffers()
}

// Update :
func Update() {
	camera.Update()
}

// GetCamera : return pounter to gg camera
func GetCamera() *Camera {
	return camera
}

// GetWindow : return pounter to gg window
func GetWindow() *glfw.Window {
	return window
}

// Enable :
func Enable(feature int, enabled bool) {
	switch feature {
	case LookEnabled:
		camera.LookEnabled = enabled
	case MoveEnabled:
		camera.MoveEnabled = enabled
	case FirstPersonMode:
		camera.LookEnabled = enabled
		camera.MoveEnabled = enabled
		if enabled {
			camera.Mode = FirstPersonMode
		} else {
			camera.Mode = NoMode
		}
	case PointerLock:
		if enabled {
			camera.EnablePointerLock()
		} else {
			camera.DisablePointerLock()
		}
	case FlyMode:
		camera.LookEnabled = enabled
		camera.MoveEnabled = enabled
		if enabled {
			camera.Mode = FlyMode
		} else {
			camera.Mode = NoMode
		}
	}
}
