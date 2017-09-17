package main

const (
	basicFragmentSRC = `
	#version 410

  in vec2 inTXT;

  out vec4 outCLR;

  uniform sampler2D tex;

  void main() {
  	//color = texture(tex, Texture);
  	outCLR = vec4(1,0,1,1);
  }` + "\x00"

	basicVertexSRC = `
	#version 410 core

	in vec4 inPOS;
  in vec2 inTXT;

  out vec2 outTXT;

	uniform mat4 MVP;
	uniform mat4 localRotation;

	void main(){
		gl_Position =  MVP * localRotation * inPOS;
		outTXT = inTXT;
	}` + "\x00"
)
