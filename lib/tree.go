package lib

import (
	"github.com/agnivade/levenshtein"
)

type bkresult = RecommendItemSuggestion

type bknode struct {
	entry    string
	children []struct {
		distance int
		node     *bknode
	}
}

type bktree struct {
	root *bknode
}

func (n *bknode) addChild(e string) *bknode {
	newnode := &bknode{entry: e}
loop:
	d := n.diffDistance(e)
	for _, c := range n.children {
		if c.distance == d {
			n = c.node
			goto loop
		}
	}
	n.children = append(n.children, struct {
		distance int
		node     *bknode
	}{d, newnode})
	return newnode
}

// Add node to tree
func (b *bktree) add(entry string) *bknode {
	if b.root == nil {
		b.root = &bknode{
			entry: entry,
		}
		return b.root
	}
	return b.root.addChild(entry)
}

// search in tree - tolerance it represents edit distance
func (b *bktree) search(needle string, tolerance int) []*bkresult {
	results := make([]*bkresult, 0)
	if b.root == nil {
		return results
	}
	candidates := []*bknode{b.root}
	for len(candidates) != 0 {
		c := candidates[0]
		candidates = candidates[1:]

		d := c.diffDistance(needle)
		if d <= tolerance {
			results = append(results, &bkresult{
				Distance: d,
				Entry:    c.entry,
			})
		}

		low, high := d-tolerance, d+tolerance
		for _, c := range c.children {
			if low <= c.distance && c.distance <= high {
				candidates = append(candidates, c.node)
			}
		}
	}
	return results
}

// search in tree starting in node - tolerance it represents edit distance
func (n *bknode) search(needle string, tolerance int) []*bkresult {
	results := make([]*bkresult, 0)
	candidates := []*bknode{n}
	for len(candidates) != 0 {
		c := candidates[0]
		candidates = candidates[1:]

		d := c.diffDistance(needle)
		if d <= tolerance {
			results = append(results, &bkresult{
				Distance: d,
				Entry:    c.entry,
			})
		}

		low, high := d-tolerance, d+tolerance
		for _, c := range c.children {
			if low <= c.distance && c.distance <= high {
				candidates = append(candidates, c.node)
			}
		}
	}
	return results
}

func (n *bknode) diffDistance(str string) int {
	return levenshtein.ComputeDistance(n.entry, str)
}
