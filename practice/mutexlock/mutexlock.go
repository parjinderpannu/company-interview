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

	fmt.Printf("Worker %d is trying to increment counter.....\n", id)

	mutex.Lock()
	counter++
	fmt.Printf("Worker %d incremented counter to %d\n", id, counter)
	mutex.Unlock()

	time.Sleep(time.Millisecond * 500)

}

func main() {
	var wg sync.WaitGroup
	numberWorkers := 5
	wg.Add(numberWorkers)
	for i := 0; i < numberWorkers; i++ {
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}
