package main

import "fmt"

func main() {
	// make a channel that will hold an int type and name it 'c'
	// if we dont declare a size of the buffer it will be an unbuffered channel
	// lets create a buffer of 5
	// this will allow the goroutine to keep writing to the channel without waiting for the main function to continously call it
	var c = make(chan int, 5)
	// we can add a value to a channel using an <- operator
	// a channel has an underlying array
	// this is an unbuffered channel so it only has enough room for 1 value
	// when we write to an unbuffered channel - the code will block until something reads from it.
	//c <- 1
	// we can assign a channel to a variable uing the following syntax
	// this will pop the value out of the channel and store it in the variable i
	// the problem is the thread will lock on the line above and never make it to this line
	//var i = <- c
	// if we just ran this - this would reach a deadlock - so it we should use these with goroutines
	//fmt.Println(i)

	// use a goroutine thread to ascend into the process function and set the channel
	go process(c)

	// print all the set values by a chan
	// since I'm writing to the chan in a loop - I need to read from it in a loop as well
	for i := range c {
		fmt.Println(i)
	}
}

func process(c chan int) {
	// once the for loop is done we need to inform all other functions that are reading that the channel is closed
	// since the upstream loop is just waiting for more data as long as the channel is active
	// once we close the channel we tell readers that no more data is being sent by the channel.
	// we can use the defer statement to say "make sure this is ran before the function exits"
	defer close(c)
	//c <- 123
	// lets add some loops
	for i := 0; i < 5; i++ {
		c <- i
	}

}
