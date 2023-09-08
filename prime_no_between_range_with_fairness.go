package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var concurrency int = 10
var currentNum int32 = 2
var limit int = 100000000
var totalPrimeNumbers int32 = 2

func isPrime(num int) {
	if num&1 == 0 {
		return
	}
	for i := 2; i < int(math.Ceil(math.Sqrt(float64(num)))); i++ {
		if num%i == 0 {
			return
		}
	}
	atomic.AddInt32(&totalPrimeNumbers, 1)
}
func calculate(i int, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	for {
		x := atomic.AddInt32(&currentNum, 1)
		if x > int32(limit) {
			break
		}
		isPrime(int(x))
	}
	fmt.Printf("Threads %d completed in %s\n", i, time.Since(start))
}
func main() {

	start := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go calculate(i, &wg)
	}
	wg.Wait()

	fmt.Printf("The program took %s time to complete\n", time.Since(start))

}
