package lib

import (
	"unicode"
)

// checkError - if error launch a panic
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// getCharacterCategory - return the category of a character
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

// isArticle - return true if value is a valid article
func isArticle(value string) (isArticle bool) {
	for _, article := range articles {
		if article == value {
			isArticle = true
			break
		}
	}
	return
}

// isRepeatedCharacter - returns true if a category is repeating consecutively
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
