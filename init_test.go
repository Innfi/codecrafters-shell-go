package test

import (
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	input := "hello"
	assert.Equal(t, input[1:], "ello")
}

func TestIsNumber(t *testing.T) {
	assert.Equal(t, unicode.IsDigit('1'), true)
	assert.Equal(t, unicode.IsDigit('"'), false)
}

func TestIsLetter(t *testing.T) {
	assert.Equal(t, unicode.IsLetter('A'), true)
	assert.Equal(t, unicode.IsLetter('r'), true)
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

			if quote == 0 {
				if next == '\\' || next == '$' || next == 'n' || next == '"' {
					runeArray = append(runeArray, rune(input[1]))
					input = input[2:]
					continue
				}

				input = input[1:]
				continue
			} else {
				runeArray = append(runeArray, rune(input[1]))

				input = input[2:]
				continue
			}
		}

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

func TestToTokenArrayRevisedNoQuotationSingleWord(t *testing.T) {
	result := ToTokenArrayRevised("hello")
	assert.Equal(t, len(result), 1)
	assert.Equal(t, result[0], "hello")
}

func TestToTokenArrayRevisedNoQuotationSpacedWord(t *testing.T) {
	result := ToTokenArrayRevised("hello world")
	assert.Equal(t, len(result), 3)
	assert.Equal(t, result[0], "hello")
	assert.Equal(t, result[1], " ")
	assert.Equal(t, result[2], "world")
}

func TestToTokenArrayRevisedNoQuotation3SpacedWord(t *testing.T) {
	result := ToTokenArrayRevised("hello   world")
	assert.Equal(t, len(result), 3)
	assert.Equal(t, result[0], "hello")
	assert.Equal(t, result[1], " ")
	assert.Equal(t, result[2], "world")
}

func TestToTokenArrayRevisedSingleQuotationSpaced(t *testing.T) {
	result := ToTokenArrayRevised("'hello' world")
	assert.Equal(t, len(result), 3)
	assert.Equal(t, result[0], "hello")
	assert.Equal(t, result[1], " ")
	assert.Equal(t, result[2], "world")
}

func TestToTokenArrayRevisedSpaceWithinSingleQuotationSpaced(t *testing.T) {
	result := ToTokenArrayRevised("'hello oh' world")
	assert.Equal(t, len(result), 3)
	assert.Equal(t, result[0], "hello oh")
	assert.Equal(t, result[1], " ")
	assert.Equal(t, result[2], "world")
}

func TestToTokenArrayRevisedSingleQuotationNoSpaced(t *testing.T) {
	result := ToTokenArrayRevised("'hello'world")
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result[0], "hello")
	assert.Equal(t, result[1], "world")
}

func TestToTokenArrayRevisedDoubleQuotationSpaced(t *testing.T) {
	result := ToTokenArrayRevised("\"hello\" world oh")
	assert.Equal(t, len(result), 5)
	assert.Equal(t, result[0], "hello")
	assert.Equal(t, result[1], " ")
	assert.Equal(t, result[2], "world")
	assert.Equal(t, result[3], " ")
	assert.Equal(t, result[4], "oh")
}

func TestToTokenArrayRevicedDoubleQuotationWithSingleQuoteNoSpaced(t *testing.T) {
	result := ToTokenArrayRevised("\"hello'world\"oh")
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result[0], "hello'world")
	assert.Equal(t, result[1], "oh")
}

func TestToTokenArrayRevicedBackSlash(t *testing.T) {
	result := ToTokenArrayRevised("world\\ \\ \\ script")
	assert.Equal(t, len(result), 5)
	assert.Equal(t, result[0], "world")
	assert.Equal(t, result[1], " ")
	assert.Equal(t, result[2], " ")
	assert.Equal(t, result[3], " ")
	assert.Equal(t, result[4], "script")
}

func TestToTokenArrayRevicedBackSlashWithinQuote(t *testing.T) {
	result := ToTokenArrayRevised("\"hello\\\\ world\"")
	assert.Equal(t, len(result), 1)
	assert.Equal(t, result[0], "hello\\ world")
}
