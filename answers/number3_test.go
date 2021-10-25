package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindStringInsideFirstBracket(t *testing.T) {
	t.Run("When Successfully Find String", func(t *testing.T) {
		result := FindStringInsideFirstBracket("disini (ada) kurungan")
		assert.Equal(t, result, "ada", fmt.Sprintf("Wrong result : %s, expected = 'ada'", result))
	})

	t.Run("When fail to  Find String", func(t *testing.T) {
		result := FindStringInsideFirstBracket("disini (ada)urungan")
		assert.NotEqual(t, result, "adaa", fmt.Sprintf("Wrong result : %s, expected = 'ada k'", result))
	})
}
