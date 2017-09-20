#version 410 core

uniform mat4 MVP, MODEL, NormalMatrix;

in vec3 vert;

void main(){
  gl_Position =  MVP * MODEL * vec4(vert, 1.0);
}
