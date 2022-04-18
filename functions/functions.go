package main

import (
	"fmt"
	"math"
	"strings"
)

func sayHi(name string) {
	fmt.Printf("Hey there, i am saying hi because my name is %q \n", name)
}

func autoSayHi(names[] string, callHiFunction func(string)){
	for _, value := range names {
		callHiFunction(value)
	}
}

func getTheLength(array[]string) int{
	return len(array)
}

func getTheArea(radius float64) float64 {
	area:= math.Pi * math.Pow(radius,2)
	return area
}

func getInitials(name string) (string, string){
	strings.ToUpper(name)
	splitedArray := strings.Split(name, " ")
	var returenedArray [] string
	for _, value := range splitedArray {
		returenedArray = append(returenedArray, value[:1])
	}
	return returenedArray[0] , returenedArray[1]
}

func main() {
	names:= [] string {"joe", "youssef", "maged", "gamed", "awy"}
	autoSayHi(names, sayHi)
	length:= getTheLength(names)
	fmt.Println("the length", length)
	area:= getTheArea(3.14)
	fmt.Printf("the area %.1f \n", area)

	s1,s2:= getInitials("Youssef Maged")
	fmt.Println("the strings are::", s1, s2)
}