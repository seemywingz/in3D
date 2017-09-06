#version 330 core

// Input vertex data, different for all executions of this shader.
layout(location = 0) in vec3 pos;

//vales that stay constant for the whole mesh
uniform mat4 MVP;
uniform vec4 translation;

void main(){
  gl_Position = MVP * (vec4(pos,1.0) + translation);
}
