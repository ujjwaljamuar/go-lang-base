package main

import "fmt"

func main() {
	// var
	var x int = 42
	fmt.Println("value of x is ", x)
	var y string  = "My name is ujjwal"
	// implicit typing
	a := 42
	b := "My name is Ujjwal Jamuar."


	// const
	const c = "Hello, How are you?"

	d := true


	fmt.Println(b)
	
	fmt.Println(y)
	fmt.Println(a)
	fmt.Println(d)
	fmt.Println("Hey Go")


	const e float32 = 31.4
	const f float64 = 12134.324
	fmt.Println(e)
	fmt.Println(f)

	const g complex64 = 21345.23
	const h complex128 = 2134533245.1234
	fmt.Println(g)
	fmt.Println(h)


	stringLength, err := fmt.Println("value of x is ", x)
	if err == nil {
		fmt.Println("String length: ", stringLength)
	}

	// %v = value
	// %T = type
	fmt.Printf("Value of x is %v\n", x)
	fmt.Printf("Type of x is %T\n", x)


	

	// fixed integer types
	// uint8
	// uint16
	// uint32
	// uint64

	// int8
	// int16
	// int32
	// int64

	// aliases
	// byte
	// uint
	// int 
	// uintptr

	// Arrays
	// Slices
	// Maps
	// Structs

	// Functions
	// Interfaces
	// Channels
	
	// Pointers
}