package main

//? add package main to compile the file

import (
	"fmt"
	"unicode/utf8"
) //? to call print functions

func main() {
	var intNum int32 = 32767
	intNum = intNum + 1
	//? var abbreviation of variable
	//? int16 int32 int64 according to bit of int
	//? same thing valid for float
	fmt.Println(intNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	//! we can typecasting like that
	fmt.Println(result)

	var myString string = "Hello \nWorld"
	fmt.Println(myString)

	fmt.Println(len("Y"))
	//* normally it must return a length according to binary but
	//* it works as we expected, I don't know why
	fmt.Println(utf8.RuneCountInString("Y"))
	//* we should have to use this, but not necessary ?

	var myRune rune = 'a'
	fmt.Println(myRune)
	//! rune represent a Unicode of that string, it's like char in Java but still don't know exactly

	//?      ---Default values---
	//? for int, unsigned int, float, rune ... is 0
	//? for string is '' (empty string)
	//? for boolean is false

	//* you don't have to write var
	myVar := "text"
	fmt.Println(myVar)
	//* it could understand which type is it

	const myConst string = "Const values cannot be changed!"
	//! myConst="as" is wrond, consts cannot be changed
	fmt.Println(myConst)

}
