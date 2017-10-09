#version 410 core

uniform mat4 MVP, MODEL, NormalMatrix;

in vec3 vert;
in vec2 vertTexCoord;
in vec3 vertNormal;
in vec3 vertTangent;
// in vec3 vertBitangent;

out vec3 fragPos;
out vec2 fragTexCoord;
out vec3 fragNoraml;
out mat3 fragTBN;

void main(){
  vec4 fragPos4 = MODEL * vec4(vert, 1.0);
  fragTexCoord = vertTexCoord;
  fragPos =  fragPos4.xyz / fragPos4.w;
  fragNoraml = (NormalMatrix * vec4(vertNormal, 0.0)).xyz;

  vec3 N = normalize(vec3(MODEL * vec4(vertNormal,    0.0)));
  vec3 T = normalize(vec3(MODEL * vec4(vertTangent,   0.0)));
  // re-orthogonalize T with respect to N
  T = normalize(T - dot(T, N) * N);
  // then retrieve perpendicular vector B with the cross product of T and N
  vec3 B = cross(N, T);
  fragTBN = mat3(T, B, N);


  gl_Position =  MVP * MODEL * vec4(vert, 1.0);
}
