package chain

import (
	"unicode"

	"github.com/machester4/jane/constants"
)

func isRepeatedCharacter(category string, repeater int) (isRepeated bool) {
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

func getCategory(r rune) (category string) {
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

func incrementRepeater(repeater *int, current rune, last rune) {
	if last == current {
		*repeater++
	} else {
		*repeater = 0
	}
}

func (c *Chain) add(value string, offset int, category string) {
	field := field{
		Start:  offset - len(value),
		Offset: offset,
		Value:  value,
	}

	switch category {
	case constants.FieldTypeLetter:
		var isArticle bool
		for _, article := range constants.Articles {
			if article == value {
				isArticle = true
				break
			}
		}
		if isArticle {
			art := Article{
				Start:  field.Start,
				Offset: field.Offset,
				Value:  field.Value,
			}
			c.Articles = append(c.Articles, &art)
			c.headArticle = &art
		} else {
			if c.headArticle != nil {
				c.headArticle.Noun = &field
				c.headArticle = nil
			}
			c.Words = append(c.Words, &field)
		}
	case constants.FieldTypePunct:
		c.Pucts = append(c.Pucts, &field)
	}
}

func New(text string) *Chain {
	var chain Chain
	var repeater int
	var last rune
	var field string

	for i, r := range text {
		category := getCategory(r)
		incrementRepeater(&repeater, r, last)

		if isRepeatedCharacter(category, repeater) || category == "" {
			continue
		}
		switch category {
		case constants.FieldTypeSpace:
			chain.add(field, i, getCategory(last))
			field = ""
		case constants.FieldTypePunct:
			chain.add(string(r), i, category)
			field = ""
		default:
			field += string(r)
		}
		last = r
	}

	// Add last field
	if field != "" {
		chain.add(field, len(text), getCategory(last))
	}

	return &chain
}
