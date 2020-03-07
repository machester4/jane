package chain

type Chain struct {
	head *Block
	tail *Block
}

type Block struct {
	IndexInChain int
	IndexInText  int
	Value        rune
	Category     string
	Previous     *Block
	Next         *Block
}

type Word struct {
	Start           int
	Length          int
	Value           string
	recommendations []string
}
