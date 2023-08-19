package main

import (
	"fmt"
	"time"
)

func main() {
	n := 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("The total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("The total time with preallocation: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
