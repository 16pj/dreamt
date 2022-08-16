package controller

import (
	"dreamt/pkg/models"

	textrank "github.com/DavidBelicza/TextRank/v2"
)

func extractKwFromDream(dream models.Dream, top int) []string {
	return mostFrequentWords(dream.Title+" "+dream.Content, top)
}

// return the most frequent words from string
func mostFrequentWords(sentence string, top int) []string {
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

func contains[T comparable](slice []T, kw T) bool {
	for _, s := range slice {
		if s == kw {
			return true
		}
	}
	return false
}
