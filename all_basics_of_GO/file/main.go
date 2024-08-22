package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file")
	content := "this needs to be in a Go file"

	// Create a new file
	file, err := os.Create("./mylcogofile.txt")
	if err != nil {
		panic(err)
	}

	// Write content to the file
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is:", length)
	file.Close()

	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	// Read the content of the file
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Convert byte slice to string and print
	fmt.Println("Text data inside the file:\n", string(databyte))
}
