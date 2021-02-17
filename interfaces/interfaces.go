package main

import "fmt"

type voicer interface {
	sound()
}

type bird struct{}

func (b *bird) sound() {
	fmt.Println("Chirp, Chirp")
}

// Car Item
type car struct{}

func (c *car) sound() {
	fmt.Println("Honk, Honk")
}

// human Item
type human struct{}

func (h *human) sound() {
	fmt.Println("Hello, Hello")
}

func main() {
	fmt.Println("vim-go")
	var v voicer

	// bird
	v = &bird{}
	v.sound()

	// car
	v = &car{}
	v.sound()

}
