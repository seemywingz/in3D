#version 410
const float pi = 3.14159265;
const int maxLights = 11;

uniform sampler2D TEXTURE;
uniform sampler2D NORMAL_MAP;
uniform mat4 MVP, MODEL;
uniform vec4 COLOR;

in vec3 fragPos;
in vec3 fragNormal;
in vec2 fragTexCoord;
in mat3 fragTBN;

out vec4 finalColor;

struct MaterialData{
  float shininess;
  vec3 Iamb, Idif, Ispec;
};
uniform MaterialData Material;

struct LightData {
  float lightRad;
  vec3 lightPos, Iamb, Idif, Ispec;
};
uniform LightData Light[maxLights];

void main() {
  vec3 textr = texture(TEXTURE, fragTexCoord).rgb;
  if (textr == vec3(0,0,0)) {
    textr = vec3(1,1,1);
  }

  vec3 normal = texture(NORMAL_MAP, fragTexCoord).rgb;
  if (normal == vec3(0,0,0)) {
      normal = vec3(0.5, 0.5, 1.0); // Default normal if missing
  }
  normal = normalize(normal * 2.0 - 1.0);


  vec3 N = normalize(fragTBN * normal);
  vec3 V = normalize(-fragPos);
  
  vec4 color = vec4(0.0);
  for(int i=0; i<maxLights; ++i) {
    vec3 L = normalize(Light[i].lightPos - fragPos);
    float lambertian = max(dot(N, L), 0.0);
    float specular = 0.0;

    if(lambertian > 0.0) {
      vec3 H = normalize(L + V);
      float specAngle = max(dot(H, N), 0.0);
      float eConservation = (100.0 + Material.shininess) / (100.0 * pi);
      specular = eConservation * pow(specAngle, Material.shininess);
    }

    float dist = distance(fragPos, Light[i].lightPos);
    float att = clamp(1.0 - dist*dist/(Light[i].lightRad*Light[i].lightRad), 0.0, 1.0); 
    att *= att;
    
    color += vec4(att * textr * (Light[i].Iamb + lambertian * Light[i].Idif + specular * Light[i].Ispec), 1);
  }
  vec4 matColor = vec4(Material.Iamb * Material.Idif * Material.Ispec, 1);
  finalColor = matColor * color;
}
