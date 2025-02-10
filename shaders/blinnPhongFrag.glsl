#version 410 core

const float pi = 3.14159265;
const int maxLights = 45;

uniform sampler2D tex;
uniform mat4 MVP, MODEL;
uniform vec4 COLOR;

in vec3 fragPos;
in vec3 fragNormal;  // Corrected typo
in vec2 fragTexCoord;

out vec4 finalColor;

struct MaterialData {
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
    vec3 texColor = texture(tex, fragTexCoord).rgb;
    if(length(texColor) < 0.01) {
        texColor = vec3(1.0); // Default to white
    }

    vec3 resultColor = vec3(0.0);

    for(int i = 0; i < maxLights; ++i) {
        vec3 L = normalize(Light[i].lightPos - fragPos);
        vec3 N = normalize(fragNormal); // Ensure we are using the corrected name
        vec3 V = normalize(-fragPos);

        float lambertian = max(dot(N, L), 0.0);
        float specular = 0.0;

        if(lambertian > 0.0) {
            vec3 H = normalize(L + V);
            float specAngle = max(dot(H, N), 0.0);
            float energyConservation = (100.0 + Material.shininess) / (100.0 * pi);
            specular = energyConservation * pow(specAngle, Material.shininess);
        }

        // Improved Light Attenuation
        float dist = distance(fragPos, Light[i].lightPos);
        float att = 1.0 / (1.0 + 0.09 * dist + 0.032 * dist * dist);

        // Compute final color with light attenuation
        resultColor += att * texColor * (Light[i].Iamb + lambertian * Light[i].Idif + specular * Light[i].Ispec);
    }

    finalColor = vec4(Material.Iamb * Material.Idif * Material.Ispec * resultColor, 1.0);
}
