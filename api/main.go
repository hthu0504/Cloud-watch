package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Cloud-watch API is running")
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/events", telemetryHandler)

	fmt.Println("API running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}