package lib

import (
	"strings"
)

type Field struct {
	Start      int
	Offset     int
	Value      string
	Forbidden  bool
	Recommends []*Result
}

type Article struct {
	Start      int
	Offset     int
	Value      string
	Noun       *Field
	Recommends []string
}

type chain struct {
	Words       []*Field
	Pucts       []*Field
	Articles    []*Article
	headArticle *Article // For set Noun (sustantivo luego del articulo)
}

func incrementRepeater(repeater *int, current rune, last rune) {
	if last == current {
		*repeater++
	} else {
		*repeater = 0
	}
}

func (c *chain) addWord(f *Field) {
	isArticle := isArticle(f.Value)
	if isArticle {
		c.addArticle(f)
	} else {
		if c.headArticle != nil {
			c.headArticle.Noun = f
			c.headArticle = nil
		}
		c.Words = append(c.Words, f)
	}
}

func (c *chain) addArticle(f *Field) {
	art := Article{
		Start:  f.Start,
		Offset: f.Offset,
		Value:  f.Value,
	}
	c.Articles = append(c.Articles, &art)
	c.headArticle = &art
}

func (c *chain) addPunct(f *Field) {
	c.Pucts = append(c.Pucts, f)
}

func (c *chain) add(value string, offset int, category string) {
	field := &Field{
		Start:  offset - len(value),
		Offset: offset,
		Value:  strings.ToLower(value),
	}
	switch category {
	case fieldTypeLetter:
		c.addWord(field)
	case fieldTypePunct:
		c.addPunct(field)
	}
}

func createChain(text string) *chain {
	var chain chain
	var repeater int
	var last rune
	var field string

	for i, r := range text {
		category := getCharacterCategory(r)
		incrementRepeater(&repeater, r, last)

		if isRepeatedCharacter(category, repeater) || category == "" {
			continue
		}
		switch category {
		case fieldTypeSpace:
			chain.add(field, i-1, getCharacterCategory(last))
			field = ""
		case fieldTypePunct:
			chain.add(string(r), i, category)
			chain.add(field, i-1, getCharacterCategory(last))
			field = ""
		default:
			field += string(r)
		}
		last = r
	}

	// Add last field
	if field != "" {
		chain.add(field, len(text), getCharacterCategory(last))
	}

	return &chain
}
