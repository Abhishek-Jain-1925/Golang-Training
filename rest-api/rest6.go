package main

import (
	"fmt"
	"net/http"
)

// Form Data Access using GET -parseForm
func main() {
	fmt.Println("Starting Server...")

	http.HandleFunc("/welcome", MyHandler)

	http.ListenAndServe("127.0.0.1:1925", nil)
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) //200

	fmt.Println("Data Before parseForm : ", r.Form)

	r.ParseForm()

	fmt.Println("Data After parseForm : ", r.Form)

	fmt.Fprint(w, "Done !!")
}
