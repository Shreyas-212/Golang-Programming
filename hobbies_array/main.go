package main

import "fmt"

type products struct {
	title string
	id string
	price float64
}

func main() {
	var hobbies = [3]string{"Cycling" , "Running" , "Swimming"}
	var goals = []string{"Package", "Arrays"}
	product1 := products{title: "Computers", id: "bbr57rd", price: 22000}
	product2 := products{title: "Mobiles", id: "n33nmms", price: 16000.03}
	var product = []products{product1, product2}

	slice1 := hobbies
	fmt.Println(slice1)

	slice2 := hobbies[:1]
	fmt.Println(slice2)

	slice3 := slice2[1:3]
	fmt.Println(slice3)

	slice4 := hobbies[0:2]
	//slice4 := hobbies[:2]
	fmt.Println(slice4)
		
	slice5 := hobbies[1:3]
	fmt.Println(slice5)

	slice6 := goals
	fmt.Println(slice6)

	goals[1] = "Types"
	goals = append(goals, "slices")
	fmt.Println(goals)

	fmt.Println(product)
	product3 := products{title: "Speakers", id: "h33hggf", price: 1500.45}
	product = append(product, product3)
	fmt.Println(product)
}