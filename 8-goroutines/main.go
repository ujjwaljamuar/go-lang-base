package main

import (
	"fmt"
	"time"
)

func main() {
	go say("Hi from goroutine")
	fmt.Println("Hi from main")

	go func(message string) {
		fmt.Println(message)
	}("Hello from anaonymous function")

	time.Sleep(2 * time.Second)
	fmt.Println("All done")
}

func say(message string) {
	time.Sleep(1 * time.Second)
	fmt.Println(message)
}
