// This is an executable
package main

import (
	norm "golang.org/x/text/unicode/norm"
)

func main() {
	// TODO
	// Get input from file or standard in (pipe to this binary)
	// Print the map of ngrams and counts
}

func ngrams(s string, n int) map[string]int {
	rs := []rune(s)
	counts := make(map[string]int)
	for i := range rs {
		if i+n > len(rs) {
			break
		}
		counts[string(rs[i:i+n])]++
	}
	return counts
}

func normalize(s string, f norm.Form) string {
	var bs []byte
	iter := &norm.Iter{}
	iter.InitString(f, s)
	for !iter.Done() {
		bs = append(bs, iter.Next()...)
	}
	// fmt.Println(string(bs))
	return string(bs)
}
