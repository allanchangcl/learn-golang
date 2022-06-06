package main

import "fmt"

type Employee struct {
	FirstName, LastName string
}

func fullName(firstName string, lastName string) (fullName string) {
	fullName = firstName + " " + lastName
	return
}

// receiver function
func (e Employee) print() {
	fmt.Printf("%s is the firstname, %s is the lastname", e.FirstName, e.LastName)
}

func main() {
	// fmt.Println("vim-go")
	e := Employee{
		FirstName: "Allan",
		LastName:  "Chang",
	}

	fmt.Println(fullName(e.FirstName, e.LastName))

	// use receiver function
	e.print()
}
