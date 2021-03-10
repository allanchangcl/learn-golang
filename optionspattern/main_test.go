package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	got := 1
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

func TestThreeFloorHouse(t *testing.T) {
	r := &house{
		material:     "concrete",
		hasFireplace: false,
		floors:       3,
	}

	h := newHouse(
		withConcrete(),
		withFloors(3),
		withoutFireplace(),
	)

	fmt.Printf("%+v\n", *h)
	fmt.Printf("%+v\n", *r)

	if *h != *r {
		t.Errorf("Structs not the same")
	}

}
