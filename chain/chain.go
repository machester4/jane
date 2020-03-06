package chain

import (
	"errors"
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

func createMainChain(text string) *Chain {
	var repeater int
	var last rune
	var c Chain

	for i, r := range text {
		category := getCategory(r)
		incrementRepeater(&repeater, r, last)

		if validateRepeaters(category, repeater) {
			// fmt.Println("repeating", &r)
			continue
		}
		c.addBlock(i, r, category)
		last = r
	}

	return &c
}

func (c *Chain) addBlock(index int, value rune, category string) {
	newBlock := &Block{
		Index:    index,
		Value:    value,
		Category: category,
	}
	if c.head == nil {
		c.head = newBlock
	} else {
		lastBlock := c.tail
		lastBlock.Next = newBlock
		newBlock.Previous = c.tail
	}
	c.tail = newBlock
}

func (c *Chain) Walk(callback func(b *Block)) error {
	currentBlock := c.head
	if currentBlock == nil {
		return errors.New("chain is empty")
	}
	for currentBlock.Next != nil {
		callback(currentBlock)
		currentBlock = currentBlock.Next
	}
	return nil
}

func New(text string) *Chain {
	return createMainChain(text)
}
