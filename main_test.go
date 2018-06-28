package main

import (
	"reflect"
	"testing"
)

func TestNgrams(t *testing.T) {

	ds := []struct {
		n    int
		s    string
		want map[string]int
	}{
		{1, "ääbñbbñ", map[string]int{"ä": 2, "b": 3, "ñ": 2}},
		{1, "äbʋ̈baíä", map[string]int{"ä": 1, "b": 2, "ʋ̈": 1, "a": 1, "í": 1, "ä": 1}},
		{2, "ääbäbäääbbä", map[string]int{"ää": 3, "äb": 3, "bä": 3, "bb": 1}},
		{2, "äbbb", map[string]int{"äb": 1, "bb": 2}},
		{2, "äébbb", map[string]int{"äé": 1, "éb": 1, "bb": 2}},
		{2, "äébbbñ", map[string]int{"äé": 1, "éb": 1, "bb": 2, "bñ": 1}},
		{2, "äébbñ", map[string]int{"äé": 1, "éb": 1, "bb": 1, "bñ": 1}},
		{3, "äébäéb", map[string]int{"äéb": 2, "ébä": 1, "bäé": 1}},
		{3, "ébñäébäébñ", map[string]int{"äéb": 2, "ébä": 1, "bäé": 1, "ébñ": 2, "bñä": 1, "ñäé": 1}},
		{4, "ééééééééé", map[string]int{"éééé": 6}},
		{4, "ééééñäñ", map[string]int{"éééé": 1, "éééñ": 1, "ééñä": 1, "éñäñ": 1}},
		{5, "äbʋ̈baíä", map[string]int{"äbʋ̈ba": 1, "bʋ̈baí": 1, "ʋ̈baíä": 1}},
	}

	for i, d := range ds {
		got := ngrams(d.s, d.n)
		if !reflect.DeepEqual(got, d.want) {
			t.Errorf("TestNgrams %d: got %v, want %v", i, got, d.want)
		}
	}

}

// nfc, nfd
// func TestTidy(t *testing.T) {
// 	ds := []struct {
// 		s    string
// 		want string
// 		f    norm.Form
// 	}{
// 		{"̈ʋ̈baíä", "̈ʋ̈baíä", norm.NFD},
// 		{"̈ʋ̈baíä", "̈ʋ̈baíä", norm.NFC},
// 	}

// 	for i, d := range ds {
// 		got := normalize(d.s, d.f)
// 		if got != d.want {
// 			t.Errorf("TestTidy %d: got %v, want %v", i, got, d.want)
// 		}
// 	}
// }

// TODO test the sorting function
// func TestOrderByOccuranceCount(t *testing.T) {
// 	ds := []struct {
// 		gs   map[string]int
// 		want Ngrams
// 	}{
// 		{map[string]int{"a": 1, "b": 2}, Ngrams{{"b", 2}, {"a", 1}}},
// 	}

// 	for i, d := range ds {
// 		got := orderByOccuranceCount(d.gs)
// 		if got != d.want {
// 			t.Errorf("TestOrderByOccuranceCount %d: got %v, want %v", i, got, d.want)
// 		}
// 	}
// }
