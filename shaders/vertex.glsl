#version 410

// Input vertex data, different for all executions of this shader.
layout(location = 0) in vec4 vertexPos;

//vales that stay constant for the whole mesh
uniform mat4 transformationMatrix;
uniform mat4 projectionMatrix;

void main() {
  gl_Position = vertexPos;
}
