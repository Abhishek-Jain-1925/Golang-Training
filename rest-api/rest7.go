package main

import (
	"fmt"
	"net/http"
)

// Form Data Access using POST -parseForm
func main() {
	fmt.Println("Starting Server...")

	http.HandleFunc("/welcome", MyHandler)

	http.ListenAndServe("127.0.0.1:1925", nil)
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) //200

	fmt.Println("Data Before parseForm : ", r.PostForm)

	r.ParseForm()

	fmt.Println("Data After parseForm : ", r.PostForm)

	fmt.Fprint(w, "Done !!")
}
