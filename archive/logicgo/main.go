package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	if x, err := demo(); err != nil {
		// handle the error
		fmt.Printf("Value inside handle error %v", x)
	}
	// this doesn't work because x is scoped to the if block
	// fmt.Println(x)
}

func demo(args ...string) (int, error) {
	// return when no errors
	// return 1, nil

	// return when have errors
	return 0, errors.New("Error encountered")
}
