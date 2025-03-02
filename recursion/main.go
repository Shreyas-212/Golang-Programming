package main

import "fmt"

func main() {
	fact := factorial(5)

	fmt.Printf("Factorial of 5 is %d\n", fact)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number - 1)
}