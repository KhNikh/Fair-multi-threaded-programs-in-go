package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)
var cuncurrency int = 10
var limit int = 100000000
var batch_size int = 10000000

func is_prime(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 || n == 3 {
		return true
	}

	for i := 2; i <= int(math.Ceil(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func calculate(start int, end int, thread_no int) {
	defer wg.Done()
	thread_start_time := time.Now()
	primes := 0
	for i := start; i < end; i++ {
		if x := is_prime(i); x == true {
			primes++
		}
	}
	fmt.Printf("there are %d primes in range [%d, %d), thread %d it took %f secs\n", primes, start, end, thread_no, time.Since(thread_start_time).Seconds())
}

func main() {
	start_time := time.Now()
	thread_no := 1
	for i := 0; i <= limit && i+batch_size <= limit; i += batch_size {
		wg.Add(1)
		go calculate(i, i+batch_size, thread_no)
		thread_no++
	}
	wg.Wait()
	fmt.Printf("The full program took %f secs\n", time.Since(start_time).Seconds())
}
