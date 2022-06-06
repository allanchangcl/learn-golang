package main

import (
	"fmt"
)

type house struct {
	material     string
	hasFireplace bool
	floors       int
}

// initial constructor
// func newHouse() *house {
// 	const (
// 		defaultFloors       = 2
// 		defaultHasFireplace = true
// 		defaultMaterial     = "wood"
// 	)

// 	h := &house{
// 		material:     defaultMaterial,
// 		hasFireplace: defaultHasFireplace,
// 		floors:       defaultFloors,
// 	}
// 	return h
// }

// add fuctional options to constructor
func newHouse(opts ...houseOption) *house {
	const (
		defaultFloors       = 2
		defaultHasFireplace = true
		defaultMaterial     = "wood"
	)

	h := &house{
		material:     defaultMaterial,
		hasFireplace: defaultHasFireplace,
		floors:       defaultFloors,
	}

	// loop through each option
	for _, opt := range opts {
		opt(h)
	}
	return h
}

// function type
type houseOption = func(*house)

func withConcrete() houseOption {
	return func(h *house) {
		h.material = "concrete"
	}
}

func withoutFireplace() houseOption {
	return func(h *house) {
		h.hasFireplace = false
	}
}

func withFloors(floors int) houseOption {
	return func(h *house) {
		h.floors = floors
	}
}

func main() {
	// fmt.Println("vim-go")
	threeFloorHouse := newHouse(
		withConcrete(),
		withFloors(3),
	)

	threeFloorHouseNofireplace := newHouse(
		withConcrete(),
		withFloors(3),
		withoutFireplace(),
	)

	fmt.Printf("%+v\n", *threeFloorHouse)
	fmt.Printf("%+v\n", *threeFloorHouseNofireplace)
}
