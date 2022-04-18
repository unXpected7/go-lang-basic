package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	greetings := "Hey there i am greeting you man!"

	fmt.Println(strings.Contains(greetings, "hoe"))
	fmt.Println(strings.ReplaceAll(greetings, "Hey", "Bitch"))
	fmt.Println((greetings))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Index(greetings, "e"))
	fmt.Println(strings.LastIndex(greetings, "e"))
	fmt.Println(strings.Split(greetings, "e"))


	ages:= [] int {30, 40, 10, 3, 1230, 23, 312, 432,4, 43, 4}
	sort.Ints((ages))
	fmt.Println(ages)

	names:= [] string {"joe", "youssef", "maged", "gamed", "awy"}
	sort.Strings(names)
	fmt.Println(sort.SearchStrings(names, "joe"))

}
