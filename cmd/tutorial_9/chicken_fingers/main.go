package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
	// create our unbuffered channels
	// unbuffered meaning size not defined
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)

	// create a Slice with each site
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}

	// loop through each site
	// pass in the current website
	// start both concurrent goroutines to check for the chicken and TOFU prices
	// each function will take in a the current website selection and channel
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	// when the thread is unblocked pass in both channels
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		// simuate 'fetch' time
		time.Sleep(time.Second * 1)
		// set the price of the chicken
		var chickenPrice = rand.Float32() * 20
		// check to see if the price is under or equal to the MAX price
		if chickenPrice <= MAX_CHICKEN_PRICE {
			// if it is - write the website to the channel
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		// simuate 'fetch' time
		time.Sleep(time.Second * 1)
		// set the price of the chicken
		var tofuPrice = rand.Float32() * 20
		// check to see if the price is under or equal to the MAX price
		if tofuPrice <= MAX_CHICKEN_PRICE {
			// if it is - write the website to the channel
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	// select is like an if statement for channels
	// the first channel to be assigned a value will take the website variable - then print a value and exit
	select {
	case website := <-chickenChannel:
		fmt.Printf("Found a deal on chicken at %s\n", website)
	case website := <-tofuChannel:
		fmt.Printf("Found a deal on tofu at %s\n", website)
	}

}
