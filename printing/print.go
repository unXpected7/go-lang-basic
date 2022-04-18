package main

import "fmt"

func main() {
	// fmt.Print("New line")
	// fmt.Print("No new line \n")
	// fmt.Print("now we have a new line")

	ageInInt := 23
	ageInFloat := 23.56
	name:= "William"
	isTrue:= false

	fmt.Println("My name is: ", name, " and my age is: ", ageInInt)
	fmt.Printf("Formatting my name %q and the age also gonna be %v \n", name,ageInInt)
	fmt.Printf("booleans is of type %T \n", isTrue)
	fmt.Printf("name is of type %T \n", name)
	fmt.Printf("float is of type %T \n", ageInFloat)
	fmt.Printf("Age is of type %T \n", ageInInt)
	fmt.Printf("Print a float with %f \n", ageInFloat)
	fmt.Printf("Print a float with %0.1f with a limit \n", ageInFloat)

	savedString := fmt.Sprintf("The saved string contains a name of: %v and an age of: %v  \n", name, ageInInt)
	fmt.Println(savedString)


}