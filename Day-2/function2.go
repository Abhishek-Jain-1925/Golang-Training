package main

import "fmt"

func Factorial(no int) (res string, err error) {
	fact := 1
	if no == 0 {
		err = fmt.Errorf("Error !! or Cant Find Factorial")
		return
	} else {
		fact = 1
		for i := 1; i <= no; i++ {
			fact *= i
		}
		res = fmt.Sprintf("Factorial Of %d = %d", no, fact)
		return
	}
}

func main() {
	fmt.Println("**** Function with Return Some Value **** \n")

	fmt.Println("Enter Number To Find Factorial : ")
	var no int
	fmt.Scan(&no)

	fmt.Println(Factorial(no))
}
