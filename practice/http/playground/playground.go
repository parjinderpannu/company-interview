package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from helloHandler")
}

func main() {
	// routing
	http.HandleFunc("/hello", helloHandler)

	// start the server
	go func() {
		log.Println("Starting a Server...")
		if err := http.ListenAndServe("localhost:9090", nil); err != nil {
			log.Println("Stopping a server... and error: ", err)
		}
	}()

	time.Sleep(1 * time.Second)

	res, err := http.Get("http://localhost:9090/hello")
	if err != nil {
		log.Fatalf("http.Get error: %v", err)
	}
	defer res.Body.Close()

	log.Println("Reading Response...")
	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		log.Fatalf("io.Copy to stdout error: %v", err)
	}

}
