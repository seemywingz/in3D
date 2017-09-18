package main

const (
	basicFragmentSRC = `
	#version 410
  precision mediump float;

	uniform sampler2D tex;
	uniform mat4 MVP;
	uniform mat4 MODEL;

	const vec3 lightPos = vec3(1.0,1.0,1.0);
	const vec3 ambientColor = vec3(0.1, 0.0, 0.0);
	const vec3 diffuseColor = vec3(0.5, 0.0, 0.0);
	const vec3 specColor = vec3(1.0, 1.0, 1.0);

	in vec3 normalInterp;
	in vec3 fragPos;
	in vec2 fragTexCoord;

  out vec4 finalColor;

  void main() {
	  finalColor = texture(tex, fragTexCoord);
    //finalColor = vec4(1,0,1,1);
  }` + "\x00"

	basicVertexSRC = `
	#version 410 core

	uniform mat4 MVP, MODEL;

	in vec3 vert;
	in vec2 vertTexCoord;
	in vec3 vertNormal;

	out vec3 normalInterp;
	out vec3 fragPos;
	out vec2 fragTexCoord;

	void main(){
	  // vec4 fragPos4 = MODEL * vec4(vert, 1.0);
	  // fragPos = vec3(fragPos4) / fragPos4.w;
	  // normalInterp = vec3(normalMat * vec4(vertNormal, 0.0));
    fragTexCoord = vertTexCoord;
		gl_Position =  MVP * MODEL * vec4(vert, 1);
	}` + "\x00"
)
