package chain

type Chain struct {
	Blocks []Block
}

type Block struct {
	Index    int
	Value    rune
	Category string
	Prev     *Block
	Next     *Block
}
