package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	// strings have indexes just like arrays (arrray of chars)
	var indexed = myString[0]
	// when printing the indexed value we get a number
	// we can see that it is an uint8
	fmt.Printf("%v %T\n", indexed, indexed)
	// iterate over the string
	for i, v := range myString {
		fmt.Printf("%v %v\n", i, v)
	}
	// utf-8 does some magic with non standard chars in arrays
	// in this case é doesn't fit in one index
	// the output of the above index is:
	// 0 114 -> 01110010
	// 1 233 -> 11000011 10101001 < note that this takes the index of 2 as well to represent this char
	// 3 115 -> 01110011
	// 4 117 -> 01110101
	// 5 109 -> 01101101
	// 6 233 -> 11000011 10101001 <- this takes the index of 7 just like it did with 2

	// taking the length of the string will show the bytes NOT the char length - for example below will be 8
	fmt.Printf("The length of myString is %v\n", len(myString))

	// an easier method is just to cast an array of runes
	// runes are uncode point numbers that encode the char
	// runes are an alias for int32
	var myString2 = []rune("résumé")
	var indexed2 = myString2[1]
	fmt.Printf("%v %T\n", indexed2, indexed2)
	for i, v := range myString2 {
		fmt.Printf("%v %v\n", i, v)
	}
	// because we are using runes this will be the actual length
	fmt.Printf("The length of myString2 is %v\n", len(myString2))

	// we can declare a rune using single quotes as well
	var myRune = 'a'
	fmt.Printf("myRune = %v\n", myRune)

	// DIY string
	// Strings are immutible in go - meaning they cannot be changed after being declared.
	// for example catStr[0] = 'a' would not work
	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	//var catStr = ""
	var strBuilder strings.Builder
	// build a new string builder

	for i := range strSlice {
		// we can concatinate with the += operator
		// this is inefficient because we build a new string each time.
		//catStr += strSlice[i]
		// instead we can use the string builder
		strBuilder.WriteString(strSlice[i])
		// This allocates an array and appends as new values are added
	}
	// define catStr as the output string here
	// The string is only created once here
	var catStr = strBuilder.String()
	fmt.Printf("%v\n", catStr)

}
