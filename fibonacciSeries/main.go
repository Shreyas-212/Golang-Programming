package main 

import "fmt"

func PrintFibonacci(num int) {
	a := 0; b := 1
	c := b
	for true {
		c = b; b = a + b
		if b >= num {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf(" %d", b)
	}
}

func main() {
	PrintFibonacci(10)
	PrintFibonacci(100)
}