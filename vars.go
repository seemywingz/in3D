package main

import "github.com/go-gl/glfw/v3.2/glfw"

// Position : struct to store 3D coords
type Position struct {
	X float32
	Y float32
	Z float32
}

// Color : struct to store RGBA values
type Color struct {
	R float32
	G float32
	B float32
	A float32
}

// DrawLogic : extra logic to perform durring DrawnObject Draw phase
type DrawLogic func(d *DrawnObjectData)

var (

	// Graphics
	window       *glfw.Window
	camera       Camera
	shader       map[string]uint32
	texture      map[string]uint32
	drawnObjects []DrawnObject

	// Shapes ....
	triangle = []float32{
		-1.0, -1.0, 0, 1.0, 0.0,
		1.0, -1.0, 0, 0.0, 0.0,
		-1.0, 1.0, 0, 1.0, 1.0,
	}

	square = []float32{
		//  X, Y, Z, U, V, normal(3)
		-1.0, -1.0, 0, 0.0, 1.0, 0.0, 0.0, 1.0,
		1.0, -1.0, 0, 1.0, 1.0, 0.0, 0.0, 1.0,
		-1.0, 1.0, 0, 0.0, 0.0, 0.0, 0.0, 1.0,

		-1.0, 1.0, 0, 0.0, 0.0, 0.0, 0.0, 1.0,
		1.0, 1.0, 0, 1.0, 0.0, 0.0, 0.0, 1.0,
		1.0, -1.0, 0, 1.0, 1.0, 0.0, 0.0, 1.0,
	}

	cardFront = []float32{
		//  X, Y, Z, U, V, normal(3)
		-1.25, -1.75, 0, 0.0, 1.0, 0.0, 0.0, 1.0,
		1.25, -1.75, 0, 1.0, 1.0, 0.0, 0.0, 1.0,
		-1.25, 1.75, 0, 0.0, 0.0, 0.0, 0.0, 1.0,

		-1.25, 1.75, 0, 0.0, 0.0, 0.0, 0.0, 1.0,
		1.25, 1.75, 0, 1.0, 0.0, 0.0, 0.0, 1.0,
		1.25, -1.75, 0, 1.0, 1.0, 0.0, 0.0, 1.0,
	}

	cardBack = []float32{
		-1.25, 1.75, -0.01, 1.0, 0.0, 0.0, 0.0, -1.0, // left top
		-1.25, -1.75, -0.01, 1.0, 1.0, 0.0, 0.0, -1.0, // left bottom
		1.25, -1.75, -0.01, 0.0, 1.0, 0.0, 0.0, -1.0, // right bottom

		-1.25, 1.75, -0.01, 1.0, 0.0, 0.0, 0.0, -1.0, // left top
		1.25, 1.75, -0.01, 0.0, 0.0, 0.0, 0.0, -1.0, // right top
		1.25, -1.75, -0.01, 0.0, 1.0, 0.0, 0.0, -1.0, //right bottom
	}

	cube = []float32{
		//  X, Y, Z, U, V, normal(3)
		// Bottom
		-1.0, -1.0, -1.0, 0.0, 0.0, 0.0, -1.0, 0.0,
		1.0, -1.0, -1.0, 1.0, 0.0, 0.0, -1.0, 0.0,
		-1.0, -1.0, 1.0, 0.0, 1.0, 0.0, -1.0, 0.0,
		1.0, -1.0, -1.0, 1.0, 0.0, 0.0, -1.0, 0.0,
		1.0, -1.0, 1.0, 1.0, 1.0, 0.0, -1.0, 0.0,
		-1.0, -1.0, 1.0, 0.0, 1.0, 0.0, -1.0, 0.0,

		// Top
		-1.0, 1.0, -1.0, 0.0, 0.0, 0.0, 1.0, 0.0,
		-1.0, 1.0, 1.0, 0.0, 1.0, 0.0, 1.0, 0.0,
		1.0, 1.0, -1.0, 1.0, 0.0, 0.0, 1.0, 0.0,
		1.0, 1.0, -1.0, 1.0, 0.0, 0.0, 1.0, 0.0,
		-1.0, 1.0, 1.0, 0.0, 1.0, 0.0, 1.0, 0.0,
		1.0, 1.0, 1.0, 1.0, 1.0, 0.0, 1.0, 0.0,

		// Front
		-1.0, -1.0, 1.0, 1.0, 0.0, 0.0, 0.0, 1.0,
		1.0, -1.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0,
		-1.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0.0, 1.0,
		1.0, -1.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0,
		1.0, 1.0, 1.0, 0.0, 1.0, 0.0, 0.0, 1.0,
		-1.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0.0, 1.0,

		// Back
		-1.0, -1.0, -1.0, 0.0, 0.0, 0.0, 0.0, -1.0,
		-1.0, 1.0, -1.0, 0.0, 1.0, 0.0, 0.0, -1.0,
		1.0, -1.0, -1.0, 1.0, 0.0, 0.0, 0.0, -1.0,
		1.0, -1.0, -1.0, 1.0, 0.0, 0.0, 0.0, -1.0,
		-1.0, 1.0, -1.0, 0.0, 1.0, 0.0, 0.0, -1.0,
		1.0, 1.0, -1.0, 1.0, 1.0, 0.0, 0.0, -1.0,

		// Left
		-1.0, -1.0, 1.0, 0.0, 1.0, -1.0, 0.0, 0.0,
		-1.0, 1.0, -1.0, 1.0, 0.0, -1.0, 0.0, 0.0,
		-1.0, -1.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0,
		-1.0, -1.0, 1.0, 0.0, 1.0, -1.0, 0.0, 0.0,
		-1.0, 1.0, 1.0, 1.0, 1.0, -1.0, 0.0, 0.0,
		-1.0, 1.0, -1.0, 1.0, 0.0, -1.0, 0.0, 0.0,

		// Right
		1.0, -1.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0.0,
		1.0, -1.0, -1.0, 1.0, 0.0, 1.0, 0.0, 0.0,
		1.0, 1.0, -1.0, 0.0, 0.0, 1.0, 0.0, 0.0,
		1.0, -1.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0.0,
		1.0, 1.0, -1.0, 0.0, 0.0, 1.0, 0.0, 0.0,
		1.0, 1.0, 1.0, 0.0, 1.0, 1.0, 0.0, 0.0,
	}
)

const (
	blinnPhongFragmentSRC = `
	#version 410
  precision mediump float;

	uniform sampler2D tex;
	uniform mat4 MVP;
	uniform vec3 CPOS;
	uniform mat4 MODEL;

	vec3 lightPos = vec3(-2.0, 0.0, 0.0);
	const vec3 ambientColor = vec3(0.1, 0.1, 0.1)*1;
	const vec3 diffuseColor = vec3(0.1, 0.1, 0.1)*7;
	const vec3 specColor = vec3(0.1, 0.1, 0.1)*10;

	in vec3 normalInterp;
	in vec3 fragPos;
	in vec2 fragTexCoord;

  out vec4 finalColor;

  void main() {
		vec3 normal = normalize(normalInterp);
	  vec3 lightDir = normalize(lightPos - fragPos);
		vec3 camDir = normalize(CPOS - fragPos);

	  float lambertian = max(dot(lightDir,normal), 0.0);
	  float specular = 0.0;

	  if(lambertian > 0.0) {

	    vec3 viewDir = normalize(-fragPos);

	    // this is blinn phong
	    vec3 halfDir = normalize(lightDir + viewDir);
	    float specAngle = max(dot(halfDir, normal), 0.0);
	    specular = pow(specAngle, 20.0);

	  }

    vec4 surfaceColor = texture(tex, fragTexCoord);
		finalColor = vec4(ambientColor +
		                  lambertian * diffuseColor +
		                  specular * specColor, 1.0) * surfaceColor;

	  // finalColor = texture(tex, fragTexCoord);
    //finalColor = vec4(1,0,1,1);
  }` + "\x00"

	blinnPhongVertexSRC = `
	#version 410 core

	uniform mat4 MVP, MODEL, NormalMatrix;

	in vec3 vert;
	in vec2 vertTexCoord;
	in vec3 vertNormal;

	out vec3 fragPos;
	out vec2 fragTexCoord;
	out vec3 normalInterp;

	void main(){
	  vec4 fragPos4 = MODEL * vec4(vert, 1.0);
	  fragPos = vec3(fragPos4) / fragPos4.w;
		fragTexCoord = vertTexCoord;

		// mat4 normalMatrix = transpose(inverse(MODEL));
	  // normalInterp = vec3(normalMatrix * vec4(vertNormal, 0.0));
	  normalInterp = vec3(NormalMatrix * vec4(vertNormal, 0.0));

		gl_Position =  MVP * MODEL * vec4(vert, 1.0);
	}` + "\x00"
)
