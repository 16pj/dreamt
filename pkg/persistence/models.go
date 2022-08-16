package persistence

var TablesToMigrate = []interface{}{
	Tag{},
	DreamPG{},
}

type DreamPG struct {
	ID      int64  `gorm:"primary_key, auto_increment" bson:"id"`
	Title   string `bson:"title"`
	Content string `bson:"content"`
	Tags    []Tag  `gorm:"foreignkey:DreamID"`
}

func (d DreamPG) TableName() string {
	return "dream"
}

type Tag struct {
	Name    string `gorm:"primary_key" bson:"name"`
	DreamID int64  `gorm:"primary_key" bson:"dreamID"`
}

type DreamMG struct {
	ID      int64    `gorm:"primary_key, auto_increment" bson:"id"`
	Title   string   `bson:"title"`
	Content string   `bson:"content"`
	Tags    []string `bson:"tags"`
}
