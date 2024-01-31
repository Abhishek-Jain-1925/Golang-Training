package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("\n Starting Server...")

	http.ListenAndServe("127.0.0.1:1925", nil)
}
