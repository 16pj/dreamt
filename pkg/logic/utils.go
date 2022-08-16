package logic

import (
	textrank "github.com/DavidBelicza/TextRank/v2"
)

// return the most frequent words from string
func MostFrequentWords(sentence string, top int) []string {
	kw := []string{}

	tr := textrank.NewTextRank()
	rule := textrank.NewDefaultRule()
	language := textrank.NewDefaultLanguage()
	algorithmDef := textrank.NewDefaultAlgorithm()
	tr.Populate(sentence, language, rule)
	tr.Ranking(algorithmDef)
	rankedWords := textrank.FindSingleWords(tr)

	for idx, w := range rankedWords {
		if idx >= top {
			break
		}

		kw = append(kw, w.Word)
	}

	return kw
}

func Contains[T comparable](slice []T, kw T) bool {
	for _, s := range slice {
		if s == kw {
			return true
		}
	}
	return false
}
