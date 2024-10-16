package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var INT_MAX int32 = 100000000

var CONC int32 = 10

var TOTALPRIMES int32 = 0

var CURR int32 = 2

func countP(x int) {
	if x&1 == 0 {
		return
	}

	for i := 3; i < int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return
		}
	}
	atomic.AddInt32(&TOTALPRIMES, 1)
}

func doBatch(name string, wg *sync.WaitGroup) {
	start := time.Now()

	defer wg.Done()
	for {
		x := atomic.AddInt32(&CURR, 1)
		if x > INT_MAX {
			break
		}
		countP(int(x))
	}

	fmt.Printf("Time taken by %s: %v\n", name, time.Since(start))

}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < int(CONC); i++ {
		wg.Add(1)
		go doBatch(fmt.Sprintf("Worker %d", i), &wg)
	}
	wg.Wait()
	fmt.Printf("Total primes: %d\n", TOTALPRIMES)
	fmt.Printf("Time taken by main: %v\n", time.Since(start))
}
