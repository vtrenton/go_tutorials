package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// The issue with concurrent threads writing data to something concurrently means there will be some overlap
// This overlap can introduce corruption and dataloss
// The safe method is to use a Mutex
// Mutex is short for "mutual exclusion"
//var m = sync.Mutex{}

// using an RWMutex allows us to use RLock and RUnlocks
// as well as the traditonal mutex loocks
var m = sync.RWMutex{}

// The issue with just calling a go on this function is it runs the entire program before getting a return value
// This means that the main program code never runs
// Fix this we can use a WaitGroup
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		// a waitgroup is a counter
		// before the ascending into this function we increment the counter
		wg.Add(1)
		// The go keyword tells the program to continue and not wait for the return of the function
		go dbCall(i)
	}
	// this will assure the counter is down to zero before continuing
	wg.Wait()
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
	fmt.Printf("The results are: %v\n", results)
}

func dbCall(i int) {
	// similate db call delay
	var delay float32 = rand.Float32() * 2000
	//var delay float32 = 2000
	// it's very important where you put the mutex lock and unlock statements
	// if we put the lock statement here for example this would completly ruin any paralellization benifits
	// This is because the compute (data procession portion) would be locked by one thread.
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	save(dbData[i])
	log()
	// this will decrement our waitgroup counter to zero
	wg.Done()
}

func save(result string) {
	// this protects the writes to the Slice by putting a lock on it
	// when a goroutine reaches this point it will lock the call.
	// all other goroutines will need to wait for this the current thread to finish before accessing the function
	m.Lock()
	results = append(results, result)
	// of course, once the thread is done with this task the next step should be to unlock.
	m.Unlock()
}

func log() {
	// this will block only when there is an active full lock (m.lock()) that is active on another thread
	// this will assure that data is not being printed out while it is being changed by another goroutine.
	m.RLock()
	fmt.Printf("The current results are %v\n", results)
	m.RUnlock()
}
