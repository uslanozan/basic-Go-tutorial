package main

import (
	"fmt"
	"strings"
)

// ! utf-8 characters doesn't work and idk why
func main() {
	var myString = "cbx7"
	var indexed = myString[0]
	fmt.Printf("%v, %T", indexed, indexed)

	for index, value := range myString {
		fmt.Println(index, value)
	}
	fmt.Printf("\nThe length of 'myString' is %v", len(myString))

	var myRune = '8'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.Write([]byte(strSlice[i]))
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)

}
