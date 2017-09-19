package main

import "github.com/go-gl/glfw/v3.2/glfw"

// Position : struct to store 3D coords
type Position struct {
	X float32
	Y float32
	Z float32
}

// Logic : extra logic to perform durring DrawnObject Draw phase
type Logic func(d *DrawnObjectData)

var (

	// Graphics
	window       *glfw.Window
	camera       Camera
	shaders      map[string]uint32
	texture      map[string]uint32
	drawnObjects []DrawnObject

	// Shapes ....
	triangle = []float32{
		-1.0, -1.0, 0, 1.0, 0.0,
		1.0, -1.0, 0, 0.0, 0.0,
		-1.0, 1.0, 0, 1.0, 1.0,
	}

	square = []float32{
		//  X, Y, Z, U, V, normal(3)
		-1.0, -1.0, 0, 0.0, 1.0, 0.0, 0.0, 1.0,
		1.0, -1.0, 0, 1.0, 1.0, 0.0, 0.0, 1.0,
		-1.0, 1.0, 0, 0.0, 0.0, 0.0, 0.0, 1.0,

		-1.0, 1.0, 0, 0.0, 0.0, 0.0, 0.0, 1.0,
		1.0, 1.0, 0, 1.0, 0.0, 0.0, 0.0, 1.0,
		1.0, -1.0, 0, 1.0, 1.0, 0.0, 0.0, 1.0,
	}

	card = []float32{
		//  X, Y, Z, U, V, normal(3)
		-1.0, -1.75, 0, 0.0, 1.0, 0.0, 0.0, 1.0,
		1.0, -1.75, 0, 1.0, 1.0, 0.0, 0.0, 1.0,
		-1.0, 1.0, 0, 0.0, 0.0, 0.0, 0.0, 1.0,

		-1.0, 1.0, 0, 0.0, 0.0, 0.0, 0.0, 1.0,
		1.0, 1.0, 0, 1.0, 0.0, 0.0, 0.0, 1.0,
		1.0, -1.75, 0, 1.0, 1.0, 0.0, 0.0, 1.0,
	}

	cube = []float32{
		//  X, Y, Z, U, V, normal(3)
		// Bottom
		-1.0, -1.0, -1.0, 0.0, 0.0, 0.0, -1.0, 0.0,
		1.0, -1.0, -1.0, 1.0, 0.0, 0.0, -1.0, 0.0,
		-1.0, -1.0, 1.0, 0.0, 1.0, 0.0, -1.0, 0.0,
		1.0, -1.0, -1.0, 1.0, 0.0, 0.0, -1.0, 0.0,
		1.0, -1.0, 1.0, 1.0, 1.0, 0.0, -1.0, 0.0,
		-1.0, -1.0, 1.0, 0.0, 1.0, 0.0, -1.0, 0.0,

		// Top
		-1.0, 1.0, -1.0, 0.0, 0.0, 0.0, 1.0, 0.0,
		-1.0, 1.0, 1.0, 0.0, 1.0, 0.0, 1.0, 0.0,
		1.0, 1.0, -1.0, 1.0, 0.0, 0.0, 1.0, 0.0,
		1.0, 1.0, -1.0, 1.0, 0.0, 0.0, 1.0, 0.0,
		-1.0, 1.0, 1.0, 0.0, 1.0, 0.0, 1.0, 0.0,
		1.0, 1.0, 1.0, 1.0, 1.0, 0.0, 1.0, 0.0,

		// Front
		-1.0, -1.0, 1.0, 1.0, 0.0, 0.0, 0.0, 1.0,
		1.0, -1.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0,
		-1.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0.0, 1.0,
		1.0, -1.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0,
		1.0, 1.0, 1.0, 0.0, 1.0, 0.0, 0.0, 1.0,
		-1.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0.0, 1.0,

		// Back
		-1.0, -1.0, -1.0, 0.0, 0.0, 0.0, 0.0, -1.0,
		-1.0, 1.0, -1.0, 0.0, 1.0, 0.0, 0.0, -1.0,
		1.0, -1.0, -1.0, 1.0, 0.0, 0.0, 0.0, -1.0,
		1.0, -1.0, -1.0, 1.0, 0.0, 0.0, 0.0, -1.0,
		-1.0, 1.0, -1.0, 0.0, 1.0, 0.0, 0.0, -1.0,
		1.0, 1.0, -1.0, 1.0, 1.0, 0.0, 0.0, -1.0,

		// Left
		-1.0, -1.0, 1.0, 0.0, 1.0, -1.0, 0.0, 0.0,
		-1.0, 1.0, -1.0, 1.0, 0.0, -1.0, 0.0, 0.0,
		-1.0, -1.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0,
		-1.0, -1.0, 1.0, 0.0, 1.0, -1.0, 0.0, 0.0,
		-1.0, 1.0, 1.0, 1.0, 1.0, -1.0, 0.0, 0.0,
		-1.0, 1.0, -1.0, 1.0, 0.0, -1.0, 0.0, 0.0,

		// Right
		1.0, -1.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0.0,
		1.0, -1.0, -1.0, 1.0, 0.0, 1.0, 0.0, 0.0,
		1.0, 1.0, -1.0, 0.0, 0.0, 1.0, 0.0, 0.0,
		1.0, -1.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0.0,
		1.0, 1.0, -1.0, 0.0, 0.0, 1.0, 0.0, 0.0,
		1.0, 1.0, 1.0, 0.0, 1.0, 1.0, 0.0, 0.0,
	}
)