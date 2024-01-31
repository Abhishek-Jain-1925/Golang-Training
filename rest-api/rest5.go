package main

import (
	"fmt"
	"net/http"
)

// Authorization using API key
func main() {
	fmt.Println("Starting Server...")

	http.HandleFunc("/welcome", WelcomeInGoRestHandler)
	//http.ListenAndServe("localhost:1925", nil)
	http.ListenAndServe("127.0.0.1:1925", nil)
}

func WelcomeInGoRestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest) //200
	apikey := r.Header.Get("Authorization")
	fmt.Println(apikey)

	if apikey == "123" {
		fmt.Fprint(w, "Hello World !! Jay Jinendra !!")
		return
	}
	fmt.Fprint(w, "Invalid Token !!")
}
