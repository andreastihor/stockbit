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
			remainingStr := str[firstOpeningBracketIndex:]
			firstClosingBracketIndex := strings.Index(remainingStr, ")")
			if firstClosingBracketIndex > 0 {
				lenOfStrInsideBracket := firstOpeningBracketIndex + firstClosingBracketIndex
				return str[firstOpeningBracketIndex+1 : lenOfStrInsideBracket]
			}

		}

	}

	return res
}
