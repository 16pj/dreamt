package controller

import (
	"dreamt/pkg/logic"
	"dreamt/pkg/models"
)

func (c Controller) GetKeywords(top int) ([]string, error) {
	dreams, err := c.DBController.GetDreamsFull()
	if err != nil {
		return nil, err
	}

	blob := ""
	for _, dream := range dreams {
		blob += dream.Content + " " + dream.Title + " "
	}

	return logic.MostFrequentWords(blob, top), nil
}

func (c Controller) GetInterpret(kw string) (models.Interpretation, error) {
	return models.Interpretation{}, nil
}
