package trie

type Cursor int

func (t *Trie) Cursor() Cursor {
	return Cursor(1)
}

func (t *Trie) advance(c int, ch int) int {
	for {
		code := ch + t.base[c]
		if code < len(t.check) && t.check[code] == c {
			return code
		} else if c == 1 {
			return 1
		}
		c = t.fail[c]
	}
}

func (t *Trie) Next(c Cursor, ch byte) Cursor {
	return Cursor(t.advance(int(c), t.charset.m[ch]))
}

func (t *Trie) Fetch(c Cursor, f func(idx int, s string)) {
	for c != 1 && c != 0 {
		if idx := t.output[c] - 1; idx != -1 {
			f(idx, t.patterns[idx])
		}
		c = Cursor(t.dict[int(c)])
	}
}

func (t *Trie) buildAhoCorasick(queue []node) {
	t.fail = make([]int, len(t.check))
	t.dict = make([]int, len(t.check))
	t.fail[1] = 1
	for i := 1; i <= t.charset.last; i++ {
		if code := t.base[1] + i; t.check[code] == 1 {
			t.fail[code] = 1
		}
	}
	for _, curr := range queue[1:] {
		for i := 1; i <= t.charset.last; i++ {
			if code := t.base[curr.code] + i; t.check[code] == curr.code {
				out := t.advance(t.fail[curr.code], i)
				t.fail[code] = out
				if t.output[out] != 0 {
					t.dict[code] = out
				} else if t.dict[out] != 0 {
					t.dict[code] = t.dict[out]
				}
			}
		}
	}
}
