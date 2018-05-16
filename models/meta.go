package models

type Meta struct {
	Mid         uint   `gorm:"primary_key" json:"mid"`
	Name        string `gorm:"type:varchar(200)" json:"name"`
	Slug        string `gorm:"type:varchar(200)" json:"slug"`
	Type        string `gorm:"type:varchar(32)" json:"type"`
	Description string `gorm:"type:varchar(200)" json:"description"`
	Count       uint   `gorm:"type:int(10) unsigned" json:"count"`
	Order       uint   `gorm:"type:int(10) unsigned" json:"order"`
}
