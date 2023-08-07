package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// if set to int16 this will overflow
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	// using float32 this result will be 1.2345679e+07
	// so best to use int64 to give enough memory
	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	// go is strong typed - ints cannot be added to floats
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	// casting "()"" is needed for type conversion
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	// division will floor by default with ints
	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	// get the remainder with a modulo
	fmt.Println(intNum1 % intNum2)

	// basic string
	var myString string = "Hello World!"
	fmt.Println(myString)
	// multi-line
	myString = "Hello\nWorld!"
	fmt.Println(myString)
	// backtick multi-line
	myString = `Hello
World!`
	fmt.Println(myString)
	// concatination
	myString = "Hello" + " " + "World"
	fmt.Println(myString)

	// NOTE: len is a count of bytes NOT chars
	fmt.Println(len("test"))
	// this will print a len of 2
	fmt.Println(len("ÿ"))

	// This is more accurate - make sure to import unicode/utf8
	fmt.Println(utf8.RuneCountInString("ÿ"))

	// Runescape
	var myRune rune = 'a'
	// prints the ascii table number - base 10
	fmt.Println(myRune)

	// Booleans
	var myBool bool = true
	fmt.Println(myBool)

	// Type inference
	myVar := "test"
	fmt.Println(myVar)

	// multi-var and mixed types
	var1, var2 := 1, "test"
	fmt.Println(var1, var2)

	//specifying type when its not obvious is good practice
	//example:
	//var var3 string = foo()

	// Constants are immutable and cannot be initialized without a value
	// example:
	// const myConst string
	// will not work
	const myConst string = "this is a string"
	const pi float32 = 3.1415
}
