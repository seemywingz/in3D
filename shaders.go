package main

const (
	basicFragmentSRC = `
	#version 410

  in vec2 fragTexCoord;

  out vec4 outCLR;

  uniform sampler2D tex;

  void main() {
  	outCLR = texture(tex, fragTexCoord);
  	//outCLR = vec4(1,0,1,1);
  }` + "\x00"

	basicVertexSRC = `
	#version 410 core

	in vec4 vert;

  in vec2 vertTexCoord;
  out vec2 fragTexCoord;

	uniform mat4 MVP;
	uniform mat4 MODEL;

	void main(){
		fragTexCoord = vertTexCoord;
		gl_Position =  MVP * MODEL * vert;
	}` + "\x00"
)
