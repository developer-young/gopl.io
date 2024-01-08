// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
	"strings"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// 交集
func (s *IntSet) IntersectWith(t* IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// 差集
func (s *IntSet) DifferenceWith(t *IntSet) {
	u := &IntSet{words: make([]uint64, len(s.words))}
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &= ^t.words[i]
		}
	}
	s = u.Copy()
}

// return the number of elements
func (s *IntSet) Len() int {
	str := s.String()[1:]
	items := strings.Fields(str[:len(str)-1])
	fmt.Printf("Len debug: items %v\n", items)
	return len(items)
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	if !s.Has(x) {
		return
	}
	word, bit := x/64, uint(x%64)
	s.words[word] &= (^(1 << bit))
}

// remove all elements from the set
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	dst := &IntSet{words: make([]uint64, len(s.words))}
	copy(dst.words, s.words)
	fmt.Printf("debug Copy: dst %v\n", dst.String())
	return dst
}

func (s *IntSet) AddAll(nums ...int) {
	for _, x := range nums {
		s.Add(x)
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string
