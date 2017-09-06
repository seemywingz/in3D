#version 410

// Hard Code Color for Now
uniform vec4 inputColour = vec4(1,1,1,1);
out vec4 fragColour;

void main() {
  fragColour = inputColour;
}
