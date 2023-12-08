// Be warned this will peg all cores on CPU
// may cause application crashes

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go count()
	}
	wg.Wait()
	fmt.Printf("The total execution time is: %v", time.Since(t0))
}

func count() {
	var res int
	for i := 0; i < 100000000; i++ {
		res += 1
	}
	wg.Done()
}
