package lib

import (
	"strings"
	"unicode"
)

func Sanatize(text *string) *string {
	last := new(rune)
	*text = strings.Map(func(r rune) rune {
		if (unicode.IsLetter(r) || unicode.IsSpace(r) || unicode.IsPunct(r)) && r != *last {
			return r
		}
		return rune(0)
	}, *text)
	return text
}
