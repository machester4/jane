package lib

type RecommendItemSuggestion struct {
	Entry    string `json:"value"`
	Distance int    `json:"distance"`
}

type RecommendResultItem struct {
	Start      int                        `json:"start"`
	Offset     int                        `json:"offset"`
	Value      string                     `json:"value"`
	Same       bool                       `json:"same"`
	Forbidden  bool                       `json:"forbidden"`
	Recommends []*RecommendItemSuggestion `json:"recommends"`
}

type RecommendResult struct {
	Words []*RecommendResultItem `json:"words"`
}
