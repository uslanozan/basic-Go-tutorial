package main

import (
	"fmt"
	"log"
)

//? new:
//? new function is used to allocate memory for a new value of a given type.
//? It allocates memory space and initializes it to the zero value of the type (e.g., 0 for an integer or an empty string).
//? new function returns a pointer to the newly allocated memory. It provides a pointer pointing to the value.
//? new function is commonly used to create new objects of primitive data types or structure types.

//? make:
//? make function is used only for certain special data types like slices, maps, and channels.
//? It creates a new instance of a specific type and allocates the necessary memory for its initialization.
//? make function sets up specific properties of the allocated memory based on the type (e.g., length and capacity for a slice or buffer size for a channel).
//? make function returns a value of the specific data type. It does not return a pointer.

//-----------------------------------------------------------\\

//* we can assign a variable in if statement
//? 	func main(){
//? 		err := foo()
//? 		if err!= nil {
//? 			log.Print(err)
//? 		}
//?
//? 	}

func main() {
	if err := foo(); err != nil {
		log.Print(err)
	}


	i,_:=foo2()
	//? normally foo2 returns 2 variable, but we use ,
	//? it means to ignore after i 
	fmt.Println(i)
}



func foo() error {
	return nil
}

func foo2() (int,error){
	return 1,nil
}