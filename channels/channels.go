package channels

import (
	"fmt"
	"time"
)

func Main() {
	var c = make(chan int,5)
	//* it holds single int value

	go Process(c)
	for i:= range c{
		fmt.Println(i)
		time.Sleep(time.Second*1)
	}
}

func Process(c chan int){
	defer close(c)
	//? defer means run the process (close(c) in this example)
	//? after finishing all process in the located function (func process(c chan int) in this example)
	//? or it may be in main function
	for i:=0; i<5; i++{
		c <- i
	}
	fmt.Println("Existing process")
}