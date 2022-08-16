package controller

import "dreamt/pkg/models"

func (c Controller) GetKeywords(top int) ([]string, error) {
	dreams, err := c.DBController.GetDreamsFull()
	if err != nil {
		return nil, err
	}

	finalKeywords := []string{}

	for _, dream := range dreams {
		keywords := extractKwFromDream(dream, top)
		for _, kw := range keywords {
			if !contains(finalKeywords, kw) {
				finalKeywords = append(finalKeywords, kw)
			}

			if len(finalKeywords) >= top {
				return finalKeywords, nil
			}
		}
	}

	return finalKeywords, nil
}

func (c Controller) GetInterpret(kw string) (models.Interpretation, error) {
	return models.Interpretation{}, nil
}
