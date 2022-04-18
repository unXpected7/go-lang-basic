package main

import "fmt"

func main() {

	var myArray [3] int = [3] int{0,1,2}

	normalArray := [4] int {0,1,2,3}

	fmt.Println(myArray)
	fmt.Println(normalArray)

	stringArray := [2] string{"Youssef", "William"}

	fmt.Println(len(stringArray))

	var mySlice = [] int {0,1,2}
	fmt.Println(mySlice)

	slice1 := [] int {1,2,3}
	slice2 := [] int {4,5,6}

	appendSlice := append(slice1, slice2...)
	fmt.Println("appending:::" , appendSlice)
	
}