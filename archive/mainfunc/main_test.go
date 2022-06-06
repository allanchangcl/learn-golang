package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/kr/pretty"
	"github.com/matryer/is"
)

func Test(t *testing.T) {
	is := is.New(t)

	args := []string{"greeter", "David", "Kat", "Jon", "Natalie", "Mark"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.NoErr(err)

	out := stdout.String()
	// fmt.Println(fmt.Sprintf("%#v", out))
	fmt.Printf("%# v", pretty.Formatter(out))
	is.True(strings.Contains(out, "Hi David"))
	is.True(strings.Contains(out, "Hi Kat"))
	is.True(strings.Contains(out, "Hi Jon"))
	is.True(strings.Contains(out, "Hi Natalie"))
	is.True(strings.Contains(out, "Hi Mark"))

}

func TestNoNames(t *testing.T) {
	is := is.New(t)

	args := []string{"greeter"}
	var stdout bytes.Buffer
	err := run(args, &stdout)
	is.True(err != nil)
}