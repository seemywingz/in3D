package in3d

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"

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
