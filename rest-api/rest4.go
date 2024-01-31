package main

import (
	"fmt"
	"net/http"
)

// Printing type of req.
func main() {
	fmt.Println("Starting Server...")

	http.HandleFunc("/welcome", WelcomeInGoRestHandler)
	//http.ListenAndServe("localhost:1925", nil)
	http.ListenAndServe("127.0.0.1:1925", nil)
}

func WelcomeInGoRestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest) //200
	fmt.Println("The req is of type : ", r.Method)
	fmt.Fprint(w, "Hello World !! Jay Jinendra !!")
}
