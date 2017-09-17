package main

const (
	basicFragmentSRC = `
	#version 410

  in vec2 fragTXT;

  out vec4 outCLR;

  uniform sampler2D texSampler;

  void main() {
  	//outCLR = texture(texSampler, fragTXT);
  	outCLR = vec4(1,0,1,1);
  }` + "\x00"

	basicVertexSRC = `
	#version 410 core

	in vec4 inPOS;

  in vec2 vertTXT;
  out vec2 fragTXT;

	uniform mat4 MVP;
	uniform mat4 MODEL;

	void main(){
		fragTXT = vertTXT;
		gl_Position =  MVP * MODEL * inPOS;
	}` + "\x00"
)
