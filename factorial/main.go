package main

import (
	"fmt"
)

func Factorial(num int) {
	if num < 0 {
		fmt.Println("Factorial is not defined for negative numbers")
		return
	}
	fac := 1 
	for i := num; i > 0; i-- {
		fac *= i
	}
	fmt.Printf("Factorial of %d is %d\n", num, fac)
}

func main() {
	Factorial(3) 
	Factorial(10)
	Factorial(-4)
}