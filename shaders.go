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
	uniform mat4 uRotation;
	uniform vec4 uTranslation;

	void main(){
		gl_Position =  MVP * uRotation * (vec4(pos, 1.0) + uTranslation);
	}` + "\x00"
)
