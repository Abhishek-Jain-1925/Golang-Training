package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var inpStr string
var m sync.Mutex

func reverseString(inpStr string) {

	defer wg.Done()
	
	var result string
	for _, v := range inpStr {
		result = string(v) + result
	}
	fmt.Println("Reversed String : ", result)
	fmt.Println("Number of Goroutines : ", runtime.NumGoroutine())
	//wg.Done()
}

func main() {

	fmt.Println("Enter String to Reverse : ")
	m.Lock() //Used mutex concept for threadSafe value assignment!!
	_, err := fmt.Scan(&inpStr)
	m.Unlock()
	if err != nil {
		fmt.Println(fmt.Errorf("error !! Plz,Enter valid input !!"))
		return
	}

	wg.Add(1)
	go reverseString(inpStr)
	wg.Wait()

}
