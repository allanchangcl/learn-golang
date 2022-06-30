package main

import "fmt"

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
)

func main() {
	fmt.Println(Paracetamol)

}
