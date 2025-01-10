package main

import (
	"bufio"
	"strings"
)

type Argument struct {
	command string
	params  string
}

func ReadArguments(reader *bufio.Reader) (*Argument, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	input = strings.Replace(input, "\r", "", -1)
	input = strings.Replace(input, "\n", "", -1)

	command, params, _ := strings.Cut(input, " ")

	return &Argument{
		command,
		params,
	}, nil
}
