package main

import (
	"fmt"
	"strings"
)

func romanEquivalent(inp string) string {

	var curVal int
	res := 0
	prevVal := -1
	flag := false

	for i := len(inp) - 1; i >= 0; i-- {
		switch inp[i] {
		case 'I':
			curVal = 1
		case 'V':
			curVal = 5
		case 'X':
			curVal = 10
		case 'L':
			curVal = 50
		case 'C':
			curVal = 100
		case 'D':
			curVal = 500
		case 'M':
			curVal = 1000
		default:
			flag = true
		}

		if prevVal > curVal {
			res -= curVal
		} else {
			res += curVal
		}

		prevVal = curVal
	}

	if flag {
		return "Invalid Roman Number !!"
	} else {
		output := fmt.Sprintf("Decimal Equivalent of Given Roman Number is : %d", res)
		return output
	}
}

func main() {

	var inp string
	fmt.Println("Enter Roman Value to Find it's Decimal Value (IN ROMAN,ALL LETTERS ARE IN CAPITAL): ")
	fmt.Scan(&inp)

	inp = strings.ToUpper(inp) //Making it Capital If in case User Enters Inp in Lowercase

	result := romanEquivalent(inp)
	fmt.Println(result)
}
