package lib

import (
	"unicode"
)

// Error Helper
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Chain Helpers
func getCharacterCategory(r rune) (category string) {
	switch {
	case unicode.IsLetter(r):
		category = fieldTypeLetter
	case unicode.IsPunct(r):
		category = fieldTypePunct
	case unicode.IsSpace(r):
		category = fieldTypeSpace
	}
	return
}

func isArticle(value string) (isArticle bool) {
	for _, article := range articles {
		if article == value {
			isArticle = true
			break
		}
	}
	return
}

func isRepeatedCharacter(category string, repeater int) (isRepeated bool) {
	switch category {
	case fieldTypeLetter:
		isRepeated = repeater > maxRepeatLetter
	case fieldTypePunct:
		isRepeated = repeater > maxRepeatPunct
	case fieldTypeSpace:
		isRepeated = repeater > maxRepeatSpace
	}
	return
}
