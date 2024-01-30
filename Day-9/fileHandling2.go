package main

import (
	"fmt"
	"os"
)

func ReadFromFile(file string) {
	_, err := os.Stat(file)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(fmt.Errorf(" *** File not Exists !! Plz, Enter valid name !! *** \n"))
			return
		}
	}
	databyte, err := os.ReadFile(file)
	notNillError(err)
	fmt.Println("\n Your Data is : ", string(databyte))
}
func WriteIntoFile() {
	var file string
	fmt.Print("Enter File name in which you want write : ")
	_, err := fmt.Scan(&file)
	notNillError(err)
	file = "./" + file

	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var content string
	fmt.Println("Enter Your Data to be Add in File : ")
	fmt.Scan(&content)

	_, err = f.WriteString(content)
	notNillError(err)

	ReadFromFile(file)
}

func main() {
	fmt.Println("\n *** Hello !! Welcome in My Menu !! ***")
	var ch int

	fmt.Print("\n 1.Read File\n 2.Write File\n 3.Exit\n Enter Your Choice : ")
	_, err := fmt.Scan(&ch)
	notNillError(err)

	switch ch {
	case 1:
		var file string
		fmt.Println("Enter File Name to Read : ")
		_, err := fmt.Scan(&file)
		notNillError(err)
		file = "./" + file

		ReadFromFile(file)
	case 2:
		WriteIntoFile()
	case 3:
		return
	default:
		fmt.Println("*** Plz, enter valid choice !!")
	}

}

func notNillError(err error) {
	if err != nil {
		fmt.Println(fmt.Errorf("error !! Due to : %s", err))
		return
	}
}
