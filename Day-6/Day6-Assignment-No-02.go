package main

import (
	"errors"
	"fmt"
)

func accessSlice(slice1 []int, index int) (result int, err error) {

	if index > (len(slice1) - 1) {
		return 0, errors.New("\n Error !! length of the slice should be more than index !! Plz,Try With Valid Index.")
	}
	if index < 0 {
		return 0, errors.New("\n Error !! length of the slice should be less than index !! Plz,Try With Valid Index.")
	}

	return slice1[index], nil
}

func main() {

	var no int
	fmt.Print("\n How Many Elements You want in Slice : ")
	_, err := fmt.Scan(&no)
	if err != nil {
		fmt.Println("Error !! Plz,Enter Valid Input !!")
		return
	}

	var slice1 = make([]int, no)

	fmt.Println("\n*** Enter ", no, " Elements One by One ***")
	for i := 0; i < no; i++ {
		fmt.Print("\nElement ", (i + 1), " : ")
		_, err1 := fmt.Scan(&slice1[i])
		if err1 != nil {
			fmt.Println(fmt.Errorf("Plz, Enter Valid Inputs and Try Again !!"))
			return
		}
	}

	fmt.Println("\n Your Given Slice Will look like - ", slice1)

	var index int
	fmt.Print("\n\n Enter Index, To fetch that Value from slice : ")
	_, err2 := fmt.Scanf("%d", &index)
	if err2 != nil {
		fmt.Println(fmt.Errorf("Error Occurred !! Due To %v", err2))
	}

	result, err3 := accessSlice(slice1, (index - 1))
	if err3 != nil {
		fmt.Println(err3.Error())
		return
	}

	fmt.Printf("\n *** Item : %d, Value : %d ***", index, result)

}
