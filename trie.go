package trie

import (
	"sort"
	"strings"
)

type charset struct {
	m    [256]int
	last int
}

func (c *charset) insert(b []byte) {
	m := c.m[:]
	for i := range b {
		if m[b[i]] == 0 {
			c.last++
			m[b[i]] = c.last
		}
	}
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

func New(patterns []string) *Trie {
	t := new(Trie)
	sort.SliceStable(patterns, func(i, j int) bool {
		if strings.HasPrefix(patterns[i], patterns[j]) && len(patterns[i]) > len(patterns[j]) {
			return true
		}
		if strings.HasPrefix(patterns[j], patterns[i]) && len(patterns[j]) > len(patterns[i]) {
			return false
		}
		return patterns[i] < patterns[j]
	})
	t.patterns = patterns
	for i := range patterns {
		t.charset.insert(stringdata(patterns[i]))
	}
	queue := t.buildDart()
	t.output = make([]int, len(t.check))
	for i := range patterns {
		t.insertDATOutput(stringdata(patterns[i]), i+1)
	}
	t.buildAhoCorasick(queue)
	return t
}
