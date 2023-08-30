package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("PORT")
	fmt.Fprintf(w, "Hello from server on port: %s!\n", port)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handler)
	log.Printf("Starting server on :%s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
