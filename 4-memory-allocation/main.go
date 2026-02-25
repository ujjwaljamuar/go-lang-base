package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hey")

	// this wont work because of nil mapping 
	// var m map[string]int
	// m["key"] = 42
	// fmt.Println(m)

	m := make(map[string]int)
	m["key"] = 5

	fmt.Println(m["key"])


	// Pointers
	fmt.Println("Pointers")
	i1 := 43
	var p1 *int = &i1

	if p1 == nil {
		fmt.Println("p1 is nill")
	} else{
		fmt.Println("p1 pointer is poiting value equals to ", *p1)
	}

	f1 := 23.12
	p2 := &f1
	fmt.Println(*p2)

	// changing value through pointer
	f2 := 123.123
	p3 := &f2

	*p3 = *p3 * 100
	fmt.Println(*p3)
	fmt.Println(f2)

}