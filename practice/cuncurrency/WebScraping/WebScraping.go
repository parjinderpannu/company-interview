package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	urls := []string{"https://golang.org", "https://example.com", "https://go.dev"}
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}

	wg.Wait()
}

func fetch(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching : ", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("\n[%d] bytes Fetched from [%s]\n", len(body), url)
}
