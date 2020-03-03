package chain

import (
	"github.com/machester4/jane/constants"
	"unicode"
)

func validateRepeaters(category string, repeat int)  (isRepeated bool) {
	switch category {
	case constants.BLOCK_TYPE_LETTER:
		isRepeated = repeat == constants.MAX_REPEAT_LETER
	case constants.BLOCK_TYPE_PUNCT:
		isRepeated = repeat == constants.MAX_REPEAT_PUNCT
	case constants.BLOCK_TYPE_SPACE:
		isRepeated = repeat == constants.MAX_REPEAT_SPACE
	}

	return
}

func incrementRepeater(repeat *int, current rune, last rune) {
	if last == current {
		*repeat++
	} else {
		*repeat = 0
	}
}

func createMainChain(text string) Chain  {
	var repeat int
	var last rune
	var chain Chain

	for i, r := range text {
		var category string
		incrementRepeater(&repeat, r, last)
		if validateRepeaters(category, repeat) {
			continue
		}

		switch  {
		case unicode.IsLetter(r):
			category = constants.BLOCK_TYPE_LETTER
		case unicode.IsPunct(r):
			category = constants.BLOCK_TYPE_PUNCT
		case unicode.IsSpace(r):
			category = constants.BLOCK_TYPE_SPACE
		}

		block := Block{
			Index:    i,
			Value:    r,
			Category: category,
			Prev:     nil,
			Next:     nil,
		}
		chain.Blocks = append(chain.Blocks, block)
		last = r
	}

	return chain
}

func New(text string) (chain Chain) {
	// TODO: Falta refactoring la main chain asignando los blockes previos y siguientes
	chain = createMainChain(text)
	return
}
