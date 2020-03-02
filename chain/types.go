package chain

type Chain struct {
	blocks []Block
}

type Block struct {
	index    int
	value    string
	category string
	prev     *Block
	next     *Block
}
