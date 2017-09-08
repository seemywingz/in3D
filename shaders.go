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

	uniform vec4 rot;

	uniform vec4 trans;

	mat4 rotationMatrix(vec3 axis, float angle)	{
	    axis = normalize(axis);
	    float s = sin(angle);
	    float c = cos(angle);
	    float oc = 1.0 - c;

	    return mat4(oc * axis.x * axis.x + c,           oc * axis.x * axis.y - axis.z * s,  oc * axis.z * axis.x + axis.y * s,  0.0,
	                oc * axis.x * axis.y + axis.z * s,  oc * axis.y * axis.y + c,           oc * axis.y * axis.z - axis.x * s,  0.0,
	                oc * axis.z * axis.x - axis.y * s,  oc * axis.y * axis.z + axis.x * s,  oc * axis.z * axis.z + c,           0.0,
	                0.0,                                0.0,                                0.0,                                1.0);
	}

  float a = 10;
	void main(){
		// vec4 translation = vec4(pos, 1.0) + trans;
		vec4 translation = vec4(pos, 1.0) + trans;
		gl_Position = MVP * ( translation * rotationMatrix(vec3(1, 0, 0), rot[0]) );

	}` + "\x00"
)
