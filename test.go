package main

//? add package main to compile the file
import "fmt"

//? to call print functions

func main() {
	var intNum int32 = 32767
	intNum = intNum + 1
	//? var abbreviation of variable
	//? int16 int32 int64 according to bit of int
	//? same thing valid for float
	fmt.Println(intNum)

}
