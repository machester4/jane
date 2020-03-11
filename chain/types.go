package chain

type field struct {
	Start      int
	Offset     int
	Value      string
	Recommends []string
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
