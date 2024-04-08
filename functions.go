package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Hello World")
	var printValue string = "Hello Mars"
	printMe(printValue)

	numerator := 11
	denominator := 0
	//-------------------------------------------------------------------------
	var result, err = intDivision(numerator, denominator)
	//? we can think nil is null in Java
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(result)

	//-------------------------------------------------------------------------

	var result2, remainder, err2 = intDivision2(numerator, denominator)
	/*
	if err2 != nil {
		fmt.Printf(err.Error())
	}else if remainder==0 {
		fmt.Printf("The result of the integer division is %v ",result2)
    }else{
		fmt.Printf("The result of the integer division is %v with remainder ",result2,remainder)
    }
	*/

	//? we can write same one with switch case
	//* it doesn't need to used BREAK like Java
	switch{
		case err!=nil:
			fmt.Printf(err2.Error())
		case remainder==0:
			fmt.Printf("The result of the integer division is %v",result2)
		default:
			fmt.Printf("The result of the integer division is %v with remainder %v",result2,remainder)
	}

	//? we can write like that

	switch remainder {
	case 0:
		fmt.Printf("\nThe division is was exact")
	case 1,2:
		fmt.Printf("\nThe division was close")
	default:
		fmt.Printf("\nThe division was not close")	
	}



	//fmt.Println("The result of the integer division is %v with remainder %v", result2, remainder)
	//fmt.Printf("The result of the integer division is %v with remainder %v", result2, remainder)
	//? to use %v methods you must use printf
	//-------------------------------------------------------------------------

}

// * you must put 1. curly brace like that
func printMe(printvalue string) {
	fmt.Println(printvalue)
}

// ? it returns integer
func intDivision(numerator int, denominator int) (int, error) {

	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, err
	}

	var result int = numerator / denominator
	return result, err
}

// ? it could be returns TWO integer
func intDivision2(numerator int, denominator int) (int, int, error) {

	//* error object is really useful for catching errors
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
