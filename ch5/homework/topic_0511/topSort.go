package main

import (
	"fmt"
)

// !+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string]map[string]struct{}{
	"algorithms": {"data structures": {}},
	"calculus":   {"linear algebra": {}},

	"compilers": {
		"data structures":       {},
		"formal languages":      {},
		"computer organization": {},
	},

	"data structures": {"discrete math": {}},
	"databases":       {"data structures": {}},
	// "discrete math":         {"intro to programming": {}},
	"formal languages":      {"discrete math": {}},
	"networks":              {"operating systems": {}},
	"operating systems":     {"data structures": {}, "computer organization": {}},
	"programming languages": {"data structures": {}, "computer organization": {}},

	"discrete math": {"programming languages": {}, "intro to programming": {}}, // mock ring
}

//!-table

// !+main
func main() {
	fmt.Printf("%v\n", topoSort(prereqs))
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// 现在线性代数的老师把微积分设为了前置课程。完善topSort，使其能检测有向图中的环
func topoSort(m map[string]map[string]struct{}) []string {
	var order []string
	seen := make(map[string]bool)
	parents := make(map[string]struct{})
	var visitOne func(items string)
	hasRing := false
	visitOne = func(item string) {
		_, exist := parents[item]
		if exist {
			hasRing = true
			return
		}
		if seen[item] {
			return
		}

		seen[item] = true
		for child := range m[item] {
			parents[item] = struct{}{}
			visitOne(child)
			delete(parents, item)
		}
		order = append(order, item)
	}

	// var keys []string
	// keys := make(map[string]struct{})
	for key := range m {
		visitOne(key)
	}

	if hasRing {
		order = nil
	}
	// sort.Strings(keys)
	return order
}
