package lib

import "strings"

type field = RecommendResultItem
type recommend = RecommendItemSuggestion

type article struct {
	start      int
	offset     int
	value      string
	noun       *field
	recommends []*recommend
}

type chain struct {
	words       []*field
	pucts       []*field
	articles    []*article
	headArticle *article // For set Noun (sustantivo luego del articulo)
}

func incrementRepeater(repeater *int, current rune, last rune) {
	if last == current {
		*repeater++
	} else {
		*repeater = 0
	}
}

func (c *chain) addWord(f *field) {
	isArticle := isArticle(f.Value)
	if isArticle {
		c.addArticle(f)
	} else {
		if c.headArticle != nil {
			c.headArticle.noun = f
			c.headArticle = nil
		}
		c.words = append(c.words, f)
	}
}

func (c *chain) addArticle(f *field) {
	art := article{
		start:  f.Start,
		offset: f.Offset,
		value:  f.Value,
	}
	c.articles = append(c.articles, &art)
	c.headArticle = &art
}

func (c *chain) addPunct(f *field) {
	c.pucts = append(c.pucts, f)
}

func (c *chain) add(value string, offset int, category string) {
	field := &field{
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
	var c chain
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
			c.add(field, i, getCharacterCategory(last))
			field = ""
		case fieldTypePunct:
			c.add(string(r), i+1, category)
			c.add(field, i, getCharacterCategory(last))
			field = ""
		default:
			field += string(r)
		}
		last = r
	}

	// Add last field
	if field != "" {
		c.add(field, len(text), getCharacterCategory(last))
	}

	return &c
}
