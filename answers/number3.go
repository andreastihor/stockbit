package main

import "strings"

// FindStringInsideFirstBracket return string inside the first () found
/*
 e.g : findStringInsideFirstBracket("there is a (word)a) inside this the bracket")
		return : word
*/
func FindStringInsideFirstBracket(str string) (res string) {
	if len(str) > 0 {
		firstOpeningBracketIndex := strings.Index(str, "(")
		if firstOpeningBracketIndex >= 0 {
			runes := []rune(str)
			insideWord := string(runes[firstOpeningBracketIndex+1 : len(str)])
			firstClosingBracketIndex := strings.Index(insideWord, ")")
			if firstClosingBracketIndex > 0 {
				return insideWord[0:firstClosingBracketIndex]
			}

		}

	}

	return res
}
