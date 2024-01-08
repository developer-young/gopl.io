package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "foo is a foo of all foos"
	f := func(s string) string {
		return "bar"
	}
	fmt.Println(expand(s, f))
}

func expand(s string, f func(string) string) string {
	return strings.ReplaceAll(s, "foo", f("foo"))
}
