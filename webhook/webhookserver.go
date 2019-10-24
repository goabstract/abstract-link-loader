package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	verbose := os.Getenv("VERBOSE") == "true"
	if verbose {
		fmt.Printf("\nRequested %s\n", time.Now())
		for key, value := range r.Header {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
	var payload bytes.Buffer
	payload.ReadFrom(r.Body)
	fmt.Printf("%s\n", payload.String())
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
