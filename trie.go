package trie

import (
	"sort"
)

type charset struct {
	m    [256]int
	last int
}

func (c *charset) insert(b string) {
	m, l := c.m[:], c.last
	for i := 0; i < len(b); i++ {
		if m[b[i]] == 0 {
			l++
			m[b[i]] = l
		}
	}
	c.last = l
}

type Trie struct {
	base  []int
	check []int

	output []int
	dict   []int
	fail   []int

	charset  // char code starts with 1
	patterns []string
}

func strictprefix(s, prefix string) bool {
	return len(s) > len(prefix) && s[:len(prefix)] == prefix
}

func New(patterns []string) *Trie {
	t := new(Trie)
	sort.SliceStable(patterns, func(i, j int) bool {
		if strictprefix(patterns[i], patterns[j]) {
			return true
		}
		if strictprefix(patterns[j], patterns[i]) {
			return false
		}
		return patterns[i] < patterns[j]
	})
	t.patterns = patterns
	for i := range patterns {
		t.charset.insert(patterns[i])
	}
	queue := t.buildDart()
	t.output = make([]int, len(t.check))
	for i := range patterns {
		t.insertDATOutput(patterns[i], i+1)
	}
	t.buildAhoCorasick(queue)
	return t
}
