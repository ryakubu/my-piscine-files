package main

import (
	"fmt"
	"strings"
)

func main() {

	var name string = "Rinret Yakubu"
	fmt.Println("My name is ", name)
	str := "Rinret Yakubu"
	length := len(str)
	fmt.Println("The length is ", length)

	str1 := "Go Programming"
	str2 := "programming"
	strings.EqualFold(str1, str2)
	fmt.Println(strings.EqualFold(str1, str2))

}
