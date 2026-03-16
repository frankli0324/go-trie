package trie

type node struct {
	left, right int
	depth       int
	code        int
}

func (t *Trie) buildDart() []node {
	m := t.charset.m
	p := t.patterns
	nextBase := 1

	queue := []node{{0, len(p), 0, 1}}
	cursor := 0
	var current node
	for {
		if len(queue) == cursor {
			return queue
		}
		current = queue[cursor]
		cursor++

		var count int
		count, queue = t.fetch(current, queue)
		children := queue[len(queue)-count:]

		if count == 0 {
			continue
		}
		avail, fails := t.find(nextBase, current.depth, children)
		t.base[current.code] = avail

		for i := 0; i < len(children); i++ {
			code := avail + m[p[children[i].left][current.depth]]
			children[i].code = code
			t.check[code] = current.code
		}

		if float32(fails)/float32(avail-nextBase) >= 0.95 {
			nextBase = avail
		}
	}
}

func (t *Trie) fetch(c node, to []node) (count int, res []node) {
	p := t.patterns
	start := c.left

	if len(p[c.left]) <= c.depth {
		return 0, to // ensured by sorting logic
	}
	prev := p[c.left][c.depth] // make sure first not hit

	for i := c.left; i < c.right; i++ {
		if len(p[i]) <= c.depth {
			c.right = i
			break
		}
		if ch := p[i][c.depth]; ch != prev {
			to = append(to, node{left: start, right: i, depth: c.depth + 1})
			start, prev = i, ch
			count++
		}
	}
	if start != c.right {
		to = append(to, node{left: start, right: c.right, depth: c.depth + 1})
		count++
	}
	return count, to
}

func (t *Trie) find(nextBase, depth int, ch []node) (avail, fails int) {
	m := t.charset.m
	p := t.patterns

	for {
		ok := true
		t.resize(nextBase + t.charset.last + 1)

		for i := range ch {
			if t.check[m[p[ch[i].left][depth]]+nextBase] != 0 {
				fails++
				ok = false
				break
			}
		}
		if ok {
			return nextBase, fails
		}
		nextBase++
	}
}

func (t *Trie) resize(sz int) {
	if sz < len(t.check) {
		return
	}
	if sz < cap(t.check) {
		t.check = t.check[:sz]
		t.base = t.base[:sz]
		return
	}
	t.check = append(t.check, make([]int, sz-len(t.check))...)
	t.base = append(t.base, make([]int, sz-len(t.base))...)
}

func (t *Trie) insertDATOutput(v string, i int) bool {
	m := t.charset.m
	c := 1
	for i := 0; i < len(v); i++ {
		code := t.base[c] + m[v[i]]
		if t.check[code] != c {
			return false
		}
		c = code
	}
	t.output[c] = i
	return true
}
