package lib

type ChainLink struct {
	value    string
	category string
	prev     *ChainLink
	next     *ChainLink
}

type Chain struct {
	lang       string
	chainLinks []ChainLink
}
