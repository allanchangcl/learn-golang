package main

import (
	"fmt"
	"reflect"
)

type Cat struct {
}

func (c Cat) Say() string {
	return "meow"
}

type Dog struct {
}

func (d Dog) Say() string {
	return "woof"
}

type Sayer interface {
	Say() string
}

func main() {
	c := Cat{}
	fmt.Println("Cat says: ", c.Say())
	d := Dog{}
	fmt.Println("Dog says: ", d.Say())

	animals := []Sayer{c, d}

	for i := 0; i < len(animals); i++ {
		fmt.Println(reflect.TypeOf(animals[i]).Name(), "says:", animals[i].Say())
	}
}
