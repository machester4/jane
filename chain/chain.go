package chain

import (
	"github.com/machester4/jane/constants"
	"github.com/machester4/jane/helpers"
	"strings"
)

func incrementRepeater(repeater *int, current rune, last rune) {
	if last == current {
		*repeater++
	} else {
		*repeater = 0
	}
}

func (c *Chain) addWord(field *field)  {
	isArticle := helpers.IsArticle(field.Value)
	if isArticle {
		c.addArticle(field)
	} else {
		if c.headArticle != nil {
			c.headArticle.Noun = field
			c.headArticle = nil
		}
		c.Words = append(c.Words, field)
	}
}

func (c *Chain) addArticle(field *field)  {
	art := Article{
		Start:  field.Start,
		Offset: field.Offset,
		Value:  field.Value,
	}
	c.Articles = append(c.Articles, &art)
	c.headArticle = &art
}

func (c *Chain) addPunct(field *field)  {
	c.Pucts = append(c.Pucts, field)
}

func (c *Chain) add(value string, offset int, category string) {
	field := &field{
		Start:  offset - len(value),
		Offset: offset,
		Value:  strings.ToLower(value),
	}
	switch category {
	case constants.FieldTypeLetter:
		c.addWord(field)
	case constants.FieldTypePunct:
		c.addPunct(field)
	}
}

func New(text string) *Chain {
	var chain Chain
	var repeater int
	var last rune
	var field string

	for i, r := range text {
		category := helpers.GetCharacterCategory(r)
		incrementRepeater(&repeater, r, last)

		if helpers.IsRepeatedCharacter(category, repeater) || category == "" {
			continue
		}
		switch category {
		case constants.FieldTypeSpace:
			chain.add(field, i-1, helpers.GetCharacterCategory(last))
			field = ""
		case constants.FieldTypePunct:
			chain.add(string(r), i, category)
			chain.add(field, i-1, helpers.GetCharacterCategory(last))
			field = ""
		default:
			field += string(r)
		}
		last = r
	}

	// Add last field
	if field != "" {
		chain.add(field, len(text), helpers.GetCharacterCategory(last))
	}

	return &chain
}
