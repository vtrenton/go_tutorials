package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// int
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice(intSlice))
	// float32
	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice(float32Slice))
	// float64
	var float64Slice = []float64{1, 2, 3}
	fmt.Println(sumSlice(float64Slice))

	// check empty
	var empIntSlice = []int{}
	//fmt.Println(isEmpty[int](empIntSlice))
	// I can ommit the type in the square brackets as go will infer it.
	fmt.Println(isEmpty(empIntSlice))

	var empFloat32Slice = []float32{1, 2, 3, 4}
	//fmt.Println(isEmpty[float32](empFloat32Slice))
	fmt.Println(isEmpty(empFloat32Slice))

	// Contact program
	// in this case we need to declare the struct type else the unmarshalling function wont know what struct to use
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("%+v\n", contacts)

	var purchase []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	fmt.Printf("%+v\n", purchase)

	// types of cars
	// note that engine takes both gas engine and electric engine as a generic
	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			mpg:     12.4,
			gallons: 40,
		},
	}

	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwh:   57.5,
			mpkwh: 4.17,
		},
	}

	fmt.Println(gasCar)
	fmt.Println(electricCar)
}

// T is the generic type for int, float32 and float64
// meaning I can pass any of those types into this function
func sumSlice[T int | float32 | float64](slice []T) T {
	// note that sum is of type T
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

// go cant always infer types of passed parameters for example

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name    string
	Price   float32
	Ammount int
}

// create a Generic of "contactInfo" and "purchaseInfo"
// return the Generic type T
func loadJSON[T contactInfo | purchaseInfo](filepath string) []T {
	// read the file into the data variable
	var data, _ = os.ReadFile(filepath)
	// intialize an empty struct object
	var loaded = []T{}
	//unmarshal the json string into the address of the "loaded" slice
	json.Unmarshal(data, &loaded)

	return loaded
}

// generics can be used with struct types
// for example

type gasEngine struct {
	mpg     float32
	gallons float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

// using a generic in a struct
type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}
