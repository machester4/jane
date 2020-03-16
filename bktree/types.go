package bktree

type node struct {
	entry    Entry
	children []struct {
		distance int
		node     *node
	}
}

type Result struct {
	Distance int
	Entry    Entry
}

type BKTree struct {
	root *node
}

type Entry interface {
	Distance(Entry) int
}
