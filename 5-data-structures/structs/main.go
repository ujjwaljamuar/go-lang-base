package main

import (
	"fmt"
	"sort"
)

func main() {
	// arrays has a fixed size, same element types
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Blue"
	colors[2] = "Green"

	fmt.Println(colors)

	var nums = [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	fmt.Println("Length of colors arrays is: ", len(colors))
	fmt.Println("Length of nums arrays is: ", len(nums))

	// slices, resizable, same element types

	// make(type, initital items, capacity)
	var colors1 = make([]string, 0, 3)
	fmt.Println(colors1)

	colors1 = append(colors1, "Yellow", "Orange", "Silver")
	fmt.Println(colors1)

	// append items
	colors1 = append(colors1, "Gold")
	fmt.Println(colors1)

	removeItems(colors1, 2)
	fmt.Println(colors1)

	sort.Strings(colors1)
	fmt.Println(colors1)

	// maps	
	country := make(map[string]string)
	country["IN"] = "India"
	country["USA"] = "United States of America"
	country["CH"] = "China"

	fmt.Println(country)

	delete(country, "CH")
	fmt.Println(country)

	for k, v := range country {
		fmt.Printf("%v : %v\n",k, v)
	}

	learnStruct()
}

func removeItems(slc []string, ind int) []string {
	return append(slc[:ind], slc[ind+1:]...)
}


type Dog struct {
	Breed string
	Weight int
}

func learnStruct(){
	poodle := Dog{"Poodle", 35}
	fmt.Println(poodle)

	fmt.Printf("%+v\n", poodle)

	fmt.Printf("Breed: %v, Weight: %v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 29
	fmt.Printf("Breed: %v, Weight: %v\n", poodle.Breed, poodle.Weight)
}
