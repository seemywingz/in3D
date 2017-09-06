package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	gt "github.com/seemywingz/gtils"
)

// readShaderFile : read the shader file and return null terminated string
func readShaderFile(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	gt.EoE("Error Reading Shader File", err)
	return string(data) + "\x00"
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw(width, height int, title string) *glfw.Window {
	gt.EoE("Error Initializing GLFW", glfw.Init())
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	window.SetInputMode(glfw.StickyMouseButtonsMode, 1)

	gt.EoE("Error Creating GLFW Window", err)
	window.MakeContextCurrent()

	return window
}

func initGL() {
	gt.EoE("Error Initializing OpenGL", gl.Init())
	//enable depth test
	gl.Enable(gl.DEPTH_TEST)
	//accept fragment if it close to the camera than the former one
	gl.DepthFunc(gl.LESS)

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
}

func createGLprogram(vertexShaderSource, fragmentShaderSource string) uint32 {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	gt.EoE("Error Compiling Vertex Shader", err)

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	gt.EoE("Error Compiling Fragment Shader", err)

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)

	gl.LinkProgram(program)
	return program
}

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(points)*4, gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func compileShader(source string, shaderType uint32) (uint32, error) {
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

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
