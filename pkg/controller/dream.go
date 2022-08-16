package controller

// import models
import "dreamt/pkg/models"

func (c Controller) GetDreams() ([]models.DreamHeader, error) {
	return c.DBController.GetDreams()
}

func (c Controller) GetDream(id string) (models.Dream, error) {
	return c.DBController.GetDream(id)
}

func (c Controller) WriteDreams(dream models.Dream) (int64, error) {
	return c.DBController.WriteDreams(dream)
}

func (c Controller) DeleteDream(id string) error {
	return c.DBController.DeleteDream(id)
}
