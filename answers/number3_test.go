package main

import (
	"fmt"
	"testing"
)

func TestFindStringInsideFirstBracketSuccess(t *testing.T) {
	result := FindStringInsideFirstBracket("disini (ada) kurungan")
	if result != "ada" {
		t.Errorf(fmt.Sprintf("Wrong result : %s, expected = 'ada'", result))
	}
}

func TestFindStringInsideFirstBracketFail(t *testing.T) {
	result := FindStringInsideFirstBracket("disini (ada k)urungan")
	if result == "ada" {
		t.Error(fmt.Sprintf("Wrong result : %s, expected = 'ada k", result))
	}
}
