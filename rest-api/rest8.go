package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("Starting a Server..")

	http.HandleFunc("/welcome", myHandler)

	http.ListenAndServe("127.0.0.1:1925", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {

	u := user{}

	fmt.Println(u)

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		fmt.Println(fmt.Errorf("error !!"))
	}

	fmt.Println(u)
}
