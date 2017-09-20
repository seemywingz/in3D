#version 410
precision mediump float;

uniform sampler2D tex;
uniform mat4 MVP;
uniform vec3 CPOS;
uniform mat4 MODEL;

vec3 lightPos = vec3(-2.0, 0.0, 0.0);
const vec3 ambientColor = vec3(0.1, 0.1, 0.1)*1;
const vec3 diffuseColor = vec3(0.1, 0.1, 0.1)*7;
const vec3 specColor = vec3(0.1, 0.1, 0.1)*10;

in vec3 normalInterp;
in vec3 fragPos;
in vec2 fragTexCoord;

out vec4 finalColor;

void main() {
  vec3 normal = normalize(normalInterp);
  vec3 lightDir = normalize(lightPos - fragPos);
  vec3 camDir = normalize(CPOS - fragPos);

  float lambertian = max(dot(lightDir,normal), 0.0);
  float specular = 0.0;

  if(lambertian > 0.0) {

    vec3 viewDir = normalize(-fragPos);

    // this is blinn phong
    vec3 halfDir = normalize(lightDir + viewDir);
    float specAngle = max(dot(halfDir, normal), 0.0);
    specular = pow(specAngle, 20.0);

  }

  vec4 surfaceColor = texture(tex, fragTexCoord);
  finalColor = vec4(ambientColor +
                    lambertian * diffuseColor +
                    specular * specColor, 1.0) * surfaceColor;

  // finalColor = texture(tex, fragTexCoord);
  //finalColor = vec4(1,0,1,1);
}
