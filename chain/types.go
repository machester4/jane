package chain

type field struct {
	Start      int
	Length     int
	Value      string
	Recommends [3]string
}
type Word = field
type Punct = field
type Article struct {
	Start      int
	Length     int
	Value      string
	Noun       *Word
	Recommends [3]string
}
type Chain struct {
	Words       []*Word
	Pucts       []*Punct
	Articles    []*Article
	headArticle *Article // For set Noun (sustantivo luego del articulo)
}
