package main

import (
	"fmt"
)

func main() {
	val1 := 1
	val2 := 2
	val3 := 3

	var sum int = addValues(val1, val2)
	fmt.Printf("Sum is %v\n", sum)

	sum, count, average := addAllValues(val1, val2, val3)
	fmt.Printf("Sum of all values is %v\n", sum)
	fmt.Printf("Count of all values is %v\n", count)
	fmt.Printf("Average of all values is %v\n", average)

	dog := Dog{"Poodle", "Woof"}
	dog.Speak()
}

func addValues(val1, val2 int) int {
	return val1 + val2
}

func addAllValues(values ...int) (int, int, float64) {
	sum := 0
	for _, v := range values {
		sum += v
	}

	count := len(values)
	average := float64(sum) / float64(count)

	return sum, count, average
}

type Dog struct {
	Breed string
	Sound string
}


func (d Dog) Speak() {
	fmt.Printf("The %v says %v\n", d.Breed, d.Sound)
}