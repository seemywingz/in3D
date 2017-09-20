#version 410 core

uniform mat4 MVP, MODEL, NormalMatrix;

in vec3 vert;
in vec2 vertTexCoord;
in vec3 vertNormal;

out vec3 fragPos;
out vec2 fragTexCoord;
out vec3 normalInterp;

void main(){
  fragTexCoord = vertTexCoord;
  fragPos =  MVP * MODEL * vec4(vert, 1.0);
  gl_Position =  MVP * MODEL * vec4(vert, 1.0);
}
