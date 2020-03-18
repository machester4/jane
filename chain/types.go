package chain

import (
	"github.com/agnivade/levenshtein"
	"github.com/machester4/jane/bktree"
)

type BkWord string
func (x BkWord) Distance(e bktree.Entry) int {
	a := string(x)
	b := string(e.(BkWord))

	return levenshtein.ComputeDistance(a, b)
}
type field struct {
	Start      int
	Offset     int
	Value      string
	Recommends []BkWord
}
type Word = field
type Punct = field
type Article struct {
	Start      int
	Offset     int
	Value      string
	Noun       *Word
	Recommends []string
}
type Chain struct {
	Words       []*Word
	Pucts       []*Punct
	Articles    []*Article
	headArticle *Article // For set Noun (sustantivo luego del articulo)
}
