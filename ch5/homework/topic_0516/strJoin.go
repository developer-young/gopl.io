package main

import (
	"fmt"
	"strings"
)

func strJoin(strs []string, seps ...string) string {
	res := strings.Builder{}
	for i, j := 0, 0; i < len(strs); i++ {
		res.WriteString(strs[i])
		if j < len(seps) {
			res.WriteString(seps[j])
			j++
		}
	}
	return res.String()
}

func main() {

	s := []string{"hello", "world", "bye", ""}
	seps := []string{" ", "!"}

	res := strings.Join(s, ",")
	fmt.Println(res)

	res = strJoin(s, seps...)
	fmt.Println(res)
}
