package main
import "fmt"

//* it creates a struct that gasEngine type
type gasEngine struct{
	mpg uint8
	gallons uint8	
	owner
	//* we can directly use other struct inside a struct
	int
}

type owner struct{
	name string
}

type oilEngine struct{
	mpo uint8
	gallons uint8
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

type engine interface{
	milesLeft() uint8
}

func (e oilEngine) milesLeft() uint8{
	return e.gallons*e.mpo
}

func (e electricEngine) milesLeft() uint8{
	return e.mpkwh*e.kwh
}

//! There is a big difference between function and method in Go
//! functions are independent but methods can be used for just specific type
//! for example canMakeIt is a function because it depends nothing
//! but milesLeft is a method, it can be used for just oilEngine and electricEngine types

func canMakeIt(e engine,miles uint8){
	if miles <=e.milesLeft() {
		fmt.Println("You can make it there!")
	}else{
		fmt.Println("Need to fuel up first!")
	}

}


func main(){
	var myEngine gasEngine =gasEngine{25,15,owner{"Alex"},10}
	fmt.Println(myEngine.gallons,myEngine.mpg,myEngine.owner,myEngine.name,myEngine.int)
	//* it's kinda interesting
	
	//----------------------------------------------

	var myEngine2 oilEngine = oilEngine{25,15}
	canMakeIt(myEngine2,50)




}