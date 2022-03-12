package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! HandlerFunc~")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World! HandlerFunc~")
}

func main () {
	fmt.Println("HTTP srerver is running")

	server := http.Server {
		Addr: "127.0.0.1:8080",

	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}