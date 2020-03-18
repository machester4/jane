package bktree

type Word string

type Entry interface {
	Distance(Entry) int
}

type BKNode struct {
	entry    Entry
	children []struct {
		distance int
		node     *BKNode
	}
}

type BKTree struct {
	root         *BKNode
	DistanceFunc func(string, string) int
}

type Result struct {
	Distance int
	Entry    Entry
}