package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
    buffer_size = 10
    producers = 2
    consumers = 3
)

var buffer []int
var mutex = sync.Mutex{}
var empty = sync.NewCond(&mutex)
var full = sync.NewCond(&mutex)

func producer(id int) {
    for {
        item := rand.Intn(100)
        fmt.Printf("producer %d produced item %d\n", id, item)
        mutex.Lock()
        if len(buffer) == buffer_size {
            full.Wait()
        }
        buffer = append(buffer, item)
        empty.Signal()
        mutex.Unlock()
        time.Sleep(1*time.Second)
    }

}
func consumer(id int){
    for {
        mutex.Lock()
        if len(buffer) == 0 {
            empty.Wait()
        }
        item := buffer[0]
        buffer = buffer[1:]
        full.Signal()
        mutex.Unlock()
        fmt.Printf("consumer %d consumed item %d\n", id, item)
        time.Sleep(1*time.Second)
    }
}
func main(){

rand.Seed(time.Now().UnixNano())

	for i := 0; i < producers; i++ {
		go producer(i)
	}

	for i := 0; i < consumers; i++ {
		go consumer(i)
	}

	// Run indefinitely to keep the program running
	select {}
}
