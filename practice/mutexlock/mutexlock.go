package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	mutex   sync.Mutex
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// fmt.Printf("%d Go routine started and counter = %d\n", id, counter)

	mutex.Lock()
	counter++
	// fmt.Printf("%d Go routine ended and counter = %d\n", id, counter)
	mutex.Unlock()

	time.Sleep(time.Microsecond * 500) // doing some work
}

func main() {
	var wg sync.WaitGroup

	totalGogoutines := 10000
	wg.Add(totalGogoutines)

	for i := 0; i < totalGogoutines; i++ {
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Printf("\nAll %d Go Routines finished \nCounter : %d", totalGogoutines, counter)

}
