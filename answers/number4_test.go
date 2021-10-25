package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagram(t *testing.T) {
	t.Run("When Successfully Group Anagram", func(t *testing.T) {
		groupedWords := Anagram()
		assert.Equal(t, len(groupedWords), 4, fmt.Sprintf("Wrong result : %d, expected = '4'", len(groupedWords)))
	})

	t.Run("When fail to Group Anagram", func(t *testing.T) {
		groupedWords := Anagram()
		assert.Equal(t, len(groupedWords), 4, fmt.Sprintf("Wrong result : %d, expected = '4'", len(groupedWords)))
	})

}
