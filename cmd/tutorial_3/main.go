package main

import "fmt"

func main() {
	// pass value into printMe function
	var printValue string = "Hello World!"
	// note that the type is enforced so I can only pass in a string
	printMe(printValue)

	numerator := 11
	denominator := 2

	// in a function that returns mutliple values I can inline assign multiple values similtaniously
	result, remainder := intDivision(numerator, denominator)
	fmt.Printf("The result of the integer division is %v with the remainder of %v", result, remainder)
}

// printValue is of type string
func printMe(printValue string) {
	fmt.Println(printValue)
}

// when returing a value you need specify the type of the return
// func <funcname>(<param1> <type1>, <param2> <type2>) <return_type>
// We can specify multiple types if there are multiple returns
// we can do this by putting the return types into parenthesis
// func <funcname>(<param1> <type1>, <param2> <type2>) (<return_type1>, <return_type2>))
func intDivision(numerator int, denomonator int) (int, int) {
	var result int = numerator / denomonator
	var remainder int = numerator % denomonator
	// return multiple values with the comma
	return result, remainder
}
