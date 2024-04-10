package main

import (
	"fmt"
	"math/rand"
	"time"
)

//? channels are made for synchronization and simultaneous communication so they're useful

var MAX_CHICKEN_PRICE float32=5
var MAX_TOFU_PRICE float32=3

func main(){
	var chickenChannel=make(chan string)
	var tofuChannel=make(chan string)
	var websites=[]string{"walmart.com","costco.com","wholefoods.com"}
	for i:=range websites{
		go checkChickenPrices(websites[i],chickenChannel)
		go checkTofuPrices(websites[i],tofuChannel)
	}
	sendMessage(chickenChannel,tofuChannel)
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for	{
		time.Sleep(time.Second*1)
		var tofuPrice = rand.Float32()*20
		if	tofuPrice <= MAX_TOFU_PRICE{
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string,tofuChannel chan string) {
	//? Switch statements work with constant expressions, while select statements work with channels.
	//? Switch statements are used with conditional expressions, 
	//? while select statements are used in concurrent programming to facilitate communication and synchronization between parallel processes.
	select{
		case website := <-chickenChannel:
			fmt.Printf("\nText Sent: Found a deal on chicken at %v",website)
		case website := <- tofuChannel:
			fmt.Printf("\nEmail Sent: Found a deal on tofu at %v",website)
	}
}

func checkChickenPrices(website string, chickenChannel chan string) {
	//? there is just for loop instead of while,do-while
	//? but you can do the same thing with different uses for loop
	for	{
		time.Sleep(time.Second*1)
		var chickenPrice = rand.Float32()*20
		if	chickenPrice <= MAX_CHICKEN_PRICE{
			chickenChannel <- website
			break
		}
	}
}