package persistence

import "dreamt/pkg/models"

type DatabaseController interface {
	GetDreams() ([]models.DreamHeader, error)
	GetDream(id string) (models.Dream, error)
	WriteDreams(dream models.Dream) (int64, error)
	DeleteDream(id string) error
	GetDreamsFull() ([]models.Dream, error)
}
