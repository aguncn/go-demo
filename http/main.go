package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", myHandler)

	log.Println("Starting httpserver")
	log.Fatal(http.ListenAndServe(":8088", mux))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, go net/http."))
}
