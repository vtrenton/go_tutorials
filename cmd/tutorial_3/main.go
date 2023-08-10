package main

import "fmt"

func main() {
	// pass value into printMe function
	var printValue string = "Hello World!"
	// note that the type is enforced so I can only pass in a string
	printMe(printValue)
}

// printValue is of type string
func printMe(printValue string) {
	fmt.Println(printValue)
}

// when returing a value you need specify the type of the return
// func <funcname>(<param1> <type1>, <param2> <type2>) <return_type>
func intDivision(numerator int, denomonator int) int {
	var result int = numerator / denomonator
	return result
}
