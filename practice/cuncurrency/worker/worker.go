package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Number of workers and jobs
const (
	numWorkers = 3
	numJobs    = 10
)

// Job struct (can hold any work details)
type Job struct {
	ID int
}

// Worker function (processes jobs)
func worker(id int, jobs <-chan Job, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500))) // Simulate work
		results <- fmt.Sprintf("Worker %d processed Job %d", id, job.ID)
	}
}

func main() {
	// rand.Seed(time.Now().UnixNano()) // Seed random generator

	jobs := make(chan Job, numJobs)       // Job queue
	results := make(chan string, numJobs) // Results channel
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs to the channel
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j}
	}
	close(jobs) // Close job channel (signals no more jobs)

	// Wait for workers to finish
	wg.Wait()
	close(results) // Close results channel

	// Collect and print results
	for result := range results {
		fmt.Println(result)
	}
}
