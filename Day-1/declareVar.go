package main

import "fmt"

func main() {
	var (
		no      int     = 10
		name    string  = "Abhi"
		marks   float64 = 92.77
		married bool    = false
		hobbies         = []int{10, 20, 30}
		phone           = [2]int{9595, 8888}
		myAddr          = map[string]string{
			"Lane":     "New Peth",
			"Village":  "Balamtakali",
			"Taluka":   "Shevgaon",
			"District": "Ahemadnagar",
			"state":    "MH",
		}
	)
	fmt.Println(no, name, marks, married, hobbies, phone, myAddr)
}
