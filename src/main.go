package main

import (
	"fmt"
	"os"
	"strings"
)

// this is test program to count words in file

func main() {
	fileContent, err := os.ReadFile("testfile.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error occurred")
		os.Exit(1)
	}

	words := strings.Fields(string(fileContent))

	fmt.Println("Word count is:")
	fmt.Println(len(words))
}
