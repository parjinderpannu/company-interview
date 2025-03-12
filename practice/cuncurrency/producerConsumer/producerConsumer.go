package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numProducers = 2
	numConsumers = 3
	bufferSize   = 5
)

// Producer function
func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		item := rand.Intn(100)
		fmt.Printf("[%d] Producer : Produced %d\n", id, item)
		ch <- item // send to channel
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

func consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range ch {
		fmt.Printf("[%d] Consumer : Consumed %d\n", id, item)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(800)))
	}
}

func main() {
	ch := make(chan int, bufferSize)
	var wg sync.WaitGroup

	for i := 0; i < numProducers; i++ {
		wg.Add(1)
		go producer(i, ch, &wg)
	}

	// Wait for producers to finish, then close channel
	go func() {
		wg.Wait() // Wait for all producers to finish
		close(ch) // Close the channel so consumers stop receiving
	}()

	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go consumer(i, ch, &wg)
	}
	wg.Wait()
}
