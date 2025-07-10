package main

import (
	"time"
	"fmt"
	"sync"
)

// go => GoRutine
// GoRutine -> execution thread light and virtual
// Channel -> comunicate things between GoRutine

func sayHello(channel chan <- string) {
	time.Sleep(1 * time.Second)
	channel <- "Hello from GoRoutine"
}

func printMessage(channel <-chan string) {
	fmt.Println("Waiting message...")
	msg := <-channel
	fmt.Println(msg)
}

func main() {
	channel := make(chan string)
	go sayHello(channel)
	printMessage(channel)

	channel2 := make(chan int)
	go func() {
		for i :=0; i < 5; i++ {
			channel2 <- i
		}
		close(channel2)
	}()

	// range used here like a reader
	// channel used here like a reader 
	for num := range channel2 {
		fmt.Println("num received: ", num)
	}

	// Example for Mutex
	// used for reserve resources
	var counter int
	var mu sync.RWMutex

	// Writer
	go func() {
		for i :=0; i < 5; i++ {
			mu.Lock()
			counter++
			mu.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Reader
	for i := 0 ; i < 3; i++ {
		go func(id int) {
			for j :=0; j < 5; j++ {
				mu.RLock()
				fmt.Printf("Reading from the GoRoutine %d: %d\n", id, counter)
				mu.RUnlock()
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep( 2 * time.Second )
	
}