package main

import (
	"os"
	"runtime"
	"strings"
	"unicode"
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
	effectiveInput = strings.TrimLeft(effectiveInput, " ")

	for {
		if len(effectiveInput) <= 0 {
			break
		}

		if effectiveInput[0] == ' ' {
			tokenArray = append(tokenArray, " ")
			effectiveInput = strings.TrimLeft(effectiveInput, " ")
			continue
		}

		start := strings.Index(effectiveInput, "'")
		if start == -1 {
			tokenArray = append(tokenArray, strings.Join(strings.Fields(effectiveInput), " "))
			break
		}

		if start > 0 {
			tokenArray = append(tokenArray, effectiveInput[:start])
		}

		effectiveInput = effectiveInput[start+1:]

		end := strings.Index(effectiveInput, "'")
		if end < 0 {
			break
		}

		tokenArray = append(tokenArray, effectiveInput[:end])

		effectiveInput = effectiveInput[end+1:]
	}

	return tokenArray
}

func ToTokenArrayRevised(origin string) []string {
	tokenArray := []string{}

	input := strings.Trim(origin, "\r\n")
	input = strings.TrimLeft(input, " ")

	for {
		first := rune(input[0])

		// double quote
		if first == '"' {
			// push the mark into the stack
		}

		// single quote

		if unicode.IsSpace(first) {
			tokenArray = append(tokenArray, " ")
			input = strings.TrimLeft(input, " ")
			continue
		}
	}

	return tokenArray
}
