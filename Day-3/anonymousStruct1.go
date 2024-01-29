package main

import (
	"fmt"
)

func main() {

	student := struct {
		rollno int64
		name   string
		marks  float64
	}{
		rollno: 1925,
		name:   "Mr.Abhi",
		marks:  86.50,
	}

	fmt.Print(student)

	nexon := struct {
		color    string
		engine   string
		price    float64
		features struct {
			powerWindow bool
			powerStr    bool
			sunroof     bool
		}
	}{
		color:  "red",
		engine: "1600CC",
		price:  180000000.00,
		features: struct {
			powerWindow bool
			powerStr    bool
			sunroof     bool
		}{true, true, true},
	}
	fmt.Println(nexon)
}

// package main

// import (
// 	"fmt"
// )

// // type myVar struct{

// // }

// func main() {

// 	student := struct {
// 		rollno int64
// 		name   string
// 		marks  float64
// 	}{
// 		rollno: 1925,
// 		name:   "Mr.Abhi",
// 		marks:  86.50,
// 	}

// 	fmt.Print(student)

// 	nexon := struct {
// 		color    string
// 		engine   string
// 		price    float64
// 		features struct {
// 			powerWindow bool
// 			powerStr    bool
// 			sunroof     bool
// 		}
// 	}{
// 		color:  "red",
// 		engine: "1600CC",
// 		price:  180000000.00,
// 		features: struct {
// 			powerWindow bool
// 			powerStr    bool
// 			sunroof     bool
// 		}{
// 			powerWindow: true, powerStr: true, sunroof: false,
// 		},
// 	}
// 	fmt.Println(nexon)
// }
