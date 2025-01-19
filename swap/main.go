package main

import "fmt"

func swap(a, b int) (int,int) {
	return b, a
}

func main() {
	var a, b int
    a = 10
    b = 30
    a, b = swap(a, b)
    fmt.Println("a, b", a, b)
}