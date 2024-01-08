package main

import "fmt"

func f(x *int) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("recover error: %v\n", p)
			panic(p)
		}
	}()
	*x = (*x) * (*x)
}

func main() {
	var x int
	f(nil)
	fmt.Println(x)
}
