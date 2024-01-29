package main

import "fmt"

type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

func (a author) show() {

	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.particles)
	fmt.Println("Salary: ", a.salary)
}

func main() {

	res := author{
		name:      "Abhishek",
		branch:    "CSE",
		particles: 203,
		salary:    34000,
	}

	res.show()
}
