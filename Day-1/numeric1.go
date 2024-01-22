package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter A Number You Want : ")
	//val, err := fmt.Scanf("%d", &num)
	_, err := fmt.Scanf("%d", &num)

	if err != nil {
		fmt.Println(fmt.Errorf("Error !! Plz, Enter Valid Input And Try Again !!"))
		return
	}

	//fmt.Println("You Have Entered : ", val) Here In This case,scanf will return either 0 or 1 in val i.e. for success val=1 and val=0 for failure

	fmt.Println("You Have Entered : ", num)
}
