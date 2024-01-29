package main

import "fmt"

type myInterface interface {
	hello()
}

func hello(name string) string {
	return fmt.Sprintf("Hello %s, Welcome in Interface !!", name)
}

func main() {
	fmt.Println(hello("Abhishek"))
}
