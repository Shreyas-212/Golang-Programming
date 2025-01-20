package main 

import "fmt"

func Split(num int) (x, y int) {
	x = num * 4 / 9
	y = num - x
	return 
}

func main() {
	fmt.Println(Split(22))
}