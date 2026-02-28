package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "./9-file-io/fromString.txt"
	file, err := os.Create(fileName)

	// wait till everything in this func gets executed
	defer file.Close()
	checkError(err)

	length, err := io.WriteString(file, "Hello from Go!")

	fmt.Printf("Wrote a file with %v characters\n", length)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
