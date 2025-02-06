package main

import "fmt" 

type transformFn func(int) int  // function type for transforming numbers

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	morenumbers := []int{6, 7, 8, 9, 10}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println("Original numbers:", numbers)
	fmt.Println("Doubled numbers:", doubled)
	fmt.Println("Tripled numbers:", tripled) 

	transformfn1 := getTransformerFunction(&numbers)
	transformfn2 := getTransformerFunction(&morenumbers)

	transformedNumbers := transformNumbers(&numbers, transformfn1)
	moreTransformedNumbers := transformNumbers(&morenumbers, transformfn2)

	fmt.Println("Transformed numbers:", transformedNumbers)
	fmt.Println("More transformed numbers:", moreTransformedNumbers)
}

func transformNumbers(numbers *[]int, transform transformFn) []int{
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(val int) int {
	return val * 2
}

func triple(val int) int {
	return val * 3
}