package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "OK")

}

func main() {
	// routing
	http.HandleFunc("/health", healthHandler)

	// start server
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalf("error : %s", err)
	}
}
