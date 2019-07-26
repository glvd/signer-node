package general

import (
	"errors"
	"path/filepath"
	"runtime"
	"strings"
)

// PathParse ...
func PathParse(url string) (string, error) {
	var parsedURL string
	var sys = runtime.GOOS
	absPath, _ := filepath.Abs("./")

	if sys == "windows" {
		parsedURL = absPath + strings.Replace(url, "/", "\\", -1)
	} else if sys == "linux" {
		parsedURL = absPath + url
	} else {
		return "nil", errors.New("os not match")
	}
	return parsedURL, nil
}
