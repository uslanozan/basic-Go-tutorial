package main

import (
	"fmt"
	"sync"
	//"time"
)

var wg sync.WaitGroup
//? 3 important method: add,wait,done

func main() {

	wg.Add(1)
	//? you have to wait for 2 different goroutines

	go printX()
	//fmt.Println()
	wg.Wait()
	//? wait untill the previous goroutine is finished
	printY()
	//? time.Sleep(time.Second)
	//? it waits 1 sec

}

func printY() {
	for i := 0; i < 20; i++ {
		fmt.Print("Y")
	}
}

func printX() {
	for i := 0; i < 500; i++ {
		fmt.Print("X")
	}
	wg.Done()
	//? process is done
}