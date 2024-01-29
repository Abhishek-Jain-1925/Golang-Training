package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello World..!! Welcome in GO File Handling Era !!\n")

	content := "This is content to add in File"

	file, err := os.Create("./demo.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(fmt.Errorf("error"))
	}

	length, err := io.WriteString(file, content)
	fmt.Println("Length : ", length)

	ReadFile("./demo.txt")

}

func ReadFile(filename string) {
	databyte, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(fmt.Errorf("error in reading file !!"))
		return
	}

	fmt.Println("Your File Data : ", string(databyte))
}
