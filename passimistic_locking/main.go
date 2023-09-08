package main

import (
	"fmt"
	"sync"
)

var count int32 = 0
var mu sync.Mutex
var wg sync.WaitGroup

func incrCount(wg *sync.WaitGroup) {
	mu.Lock()
	defer mu.Unlock()
	count++
	wg.Done()
}
func main() {
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go incrCount(&wg)
	}
	wg.Wait()

	fmt.Printf("count: %d\n", count)
}
