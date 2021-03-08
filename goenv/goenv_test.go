package main

import (
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	result := "Hello World"

	if result != "Hello World" {
		t.Errorf("Result was incorrect, got: %s, want: %s.",
			result, "Hello World")
	}
}

// Test go enviroments
func TestGoEnv(t *testing.T) {
	_, ok := os.LookupEnv("MONGO_PW")
	if !ok {
		// fmt.Println("error: unable to find MONGO_PW in the environment")
		t.Errorf("error: unable to find MONGO_PW in the environment")
	}
}
