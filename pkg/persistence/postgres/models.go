package postgres

var tablesToMigrate = []interface{}{
	Tag{},
	Dream{},
}

type Dream struct {
	ID      int64 `gorm:"primary_key, auto_increment"`
	Title   string
	Content string
	Tags    []Tag `gorm:"foreignkey:DreamID"`
}

type Tag struct {
	Name    string `gorm:"primary_key"`
	DreamID int64  `gorm:"primary_key"`
}
