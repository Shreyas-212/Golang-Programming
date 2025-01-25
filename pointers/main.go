package main

import "fmt"

func main() {
	var agePointer *int
	age := 32
	agePointer = &age

	fmt.Println("Age: ", *agePointer)

	adultAge := adultAgeValue(&age)
	fmt.Println("Adult Age: ", adultAge)
}

func adultAgeValue(age *int) int {
	return *age - 18
}