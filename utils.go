package gg

import (
	"bufio"
	"bytes"
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"
)

// LoE : exit with error code 1 and print if err is notnull
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

// FtoA : convert float32 to string
func FtoA(n float32) string {
	return strconv.FormatFloat(float64(n), 'f', 6, 32)
}

// Loop2D : loop through 2 dimentional slice with 2 nested for loops uwing provided width and height
func Loop2D(height, width int, logic func(row, col int)) {
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			logic(row, col)
		}
	}
}

// Mkdir : make a directory if it does not exist
func Mkdir(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Mkdir(filePath, os.ModePerm)
	}
}

// DownloadImage : download image from provided url and save to provided filelocation
func DownloadImage(url, fileName string) {
	response, err := http.Get(url)
	if err != nil {
		LoE("Error Getting Image "+url, err)
		return
	}

	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create(fileName)
	LoE("Error Creating File", err)

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	LoE("Error Saving Image File", err)
	file.Close()
}

// SendRequest : send http request to provided url
func SendRequest(req *http.Request) []byte {
	client := http.Client{}
	res, err := client.Do(req)
	EoE("Error Getting HTTP Response", err)

	resData, err := ioutil.ReadAll(res.Body)
	EoE("Error Parsing HTTP Response", err)
	return resData
}

// GetHomeDir : returns a full path to user's home dorectory
func GetHomeDir() string {
	usr, err := user.Current()
	EoE("Failed to get Current User", err)
	if usr.HomeDir != "" {
		return usr.HomeDir
	}
	return os.Getenv("HOME")
}

// Confirm : return confirmation based on user input
func Confirm(q string) bool {
	a := GetInput(q + " (Y/n) ")
	var res bool
	switch a {
	case "":
		fallthrough
	case "y":
		fallthrough
	case "Y":
		res = true
	case "n":
	case "N":
		res = false
	default:
		return Confirm(q)
	}
	return res
}

// GetInput : return string of user input
func GetInput(q string) string {
	print(q)
	reader := bufio.NewReader(os.Stdin)
	ans, _ := reader.ReadString('\n')
	return strings.TrimRight(ans, "\n")
}

// SelectFromArray : select an element in the provided array
func SelectFromArray(a []string) string {
	fmt.Println("Choices:")
	for i := range a {
		fmt.Println("[", i, "]: "+a[i])
	}
	sel, err := strconv.Atoi(GetInput("Enter Number of Selection: "))
	EoE("Error Getting Integer Input from User", err)
	if sel <= len(a)-1 {
		return a[sel]
	}
	return SelectFromArray(a)
}

// SetFromInput : set value of provided var to the value of user input
func SetFromInput(q string, a *string) {
	*a = strings.TrimRight(GetInput(q), "\n")
}

// GetIP : get local ip address
func GetIP() string {
	addrs, err := net.InterfaceAddrs()
	EoE("Failed to Get Inet Address", err)
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// LineCounter : count number of lines `\n`
func LineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
