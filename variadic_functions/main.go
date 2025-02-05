package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumUp(1, 10, 15, 40)
	anotherSum := sumUp(numbers...)

	fmt.Printf("Sum of the numbers: %d\n", sum)
	fmt.Printf("Sum of the numbers: %d\n", anotherSum)
}

func sumUp(number ...int) int {           //Variadic Function
	sum := 0

	for _, val := range number {
        sum += val
    }
	return sum
}