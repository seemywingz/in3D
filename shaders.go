package main

const (
	basicFragmentSRC = `
	#version 410

  in vec2 Texture;

  out vec4 color;

  uniform sampler2D tex;

  void main() {
  	//color = texture(tex, Texture);
  	color = vec4(1,0,1,1);
  }` + "\x00"

	basicVertexSRC = `
	#version 410 core

	in vec4 inPOS;

	uniform mat4 MVP;
	uniform mat4 localRotation;

	void main(){
		gl_Position =  MVP * localRotation * inPOS;
	}` + "\x00"
)
