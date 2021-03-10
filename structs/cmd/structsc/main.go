package main

import (
	"encoding/json"
	"fmt"
)

type simpleUser struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

var (
	curUserId = 0
	// array of structs
	users []simpleUser
)

func newSimpleUser(name string) simpleUser {
	u := simpleUser{
		Name: name,
		ID:   curUserId,
	}
	curUserId++
	users = append(users, u)

	return u
}

func main() {
	newSimpleUser("Ricky")
	newSimpleUser("Snowee")

	for _, u := range users {
		// fmt.Println(strings.ToUpper(u.Name))
		jsonBytes, err := json.Marshal(u)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonBytes))
	}
}
