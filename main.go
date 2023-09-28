package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Hello struct {
	Title string
	Desc  string
	Ver   string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Hello{Title: "Hello", Desc: "World", Ver: "0.0.1"})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
