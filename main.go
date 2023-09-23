package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Hello struct {
	Title string
	Desc  string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Hello{Title: "Hello", Desc: "World"})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
