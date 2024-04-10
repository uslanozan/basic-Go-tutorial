package main

import (
	"fmt"
	"sync"
	"time"
)

//? go can do multi tasking
//? Concurrent vs. Parallel Programming: Concurrency focuses on managing task dependencies and communication, 
//? while parallelism focuses on actual parallel execution of tasks on multiple processing units.

var wg=sync.WaitGroup{}
var dbData=[]string{"id1","id2","id3","id4","id5"}

func main(){
	t0 :=time.Now()
	for i:=0;i <1000; i++{
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v",time.Since(t0))

}

func dbCall(i int) {
	//Simulate DB call delay
	var delay float32=2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	wg.Done()
}

/*
var m=sync.RWMutex{}
var dbData=[]string{"id1","id2","id3","id4","id5"}
var wg=sync.WaitGroup{}
var results=[]string{}

func main(){
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v",time.Since(t0))
	fmt.Printf("\nThe results are %v",results)

}

func dbCall(i int)  {
	// Simulate DB call delay
	var delay float32=2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string){
	m.Lock()
	results=append(results, result)
	m.Unlock()
}

func log(){
	m.RLock()
	fmt.Printf("\nThe current results are: %v",results)
	m.RUnlock()

}
*/