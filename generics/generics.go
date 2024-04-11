package main

import (
	"fmt"
)

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))
	
	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice[float32](float32Slice))
	//* fmt.Println(sumSlice(float32Slice)) you can write like that, go compiler will understand
	
}

//* this function could take int, float32 and float64 types
func sumSlice[T int| float32|float64](slice []T) T{
	var sum T
	for _, v:= range slice{
	//? we use v for adding operator but not using directly
	//? to show that we can use _
	//? it means, it won't use directly ??	
		sum+=v
	}
	return sum
}