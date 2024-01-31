package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting Server...")

	http.HandleFunc("/welcome", WelcomeInGoRestHandler)
	//http.ListenAndServe("localhost:1925", nil)
	http.ListenAndServe("127.0.0.1:1925", nil)
}

func WelcomeInGoRestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World !! Jay Jinendra !!")
}
