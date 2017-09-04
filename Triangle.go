package main

// Triangle : a struct to hold openGL triangle data
type Triangle struct {
	Points []float32
	Vao    uint32
}

var points = []float32{
	0, 0.5, 0,
	-0.5, -0.5, 0,
	0.5, -0.5, 0,
}

// New : Create a new Triangle
func (t Triangle) New() Triangle {
	return Triangle{points, makeVao(points)}
}
