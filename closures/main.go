package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)

	// transformed := transformNumbers(&numbers, func(number int) int {  
	// 	return number * 2
	// }) 
	
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

    // fmt.Println(transformed)
	fmt.Println("Original numbers:", numbers)
	fmt.Println("Doubled numbers:", doubled)
	fmt.Println("tripled numbers:", tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {       //Factory function
	return func(number int) int {                        //Every Anonymous function is called closure function
		return number * factor
	}
}