package main

import "fmt"

func main() {
	// arrays are:
	// fixed length
	// same type
	// indexable
	// contigous in memory

	// this will initialize an empty array
	//var intArr [3]int32
	// if I want to pre-populate data i can:
	//var intArr [3]int32 = {1,2,3}
	//OR use type inference and dot notation (to detect length)
	intArr := [...]int32{1, 2, 3}
	// mutation
	//intArr[1] = 123
	// get zeroth element
	//fmt.Println(intArr[0])
	// get second element
	//fmt.Println(intArr[1:3])
	fmt.Println(intArr)

	// an int32 is 4bytes of memory
	// our array is a length of 3
	// 4*3 = 12
	// the addresses will be contigous in memory

	// use the AMP & to dref
	//fmt.Println(&intArr[0])
	//fmt.Println(&intArr[1])
	//fmt.Println(&intArr[2])

	// Slices
	// Slices are a wrapper around Arrays
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	// go is taking the original Array and copying everything from it into a new array to accomidate the new length
	// we can see this by comparing length to capacity after the append
	fmt.Printf("The length is %v with a capacity of %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with a capacity of %v\n", len(intSlice), cap(intSlice))
	// trying to access elements outside of index length will result in an index out of range error
	//fmt.Println(intSlice[4])

	// you can add multiple values to a slice with the spread operator
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// we can initialize a slice with a preset length and capacity with the 'make' keyword
	// var <name> <type> = make(<type>, length, capacity)
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	//Maps
	// key => value pairs key is type string and values are uint8
	// var <name> map[<key_type>]<value_type>
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	// if I try to get a value that doesn't exist - go will return the default value
	// the default for uint8 is 0
	// a map will ALWAYS return - even if the key doesn't exist
	fmt.Println(myMap2["Jason"])
	// maps return a second value that is a bool to determine if the value is part of the map
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Printf("Invalid Name\n")
	}
	// values can be deleted from maps using the delete function
	//delete(myMap2, "Adam")

	// loops
	// no order is preserved when looping over maps
	// No while loops but we can achieve similar functionality with for loops
	for name, age := range myMap2 {
		fmt.Printf("Name: %v Age: %v\n", name, age)
	}
	// looping Arrays
	for i, v := range intArr {
		fmt.Printf("Index: %v Value: %v\n", i, v)
	}

	// iteration
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}

	//tradional c method
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
