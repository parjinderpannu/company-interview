package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
Response Writer : Sends data back to API caller.
Request : we get from API caller
*/
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	// routing
	http.HandleFunc("/health", healthHandler)

	// run server

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}
