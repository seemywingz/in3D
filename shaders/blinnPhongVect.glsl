#version 410 core

uniform mat4 MVP, MODEL, NormalMatrix;

in vec3 vert;
in vec2 vertTexCoord;
in vec3 vertNormal;

out vec3 fragPos;
out vec2 fragTexCoord;
out vec3 fragNoraml;

void main(){
  vec4 fragPos4 = MODEL * vec4(vert, 1.0);

  fragTexCoord = vertTexCoord;
  fragPos =  fragPos4.xyz / fragPos4.w;
  fragNoraml = (NormalMatrix * vec4(vertNormal, 0.0)).xyz;

  gl_Position =  MVP * MODEL * vec4(vert, 1.0);
}
