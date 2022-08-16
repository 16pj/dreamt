package postgres

//import gorm
import (
	"dreamt/pkg/models"
	"dreamt/pkg/persistence"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PGController struct {
	db *gorm.DB
}

func NewPGController(dsn string) persistence.DatabaseController {
	if dsn == "" {
		dsn = "host=localhost port=5432 user=dbuser dbname=test3 password=dbuser sslmode=disable"
	}

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	for _, t := range persistence.TablesToMigrate {
		db.AutoMigrate(t)
	}

	return PGController{
		db: db,
	}
}

func (p PGController) GetDreams() ([]models.DreamHeader, error) {
	dreamHeaders := []models.DreamHeader{}
	var pgDoc []persistence.DreamPG
	if err := p.db.Select("id, title").Find(&pgDoc).Error; err != nil {
		return nil, err
	}

	for _, dream := range pgDoc {
		dreamHeaders = append(dreamHeaders, models.DreamHeader{
			ID:    dream.ID,
			Title: dream.Title,
		})
	}

	return dreamHeaders, nil
}

func (p PGController) GetDream(id string) (models.Dream, error) {
	var pgDoc persistence.DreamPG

	if err := p.db.Debug().
		Preload("Tags").
		First(&pgDoc, id).Error; err != nil {
		return models.Dream{}, err
	}

	dream := models.Dream{
		ID:      pgDoc.ID,
		Title:   pgDoc.Title,
		Content: pgDoc.Content,
	}

	fmt.Println(pgDoc.Tags)

	for _, tag := range pgDoc.Tags {
		dream.Tags = append(dream.Tags, tag.Name)
	}

	return dream, nil
}

func (p PGController) WriteDreams(dream models.Dream) (int64, error) {
	pgDoc := persistence.DreamPG{
		Title:   dream.Title,
		Content: dream.Content}

	for _, tag := range dream.Tags {
		pgDoc.Tags = append(pgDoc.Tags, persistence.Tag{
			Name: tag,
		})
	}

	err := p.db.Create(&pgDoc).Error
	return pgDoc.ID, err
}

func (p PGController) DeleteDream(id string) error {
	err := p.db.Delete(&persistence.DreamPG{}, id).Error
	return err
}

func (p PGController) GetDreamsFull() ([]models.Dream, error) {
	dreams := []models.Dream{}
	var pgDoc []persistence.DreamPG
	if err := p.db.Find(&pgDoc).Error; err != nil {
		return nil, err
	}

	for _, dream := range pgDoc {
		dreams = append(dreams, models.Dream{
			ID:      dream.ID,
			Title:   dream.Title,
			Content: dream.Content,
		})
	}

	return dreams, nil
}
