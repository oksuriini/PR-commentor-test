package main

import (
	"flag"
	"os"
)

// this is test program to count words in file
var fileName string

func init() {
	flag.StringVar(&fileName, "filename", "", "Name of the file from which to count words")
	flag.Parse()
}

func main() {
	readFile(fileName)
}

func readFile(fileName string) ([]byte, error) {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}
