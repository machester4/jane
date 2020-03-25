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

func (n *bknode) addChild(e string) {
	newnode := &bknode{entry: e}
loop:
	d := distance(n.entry, e)
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
}

// Add node to tree
func (b *bktree) add(entry string) {
	if b.root == nil {
		b.root = &bknode{
			entry: entry,
		}
		return
	}
	b.root.addChild(entry)
}

// search in tree - tolerance it represents edit distance
func (b *bktree) search(needle string, tolerance int) []*bkresult {
	results := make([]*bkresult, 0)
	if b.root == nil {
		return results
	}
	candidates := []*bknode{b.root}
	for len(candidates) != 0 {
		c := candidates[len(candidates)-1]
		candidates = candidates[:len(candidates)-1]
		d := distance(c.entry, needle)
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

func distance(a string, b string) int {
	return levenshtein.ComputeDistance(a, b)
}
