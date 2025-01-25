package main

import "fmt"

func main() {
	var agePointer *int
	age := 32
	agePointer = &age

	fmt.Println("Age: ", *agePointer)

	EditAdultAgeValue(&age)
	fmt.Println("Age: ", age)
}

func EditAdultAgeValue(age *int) {
	*age = *age - 18
}