package main 

import "fmt"

func main() {
	var a, i int 
	a = 8
	for i = 1; i <= a; i++ {
		if a % i == 0 {
        fmt.Printf("%d ", i)
    }
	}
}