package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("I'm in Main")
	defer fmt.Println("End of main")

	//var wg sync.WaitGroup
	ch1 := make(chan string)
	ch2 := make(chan string)

	var inpStr string
	fmt.Print("Enter Your String: ")
	_, err := fmt.Scanf("%s", &inpStr)
	if err != nil {
		fmt.Println(fmt.Errorf("error! Please try again with valid input"))
		return
	}

	//wg.Add(1)
	go func() {
		var msg string

		for _, char := range inpStr {
			switch char {
			case '$':
				ch1 <- msg
				go read(ch1)
				msg = ""
			case '#':
				ch2 <- msg
				msg = ""
			case '^':
				//wg.Done()
				return
			default:
				msg += string(char)
			}
		}
		close(ch1)
		close(ch2)
	}()
	time.Sleep(time.Second * 2)

	//wg.Wait()

	// fmt.Println("Messages from Alice:")
	// for msg := range ch1 {
	//fmt.Println("alice:", <-ch1)
	// }

	// fmt.Println("Messages from Alice:")
	// fmt.Println(len(ch1))
	// for i := 0; i < len(ch1); i++ {
	// 	fmt.Println("alice:", <-ch1)
	// }

	// fmt.Println("Messages from Bob:")
	// for msg := range ch2 {
	//fmt.Println("bob:", <-ch2)
	// }

	// fmt.Println("Messages from Bob:")
	// for {
	// 	msg, ok := <-ch2
	// 	if !ok {
	// 		break // Channel closed
	// 	}
	// 	fmt.Println("bob:", msg)
	// }

}

func read(ch1 <-chan string) {
	fmt.Println("alice,", <-ch1)
}
