package main

import "fmt"

func main() {

	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr[0])
	fmt.Println(intArr[0:3])

	//*memory addresses
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	//* we can add or remove items from slices
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	//* capacity is different than length

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	//* it makes a slice that has 3 length and 8 capacity
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice3), cap(intSlice3))

	//* it means keys type are string and values type are unsigned int8
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam":23,"Sarah":45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Sarah"])
	fmt.Println(myMap2["Jason"])
	delete(myMap2,"Jason")
	//* tough there is no key as "Jason", default value is 0
	//* maps always return something even if key doesn't exist
	var age,ok = myMap2["Michael"]
	//? second variable (ok) is optional and it returns boolean according to existing
	if ok {
		fmt.Println("The age is %v",age)
	}else{
		fmt.Println("Invalid name")
	}

	
	for name,age := range myMap2{
		fmt.Printf("Name: %v, Age: %v \n",name,age)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}



}
