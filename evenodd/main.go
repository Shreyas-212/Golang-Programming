package main 

import "fmt" 

func main() {
	var a, b int 
    a = 3
    b = 2
	if a % 2 == 0 {
		fmt.Println("a is even number")
	} else {
		fmt.Println("a is odd number")
	}
	
	if b % 2 == 0 {
        fmt.Println("b is even number")
    } else {
		fmt.Println("b is odd number")
	}
}