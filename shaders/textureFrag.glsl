#version 410
precision mediump float;

uniform sampler2D tex;
uniform mat4 MVP, MODEL;

in vec3 fragPos;
in vec2 fragTexCoord;

out vec4 finalColor;

void main() {
  // finalColor = texture(tex, fragTexCoord);
  //finalColor = vec4(1,0,1,1);
}
