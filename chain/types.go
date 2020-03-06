package chain

type Chain struct {
	head *Block
	tail *Block
}

type Block struct {
	Index    int
	Value    rune
	Category string
	Previous *Block
	Next     *Block
}
