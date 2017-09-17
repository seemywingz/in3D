package main

const (
	basicFragmentSRC = `
	#version 410

  in vec2 fragTXT;

  out vec4 outCLR;

  uniform sampler2D tex;

  void main() {
  	outCLR = texture(tex, fragTXT);
  	//outCLR = vec4(1,0,1,1);
  }` + "\x00"

	basicVertexSRC = `
	#version 410 core

	in vec4 vertPOS;

  in vec2 vertTXT;
  out vec2 fragTXT;

	uniform mat4 MVP;
	uniform mat4 MODEL;

	void main(){
		fragTXT = vertTXT;
		gl_Position =  MVP * MODEL * vertPOS;
	}` + "\x00"
)
