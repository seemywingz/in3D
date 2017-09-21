package gg

// Light : struct to hold light data
type Light struct {
	LightData
}

// LightData :
type LightData struct {
	Position
	Iamb  [3]float32
	Idif  [3]float32
	Ispec [3]float32
}

// NewLight :
func NewLight(data LightData) *Light {
	return &Light{}
}
