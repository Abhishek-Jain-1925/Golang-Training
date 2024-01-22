package main

import "fmt"

func main() {
	var val int
	fmt.Println("Enter Integer Value : ")
	_, err := fmt.Scanf("%d", &val)

	if err != nil {
		fmt.Println(fmt.Errorf("Error !! Enter Valid Input !!"))
		return
	}
	fmt.Printf("You Have Entered : %d", val)
}
