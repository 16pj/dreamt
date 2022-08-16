package postgres

//import gorm
import (
	"dreamt/pkg/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PGController struct {
	DB *gorm.DB
}

func NewPGController(dsn string) PGController {
	if dsn == "" {
		dsn = "host=localhost port=5432 user=dbuser dbname=test3 password=dbuser sslmode=disable"
	}

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	for _, t := range tablesToMigrate {
		db.AutoMigrate(t)
	}

	return PGController{
		DB: db,
	}
}

func (p PGController) GetDreams() ([]models.DreamHeader, error) {
	dreamHeaders := []models.DreamHeader{}
	var pgDoc []Dream
	if err := p.DB.Select("id, title").Find(&pgDoc).Error; err != nil {
		return nil, err
	}

	for _, dream := range pgDoc {
		dreamHeaders = append(dreamHeaders, models.DreamHeader{
			Id:    dream.ID,
			Title: dream.Title,
		})
	}

	return dreamHeaders, nil
}

func (p PGController) GetDream(id string) (models.Dream, error) {
	var pgDoc Dream

	if err := p.DB.Debug().
		Preload("Tags").
		First(&pgDoc, id).Error; err != nil {
		return models.Dream{}, err
	}

	dream := models.Dream{
		Id:      pgDoc.ID,
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
	pgDoc := Dream{
		Title:   dream.Title,
		Content: dream.Content,
	}

	for _, tag := range dream.Tags {
		pgDoc.Tags = append(pgDoc.Tags, Tag{
			Name: tag,
		})
	}

	err := p.DB.Create(&pgDoc).Error
	return pgDoc.ID, err
}

func (p PGController) DeleteDream(id string) error {
	err := p.DB.Delete(&Dream{}, id).Error
	return err
}

func (p PGController) GetDreamsFull() ([]models.Dream, error) {
	dreams := []models.Dream{}
	var pgDoc []Dream
	if err := p.DB.Find(&pgDoc).Error; err != nil {
		return nil, err
	}

	for _, dream := range pgDoc {
		dreams = append(dreams, models.Dream{
			Id:      dream.ID,
			Title:   dream.Title,
			Content: dream.Content,
		})
	}

	return dreams, nil
}
