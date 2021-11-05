package main

import "fmt"

//go:generate easytags $GOFILE
type Large struct {
	Field            string `json:"field"`
	OtherField       string `json:"other_field"`
	AnotherField     string `json:"another_field"`
	YetAnotherField  string `json:"yet_another_field"`
	UnnecessaryField string `json:"unnecessary_field"`
	ForgottenField   string `json:"forgotten_field"`
}

func main() {
	fmt.Println("vim-go")
}
