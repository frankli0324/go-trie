//go:build go1.20

package trie

import "unsafe"

func stringdata(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
