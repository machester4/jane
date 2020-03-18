package helpers

import (
	"github.com/machester4/jane/constants"
	"unicode"
)


// Error Helper
func CheckError(err error)  {
	if err != nil {
		panic(err)
	}
}

// Chain Helpers
func GetCharacterCategory(r rune) (category string) {
	switch {
	case unicode.IsLetter(r):
		category = constants.FieldTypeLetter
	case unicode.IsPunct(r):
		category = constants.FieldTypePunct
	case unicode.IsSpace(r):
		category = constants.FieldTypeSpace
	}
	return
}

func IsArticle(value string) (isArticle bool)  {
	for _, article := range constants.Articles {
		if article == value {
			isArticle = true
			break
		}
	}
	return
}

func IsRepeatedCharacter(category string, repeater int) (isRepeated bool) {
	switch category {
	case constants.FieldTypeLetter:
		isRepeated = repeater > constants.MaxRepeatLetter
	case constants.FieldTypePunct:
		isRepeated = repeater > constants.MaxRepeatPunct
	case constants.FieldTypeSpace:
		isRepeated = repeater > constants.MaxRepeatSpace
	}
	return
}