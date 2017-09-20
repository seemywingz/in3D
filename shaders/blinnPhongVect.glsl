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
}
