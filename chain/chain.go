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

		if validateRepeaters(category, repeater) && category != "" {
			// fmt.Println("repeating", &r)
			continue
		}
		c.addBlock(i, r, category)
		last = r
	}

	return &c
}

func (c *Chain) addBlock(indexInText int, value rune, category string) {
	var newBlock *Block
	var lastBlock *Block

	newBlock = &Block{
		IndexInText: indexInText,
		Value:       value,
		Category:    category,
	}
	lastBlock = c.tail

	if c.head == nil {
		c.head = newBlock
	} else {
		newBlock.IndexInChain = lastBlock.IndexInChain + 1
		newBlock.Previous = lastBlock
		lastBlock.Next = newBlock
	}
	c.tail = newBlock
}

func (c *Chain) Walk(callback func(b *Block)) error {
	currentBlock := c.head
	if currentBlock == nil {
		return errors.New("chain is empty")
	}

	for currentBlock.Next != nil || currentBlock.Previous != nil {
		callback(currentBlock)
		if currentBlock.Next != nil {
			currentBlock = currentBlock.Next
		} else {
			break
		}
	}
	return nil
}

func (c *Chain) GetAllBlocks() (blocks []*Block) {
	c.Walk(func(b *Block) {
		blocks = append(blocks, b)
	})
	return
}

func (c *Chain) GetWords() (words []Word) {
	var currentWord Word

	c.Walk(func(b *Block) {
		if b.Category == constants.BlockTypeLetter {
			currentWord.Value = append(currentWord.Value, b)
			currentWord.Start = b.IndexInText
			currentWord.Length = b.IndexInText
		}

		if b.Category == constants.BlockTypeSpace {
			words = append(words, currentWord)
			/* words = append(words, Word{
				Start:  b.IndexInText,
				Length: b.IndexInText,
				Value:  string(b.Value),
			}) */
			currentWord.Value = nil
		}
	})
	// Append last word
	if currentWord.Value != nil {
		words = append(words, currentWord)
	}
	return
}

func (w *Word) ToString() string {
	var result string
	for _, b := range w.Value {
		result += string(b.Value)
	}
	return result
}

func New(text string) *Chain {
	return createMainChain(text)
}
