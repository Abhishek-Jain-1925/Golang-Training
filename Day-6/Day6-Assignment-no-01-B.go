package main

import "fmt"

func accessSlice(slice []int, index int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("internal error:", r)
		}
	}()

	value := slice[index]
	fmt.Printf("item %d, value %d\n", index, value)
}

func main() {
	fmt.Println("*** Hello World !! Welcome in Assignment No. 01 ***")
	defer fmt.Println("\n\n *** Thank You !! Have a Nice Day !! ***")

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

	accessSlice(slice1, (index - 1))

}
