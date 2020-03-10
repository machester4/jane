package recommender

func addRecommends(word *Word) func()  {
	return func() {
		// levenshtein.ComputeDistance(word.ToString(), "word en dic")
		word.Recommendations = [3]string{word.ToString(),"c"}
	}
}

// Contexts are dictionaries with the respective words.
// have higher priority than the words of the language itself
func Recommend(chain *Chain, lang string, contexts []string) []*Word {
	// wordsDict := helpers.GetDictionary("chivito")
	words := chain.GetWords()
	var steps []func()

	for _, word := range words {
		steps = append(steps, addRecommends(word))
	}

	stages := []*Stage{
		{
			Name: "Add Recommends",
			Steps: steps,
		},
	}
	pipeline := Pipeline{Stages: stages}
	pipeline.Start(false)

	return words
}
