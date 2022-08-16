package postgres

var tablesToMigrate = []interface{}{
	DreamPG{},
	TagsPG{},
}

type DreamPG struct {
	ID      int64 `gorm:"primary_key, auto_increment"`
	Title   string
	Content string
	Tags    []TagsPG `gorm:"many2many:dream_tags;foreignkey:dream_id"`
}

type TagsPG struct {
	Name    string `gorm:"primary_key"`
	DreamID int64  `gorm:"primary_key"`
}
