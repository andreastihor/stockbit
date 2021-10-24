package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindStringInsideFirstBracketSuccess(t *testing.T) {
	result := FindStringInsideFirstBracket("disini (ada) kurungan")
	assert.Equal(t, result, "ada", fmt.Sprintf("Wrong result : %s, expected = 'ada'", result))
}

func TestFindStringInsideFirstBracketFail(t *testing.T) {
	result := FindStringInsideFirstBracket("disini (ada)urungan")
	assert.NotEqual(t, result, "adaa", fmt.Sprintf("Wrong result : %s, expected = 'ada k'", result))
}
