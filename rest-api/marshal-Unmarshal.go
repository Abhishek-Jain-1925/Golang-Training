package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{"John", 30}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	var p2 Person
	err = json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p2)
}
