package main

const (
	basicFragmentSRC = `
	#version 410

	uniform sampler2D tex;
	struct Light {
   vec3 position;
   vec3 intensities; //a.k.a the color of the light
  } light;

  in vec2 fragTexCoord;
  in vec3 fragNormal;
  in vec4 fragVert;
	in mat4 fragModel;

	out vec4 finalColor;

  void main() {
		light.position = vec3(0,0,0);
		light.intensities = vec3(1,1,1);

    //calculate normal in world coordinates
    mat3 normalMatrix = transpose(inverse(mat3(fragModel)));
    vec3 normal = normalize(normalMatrix * fragNormal);

    //calculate the location of this fragment (pixel) in world coordinates
    vec3 fragPosition = vec3(fragModel * fragVert);

    //calculate the vector from this pixels surface to the light source
    vec3 surfaceToLight = light.position - fragPosition;

    //calculate the cosine of the angle of incidence
    float brightness = dot(normal, surfaceToLight) / (length(surfaceToLight) * length(normal));
    brightness = clamp(brightness, 0, 1);

    //calculate final color of the pixel, based on:
    // 1. The angle of incidence: brightness
    // 2. The color/intensities of the light: light.intensities
    // 3. The texture and texture coord: texture(tex, fragTexCoord)
    vec4 surfaceColor = texture(tex, fragTexCoord);
    //finalColor = vec4(brightness * light.intensities * surfaceColor.rgb, surfaceColor.a);
		finalColor = vec4(brightness * light.intensities, 1) * texture(tex, fragTexCoord);
  	//finalColor = texture(tex, fragTexCoord);
  	//finalColor = vec4(1,0,1,1);
  }` + "\x00"

	basicVertexSRC = `
	#version 410 core

	uniform mat4 MVP;
	uniform mat4 MODEL;

	in vec4 vert;
  in vec2 vertTexCoord;
	in vec3 vertNormal;

  out vec4 fragVert;
  out vec2 fragTexCoord;
  out vec3 fragNormal;
	out mat4 fragModel;


	void main(){
		fragTexCoord = vertTexCoord;
    fragNormal = vertNormal;
    fragVert = vert;
		fragModel = MODEL;
		gl_Position =  MVP * MODEL * vert;
	}` + "\x00"
)
