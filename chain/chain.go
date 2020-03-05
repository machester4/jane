package chain

import (
	"fmt"
	"unicode"

	"github.com/machester4/jane/constants"
)

func validateRepeaters(category string, repeater int) (isRepeated bool) {
	switch category {
	case constants.BlockTypeLetter:
		isRepeated = repeater > constants.MaxRepeatLeter
	case constants.BlockTypePunct:
		isRepeated = repeater > constants.MaxRepeatPunct
	case constants.BlockTypeSpace:
		isRepeated = repeater > constants.MaxRepeatSpace
	}
	return
}

func getCategory(r rune) (category string) {
	switch {
	case unicode.IsLetter(r):
		category = constants.BlockTypeLetter
	case unicode.IsPunct(r):
		category = constants.BlockTypePunct
	case unicode.IsSpace(r):
		category = constants.BlockTypeSpace
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

func createMainChain(text string) Chain {
	var repeater int
	var last rune
	var chain Chain

	for i, r := range text {
		category := getCategory(r)
		incrementRepeater(&repeater, r, last)

		if validateRepeaters(category, repeater) {
			fmt.Println("repeating", &r)
			continue
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

func isOutOfRange(index int, dir string, blocks []Block) bool {
	length := len(blocks)
	if dir == "prev" {
		return index < 1
	}
	return index > length
}

func joinBlocks(chain *Chain) {
	for i, b := range chain.Blocks {
		if !isOutOfRange(i, "prev", chain.Blocks) {
			prev := &chain.Blocks[i-1]
			b.Prev = prev
		} else if !isOutOfRange(i, "next", chain.Blocks) {
			next := &chain.Blocks[i+1]
			b.Prev = next
		}
	}
}

func New(text string) (chain Chain) {
	// TODO: Falta refactoring la main chain asignando los blockes previos y siguientes
	chain = createMainChain(text)
	joinBlocks(&chain)
	return
}
