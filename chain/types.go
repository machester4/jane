package chain

type Chain struct {
	Blocks []Block
}

type Block struct {
	Index    int
	Value    string
	Category string
	Prev     *Block
	Next     *Block
}
