// This is an executable
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("⚡ Usage: ngram corp.txt 3")
		os.Exit(3)
	}

	filepath := os.Args[1]
	n := os.Args[2]

	dat, err := ioutil.ReadFile(filepath)
	check(err)

	i, err := strconv.Atoi(n)
	check(err)
	ngrams := orderByOccuranceCount(ngrams(string(dat), i))

	for _, gram := range ngrams {
		fmt.Println(gram)
	}
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

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

// Gram has the substring that forms an N-Gram and the number of occurances
type Gram struct {
	SubS string
	Occu int
}

/* TODO print control words based on flag
␉ 2409 symbol for horizontal tabulation
␊ 240A symbol for line feed
␍ 240D symbol for carriage return
*/

func (g Gram) String() string {
	return fmt.Sprintf("%s\t%d", g.SubS, g.Occu)
}

// Ngrams is a map of bigrams, or trigrams, or n-grams
type Ngrams []Gram

func (gs Ngrams) Len() int           { return len(gs) }
func (gs Ngrams) Less(i, j int) bool { return gs[i].Occu < gs[j].Occu }
func (gs Ngrams) Swap(i, j int)      { gs[i], gs[j] = gs[j], gs[i] }

func orderByOccuranceCount(ngramOccurances map[string]int) Ngrams {
	pl := make(Ngrams, len(ngramOccurances))
	i := 0
	for k, v := range ngramOccurances {
		pl[i] = Gram{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

// func normalize(s string, f norm.Form) string {
// 	var bs []byte
// 	iter := &norm.Iter{}
// 	iter.InitString(f, s)
// 	for !iter.Done() {
// 		bs = append(bs, iter.Next()...)
// 	}
// 	return string(bs)
// }
