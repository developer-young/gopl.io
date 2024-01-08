package main

import (
	"fmt"
	"math"
)

func max(nums ...int) (int, error) {
	var err error
	if len(nums) < 1 {
		err = fmt.Errorf("At least a argument")
		return 0, err
	}
	x := math.MinInt
	for _, y := range nums {
		if x < y {
			x = y
		}
	}
	return x, err
}

func min(nums ...int) (int, error) {
	var err error
	if len(nums) < 1 {
		err = fmt.Errorf("At least a argument")
		return 0, err
	}
	x := math.MaxInt
	for _, y := range nums {
		if x > y {
			x = y
		}
	}
	return x, err
}

func main() {
	s := [8]int{1, 99, 4389, -4389, 5489, -3489, 438921, 54}

	maxv, err := max(s[:]...)
	if err != nil {
		fmt.Println(err)
	}

	minv, err := min(s[:]...)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("max: %v, min: %v\n", maxv, minv)
}
