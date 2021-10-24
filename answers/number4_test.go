package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagramSuccess(t *testing.T) {
	groupedWords := Anagram()
	assert.Equal(t, len(groupedWords), 4, fmt.Sprintf("Wrong result : %d, expected = '4'", len(groupedWords)))
}

func TestAnagramFail(t *testing.T) {
	groupedWords := Anagram()
	assert.NotEqual(t, len(groupedWords), 5, fmt.Sprintf("Wrong result : %d, expected = '4'", len(groupedWords)))
}
