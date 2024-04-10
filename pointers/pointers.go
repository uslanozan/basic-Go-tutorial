package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	//* if we don't write new(int32) it gives runtime error
	//* nil pointer derefernce

	var i int32
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe address of p points to is: %v", p)
	fmt.Printf("\nThe value if is: %v", i)
	p = &i
	*p = 1
	//* p refers memomery address of i
	fmt.Printf("\nThe value p points to is: %v", *p)
	fmt.Printf("\nThe address of p points to is: %v", p)
	fmt.Printf("\nThe value if is: %v", i)
	var k int32 = 2
	i = k
	//* i is equal to value of k, not memory address
	fmt.Println("\n--------------------------------")

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	//* original slice's been changed
	//! they refers same memory address
	fmt.Println(sliceCopy)
	fmt.Println("\n--------------------------------")

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result is: %v", result)
	var result2 [5]float64 = squareV2(&thing1) //!
	fmt.Printf("\nThe result is: %v", result2) //!
	fmt.Printf("\nThe value of thing1 is: %v", thing1)

	//! the purpose of squareV2 is don't wasting memory
	//! in square function, because we don't use pointers, it will creates a new array
	//! but we use pointers in squareV2 so they use the same memory
	//! red lines is show that

}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

func squareV2(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing1 array is: %p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
