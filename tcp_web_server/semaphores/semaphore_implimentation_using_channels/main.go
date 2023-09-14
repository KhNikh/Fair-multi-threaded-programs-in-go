package main

import (
    "fmt"
    "sync"
    "time"
)

var sem = make(chan struct{}, 4)
var wg sync.WaitGroup

func task(num int) {
    defer wg.Done()

    sem <- struct{}{} // Acquire semaphore
    defer func() {
        <-sem // Release semaphore
    }()

    fmt.Println("Executing task", num)
    time.Sleep(time.Second * 1)
}

func main() {
    for i := 1; i <= 4; i++ {
        wg.Add(1)
        go task(i)
    }

    wg.Wait()
}
