package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readingInputs(promt string, reader* bufio.Reader)(string, error){
	fmt.Print(promt)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func readToAddItem(reader *bufio.Reader, bill *Bill){
	itemName, _ := readingInputs("Add item's name: ", reader)
	itemPrice, _:= readingInputs("Add item's price: ", reader)
	convertedPrice, err:= strconv.ParseFloat(itemPrice, 64)
	if err != nil {
		fmt.Println("Price must be a number")
		promtOptions(bill)
	}
	bill.addMenuItem(itemName, convertedPrice)
	fmt.Println("Your item is added - ", itemName, itemPrice)
	promtOptions(bill)

}

func readToAddTip(reader *bufio.Reader, bill *Bill){
	tip, _ := readingInputs("Add tip to the bill: ", reader)
	convertedTip, err:= strconv.ParseFloat(tip, 64)
	if err != nil{
		fmt.Println("Price must be a number")
		promtOptions(bill)
	}
	bill.updatingTips(convertedTip)
	fmt.Println("Your tip is added - ",bill)

	promtOptions(bill)
}

func promtOptions(b *Bill){
	inputReader := bufio.NewReader(os.Stdin)
	options, _ := readingInputs("Choose an option ('a' -> add item, 's' -> save bill, 't' -> add tips): \n",inputReader)
	switch options {
	case "a":
		readToAddItem(inputReader, b)
	case "t":
		readToAddTip(inputReader, b)
	case "s":
		fmt.Println("You chose to save your bill", b)
		b.saveFile()
	default:
		fmt.Print("Invalid options, please specify again")
		promtOptions(b)
	}
}

func createBill() Bill {
	inputReader := bufio.NewReader(os.Stdin)

	name, _ :=readingInputs("Create a new bill, what's your name? \n",inputReader)

	bill := newBill(name)

	fmt.Printf("Created the bill under %q name \n", name)

	promtOptions(&bill)
	return bill
}
func main() {

	myBill := createBill()
	// fmt.Println("\nmy new bill",myBill)
	// name:= "joey"
	// myBill := createBill(name)
	fmt.Println(myBill.format())

	// fmt.Println("Original tips::", myBill.tip)
	// myBill.updatingTips(4)
	// fmt.Println("Updated tips::", myBill.tip)

	// myBill.addMenuItem("beer", 20)
	// fmt.Println(myBill.format())

}