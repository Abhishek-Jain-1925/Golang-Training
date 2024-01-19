package main

import (
	"fmt"
)

type AnyType interface {
}

type Hello string

func AcceptAnything(inp AnyType) string {
	var output string
	switch inp.(type) {
	case int:
		output = fmt.Sprintf("This is a Value of type Integer, '%d'", inp)
	case string:
		output = fmt.Sprintf("This is a Value of type String, '%s'", inp)
	case bool:
		output = fmt.Sprintf("This is a Value of type Boolean, '%t'", inp)
	case Hello:
		output = fmt.Sprintf("This is a Value of type Hello, ", inp)
	default:
		output = "*** Plz, Enter Valid Choice !!"
	}
	return output
}

func main() {
	fmt.Println("**** Hello World !! I'm in Assignment NO. 01 ****")

	var no int
	fmt.Println("Enter Your Input Between 1 to 4 : ")
	fmt.Scan(&no)

	var val AnyType
	switch no {
	case 1:
		fmt.Println("Enter Any Inter Value You want : ")
		fmt.Scanf("%d \n", &val)
		// val = 10
		fmt.Printf("%v....%T", val, val)
		fmt.Println(AcceptAnything(val))
	case 2:
		fmt.Println("Enter Any String Value You Want : ")
		fmt.Scan(&val)
		fmt.Println(AcceptAnything(val))
	case 3:
		fmt.Println("Enter Any Boolean Value You Want(true or false) : ")
		fmt.Scan(&val)
		fmt.Println(AcceptAnything(val))
	case 4:
		var myVar Hello
		fmt.Println("Enter Anything You Want : ")
		fmt.Scan(&myVar)
		fmt.Println(AcceptAnything(myVar))
	}
}
