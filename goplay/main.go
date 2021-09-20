package main

import (
	"encoding/json"
	"fmt"
)

type thing struct {
	ID     int
	phrase string
}

func setID(newID int) (t thing) {
	t.ID = newID
	return t
}

func main() {
	thing := setID(4)
	s, _ := json.MarshalIndent(thing, "", "\t")
	// fmt.Println(thing)
	fmt.Println(string(s))
	fmt.Println("vim-go")
}
