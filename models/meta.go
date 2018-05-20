package models

import "github.com/salamander-mh/SalamanderBlog/db"

type Meta struct {
	Mid         uint   `gorm:"primary_key" json:"mid"`
	Name        string `gorm:"type:varchar(200)" json:"name"`
	Slug        string `gorm:"type:varchar(200)" json:"slug"`
	Type        string `gorm:"type:varchar(32)" json:"type"`
	Description string `gorm:"type:varchar(200)" json:"description"`
	Count       uint   `gorm:"type:int(10) unsigned" json:"count"`
	Order       uint   `gorm:"type:int(10) unsigned" json:"order"`
}

func GetAllCategories() ([]Meta, error) {
	var metasAll []Meta
	q := db.DB.Select("mid, name").Where("type = ?", "category").Find(&metasAll)
	return metasAll, q.Error
}

func GetCategories()  {
	
}
