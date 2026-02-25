package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	// String, error object 
	str, _ := reader.ReadString('\n')
	fmt.Printf("Your name is %v",str)
}