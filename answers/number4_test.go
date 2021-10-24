package main

import (
	"fmt"
	"testing"
)

func TestAnagramSuccess(t *testing.T) {
	groupedWords := Anagram()
	if len(groupedWords) != 4 {
		t.Errorf(fmt.Sprintf("Wrong result : %d, expected = '4'", len(groupedWords)))
	}
}
