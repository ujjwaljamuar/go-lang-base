package main

import (
	"fmt"
)

func main() {
	const i1 int = 64
	const i2 int = 64
	iSum := i1 + i2

	fmt.Printf("Sum = %v\n",iSum)

	const f1 float64 = 32.12
	const f2 float64 = 322.32
	const f3 float64 = 12.5
	const f5, f6, f7 = 11.1, 12.5, 123.3

	fSum := f1 + f2 + f3

	fmt.Printf("Sum = %v\n", fSum)

	ifSum := float64(i1) + f1
	fmt.Printf("Sum = %v\n", ifSum)

	f2Sum := f5 + f6 + f7
	fmt.Printf("Sum = %v\n", f2Sum)

}