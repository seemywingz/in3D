#version 410
const float pi = 3.14159265;
const int maxLights = 11;

uniform sampler2D tex;
uniform mat4 MVP, MODEL;
uniform vec4 COLOR;

in vec3 fragPos;
in vec3 fragNoraml;
in vec2 fragTexCoord;

out vec4 finalColor;

struct MaterialData{
  float shininess;
  vec3
    Iamb,
    Idif,
    Ispec;
};
uniform MaterialData Material;

struct LightData {
  float lightRad;
  vec3
    lightPos,
    Iamb,
    Idif,
    Ispec;
};
uniform LightData Light[maxLights];

float dinstance(vec3 p0, vec3 p1){
  float dx, dy, dz;
  dx = pow((p1.x - p0.x), 2);
  dy = pow((p1.y - p0.y), 2);
  dz = pow((p1.z - p0.z), 2);
  return sqrt((dx+dy+dz));
}

void main() {
  for(int i=0;i<maxLights;++i) {
    vec3 L = normalize(Light[i].lightPos - fragPos);
    vec3 N = normalize(fragNoraml);
    vec3 V = normalize(-fragPos);

    float lambertian = max(dot(N,L), 0.0);
    float specular = 0.0;

    if(lambertian > 0.0) {
      vec3 H = normalize(L + V);
      float specAngle = max(dot(H, N), 0.0);
      float x = 100.0;
      float eConservation = ( x + Material.shininess ) / ( x * pi );
      specular = eConservation * pow(specAngle, Material.shininess);
    }
    float diffuse = max(dot(normalize(fragNoraml), normalize(Light[i].lightPos)), 0.0);

    vec3 texture = texture(tex, fragTexCoord).rgb;
    if(texture == vec3(0,0,0)){
      texture = vec3(1,1,1);// white default
    }

    float
      dist = distance(fragPos, Light[i].lightPos),
      att = clamp(1.0 - dist*dist/(Light[i].lightRad*Light[i].lightRad), 0.0, 1.0); att *= att;
    
    finalColor +=  vec4( att * texture *(Light[i].Iamb +
                      lambertian * Light[i].Idif +
                      specular * Light[i].Ispec ) ,1);
  }
  vec4 matColor =  vec4((Material.Iamb * Material.Idif * Material.Ispec),1);
  finalColor = matColor * finalColor;
}
