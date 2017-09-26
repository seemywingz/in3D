#version 410

struct MaterialData{
  vec3
    Iamb,
    Idif,
    Ispec;
};
uniform MaterialData Material;

out vec4 finalColor;

void main() {
  finalColor = vec4(Material.Idif, 1);
}
