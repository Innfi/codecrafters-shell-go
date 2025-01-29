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

	var quote rune
	runeArray := []rune{}

	for {
		if len(input) <= 0 {
			break
		}

		elem := rune(input[0])

		if unicode.IsSpace(elem) {
			if quote == 0 {
				if len(runeArray) > 0 {
					tokenArray = append(tokenArray, string(runeArray))
					runeArray = []rune{}
				}

				tokenArray = append(tokenArray, " ")
				input = strings.TrimLeft(input, " ")

				continue
			} else {
				runeArray = append(runeArray, elem)

				input = input[1:]
				continue
			}
		}

		if elem == '\\' {
			next := rune(input[1])

			if next == '\\' || next == '$' || next == 'n' || next == '"' {
				runeArray = append(runeArray, rune(next))
				input = input[2:]
				continue
			}

			if quote != 0 {
				runeArray = append(runeArray, rune(elem))
			}

			input = input[1:]
			continue
		}
		// if elem == '\\' {
		// 	next := rune(input[1])

		// 	if quote == 0 {
		// 		if next == '\\' || next == '$' || next == 'n' || next == '"' {
		// 			runeArray = append(runeArray, rune(input[1]))
		// 			input = input[2:]
		// 			continue
		// 		}

		// 		input = input[1:]
		// 		continue
		// 	} else {
		// 		if next == '\\' || next == '$' || next == 'n' || next == '"' {
		// 			runeArray = append(runeArray, rune(next))

		// 			input = input[2:]
		// 			continue
		// 		}

		// 		runeArray = append(runeArray, rune(elem))
		// 		input = input[1:]
		// 		continue
		// 	}
		// }

		if elem == '"' || elem == '\'' {
			if quote == 0 {
				quote = elem

				input = input[1:]
				continue
			}

			if quote == elem {
				tokenArray = append(tokenArray, string(runeArray))
				runeArray = []rune{}
				quote = 0
			} else {
				runeArray = append(runeArray, elem)
			}

			input = input[1:]
			continue
		}

		runeArray = append(runeArray, elem)

		input = input[1:]
		continue
	}

	if len(runeArray) > 0 {
		tokenArray = append(tokenArray, string(runeArray))
	}

	return tokenArray
}
