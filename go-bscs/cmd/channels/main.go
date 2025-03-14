package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var MAX_CHICKEN_PRICE = 5;
var MAX_TOFU_PRICE = 3

func main() {
	var webWithChickenDealChannel = make(chan string)
	var webWithTofuDealChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefood.com"}

	for i:=0; i<5; i++ {
		go checkChickenPrices(websites[i], webWithChickenDealChannel)
		go checkTofuPrices(websites[i], webWithTofuDealChannel)

	}

	sendMessage(webWithChickenDealChannel, webWithTofuDealChannel)
}

func checkChickenPrices(website string, webWithChickenDealChannel chan string){
	for {
		time.Sleep(time.Second+1)
		var chickenPrice = rand.Int()*20
		if chickenPrice <= MAX_CHICKEN_PRICE{
			webWithChickenDealChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, webWithTofuDealChannel chan string){
	for {
		time.Sleep(time.Second+1)
		var tofuPrice = rand.Int()*20
		if tofuPrice <= MAX_TOFU_PRICE{
			webWithTofuDealChannel <- website
			break
		}
	}
}

func sendMessage(webWithChickenDealChannel chan string, webWithTofuDealChannel chan string){
	//fmt.Printf("\nFound a deal on chicken at%s", <- webWithChickenDealChannel)
	select{
		case website := <-webWithChickenDealChannel:
			fmt.Printf("\nFound a deal on chicken at %v", website)
		case website := <-webWithTofuDealChannel:
			fmt.Printf("\nFound a deal on tofu at %v", website)
	}
}