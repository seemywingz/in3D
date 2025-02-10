package in3d

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

// SetRelPath : resolves the absolute path for provided relative path.
func SetRelPath(relPath string) error {
	if _, filename, _, ok := runtime.Caller(1); ok {
		re := regexp.MustCompile("[a-zA-Z0-9-]*.go$")
		path := filepath.Join(re.ReplaceAllString(filename, ""), relPath)
		return os.Chdir(path)
	} else {
		return fmt.Errorf("Error Getting Caller Location: %s", filename)
	}
}

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
