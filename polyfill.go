//go:build !go1.20

package trie

import "unsafe"

type stringheader struct {
	Data *byte
	Len  int
}

func stringdata(s string) []byte {
	h := (*stringheader)(unsafe.Pointer(&s))
	return unsafe.Slice(h.Data, h.Len)
}
