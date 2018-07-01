package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"golang.org/x/text/unicode/runenames"
)

// ngram ---
func runNgram(cnf ngramConf) error {
	dat, err := ioutil.ReadFile(cnf.file)
	check(err)
	sDat := string(dat)

	ngrams := orderByOccuranceCount(ngrams(sDat, cnf.n))

	for _, gram := range ngrams {
		if cnf.names {
			fmt.Print(gram, "\t")
			for _, r := range gram.SubS {
				fmt.Print(runenames.Name(r), ", ")
			}
			fmt.Print("\n")
		} else {
			fmt.Println(gram)
		}

	}

	return nil
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

// Sub-commands ---
func runCommand(c *Conf) error {
	switch c.cmd {
	case c.ngram.fs.Name():
		return runNgram(c.ngram)
	default:
		return fmt.Errorf("missing/unknown command")
	}
}
