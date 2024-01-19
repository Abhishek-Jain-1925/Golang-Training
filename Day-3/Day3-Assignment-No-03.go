package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter string: ")

	scanner.Scan()
	var s string = scanner.Text()

	temp := ""

	freq := make(map[string]int, 0)

	var orderedWords []string

	values := strings.Split(s, " ")

	fmt.Println(values)

	fmt.Println("You Have Entered : ", s)

	for i := 0; i < len(); i++ {
		ch := fmt.Sprintf("%c", s[i])

		if ch == " " && temp != "" {
			freq[temp]++
			orderedWords = append(orderedWords, temp)
			temp = ""
		} else {
			temp = temp + ch
		}
	}

	if temp != " " {
		freq[temp]++
		orderedWords = append(orderedWords, temp)
	}

	var maxFreq int = 0
	for _, cnt := range freq {
		if cnt > maxFreq {
			maxFreq = cnt
		}
	}

	var maxFreqWords []string

	for i := 0; i < len(orderedWords); i++ {
		if freq[orderedWords[i]] == maxFreq {
			maxFreqWords = append(maxFreqWords, orderedWords[i])
			freq[orderedWords[i]] = -1
		}
	}

	fmt.Println(maxFreqWords)

}
