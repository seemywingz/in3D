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

	in vec3 pos;

	uniform mat4 MVP;
	uniform mat4 localRotation;

	void main(){
		vec4 p = vec4(pos, 1.0);
		gl_Position =  MVP * (localRotation) * (p);
	}` + "\x00"
)
