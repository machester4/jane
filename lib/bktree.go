package lib

import "github.com/agnivade/levenshtein"

type bkword string

type bkentry interface {
	distance(bkentry) int
}

type bknode struct {
	entry    bkentry
	children []struct {
		distance int
		node     *bknode
	}
}

type bktree struct {
	root *bknode
}

func (n *bknode) addChild(e bkentry) {
	newnode := &bknode{entry: e}
loop:
	d := n.entry.distance(e)
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

func (b *bktree) add(entry bkentry) {
	if b.root == nil {
		b.root = &bknode{
			entry: entry,
		}
		return
	}
	b.root.addChild(entry)
}

func (b *bktree) search(needle bkentry, tolerance int) []*Result {
	results := make([]*Result, 0)
	if b.root == nil {
		return results
	}
	candidates := []*bknode{b.root}
	for len(candidates) != 0 {
		c := candidates[len(candidates)-1]
		candidates = candidates[:len(candidates)-1]
		d := c.entry.distance(needle)
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

func (x bkword) distance(e bkentry) int {
	a := string(x)
	b := string(e.(bkword))

	return levenshtein.ComputeDistance(a, b)
}
