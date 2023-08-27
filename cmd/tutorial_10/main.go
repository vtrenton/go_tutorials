package main

import "fmt"

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
