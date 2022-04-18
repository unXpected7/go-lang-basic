package main

import "fmt"

func updateName(x *string) {
	*x = "Joey gamed fash5"
}
func main() {

	name:= "joe"

	fmt.Println("The momory location of joe:: ", &name)

	memory:= &name

	fmt.Println("The momory again is:: ", memory)

	orignalName:= *memory

	fmt.Println("The original name is ", orignalName)

	fmt.Println("Before updating ::", name)
	updateName(memory)
	fmt.Println("Before updating ::", name)


}