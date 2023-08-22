package main

import "fmt"

// structs

type gasEngine struct {
	mpg     uint8
	gallons uint8
	// we can pass in another parent struct
	//ownerInfo owner
	// I actually dont even need to pass the name just the object
	owner
	// I can do this with all types for example an int
	int
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

// methods
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}
func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

// an interface - used to share generic functions across structs
type engine interface {
	// common shared method between all structs
	// method name and return type is called the 'method signature'
	milesLeft() uint8
}

// right now this function is only availible to gasEngine
//in order to make it more general I can use an Interface (above)
//func canMakeIt(e gasEngine, miles uint8) {
func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("You need to fuel up first!")
	}
}

// functions
func main() {
	// we can set the values in the struct using struct litteral syntax
	//var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15}
	// we can ommit the field names and this will add them in order.
	var myEngine gasEngine = gasEngine{25, 15, owner{"Trent"}, 15}
	var myElectricEng electricEngine = electricEngine{69, 100}
	// values can be set manually as well
	myEngine.mpg = 20
	// anonymous structs
	// need to be defined and initialized in the same location
	// the big problem here is this is not reusable
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}
	// by directly calling 'owner' in the struct - i dont have to use owner.name to access the name field
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name, myEngine2.mpg, myEngine2.gallons)
	// since the milesleft() function takes in gasEngine when myEngine is instantiated it will have access to the milesLeft() method
	// in OOP languages - this is like calling the method of a class.
	fmt.Printf("Total miles left in the tank: %v\n", myEngine.milesLeft())
	// look ma I can pass in many engine types
	canMakeIt(myEngine, 50)
	canMakeIt(myElectricEng, 50)
}
