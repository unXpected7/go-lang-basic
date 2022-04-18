package main

import "fmt"

func main() {

	//strings
	var oneWay string = "Joe"
	var twoWay = "Joeyyy"
	var threeWay string

	fmt.Println(oneWay, twoWay, threeWay)

	oneWay = "no man can handle this"

	fmt.Println(oneWay)

	fourWay := "My name is joe"
	oneWayNumber := 4

	fmt.Println(fourWay)
	fmt.Println(oneWayNumber)


	//int
	var oneWayInt int = 5
	var twoWayInt = 4
	var threeWayInt int 

	fourWayInt := 90

	fmt.Println(oneWayInt, twoWayInt, threeWayInt, fourWayInt)

	var bitInt int32 = 2066765425

	fmt.Println(bitInt)


	var oneWayFloat float32 =34.4
	var twoWayFloat string

	threeWayFloat := 349.063423132

	fmt.Println(oneWayFloat,twoWayFloat, threeWayFloat)

}