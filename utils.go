package in3d

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"time"

	"github.com/seemywingz/go-toolbox"
)

// SetDir : resolves the absolute path for provided relative path.
func SetDir(relPath string) {
	if _, filename, _, ok := runtime.Caller(1); ok {
		re := regexp.MustCompile("[a-zA-Z0-9-]*.go$")
		path := filepath.Join(re.ReplaceAllString(filename, ""), relPath)
		fmt.Println("Setting Path to:", path)
		toolbox.EoE(os.Chdir(path), "Error Setting Directory")
	} else {
		toolbox.EoE(fmt.Errorf("Error Setting Directory"), "Error Setting Directory")
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
