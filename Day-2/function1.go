package main

import "fmt"

func sayNamaste() {
	fmt.Println("\n Namaste !! Welcome in GoLang World !! - From MyFunc")
}

func main() {
	fmt.Print("Hello There !! - From main")

	sayNamaste()

	fmt.Println("Bye, Bye !! Have A Good Day !! - From Main")
}
