package main

import "fmt"

type A struct {
	year int
}

func (a A) Greet() {
	fmt.Println("Hello GolangUK", a.year)
}

type B struct {
	A
}

func (b B) Greet() {
	fmt.Println("Welcome to GolangUK", b.year)
}

type Cat struct {
	Name string
}

func (c Cat) Legs() int {
	return 4
}

func (c Cat) PrintLegs() {
	fmt.Printf("I have %d legs\n", c.Legs())
}

type OctoCat struct {
	Cat
}

func (o OctoCat) Legs() int {
	return 5
}

func main() {
	var a A
	a.year = 2016
	var b B
	b.year = 2016
	a.Greet()
	b.Greet()

	var octo OctoCat
	fmt.Println(octo.Legs())
	octo.PrintLegs()
	fmt.Println("vim-go")
}
