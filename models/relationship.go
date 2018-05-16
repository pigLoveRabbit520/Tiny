package models

type Relationship struct {
	Cid  uint   `gorm:"type:int(10) unsigned;primary_key" json:"cid"`
	Mid  uint   `gorm:"type:int(10) unsigned;primary_key" json:"mid"`
}
