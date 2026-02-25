package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Blue"
	colors[2] = "Green"

	fmt.Println(colors)

	var nums = [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	fmt.Println("Length of colors arrays is: ", len(colors))
	fmt.Println("Length of nums arrays is: ", len(nums))
}
