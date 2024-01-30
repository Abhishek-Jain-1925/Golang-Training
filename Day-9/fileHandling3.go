package main

import (
	"fmt"
	"os"
)

func main() {
	var file string
	fmt.Print("Enter File name to get all Info : ")
	fmt.Scan(&file)
	file = "./" + file

	fi, err := os.Stat(file)
	if err != nil {
		fmt.Println(fmt.Errorf("error !! Something went Wrong !!"))
		return
	}

	fmt.Println("\n *** Details are As Follows : ")
	fmt.Println("Name:", fi.Name())
	fmt.Println("Size:", fi.Size())
	fmt.Println("Mode:", fi.Mode())
	fmt.Println("ModTime:", fi.ModTime())
}
