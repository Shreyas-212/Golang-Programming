package main 

import "fmt"

func main() {
	var prices = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(prices)
	fmt.Println(len(prices))

	featuredValues := prices[0:6]
    fmt.Println(featuredValues)
	fmt.Println(len(featuredValues))

	exclusiveValues := featuredValues[0:9] // Counts towards right but not left
	fmt.Println(exclusiveValues)
	fmt.Println(len(exclusiveValues))	
}
