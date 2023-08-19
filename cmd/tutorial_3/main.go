package main

import (
	"errors"
	"fmt"
)

func main() {
	// pass value into printMe function
	var printValue string = "Hello World!"
	// note that the type is enforced so I can only pass in a string
	printMe(printValue)

	numerator := 15
	denominator := 3

	// in a function that returns mutliple values I can inline assign multiple values similtaniously
	result, remainder, err := intDivision(numerator, denominator)

	//if err != nil {
	//	fmt.Printf(err.Error())
	//} else if remainder == 0 {
	//	fmt.Printf("The result of the integer division is %v\n", result)
	//} else {
	//	fmt.Printf("The result of the integer division is %v with the remainder of %v\n", result, remainder)
	//}

	// switches are more clean
	// in go breaks are implied by default so no need to specify them
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v\n", result)
	default:
		fmt.Printf("The result of the integer division is %v with the remainder of %v\n", result, remainder)
	}

	// condititonal switch
	// cases are based on value of variable.
	switch remainder {
	case 0:
		fmt.Printf("The Division was exact\n")
	case 1, 2:
		fmt.Printf("The Division was close\n")
	default:
		fmt.Printf("The Division was not close\n")
	}
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
func intDivision(numerator int, denomonator int) (int, int, error) {
	var err error
	if denomonator == 0 {
		err = errors.New("Cannot Divide by Zero\n")
		return 0, 0, err
	}
	var result int = numerator / denomonator
	var remainder int = numerator % denomonator
	// return multiple values with the comma
	// if there is no error - error will return nil back to the caller
	return result, remainder, err
}
