package main

const (
	basicFragmentSRC = `
	#version 410
	// Hard Code Color for Now
	uniform vec4 inputColour = vec4(1,1,1,1);
	out vec4 fragColour;

	void main() {
	  fragColour = inputColour;
	}` + "\x00"

	basicVertexSRC = `
	#version 410 core

	layout(location = 0) in vec3 pos;

	uniform mat4 MVP;
	uniform vec4 translation;
	uniform mat4 foo;

	void main(){
		gl_Position =  MVP * (vec4(pos, 1.0) + translation);
	}` + "\x00"
)
