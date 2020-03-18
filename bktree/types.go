package bktree

type BKNode struct {
	Str      string
	Children map[int]*BKNode
}

type BKTree struct {
	root         *BKNode
	DistanceFunc func(string, string) int
}