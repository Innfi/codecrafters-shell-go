package main

import (
	"fmt"
	"os"
)

func HandleCommandPwd() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(path)
}
