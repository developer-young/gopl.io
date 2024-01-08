// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Printf("x=%s, len(x): %d\n", x.String(), x.Len())
}

func TestRemove(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(5)
	x.Add(9)

	x.Remove(1)
	x.Remove(9)
	x.Remove(100)
	fmt.Printf("After Remove, x=%s, len(x): %d\n", x.String(), x.Len())
}

func TestClear(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(5)
	x.Add(9)
	x.Clear()

	fmt.Printf("After Clear, x=%s, len(x): %d\n", x.String(), x.Len())
}

func TestCopy(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(5)
	x.Add(9)

	y := x.Copy()
	x.Remove(9)

	fmt.Printf("x: %v\n", x.String())
	fmt.Printf("y: %v\n", y.String())
}

func TestAddAll(t *testing.T) {
	var s IntSet
	nums := []int{1, 2, 16, 32, 64}
	s.AddAll(nums...)
	fmt.Printf("s: %v\n", s.String())
}

func TestIntersectWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(1)
	y.Add(144)
	fmt.Println(y.String())

	x.IntersectWith(&y)
	fmt.Println(x.String())
}

func TestDifferenceWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(5)
	x.Add(256)
	fmt.Println(x.String())

	y.Add(1)
	y.Add(144)
	y.Add(9)
	y.Add(255)
	fmt.Println(y.String())	

	x.DifferenceWith(&y)
	fmt.Println(x.String())
}

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}
