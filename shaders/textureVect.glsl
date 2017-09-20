#version 410 core

uniform mat4 MVP, MODEL;

in vec3 vert;
in vec2 vertTexCoord;

out vec2 fragTexCoord;

void main(){
  fragTexCoord = vertTexCoord;
  gl_Position =  MVP * MODEL * vec4(vert, 1.0);
}
