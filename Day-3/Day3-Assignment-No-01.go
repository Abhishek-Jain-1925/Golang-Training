package main

import (
	"fmt"
)

func findADay(no int) (result string, err error) {
	day := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	result, ok := day[no]
	// fmt.Println(result)
	// fmt.Println(ok)
	if !ok {
		err = fmt.Errorf("Not a day!")
	}

	return result, err
}

func main() {
	var index int
	fmt.Println("Enter the number for which you want to Find an Equivalent Day:")
	fmt.Scan(&index)

	day, err := findADay(index)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Equivalent day:", day)
	}
}

// func findADay(no int) (result string, err error) {
// 	day := map[int]string{
// 		1: "Monday",
// 		2: "Tuesday",
// 		3: "Wednesday",
// 		4: "Thursday",
// 		5: "Friday",
// 		6: "Saturday",
// 		7: "Sunday",
// 	}

// 	for key, value := range day {
// 		if key == no {
// 			result = value
// 			break
// 		}
// 	}
// 	if err != nil {
// 		err = fmt.Errorf("Not a Day !!")
// 		return
// 	} else {
// 		return
// 	}
// }

// func main() {
// 	var index int
// 	fmt.Println("Enter The Number For Which You Want to Find A Equivalent Day : ")
// 	fmt.Scan(&index)

// 	fmt.Println(findADay(index))
// }
