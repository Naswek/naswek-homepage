package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./static")))

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "homepage работает")
	})

	log.Println("http://naswek.test:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
}