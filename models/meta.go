package models

import (
	"github.com/salamander-mh/SalamanderBlog/config"
	"github.com/salamander-mh/SalamanderBlog/db"
)

type Meta struct {
	Mid         uint   `gorm:"primary_key" json:"mid"`
	Name        string `gorm:"type:varchar(200)" json:"name"`
	Slug        string `gorm:"type:varchar(200)" json:"slug"`
	Type        string `gorm:"type:varchar(32)" json:"type"`
	Description string `gorm:"type:varchar(200)" json:"description"`
	Count       uint   `gorm:"type:int(10) unsigned" json:"count"`
	Order       uint   `gorm:"type:int(10) unsigned" json:"order"`
	Parent      uint   `gorm:"type:int(10) unsigned" json:"parent"`
}

func (Meta) TableName() string {
	return config.Conf.DB.Prefix + "metas"
}

func GetAllCategories() ([]Meta, error) {
	var metasAll []Meta
	q := db.DB.Select("mid, name").Where("type = ?", "category").Find(&metasAll)
	return metasAll, q.Error
}

func GetCategories(page int) ([]Meta, int, error) {
	var metas []Meta
	var count int
	q := db.DB.Select("mid, name, slug, count, parent").Where("type = ?", "category").
		Offset((page - 1) * PAGE_SIZE).Limit(PAGE_SIZE).
		Find(&metas)
	q.Offset(-1).Limit(-1).Count(&count)
	return metas, count, q.Error
}

func GetTags(page int) ([]Meta, int, error) {
	var metas []Meta
	var count int
	q := db.DB.Select("mid, name, slug, count").Where("type = ?", "tag").
		Offset((page - 1) * PAGE_SIZE).Limit(PAGE_SIZE).
		Find(&metas)
	q.Offset(-1).Limit(-1).Count(&count)
	return metas, count, q.Error
}
