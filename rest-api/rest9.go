package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	User     string
	Password string
}

func main() {
	fmt.Println("Test Client")

	buf := new(bytes.Buffer)

	u := User{
		User:     "test",
		Password: "test@gmail",
	}

	err := json.NewEncoder(buf).Encode(u)
	if err != nil {
		fmt.Println("Error !!")
	}

	client := http.Client{}

	r, err := http.NewRequest("POST", "127.0.0.1:1925/user", buf)

	resp, _ := client.Do(r)

	defer resp.Body.Close()

	response, _ := io.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode, string(response))
}
