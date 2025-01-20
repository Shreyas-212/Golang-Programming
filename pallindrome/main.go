package main

import "fmt"

func main() {
	var num ,temp, rem, rev int = 1221, 0, 0, 0
	
	rev = 0
	temp = num
	for temp != 0 {
		rem = temp % 10
		rev = rev*10 + rem
		temp /= 10
	}
	if num == rev {
		fmt.Println("The number is a palindrome")
	} else {
		fmt.Println("The number is not a palindrome")
	}
}
