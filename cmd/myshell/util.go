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

func ToTokenArray(input string) []string {
	tokenArray := []string{}

	effectiveInput := strings.Trim(input, "\r\n")

	for {
		start := strings.Index(effectiveInput, "'")
		if start == -1 {
			tokenArray = append(tokenArray, strings.Fields(effectiveInput)...)
			break
		}

		tokenArray = append(tokenArray, strings.Fields(effectiveInput[:start])...)
		effectiveInput = effectiveInput[start+1:]

		end := strings.Index(effectiveInput, "'")
		if end < 0 {
			// tokenArray = append(tokenArray, strings.Fields(effectiveInput)...)
			break
		}

		tokenArray = append(tokenArray, effectiveInput[:end])

		effectiveInput = effectiveInput[end+1:]
	}

	return tokenArray
}
