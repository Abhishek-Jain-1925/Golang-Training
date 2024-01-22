package main

import "fmt"

type ErrorMessage struct {
	Message    string
	statusCode int
}

func (e *ErrorMessage) Error() string {
	return fmt.Sprintf("Message : %s, Status Code : %d", e.Message, e.statusCode)
}

func test() error {
	return &ErrorMessage("My User Custom Msg !!", 19)
}

func main() {
	err := test()
	if err != nil {
		fmt.Println(err.Error())
	}
}
