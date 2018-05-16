package models

type Option struct {
	Name  string `gorm:"type:varchar(32);primary_key" json:"name"`
	User  uint   `gorm:"type:int(10) unsigned;primary_key" json:"user"`
	Value string `gorm:"type:text" json:"value"`
}
