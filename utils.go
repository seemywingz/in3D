package in3d

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

// Random : return pseudo random number in range
func Random(min, max int) int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// RandomF : return pseudo random float32 number in range
func RandomF() float32 {
	rand.NewSource(time.Now().UnixNano())
	return rand.Float32()
}

// ExecPath :
func ExecPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
