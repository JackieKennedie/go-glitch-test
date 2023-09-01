package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<html><body><button>buton</button></body><html>")
}

func main() {
	listenPort := ":8080"
	if val := os.Getenv("PORT"); val != "" {
		listenPort = ":" + val
	}
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(listenPort, nil)
}
