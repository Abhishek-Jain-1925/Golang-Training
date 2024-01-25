package main

import (
	"fmt"
)

func printMessage(a chan string, b chan string, flag chan bool) {
	// defer func() {
	//  close(a)
	//  close(b)
	// }()

	for {
		select {
		case val := <-a:
			fmt.Printf("alice: %s,", val)

		case val := <-b:
			fmt.Printf("bob: %s,", val)

		case <-flag:
			return
		}
	}
}

func main() {
	s := "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"

	//fmt.Println(len(s))

	alice := make(chan string)
	bob := make(chan string)
	flag := make(chan bool)

	go printMessage(alice, bob, flag)

	var temp string
	for i := 0; i < len(s); i++ {
		ch := fmt.Sprintf("%c", s[i])

		if ch == "$" {
			alice <- temp
			temp = ""
		} else if ch == "#" {
			bob <- temp
			temp = ""
		} else if ch == "^" {
			flag <- true
			break
		} else {
			temp = temp + ch
		}
	}
}
