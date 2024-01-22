package main

import (
	"errors"
	"fmt"
)

func function1() error {
	return errors.New("Something Went Wrong !!")
}

func main() {
	err := function1()
	if err != nil {
		fmt.Println(err.Error())
	}
}
