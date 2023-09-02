package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	listenPort := ":8080"
	if val := os.Getenv("PORT"); val != "" {
		listenPort = ":" + val
	}
	http.HandleFunc("/", hello)
	http.HandleFunc("/get", get)
	http.ListenAndServe(listenPort, nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<html><body><button>buton</button></body><html>")
}

func get(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, 123)
}
