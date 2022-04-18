package main

import "fmt"

func main() {

	names:= [] string {"Joe", "Youssef", "Maria", "Ginny"}

	for i := 0; i < len(names); i++ {
		fmt.Println("My name is ", names[i])
	}

	for index, value := range names {
		fmt.Printf("The index is %v and the its value is %q \n", index, value)
	}

	age:= 23

	if age > 20 {
		fmt.Println("You are an adult")
	}
	if age > 22 {
		fmt.Println("You just graduated college")
	}
	if age <= 23 {
		fmt.Println("You will be great someday")
	}
}