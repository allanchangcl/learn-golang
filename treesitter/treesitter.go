package main

import "fmt"

type person struct {
	name string
	mom  *person
}

func newPerson(name string, mom *person) person {
	return person{name: name, mom: mom}
}

func (person *person) getName() string {
	return person.name
}

func (person *person) getMom() *person {
	return person.mom
}

func main() {
	fmt.Println("vim-go")
}
