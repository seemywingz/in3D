package in3D

import (
	"fmt"
	"go/build"
	"log"
	"math/rand"
	"os"
	"time"
)

// LoE : log if err is notnull
func LoE(msg string, err error) {
	if err != nil {
		log.Printf("\n❌  %s\n   %v\n", msg, err)
	}
}

// EoE : exit with error code 1 and print, if err is not nil
func EoE(msg string, err error) {
	if err != nil {
		fmt.Printf("\n❌  %s\n   %v\n", msg, err)
		os.Exit(1)
		panic(err)
	}
}

// SetDirPath : resolves the absolute path from importPath.
// There doesn't need to be a valid Go package inside that import path, but the directory must exist.
func SetDirPath(importPath string) {
	// importPath = "github.com/seemywingz/gtils"
	path, err := build.Import(importPath, "", build.FindOnly)
	EoE("Unable to find Go package in your GOPATH, it's needed to load assets:", err)

	err = os.Chdir(path.Dir)
	EoE("Error Setting Package Dir", err)
	// println(path.Dir)
}

// Random : return pseudo random number in range
func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// Randomf : return pseudo random float32 number in range
func Randomf() float32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32()
}
