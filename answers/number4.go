package main

import (
	"fmt"
	"sort"
)

func main() {
	listOfWords := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	groupedWords := make(map[string][]string)

	for _, word := range listOfWords {
		sorted := SortString(word)
		groupedWords[sorted] = append(groupedWords[sorted], word)
	}

	for _, v := range groupedWords {
		fmt.Println(v)
	}

}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
