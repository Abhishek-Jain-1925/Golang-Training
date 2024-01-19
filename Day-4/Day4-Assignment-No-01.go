package main

import "fmt"

type AnyType interface{}

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
		output = fmt.Sprintf("This is a Value of type Hello, '%s'", inp)
	default:
		output = "*** Plz, Enter Valid Choice !!"
	}
	return output
}

func main() {
	fmt.Println("**** Hello World !! I'm in Assignment NO. 01 ****")

	var no int
	fmt.Println("Enter Your Input Between 1 to 4 : ")

	_, err := fmt.Scan(&no)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if no < 1 || no > 4 {
		fmt.Println("*** Plz, Enter Valid Choice !!")
		return
	}

	var val AnyType
	switch no {
	case 1:
		val = 10
		fmt.Println(AcceptAnything(val))
	case 2:
		val = "Namaste"
		fmt.Println(AcceptAnything(val))
	case 3:
		val = true
		fmt.Println(AcceptAnything(val))
	case 4:
		var myVar Hello
		myVar = "My Custom DT"
		fmt.Println(AcceptAnything(myVar))
	default:
		fmt.Println("*** Plz, Enter Valid Choice !!")
	}
}
