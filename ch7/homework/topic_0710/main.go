package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] > p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	// data1 := IntSlice{1, 2, 3, 2, 1}
	// sort.Sort(data1)
	// fmt.Println(data1)

	data2 := IntSlice{1, 2, 3, 2, 1}
	fmt.Println(IsPalindrome(data2)) // Output: true

	data3 := IntSlice{1, 2, 3, 2, 1, 5}
	fmt.Println(IsPalindrome(data3)) // Output: false
}
