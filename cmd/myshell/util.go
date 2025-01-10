package main

import (
	"os"
	"runtime"
	"strings"
)

func GetPathAndDivider() ([]string, string) {
	delimiter, divider := getPathTokens()
	envPath := os.Getenv("PATH")
	pathArray := strings.Split(envPath, delimiter)

	return pathArray, divider
}

func getPathTokens() (string, string) {
	os := runtime.GOOS

	if os == "windows" {
		return ";", "\\"
	}

	return ":", "/"
}
