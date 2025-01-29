package main

import "fmt"

func findsumofdigit(num int) int {
	res := 0
	for num > 0 {
        res += num % 10
        num /= 10
    }
	return res
}

func main() {
	fmt.Println(findsumofdigit(123))
	fmt.Println(findsumofdigit(456))
}

