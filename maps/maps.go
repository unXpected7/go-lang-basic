package main

import "fmt"

func main(){

	menuMap := map[string]float64{
		"pizza": 12.76,
		"burger": 10.99,
		"shawarma": 3,
		"cola": 1.99,
	}

	for key, value := range menuMap {
		fmt.Printf("The damn items are %q and their prices are %0.1f", key, value)
		fmt.Println(value)
	}
	fmt.Println("the price of cola::", menuMap["cola"])

	contacts := map[int]string{
		1002021414: "Youssef",
		1001092002: "Maged",
		1001579090: "Amal",
	}
	for number, name:= range contacts {

		fmt.Printf("The number is %v and it's %q \n", number, name)
	}
	fmt.Println("Original name:::", contacts[1002021414])

	contacts[1002021414] = "Joey"

	fmt.Println("Now the updated name:::", contacts[1002021414])
}