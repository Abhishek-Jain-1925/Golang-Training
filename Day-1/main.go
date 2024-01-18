package main

import "fmt"

var enabled bool = true

func main() {

	var enabled = false

	if enabled {

		fmt.Println("hey there... !")

		return

	}

	fmt.Println("make enabled true... !")

}
