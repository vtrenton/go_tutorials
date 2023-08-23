package main

import "fmt"

func main() {
	// the variable p will hold the memory address of an int32 value.
	// the star '*' is used to declare a new pointer - NOTE this is also used to dereference it later.
	// I can initialize the pointer using the new keyword
	// without the new keyword nil is all that is stored and you cannot deref a null pointer.
	// This will cause the pointer to store the memory address of the newly initialized int32 in memory
	// int32 is initialized with the default value of 0
	// meaning the pointer is referencing a memory address storing the value 0
	var p *int32 = new(int32)
	var i int32
	// I can get the value stored at the address space stored in p (or pointed to by p) by dereferencing p
	// I can dereference a pointer by adding a * behind it like below
	// Below will output 0 and 0 because I haven't assgined a value yet
	fmt.Printf("The dereferenced value of out pointer p is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)

	// I can make the pointer p instead point to the address of i
	// This can be accomplished using the amp '&' after the variable name
	// This stores the memory address instead of the value of the variable
	// in this case the pointer p will now store the address of the variable i
	p = &i

	// By default - The 'zeroed out' or default values for objects are used to populate the memory addresses
	// we can change this by assigning the derefereced address of the pointer
	// this will populate the memory address stored in the pointer with the assigned value
	*p = 10

	// Note that we assigned p to the address of i earlier
	// We then assigned the derefenced value of p (which is the address of i) to 10
	// this means we assigned i = 10
	fmt.Printf("The dereferenced value of out pointer p is now: %v\n", *p)
	fmt.Printf("The value of i is now: %v\n", i)

	// Pointers are different from value copies
	var k int32 = 2
	// this copies the value of k into the memory address of i
	i = k

	// the exception to this rule is when working with slices
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	// note that after we modified the copy the original value also changed
	// This is because slices are pointers to underlying arrays
	// so when copying in the value we get a pointer to the underlaying array
	// both the original and copy reference the same object in memory
	// so we are updating the one object they both reference.
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	// functions and pointers <3
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The current value of thing1 is: %v\n", thing1)
	// note that since im getting a memory address i use %p instead of %v in Printf
	fmt.Printf("The memory location of the thing1 array is: %p\n", &thing1)
	// note that im passing in the address of thing1 using &
	var result [5]float64 = square(&thing1)
	fmt.Printf("the result of the square function is: %v\n", result)
	fmt.Printf("The value of thing1 is now: %v", thing1)
}

// Normally both thing1 and thing2 have different memory addresses - I can change thing2 without effecting thing1
// But in this case thing2 is a pointer which means it takes in the memory address of the thing1 array
// this allows this changes of thing2 to modify thing1
// this is useful if you want to mutate a variable within a function
// it's also more memory efficiant as it will not need to copy a new array
func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("The memory location of the thing2 array is: %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
