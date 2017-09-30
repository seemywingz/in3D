package in3D

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
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

// SetRelPath : resolves the absolute path for provided relative path.
func SetRelPath(relPath string) {
	if _, filename, _, ok := runtime.Caller(1); ok {
		re := regexp.MustCompile("[a-zA-Z0-9-]*.go$")
		path := filepath.Join(re.ReplaceAllString(filename, ""), relPath)
		EoE("Error Accessing relPath:", os.Chdir(path))
	} else {
		EoE("Error Getting Caller Location", errors.New(filename))
	}
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

// ExecPath :
func ExecPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
