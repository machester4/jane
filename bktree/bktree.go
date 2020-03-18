package bktree

import "github.com/agnivade/levenshtein"

func (n *BKNode) addChild(e Entry) {
	newnode := &BKNode{entry: e}
loop:
	d := n.entry.Distance(e)
	for _, c := range n.children {
		if c.distance == d {
			n = c.node
			goto loop
		}
	}
	n.children = append(n.children, struct {
		distance int
		node     *BKNode
	}{d, newnode})
}

func (b *BKTree) Add(entry Entry) {
	if b.root == nil {
		b.root = &BKNode{
			entry: entry,
		}
		return
	}
	b.root.addChild(entry)
}

func (b *BKTree) Search(needle Entry, tolerance int, limit int) []*Result {
	results := make([]*Result, 0)
	if b.root == nil {
		return results
	}
	candidates := []*BKNode{b.root}
	for len(candidates) != 0 && len(results) != limit {
		c := candidates[len(candidates)-1]
		candidates = candidates[:len(candidates)-1]
		d := c.entry.Distance(needle)
		if d <= tolerance {
			results = append(results, &Result{
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

func (x Word) Distance(e Entry) int {
	a := string(x)
	b := string(e.(Word))

	return levenshtein.ComputeDistance(a, b)
}