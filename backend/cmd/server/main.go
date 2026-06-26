package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"status":"ok"}`)
}

func main() {
	http.HandleFunc("/health", healthHandler)

	fmt.Println("SnowLeo backend running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
