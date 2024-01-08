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

	"data structures":       {"discrete math": {}},
	"databases":             {"data structures": {}},
	"discrete math":         {"intro to programming": {}},
	"formal languages":      {"discrete math": {}},
	"networks":              {"operating systems": {}},
	"operating systems":     {"data structures": {}, "computer organization": {}},
	"programming languages": {"data structures": {}, "computer organization": {}},
}

//!-table

// !+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// 重写topoSort函数，用map代替切片并移除对key的排序代码。验证结果的正确性
func topoSort(m map[string]map[string]struct{}) []string {
	// order := make(map[string]struct{})
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]struct{})

	visitAll = func(items map[string]struct{}) {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	// var keys []string
	keys := make(map[string]struct{})
	for key := range m {
		keys[key] = struct{}{}
	}

	// sort.Strings(keys)
	visitAll(keys)
	return order
}
