package main

import (
	"fmt"
	"strings"
)

func HandleCommandEcho(argument *Argument) {
	toTokenArray(argument.params)

	fmt.Println(argument.params)
}

func toTokenArray(input string) []string {
	tokenArray := []string{}

	effectiveInput := input

	for len(effectiveInput) > 0 {
		effectiveInput = strings.TrimLeft(effectiveInput, " ")

		if effectiveInput[0] == '\'' {
			index := strings.IndexRune(effectiveInput[1:], '\'')
			if index < 0 {
				break
			}

			tokenArray = append(tokenArray, effectiveInput[1:index+1])

			if len(effectiveInput) >= index+2 {
				break
			}
			effectiveInput = effectiveInput[index+2:]
			fmt.Println("after cut: ", effectiveInput)

			continue
		}

		before, after, _ := strings.Cut(effectiveInput, " ")

		if len(before) > 0 {
			tokenArray = append(tokenArray, before)
			effectiveInput = after
			fmt.Println("effectiveInput: ", effectiveInput)
		} else {
			tokenArray = append(tokenArray, after)
			break
		}
	}

	fmt.Println("tokenArray: ", tokenArray)

	return tokenArray
}
