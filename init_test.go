package test

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestIsNumber(t *testing.T) {
	assert.Equal(t, unicode.IsDigit('1'), true)
	assert.Equal(t, unicode.IsDigit('"'), false)
}

func TestIsLetter(t *testing.T) {
	assert.Equal(t, unicode.IsLetter('A'), true)
	assert.Equal(t, unicode.IsLetter('r'), true)
}
